package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	text := request.QueryStringParameters["text"]
	args := strings.Split(text, " ")
	command := args[0]
	emoji := ""
	if len(args) > 1 {
		emoji = args[1]
	}

	var responseBody []byte
	switch command {
	case "list":
		shapeKeys := getShapeKeys()
		response := ""
		for _, key := range shapeKeys {
			response += "`/love " + key + " [emoji]` \n"
		}

		responseBody, _ = json.Marshal(map[string]string{
			"response_type": "in_channel",
			"text":          response,
			"mrkdwn":        "true",
		})
	default:
		shape := getShape(command)
		newHeart := strings.Replace(shape, ":x:", emoji, -1)

		responseBody, _ = json.Marshal(map[string]string{
			"response_type": "in_channel",
			"text":          newHeart,
		})
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(handler)
}
