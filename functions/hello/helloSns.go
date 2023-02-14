package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		name := snsRecord.Message
		log.Printf("Hello %s, from AWS Lambda!", name)
	}
}

func main() {
	lambda.Start(handler)
}
