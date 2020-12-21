
package petstore

// Bool returns a pointer to the given bool
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer to the given int.
func Int(v int) *int {
	return &v
}

// String returns a pointer to the given string.
func String(v string) *string {
	return &v
}