package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

func MakeArithmeticEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)
		a := req.A
		b := req.B
		var res int
		var callError error

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Substract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, callError = svc.Divide(a, b)
		}

		return ArithmeticResponse{res, callError}, nil
	}
}