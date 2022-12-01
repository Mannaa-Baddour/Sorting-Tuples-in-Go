/*
 * Sorting Tuples in Go
 *
 * API that allows users to signup/login to the server, giving them the ability to create input files of type csv, and sort the column data in them based on a certain column.
 *
 * API version: 0.1.0
 * Contact: mannaabaddour29699@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
