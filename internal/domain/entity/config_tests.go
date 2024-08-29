package entity

type test[T any] struct {
	description string
	fields T
	wantError bool
	errorQuantityExpected int
} 