package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("empfetcher", func() {
	Title("Employee Data Fetcher")
})

var _ = Service("empfetcher", func() {
	Error("unauthorized", func() {
		Description("This error occurs if the request doesn't have a valid JWT token")
	})
	Error("forbidden", func() {
		Description("This error occurs if the JWT token included in the request doesn't have the necessary scopes")
	})
	Error("bad_gateway", func() {
		Description("This error occurs when one of the downstream services is unavailable")
	})
	Error("bad_request", func() {
		Description("This error is returned if bad request is sent")
	})
	Error("internal_error", func() {
		Description("This error is returned if an unexpected error happens")
	})
	HTTP(func() {
		Path("api/v1/employees")
		Response("unauthorized", StatusUnauthorized)
		Response("forbidden", StatusForbidden)
		Response("bad_gateway", StatusBadGateway)
		Response("bad_request", StatusBadRequest)
		Response("internal_error", StatusInternalServerError)

	})
	Method("add", func() {
		Description(`Adds an Employee Details`)
		Payload(EmployeeData)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})

	})

	Method("update", func() {
		Description("Updates an Employee Details")
		Payload(EmployeeData)
		Error("not_found")
		HTTP(func() {
			PUT("/")
			Response(StatusNoContent)
		})

	})

	Method("list", func() {
		Description("List All Employee Details")
		Result(ArrayOf(EmployeeData))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Response(StatusNotFound)
		})

	})

	Method("show", func() {
		Description("Show Employee Details based on ID")
		Payload(func() {
			Field(1, "id", String, "ID is the unique id of an employee")
			Required("id")
		})
		Result(EmployeeData)
		Error("not_found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response(StatusNotFound)
		})

	})

	Method("delete", func() {
		Description("Delete Employee Details")
		Payload(func() {
			Field(1, "id", String, "ID is the unique id of an employee")
			Field(2, "permdelete", Boolean, "Delete Permanently if this is yes")
			Required("id")
		})
		Error("not_found")
		HTTP(func() {
			DELETE("/{id}")
			Param("permdelete", func() {
				Default(0)
			})
			Response(StatusNoContent)
		})
	})

	Method("restore", func() {
		Description("Restore Deleted Employee Details based on ID")
		Description("Restores an Employee Details")
		Payload(func() {
			Field(1, "id", String, "ID is the unique id of an employee")
			Required("id")
		})
		Error("not_found")
		HTTP(func() {
			PUT("/{id}")
			Response(StatusNoContent)
		})
	})

	Method("viewdeleted", func() {
		Description("View All deactivated Employee Details")
		Result(ArrayOf(EmployeeData))
		HTTP(func() {
			GET("/deactivated")
			Response(StatusOK)
			Response(StatusNotFound)
		})

	})

	Method("search", func() {
		Description("Search employees by name")
		Payload(func() {
			Field(1, "name", String, "Name of an employee")
			Required("name")
		})
		Result(ArrayOf(EmployeeData))
		HTTP(func() {
			GET("/search/{name}")
			Response(StatusOK)
			Response(StatusNotFound)
		})
	})

})
