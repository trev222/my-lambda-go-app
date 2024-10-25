package handlers

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
)

func HelloHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    body, _ := json.Marshal(map[string]string{"message": "Hello from Lambda!"})
    return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}
