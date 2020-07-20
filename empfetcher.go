package empfetcherapi

import (
	"context"
	"log"

	empfetcher "github.com/flexera/empfetcher/gen/empfetcher"
	mssql "github.com/flexera/empfetcher/services"
	guuid "github.com/google/uuid"
)

// empfetcher service example implementation.
// The example methods log the requests and return zero values.
type empfetchersrvc struct {
	mssqlc mssql.Client
}

// NewEmpfetcher returns the empfetcher service implementation.
func NewEmpfetcher(mssqlc mssql.Client) empfetcher.Service {
	return &empfetchersrvc{mssqlc}
}

// Adds an Employee Details
func (s *empfetchersrvc) Add(ctx context.Context, p *empfetcher.EmployeePayload) (err error) {
	log.Print("empfetcher.add")
	p.ID = guuid.New().String()
	err = s.mssqlc.Add(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

// Updates an Employee Details
// TO DO: Yet to implement the code for Update
func (s *empfetchersrvc) Update(ctx context.Context, p *empfetcher.EmployeePayload) (err error) {
	log.Print("empfetcher.update")
	return nil
}

// List All Employee Details
func (s *empfetchersrvc) List(ctx context.Context) (res []*empfetcher.EmployeePayload, err error) {
	log.Print("empfetcher.list")
	res, err = s.mssqlc.List(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Show Employee Details based on ID
func (s *empfetchersrvc) Show(ctx context.Context, p *empfetcher.ShowPayload) (res *empfetcher.EmployeePayload, err error) {
	log.Print("empfetcher.show")
	res, err = s.mssqlc.Show(ctx, *p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete Employee Details
func (s *empfetchersrvc) Delete(ctx context.Context, p *empfetcher.DeletePayload) (err error) {
	log.Print("empfetcher.delete")
	err = s.mssqlc.Delete(ctx, *p)
	if err != nil {
		return err
	}
	return nil
}

// Restores an Employee Details
func (s *empfetchersrvc) Restore(ctx context.Context, p *empfetcher.RestorePayload) (err error) {
	log.Print("empfetcher.restore")
	err = s.mssqlc.Restore(ctx, *p)
	if err != nil {
		return err
	}
	return nil
}

// View All deactivated Employee Details
func (s *empfetchersrvc) Viewdeleted(ctx context.Context) (res []*empfetcher.EmployeePayload, err error) {
	log.Print("empfetcher.viewdeleted")
	res, err = s.mssqlc.Viewdeleted(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Search employees by name
// TO DO: Yet to implement the code for Search
func (s *empfetchersrvc) Search(ctx context.Context, p *empfetcher.SearchPayload) (res []*empfetcher.EmployeePayload, err error) {
	log.Print("empfetcher.search")
	return
}
