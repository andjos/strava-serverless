package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var response events.APIGatewayProxyResponse

	if request.QueryStringParameters["hub.verify_token"] == os.Getenv("hubverifytoken") {
		hubC := request.QueryStringParameters["hub.challenge"]
		log.Println("hub.challenge:", hubC)

		body, _ := json.Marshal(map[string]interface{}{
			"hub.challenge": hubC,
		})

		response = events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

	} else {
		body, _ := json.Marshal(map[string]interface{}{
			"message": "hub.verify_token not valid",
		})

		response = events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

	}
	return response, nil
}
func main() {
	lambda.Start(Handler)
}
