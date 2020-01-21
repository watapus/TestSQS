package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateMessage(documentType, correlationID, borrowerName string) {
	fmt.Println("We are creating a message to send")

	messageAttributes := map[string]*sqs.MessageAttributeValue{
		"CorrelationID": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(correlationID),
		},
		"BorrowerName": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(borrowerName),
		},
		"DocumentType": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(documentType),
		},
	}

	result, err := service.SQS.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      aws.Int64(60),
		QueueUrl:          &queue_url,
		MessageAttributes: messageAttributes,
		MessageBody:       aws.String("{\"TEST\":\"Test A Message\"}"),
	})

	fmt.Println("Message Created and Sent")
	if err != nil {
		fmt.Println("We had other issues")
	}

	fmt.Println("Success" + *result.MessageId)
}

func RecieveMessage() string {

	for ok := true; ok; ok = true {
		result, err := service.SQS.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &queue_url,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(20),
			WaitTimeSeconds:     aws.Int64(20),
		})

		if err != nil {
			fmt.Println("We have issues in Recieve")
		}

		if len(result.Messages) == 0 {
			fmt.Println("Recieved No Messages")
		}

		fmt.Println("Found")
		fmt.Println(result)
	}
	return ""
}
