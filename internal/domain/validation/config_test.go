package validation

type test struct {
	description string
	field string
	wantError bool
	errorExpected error
}
