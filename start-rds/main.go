package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

const (
	AwsRegion = "us-east-1"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

var awsSession *session.Session

func getAwsSession() *session.Session {
	if awsSession == nil {
		awsSession = session.Must(session.NewSession(&aws.Config{Region: aws.String(AwsRegion)}))
	}
	return awsSession
}

// Start AWS RDS
func startRds(session *session.Session) error {
	// check databases
	svc := rds.New(session)

	// stop rds instance or cluster
	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(os.Getenv("DATABASE")),
	}

	_, err := svc.StartDBInstance(input)

	if err != nil {
		return err
	}
	return nil
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	err := startRds(getAwsSession())

	//if err exist, then print and return code 500
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "{\"message\": \"RDS started\"}",
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
