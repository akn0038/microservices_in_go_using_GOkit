package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func reverseString(s string) string {
	strRev := ""
	for _, c := range s {
		strRev = string(c) + strRev
	}
	return strRev
}

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
	Reverse(string) (string, error)
	isPelindrome(string) string
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

func (stringService) Reverse(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return reverseString(s), nil
}

func (stringService) isPelindrome(s string) string {
	revStr := reverseString(s)
	if s == revStr {
		return "True"
	}
	return "False"
}

var ErrEmpty = errors.New("empty string")

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

type reverseRequest struct {
	S string `json:"s"`
}

type reverseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type isPelindromeRequest struct {
	S string `json: "s"`
}

type isPelindromeResponse struct {
	Result string `json: "result"`
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func makeReverseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(reverseRequest)
		v, err := svc.Reverse(req.S)
		if err != nil {
			return reverseResponse{v, err.Error()}, nil
		}
		return reverseResponse{v, ""}, nil
	}
}

func makeIsPelindromeEndpoint(svc stringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(isPelindromeRequest)
		v := svc.isPelindrome(req.S)
		return isPelindromeResponse{v}, nil
	}
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	reverseHandler := httptransport.NewServer(
		makeReverseEndpoint(svc),
		decodeReverseRequest,
		encodeResponse,
	)

	ispelindrome := httptransport.NewServer(
		makeIsPelindromeEndpoint(svc),
		decodeIsPelindromeRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/reverse", reverseHandler)
	http.Handle("/ispelindrome", ispelindrome)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeReverseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request reverseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeIsPelindromeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request isPelindromeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
