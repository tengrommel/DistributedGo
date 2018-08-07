package main

import "github.com/aws/aws-lambda-go/lambda"

func hello() (string, error) {
	return "Hello ƛ!", nil
}

func main() {
	lambda.Start(hello)
}
