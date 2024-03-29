// Code generated by goa v3.2.0, DO NOT EDIT.
//
// empfetcher HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/flexera/empfetcher/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	empfetcher "github.com/flexera/empfetcher/gen/empfetcher"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "empfetcher" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddEmpfetcherPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the empfetcher add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*empfetcher.EmployeePayload)
		if !ok {
			return goahttp.ErrInvalidType("empfetcher", "add", "*empfetcher.EmployeePayload", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("empfetcher", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the empfetcher
// add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeAddResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body AddUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "add", err)
			}
			err = ValidateAddUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "add", err)
			}
			return nil, NewAddUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body AddForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "add", err)
			}
			err = ValidateAddForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "add", err)
			}
			return nil, NewAddForbidden(&body)
		case http.StatusBadGateway:
			var (
				body AddBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "add", err)
			}
			err = ValidateAddBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "add", err)
			}
			return nil, NewAddBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body AddBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "add", err)
			}
			err = ValidateAddBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "add", err)
			}
			return nil, NewAddBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body AddInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "add", err)
			}
			err = ValidateAddInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "add", err)
			}
			return nil, NewAddInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "empfetcher" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*empfetcher.EmployeePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("empfetcher", "update", "*empfetcher.EmployeePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateEmpfetcherPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the empfetcher
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*empfetcher.EmployeePayload)
		if !ok {
			return goahttp.ErrInvalidType("empfetcher", "update", "*empfetcher.EmployeePayload", v)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("empfetcher", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// empfetcher update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body UpdateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "update", err)
			}
			err = ValidateUpdateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "update", err)
			}
			return nil, NewUpdateUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body UpdateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "update", err)
			}
			err = ValidateUpdateForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "update", err)
			}
			return nil, NewUpdateForbidden(&body)
		case http.StatusBadGateway:
			var (
				body UpdateBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "update", err)
			}
			err = ValidateUpdateBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "update", err)
			}
			return nil, NewUpdateBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body UpdateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "update", err)
			}
			err = ValidateUpdateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "update", err)
			}
			return nil, NewUpdateInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "empfetcher" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListEmpfetcherPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the
// empfetcher list endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateEmployeePayloadResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			res := NewListEmployeePayloadOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			err = ValidateListUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			return nil, NewListUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ListForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			err = ValidateListForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			return nil, NewListForbidden(&body)
		case http.StatusBadGateway:
			var (
				body ListBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			err = ValidateListBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			return nil, NewListBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body ListInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "list", err)
			}
			err = ValidateListInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "list", err)
			}
			return nil, NewListInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "empfetcher" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*empfetcher.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("empfetcher", "show", "*empfetcher.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowEmpfetcherPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the
// empfetcher show endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowOKResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			res := NewShowEmployeePayloadOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ShowUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			return nil, NewShowUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ShowForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			return nil, NewShowForbidden(&body)
		case http.StatusBadGateway:
			var (
				body ShowBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			return nil, NewShowBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body ShowBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			return nil, NewShowBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body ShowInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "show", err)
			}
			err = ValidateShowInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "show", err)
			}
			return nil, NewShowInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "empfetcher" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*empfetcher.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("empfetcher", "delete", "*empfetcher.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteEmpfetcherPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the empfetcher
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*empfetcher.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("empfetcher", "delete", "*empfetcher.DeletePayload", v)
		}
		values := req.URL.Query()
		if p.Permdelete != nil {
			values.Add("permdelete", fmt.Sprintf("%v", *p.Permdelete))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// empfetcher delete endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body DeleteUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "delete", err)
			}
			err = ValidateDeleteUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "delete", err)
			}
			return nil, NewDeleteUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body DeleteForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "delete", err)
			}
			err = ValidateDeleteForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "delete", err)
			}
			return nil, NewDeleteForbidden(&body)
		case http.StatusBadGateway:
			var (
				body DeleteBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "delete", err)
			}
			err = ValidateDeleteBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "delete", err)
			}
			return nil, NewDeleteBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body DeleteInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "delete", err)
			}
			err = ValidateDeleteInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "delete", err)
			}
			return nil, NewDeleteInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "delete", resp.StatusCode, string(body))
		}
	}
}

