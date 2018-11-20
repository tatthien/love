package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

var heart = `:heart::heart::heart::heart::heart::heart::heart::heart::heart::heart::heart:
:heart::heart::heart:            :heart:           :heart::heart::heart:
:heart::heart:                                        :heart::heart:
:heart::heart::heart:                             :heart::heart::heart:
:heart::heart::heart::heart:                  :heart::heart::heart::heart:
:heart::heart::heart::heart::heart:      :heart::heart::heart::heart::heart:
:heart::heart::heart::heart::heart::heart::heart::heart::heart::heart::heart:`

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	text := request.QueryStringParameters["text"]
	newHeart := strings.Replace(heart, ":heart:", text, -1)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: newHeart,
	}, nil
}

func main() {
	lambda.Start(handler)
}
