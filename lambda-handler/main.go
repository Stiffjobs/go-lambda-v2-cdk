package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Message string `json:"greeting"`
}

func init() {}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp := &response{
		Message: "Hello world",
	}

	body, err := json.Marshal(resp)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string("Error parsing payload"),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body: string(body), StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
