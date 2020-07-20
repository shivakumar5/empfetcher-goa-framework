// Code generated by goa v3.2.0, DO NOT EDIT.
//
// empfetcher HTTP client CLI support package
//
// Command:
// $ goa gen github.com/flexera/empfetcher/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	empfetcher "github.com/flexera/empfetcher/gen/empfetcher"
)

// BuildAddPayload builds the payload for the empfetcher add endpoint from CLI
// flags.
func BuildAddPayload(empfetcherAddBody string) (*empfetcher.EmployeePayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(empfetcherAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"address\": \"Bangalore\",\n      \"department\": \"development\",\n      \"id\": \"fgfhjsddctybnjgjh\",\n      \"name\": \"shiva\",\n      \"skills\": \"golang, docker\"\n   }'")
		}
	}
	v := &empfetcher.EmployeePayload{
		ID:         body.ID,
		Name:       body.Name,
		Department: body.Department,
		Address:    body.Address,
		Skills:     body.Skills,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the empfetcher update endpoint
// from CLI flags.
func BuildUpdatePayload(empfetcherUpdateBody string, empfetcherUpdateID string) (*empfetcher.EmployeePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(empfetcherUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"address\": \"Bangalore\",\n      \"department\": \"development\",\n      \"name\": \"shiva\",\n      \"skills\": \"golang, docker\"\n   }'")
		}
	}
	var id string
	{
		id = empfetcherUpdateID
	}
	v := &empfetcher.EmployeePayload{
		Name:       body.Name,
		Department: body.Department,
		Address:    body.Address,
		Skills:     body.Skills,
	}
	v.ID = id

	return v, nil
}

// BuildShowPayload builds the payload for the empfetcher show endpoint from
// CLI flags.
func BuildShowPayload(empfetcherShowID string) (*empfetcher.ShowPayload, error) {
	var id string
	{
		id = empfetcherShowID
	}
	v := &empfetcher.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the empfetcher delete endpoint
// from CLI flags.
func BuildDeletePayload(empfetcherDeleteID string, empfetcherDeletePermdelete string) (*empfetcher.DeletePayload, error) {
	var err error
	var id string
	{
		id = empfetcherDeleteID
	}
	var permdelete bool
	{
		if empfetcherDeletePermdelete != "" {
			permdelete, err = strconv.ParseBool(empfetcherDeletePermdelete)
			if err != nil {
				return nil, fmt.Errorf("invalid value for permdelete, must be BOOL")
			}
		}
	}
	v := &empfetcher.DeletePayload{}
	v.ID = id
	v.Permdelete = &permdelete

	return v, nil
}

// BuildRestorePayload builds the payload for the empfetcher restore endpoint
// from CLI flags.
func BuildRestorePayload(empfetcherRestoreID string) (*empfetcher.RestorePayload, error) {
	var id string
	{
		id = empfetcherRestoreID
	}
	v := &empfetcher.RestorePayload{}
	v.ID = id

	return v, nil
}

// BuildSearchPayload builds the payload for the empfetcher search endpoint
// from CLI flags.
func BuildSearchPayload(empfetcherSearchName string) (*empfetcher.SearchPayload, error) {
	var name string
	{
		name = empfetcherSearchName
	}
	v := &empfetcher.SearchPayload{}
	v.Name = name

	return v, nil
}
