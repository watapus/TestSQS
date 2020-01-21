package main

import (
	"fmt"
	"net/http"
)

var access_key = ""
var secret_key = ""
var queue_url = "https://sqs.us-east-1.amazonaws.com/301279725147/thunder_test_queue"
var region_name = "us-east-1"
var service *SQSService
var err error

func main() {
	service, err = newSQSService(queue_url, access_key, secret_key, region_name)

	if err != nil {
		fmt.Println("We have issues")
	}

	r := CreateRouter()

	out1 := make(chan error)
	out2 := make(chan string)

	go func() {
		out1 <- http.ListenAndServe(":8000", r)
	}()
	go func() {
		out2 <- RecieveMessage()
	}()
	f2(<-out1, <-out2)

	go RecieveMessage()
}

func f2(arg1 error, arg2 string) {
	fmt.Println("Fired off both channels")
}
