package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast1RegionID),
		Credentials: credentials.NewSharedCredentials("", "dynamo-local-profile"),
	})
	_, err = sess.Config.Credentials.Get()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("We're good!")
		fmt.Printf("AWS Config = %v", sess.Config)
	}
}