// BuildRestoreRequest instantiates a HTTP request object with method and path
// set to call the "empfetcher" service "restore" endpoint
func (c *Client) BuildRestoreRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*empfetcher.RestorePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("empfetcher", "restore", "*empfetcher.RestorePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RestoreEmpfetcherPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "restore", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRestoreResponse returns a decoder for responses returned by the
// empfetcher restore endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRestoreResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeRestoreResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body RestoreUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "restore", err)
			}
			err = ValidateRestoreUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "restore", err)
			}
			return nil, NewRestoreUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body RestoreForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "restore", err)
			}
			err = ValidateRestoreForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "restore", err)
			}
			return nil, NewRestoreForbidden(&body)
		case http.StatusBadGateway:
			var (
				body RestoreBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "restore", err)
			}
			err = ValidateRestoreBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "restore", err)
			}
			return nil, NewRestoreBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body RestoreBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "restore", err)
			}
			err = ValidateRestoreBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "restore", err)
			}
			return nil, NewRestoreBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body RestoreInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "restore", err)
			}
			err = ValidateRestoreInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "restore", err)
			}
			return nil, NewRestoreInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "restore", resp.StatusCode, string(body))
		}
	}
}

// BuildViewdeletedRequest instantiates a HTTP request object with method and
// path set to call the "empfetcher" service "viewdeleted" endpoint
func (c *Client) BuildViewdeletedRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ViewdeletedEmpfetcherPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "viewdeleted", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeViewdeletedResponse returns a decoder for responses returned by the
// empfetcher viewdeleted endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeViewdeletedResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeViewdeletedResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ViewdeletedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateEmployeePayloadResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			res := NewViewdeletedEmployeePayloadOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ViewdeletedUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			err = ValidateViewdeletedUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			return nil, NewViewdeletedUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ViewdeletedForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			err = ValidateViewdeletedForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			return nil, NewViewdeletedForbidden(&body)
		case http.StatusBadGateway:
			var (
				body ViewdeletedBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			err = ValidateViewdeletedBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			return nil, NewViewdeletedBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body ViewdeletedBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			err = ValidateViewdeletedBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			return nil, NewViewdeletedBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body ViewdeletedInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "viewdeleted", err)
			}
			err = ValidateViewdeletedInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "viewdeleted", err)
			}
			return nil, NewViewdeletedInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "viewdeleted", resp.StatusCode, string(body))
		}
	}
}

// BuildSearchRequest instantiates a HTTP request object with method and path
// set to call the "empfetcher" service "search" endpoint
func (c *Client) BuildSearchRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SearchEmpfetcherPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("empfetcher", "search", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSearchRequest returns an encoder for requests sent to the empfetcher
// search server.
func EncodeSearchRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*empfetcher.SearchPayload)
		if !ok {
			return goahttp.ErrInvalidType("empfetcher", "search", "*empfetcher.SearchPayload", v)
		}
		body := NewSearchRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("empfetcher", "search", err)
		}
		return nil
	}
}

// DecodeSearchResponse returns a decoder for responses returned by the
// empfetcher search endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSearchResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeSearchResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SearchResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateEmployeePayloadResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			res := NewSearchEmployeePayloadOK(body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body SearchUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			err = ValidateSearchUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			return nil, NewSearchUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body SearchForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			err = ValidateSearchForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			return nil, NewSearchForbidden(&body)
		case http.StatusBadGateway:
			var (
				body SearchBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			err = ValidateSearchBadGatewayResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			return nil, NewSearchBadGateway(&body)
		case http.StatusBadRequest:
			var (
				body SearchBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			err = ValidateSearchBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			return nil, NewSearchBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body SearchInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("empfetcher", "search", err)
			}
			err = ValidateSearchInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("empfetcher", "search", err)
			}
			return nil, NewSearchInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("empfetcher", "search", resp.StatusCode, string(body))
		}
	}
}

// unmarshalEmployeePayloadResponseToEmpfetcherEmployeePayload builds a value
// of type *empfetcher.EmployeePayload from a value of type
// *EmployeePayloadResponse.
func unmarshalEmployeePayloadResponseToEmpfetcherEmployeePayload(v *EmployeePayloadResponse) *empfetcher.EmployeePayload {
	res := &empfetcher.EmployeePayload{
		ID:         *v.ID,
		Name:       *v.Name,
		Department: *v.Department,
		Address:    *v.Address,
		Skills:     *v.Skills,
	}

	return res
}
