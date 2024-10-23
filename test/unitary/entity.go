package unitary

import "github.com/jaswdr/faker"

var Fake = faker.New()

type TestField[T any] struct {
	TestDescription string
	Handler         func(T)
}
