package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"

	"fmt"
	"os"
)

// events should describe trails
func events(trail string, sess *session.Session) {
	// Create CloudTrail client
	svc := cloudtrail.New(sess)

	input := &cloudtrail.LookupEventsInput{EndTime: aws.Time(time.Now())}

	resp, err := svc.LookupEvents(input)
	if err != nil {
		fmt.Println("Got error calling CreateTrail:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Found", len(resp.Events), "events before now")
	fmt.Println("")

	for _, event := range resp.Events {
        if true {
            fmt.Println("Event:")
            fmt.Println(aws.StringValue(event.CloudTrailEvent))
            fmt.Println("")
        }

        fmt.Println("Name    ", aws.StringValue(event.EventName))
        fmt.Println("ID:     ", aws.StringValue(event.EventId))
        fmt.Println("Time:   ", aws.TimeValue(event.EventTime))
        fmt.Println("User:   ", aws.StringValue(event.Username))

        fmt.Println("Resourcs:")

        for _, resource := range event.Resources {
            fmt.Println("  Name:", aws.StringValue(resource.ResourceName))
            fmt.Println("  Type:", aws.StringValue(resource.ResourceType))
        }

        fmt.Println("")
    }
}
