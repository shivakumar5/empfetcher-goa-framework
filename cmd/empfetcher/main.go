package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	empfetcherapi "github.com/flexera/empfetcher"
	empfetcher "github.com/flexera/empfetcher/gen/empfetcher"
	mssql "github.com/flexera/empfetcher/services"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")

		sqlServer = flag.String("sql-server", "localhost", "SQL server")
		userName  = flag.String("username", "sa", "Username for SQL server")
		password  = flag.String("password", "", "Password for SQL server")
		database  = flag.String("database", "employee", "Name of the SQL database")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[empfetcherapi] ", log.Ltime)
	}

	// Create Context
	var ctx context.Context

	cred := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", *sqlServer, *userName, *password, *database)
	mssqlconn := mssql.NewClient(cred)
	err := mssqlconn.Healthy(ctx)
	if err != nil {
		log.Fatal("SQLClient - Health check fail!", err)
		panic(err)
	}

	// Initialize the services.
	var (
		empfetcherSvc empfetcher.Service
	)
	{
		empfetcherSvc = empfetcherapi.NewEmpfetcher(mssqlconn)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		empfetcherEndpoints *empfetcher.Endpoints
	)
	{
		empfetcherEndpoints = empfetcher.NewEndpoints(empfetcherSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:80"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, empfetcherEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
