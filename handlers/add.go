package handlers

import (
    "context"
    "encoding/json"
    "strconv"
    "github.com/aws/aws-lambda-go/events"
)

func AddHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    a, _ := strconv.Atoi(req.QueryStringParameters["a"])
    b, _ := strconv.Atoi(req.QueryStringParameters["b"])
    body, _ := json.Marshal(map[string]int{"result": a + b})
    return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

