package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
)

// Post represents a Post object in the DynamoDb table
type Post struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Text   string `json:"text"`
	S3Url  string `json:"url"`
	Voice  string `json:"voice"`
}

var posts []Post

// GetPosts route handler to return all the Posts in the posts slice
func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
	})

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("We're good!")
	}

	// Let's get that dynamoDb client
	dynamoClient := dynamodb.New(sess)
	if err != nil {
		log.Fatalln(err)
	}

	// Query the table and return the results
	result, err := queryTable("posts", dynamoClient)
	if err != nil {
		log.Fatalln(err)
	}

	// Popualate posts array
	err = populateTableItemsFromResults(result)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new router with a 'posts' endooint for GET
	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Make the DynamoDB Query API call
func queryTable(tableName string, dyna *dynamodb.DynamoDB) (*dynamodb.ScanOutput, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dyna.Scan(params)
	return result, err
}

func populateTableItemsFromResults(results *dynamodb.ScanOutput) error {

	for _, i := range results.Items {
		post := Post{}

		err := dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return err
		}

		posts = append(posts, post)
	}

	return nil
}
