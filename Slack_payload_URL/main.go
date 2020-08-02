package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "fmt"
)
type Request struct {
    Body1 string `json:"body"`
}
 

type Response struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

func handler(abc Request) (Response, error) {
    fmt.Printf("%s", abc.Body1)
	return Response{
		StatusCode: 200,
		Body:       abc.Body1,
	}, nil
}

func main() {
	lambda.Start(handler)
}
