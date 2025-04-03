package main

// go run aws/aws-ecr/main.go

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type ECRAuthenticator struct{}

const (
	region = "eu-west-1"
)

var (
	accounts = map[string]string{
		"010438477961": "arn:aws:ecr:eu-west-1:010438477961:repository/", // container registry account ID
		"048958980748": "arn:aws:ecr:eu-west-1:048958980748:repository/", // sandbox account ID
	}
)

func main() {
	repository := "hnbi/platform/paas/hello-ecr"
	arn := accounts["010438477961"] + repository
	fmt.Println(arn)
	Tags()
	c, err := Authenticate()
	if err != nil {
		fmt.Println("UPS:", err)
	} else {
		fmt.Println("result:", c)
	}
}

func Tags() {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an ECR client
	svc := ecr.NewFromConfig(cfg)

	// Define the resource ARN
	resourceArn := "arn:aws:ecr:eu-west-1:010438477961:repository/hnbi/platform/paas/hello-ecr"

	// Call the ListTagsForResource API
	input := &ecr.ListTagsForResourceInput{
		ResourceArn: aws.String(resourceArn),
	}

	result, err := svc.ListTagsForResource(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to list tags for resource, %v", err)
	}

	// Print the tags
	for _, tag := range result.Tags {
		fmt.Printf("Key: %s, Value: %s\n", aws.ToString(tag.Key), aws.ToString(tag.Value))
	}
}

func Authenticate() (string, error) {
	// Placeholder for ECR authentication logic
	// This should return the authentication token for ECR
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ecr.NewFromConfig(cfg)

	// Call the GetAuthorizationToken API
	input := &ecr.GetAuthorizationTokenInput{
		// container registry account ID
		// RegistryIds: []string{"010438477961"},
	}

	result, err := svc.GetAuthorizationToken(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to get authorization token, %v", err)
	}

	// Print the authorization token
	// fmt.Printf("Authorization Token: %s\n", aws.ToString(result.AuthorizationData[0].AuthorizationToken))
	var decoded string
	for _, authData := range result.AuthorizationData {
		token, err := base64.StdEncoding.DecodeString(aws.ToString(authData.AuthorizationToken))
		if err != nil {
			log.Fatalf("failed to decode authorization token, %v", err)
		}
		decoded = string(token)
	}
	tokenParts := strings.SplitN(decoded, ":", 2)
	if len(tokenParts) != 2 {
		log.Fatalf("invalid authorization token format")
	}
	username := tokenParts[0]
	password := tokenParts[1]

	fmt.Printf("Authorization Token: export DOCKER_TOKEN=\"%s\"\n", password)
	fmt.Println(fmt.Sprintf("echo $DOCKER_TOKEN | docker login --password-stdin --username %s 010438477961.dkr.ecr.eu-west-1.amazonaws.com", username))

	return "", nil
}
