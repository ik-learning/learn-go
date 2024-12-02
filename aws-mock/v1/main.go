package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Request struct{}
type Response struct{}

type Lambda struct {
	s3Client *s3.Client
}

func New() *Lambda {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	m := new(Lambda)
	m.SetS3Client(s3.NewFromConfig(cfg))
	return m
}

func (x *Lambda) SetS3Client(client *s3.Client) {
	x.s3Client = client
}

func (x *Lambda) Handler(ctx context.Context, request Request) (Response, error) {
	// Your lambda code goes here
	if x.s3Client == nil {
		fmt.Println("s3Client is nil")
	}
	fmt.Println("MAKE A request", ctx)
	out, err := x.s3Client.PutObject(ctx, &s3.PutObjectInput{})
	fmt.Println("handler err:", err)
	fmt.Println("handler out:", out)
	return Response{}, nil
}
