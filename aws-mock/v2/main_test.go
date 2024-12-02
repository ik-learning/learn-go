package v2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func TestHandler(t *testing.T) {
	var ctx = context.Background()
	var event Request

	t.Run("Upload a file to S3", func(t *testing.T) {
		stubber := testtools.NewStubber()
		lambda := New(*stubber.SdkConfig)

		stubber.Add(testtools.Stub{
			OperationName: "PutObject",
			Input: &s3.PutObjectInput{
				Bucket: aws.String("my-sample-bucket"),
				Key:    aws.String("my/object.json"),
				Body:   bytes.NewReader([]byte{}),
			},
			Output: &s3.PutObjectOutput{},
		})

		response, err := lambda.Handler(ctx, event)
		testtools.ExitTest(stubber, t)
		fmt.Println(response, err)
		// Perform Assertions
	})

	t.Run("Fail on upload", func(t *testing.T) {
		stubber := testtools.NewStubber()
		lambda := New(*stubber.SdkConfig)
		raiseErr := &testtools.StubError{Err: errors.New("ClientError")}

		stubber.Add(testtools.Stub{
			OperationName: "PutObject",
			Input: &s3.PutObjectInput{
				Bucket: aws.String("my-sample-bucket"),
				Key:    aws.String("my/object.json"),
				Body:   bytes.NewReader([]byte{}),
			},
			Error: raiseErr,
		})

		_, err := lambda.Handler(ctx, event)
		testtools.VerifyError(err, raiseErr, t)
		testtools.ExitTest(stubber, t)
	})
}
