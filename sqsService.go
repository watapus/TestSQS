package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSService struct {
	Session *session.Session
	SQS     *sqs.SQS
	SQSURL  string
}

func newSQSService(queueURL, access_key, secret_key, region_name string) (*SQSService, error) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region_name),
		MaxRetries:  aws.Int(5),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:    aws.String(queueURL),
	})

	svc := sqs.New(sess)

	builder := &SQSService{
		Session: sess,
		SQS:     svc,
		SQSURL:  queueURL,
	}

	return builder, nil
}
