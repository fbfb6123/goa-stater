// Code generated by goa v3.4.3, DO NOT EDIT.
//
// HTTP request path constructors for the term_limit service.
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"fmt"
)

// AddTermLimitPath returns the URL path to the term_limit service add HTTP endpoint.
func AddTermLimitPath(c2 int, d int) string {
	return fmt.Sprintf("/term_limit/%v/%v", c2, d)
}
