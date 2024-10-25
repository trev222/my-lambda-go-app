package main

import (
	"context"
	"net/http"
	"os"
	"log"

	"my-lambda-go-app/handlers" 

	"github.com/gin-gonic/gin"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the homepage! 127"})
	})

	router.GET("/ping", func(c *gin.Context) {
		resp, _ := handlers.PingHandler(context.Background(), events.APIGatewayProxyRequest{})
		c.JSON(resp.StatusCode, gin.H{"message": resp.Body})
	})

	router.GET("/hello", func(c *gin.Context) {
		resp, _ := handlers.HelloHandler(context.Background(), events.APIGatewayProxyRequest{})
		c.JSON(resp.StatusCode, gin.H{"message": resp.Body})
	})

	router.GET("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")

		resp, _ := handlers.AddHandler(context.Background(), events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"a": a, "b": b},
		})

		c.JSON(resp.StatusCode, gin.H{"result": resp.Body})
	})

	return router
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		ginLambda = ginadapter.New(SetupRouter())
	}
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	if os.Getenv("_LAMBDA_SERVER_PORT") != "" {
		logger.Println("Running in Lambda mode")
		lambda.Start(LambdaHandler)
	} else {
		logger.Println("Running locally on http://localhost:8080")
		router := SetupRouter()
		router.Run(":8080")
	}
}
