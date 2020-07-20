package mssql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	types "github.com/flexera/empfetcher/gen/empfetcher"
)

type (

	// Client interface
	Client interface {
		Healthy(context.Context) error
		Add(context.Context, *types.EmployeePayload) error
		Update(context.Context, types.EmployeePayload) error
		List(context.Context) ([]*types.EmployeePayload, error)
		Show(context.Context, types.ShowPayload) (*types.EmployeePayload, error)
		MarkAsDeleted(context.Context, types.DeletePayload) error
		Delete(context.Context, types.DeletePayload) error
		Restore(context.Context, types.RestorePayload) error
		Viewdeleted(context.Context) ([]*types.EmployeePayload, error)
		Search(context.Context, types.SearchPayload) ([]*types.EmployeePayload, error)
	}
	client struct {
		mssql string
	}
)

//OpenConnection connects to the MSSQL database
func (c *client) OpenConnection() (*sql.DB, error) {
	svc, err := sql.Open("mssql", c.mssql)
	if err != nil {
		return nil, err
	}
	return svc, err
}

//NewClient Definition for MSSQL Connection ...
func NewClient(credentials string) Client {
	return &client{mssql: credentials}
}

//Check the Health of MSSQL Connection and Notify Error
func (c *client) Healthy(context.Context) error {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	defer svc.Close()
	err = svc.Ping()
	if err != nil {
		return err
	}
	return err
}

// Add Employee Data
func (c *client) Add(ctx context.Context, empData *types.EmployeePayload) error {
	fmt.Println("Inserting")

	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	defer svc.Close()
	sqlStatement := fmt.Sprintf("INSERT INTO EmployeeData (id,name,department,address,skills) VALUES ('%s','%s','%s','%s','%s');",
		empData.ID, empData.Name, empData.Department, empData.Address, empData.Skills)
	_, err = svc.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

// Update - update employee data
func (c *client) Update(ctx context.Context, empData types.EmployeePayload) error {
	fmt.Println("Updating")
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	defer svc.Close()
	sqlStatement := `update EmployeeData set name = $1, department= $2, address = $3, skills = $4 where ID = $5 `
	var rows sql.Result
	rows, err = svc.Exec(sqlStatement, &empData.Name, &empData.Department, &empData.Address, &empData.Skills, &empData.ID)

	var updatedRow int64
	updatedRow, err = rows.RowsAffected()
	if updatedRow == 0 {
		noRecords := errors.New("No employee data to be updated")
		log.Fatal("error: No employee data to be updated", err)
		return noRecords
	}
	if err != nil {
		unHandledEr := errors.New("Internal Error")
		log.Fatal("error: Internal Error", err)
		return unHandledEr
	}
	return nil

}

// List returns all the employee data
func (c *client) List(ctx context.Context) (empData []*types.EmployeePayload, err error) {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return nil, ConnectionEr
	}
	defer svc.Close()

	sqlStatement := `select id, name, department, address, skills from EmployeeData where softDelete is null`
	rows, err := svc.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item types.EmployeePayload
		err = rows.Scan(&item.ID, &item.Name, &item.Department, &item.Address, &item.Skills)
		if err != nil {
			return nil, err
		}
		empData = append(empData, &item)
	}
	return empData, nil
}

// Show - Employee data based on Unique ID
func (c *client) Show(ctx context.Context, empID types.ShowPayload) (*types.EmployeePayload, error) {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return nil, ConnectionEr
	}
	defer svc.Close()
	sqlStatement := `Select id, name, department, address, skills from EmployeeData where id = $1 and softDelete is null `
	var item types.EmployeePayload
	rows, err := svc.Query(sqlStatement, &empID.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Name, &item.Department, &item.Address, &item.Skills)
		if err != nil {
			return nil, err
		}
	}
	return &item, nil
}

// MarkAsDeleted -- do softDelete
func (c *client) MarkAsDeleted(ctx context.Context, empID types.DeletePayload) error {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	sqlStatement := `update EmployeeData set softDelete = 1 where id = $1 `
	var rows sql.Result
	rows, err = svc.Exec(sqlStatement, &empID.ID)

	var updatedRow int64
	updatedRow, err = rows.RowsAffected()
	if updatedRow == 0 {
		noRecords := errors.New("No employee data to be deleted")
		log.Fatal("error: No employee data to be deleted", err)
		return noRecords
	}
	if err != nil {
		unHandledEr := errors.New("Internal Error")
		log.Fatal("error: Internal Error", err)
		return unHandledEr
	}
	return nil
}

// Delete - delete employee record
func (c *client) Delete(ctx context.Context, empID types.DeletePayload) error {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	defer svc.Close()
	var sqlStatement string

	if !*empID.Permdelete {
		c.MarkAsDeleted(ctx, empID)
	} else {

		sqlStatement = `Delete from EmployeeData where id = $1 `
		var rows sql.Result
		rows, err = svc.Exec(sqlStatement, &empID.ID)

		var deletedRow int64
		deletedRow, err = rows.RowsAffected()
		if deletedRow == 0 {
			noRecords := errors.New("No employee data to be deleted")
			log.Fatal("error: No employee data to be deleted", err)
			return noRecords
		}
		if err != nil {
			unHandledEr := errors.New("Internal Error")
			log.Fatal("error: Internal Error", err)
			return unHandledEr
		}
	}
	return nil
}

// Restore -- restore the softDeleted record
func (c *client) Restore(ctx context.Context, empID types.RestorePayload) error {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return ConnectionEr
	}
	sqlStatement := `update EmployeeData set softDelete = NULL where id = $1 `
	_, err = svc.Exec(sqlStatement, &empID.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Viewdeleted(ctx context.Context) (empData []*types.EmployeePayload, err error) {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return nil, ConnectionEr
	}
	defer svc.Close()

	sqlStatement := `select id, name, department, address, skills from EmployeeData where softDelete = 1`
	rows, err := svc.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item types.EmployeePayload
		err = rows.Scan(&item.ID, &item.Name, &item.Department, &item.Address, &item.Skills)
		if err != nil {
			return nil, err
		}
		empData = append(empData, &item)
	}
	return empData, nil

}

// Show - Employee data based on Unique ID
func (c *client) Search(ctx context.Context, matchParam types.SearchPayload) (empData []*types.EmployeePayload, err error) {
	svc, err := c.OpenConnection()
	if err != nil {
		ConnectionEr := errors.New("Connection to MSSQL Database failed")
		return nil, ConnectionEr
	}
	defer svc.Close()
	sqlStatement := `Select id, name, department, address, skills from EmployeeData where name like '%` + matchParam.SearchString + `%' or department like '%` + matchParam.SearchString + `%' or address like '%` + matchParam.SearchString + `%' or skills like '%` + matchParam.SearchString + `%'`
	rows, err := svc.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item types.EmployeePayload
		err = rows.Scan(&item.ID, &item.Name, &item.Department, &item.Address, &item.Skills)
		if err != nil {
			return nil, err
		}
		empData = append(empData, &item)
	}
	return empData, nil
}
