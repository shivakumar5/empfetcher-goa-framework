package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("empfetcher", func() {
	Title("Employee Data Fetcher")
})

// // JWTAuth is the JWTSecurity DSL function for adding JWT support in the API
// var JWTAuth = JWTSecurity("jwt", func() {
// 	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
// 	Scope("api:read", "Read-only access")
// 	Scope("api:write", "Read and write access")
// })
