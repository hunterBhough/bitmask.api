package locatorerrors

import "github.com/aws/aws-lambda-go/events"

// MarshalError is returned when there is a JSON Marshal error along the way
func MarshalError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Service Error",
		StatusCode: 500,
	}, err
}
