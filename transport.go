package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

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

//-------------------------------

func makeBooksEndpoint(svc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//		?\request.(booksRequest)
		b := svc.Books()
		//if err != nil {
		//	return booksResponse{Books:b}, nil
		//}
		return booksResponse{Books:b}, nil
	}
}

func makeBookEndpoint(svc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getBookRequest)
		//idString := req.ID
		//id, _  := strconv.Atoi(idString)
		//book, err := svc.GetBook(id)
		book, err := svc.GetBook(req.ID)
		if err != nil {
			return getBookResponse{Book:book}, nil
		}
		return getBookResponse{Book:book}, nil
	}
}


func makeSetBookEndpoint(svc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(setBookRequest)
		//idString := req.ID
		//id, _  := strconv.Atoi(idString)
		//book, err := svc.GetBook(id)
		err := svc.SetBook(req.Book)
		if err != nil {
			return setBookResponse{ok:true}, nil
		}
		return setBookResponse{ok:true}, nil
	}
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


//----------------------
func decodeBooksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request booksRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeSetBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request setBookRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
// -----------------------------------------

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

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


//------------------------------


type getBookRequest struct {
	ID int `json:"id"`
}

type getBookResponse struct {
	Book Book `json:"book"`
}

type setBookRequest struct {
	Book Book `json:"book"`
}

type setBookResponse struct {
	ok bool `json:"ok"`
}


type booksRequest struct {

}

type booksResponse struct {
	Books 	map[string]Book `json:"books"`
}
