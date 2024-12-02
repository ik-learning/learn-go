package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Request struct{}
type Response struct{}

type Lambda struct {
	s3Client *s3.Client
}
