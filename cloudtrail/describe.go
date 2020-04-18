package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"

	"fmt"
	"os"
)

type NameBucket struct {
	Name   string
	Bucket string
}

// describe should describe trails
func describe(region string, sess *session.Session) []NameBucket {
	// Create CloudTrail client
	svc := cloudtrail.New(sess)

	resp, err := svc.DescribeTrails(&cloudtrail.DescribeTrailsInput{TrailNameList: nil})
	if err != nil {
		fmt.Println("Got error calling DescribeTrail:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Found", len(resp.TrailList), "trail(s) in", region)
	fmt.Println("")

	var result []NameBucket

	for _, trail := range resp.TrailList {
		result = append(result, NameBucket{Name: *trail.Name, Bucket: *trail.S3BucketName})
	}
	return result
}
