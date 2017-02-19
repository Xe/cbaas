package main

import (
	"log"

	"github.com/apex/invoke"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func init() {
	// Create a Session with a custom region
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		log.Fatal(err)
	}

	invoke.DefaultClient = lambda.New(sess)
}
