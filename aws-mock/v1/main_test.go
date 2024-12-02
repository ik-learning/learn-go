package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockS3Client struct {
	s3.Client
	Error error
}

func (m *mockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	fmt.Println("IN DUMMY PutObject")
	return &s3.PutObjectOutput{}, nil
}

func TestHandler(t *testing.T) {
	lambda := New()
	mock := mockS3Client{}
	lambda.SetS3Client(&mock.Client)
	var ctx = context.Background()
	event := Request{}

	t.Run("Invoke Handler", func(t *testing.T) {
		response, err := lambda.Handler(ctx, event)

		fmt.Println("----------")
		fmt.Println(err)
		fmt.Println(response)
		// Perform Assertions
	})
}
