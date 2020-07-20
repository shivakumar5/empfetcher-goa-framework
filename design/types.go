package design

import (
	. "goa.design/goa/v3/dsl"
)

// EmployeeData - Employee Information
var EmployeeData = Type("EmployeePayload", func() {

	Field(1, "id", String, "Unique ID of an Employee", func() {
		Example("fgfhjsddctybnjgjh")
	})
	Field(2, "name", String, "Name of an Employee", func() {
		Example("shiva")
	})
	Field(3, "department", String, "The Department of an Employee", func() {
		Example("development")
	})
	Field(4, "address", String, "Address of an Employee", func() {
		Example("Bangalore")
	})
	Field(5, "skills", String, "Skillsets of an Employee", func() {
		Example("golang, docker")
	})
	Required("id", "name", "department", "address", "skills")
})
