package v2

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Request struct{}
type Response struct{}

type Lambda struct {
	ctx      context.Context
	s3Client *s3.Client
}

func New(cfg aws.Config) *Lambda {
	m := new(Lambda)
	m.s3Client = s3.NewFromConfig(cfg)
	return m
}

func (x *Lambda) Handler(ctx context.Context, request Request) (Response, error) {
	// Your lambda code goes here

	return Response{}, nil
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	lambda.Start(New(cfg).Handler)
}
