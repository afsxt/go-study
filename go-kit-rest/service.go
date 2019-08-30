package main

import "errors"

type Service interface {
	Add(a, b int) int

	Subtract(a, b int) int

	Multiply(a, b int) int

	Divide(a, b int) (int, error)
}

type arithmeticService struct {
}

func (s arithmeticService) Add(a, b int) int {
	return a + b
}

func (s arithmeticService) Subtract(a, b int) int {
	return a - b
}

func (s arithmeticService) Multiply(a, b int) int {
	return a * b
}

func (s arithmeticService) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero")
	}
	return a / b, nil
}
