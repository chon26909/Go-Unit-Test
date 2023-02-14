package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}

	c.On("GetCustomer", 1).Return("bond", 19, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("no found"))

	name, age, err := c.GetCustomer(2)
	if err != nil {
		return
	}

	fmt.Println(name, age)
}

type CustomerRepository interface {
	GetCustomer(id int) (name string, age int, err error)
}
type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
