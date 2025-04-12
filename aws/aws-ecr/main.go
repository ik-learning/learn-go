package main

// go run aws/aws-ecr/main.go
// https://github.com/awslabs/amazon-ecr-credential-helper

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

type ECRAuthenticator struct {
	stsClient *sts.Client
}

const (
	region = "eu-west-1"
)

var (
	accounts = map[string]string{
		"010438477961": "arn:aws:ecr:eu-west-1:010438477961:repository/", // container registry account ID
		"048958980748": "arn:aws:ecr:eu-west-1:048958980748:repository/", // sandbox account ID
	}
)

const (
	roleArn = "arn:aws:iam::048958980748:role/GitlabCI"
)

const (
	REGION    = "eu-west-1"
	LOGIN_CMD = `
login=$(echo %s | docker login --password-stdin --username %s %s)
`
)

func main() {
	repository := "hnbi/platform/paas/hello-ecr"
	arn := accounts["010438477961"] + repository
	// Tags()
	a := ECRAuthenticator{}
	c, err := a.AuthenticateWithAssume(arn)
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

func (a *ECRAuthenticator) Authenticate(input string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(REGION))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	svc := ecr.NewFromConfig(cfg)
	token, err := svc.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return "", fmt.Errorf("failed to get authorization token, %v", err)
	}

	var decoded string
	for _, authData := range token.AuthorizationData {
		token, err := base64.StdEncoding.DecodeString(aws.ToString(authData.AuthorizationToken))
		if err != nil {
			return "", fmt.Errorf("failed to decode authorization token, %v", err)
		}
		decoded = string(token)
	}
	tokenParts := strings.SplitN(decoded, ":", 2)
	if len(tokenParts) != 2 {
		return "", fmt.Errorf("invalid authorization token format")
	}
	username := tokenParts[0]
	password := tokenParts[1]

	result := fmt.Sprintf(LOGIN_CMD, password, username, "010438477961.dkr.ecr.eu-west-1.amazonaws.com")

	return result, nil
}

func (e *ECRAuthenticator) AuthenticateWithAssume(input string) (string, error) {
	if e.stsClient == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			return "", fmt.Errorf("unable to load SDK config, %v", err)
		}
		e.stsClient = sts.NewFromConfig(cfg)
	}

	roleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String("session"),
	}
	response, err := e.stsClient.AssumeRole(context.Background(), &roleInput)
	if err != nil {
		return "", fmt.Errorf("failed to assume role, %v", err)
	}

	fmt.Println("response creds:", response)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(REGION))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	svc := ecr.NewFromConfig(cfg)
	token, err := svc.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return "", fmt.Errorf("failed to get authorization token, %v", err)
	}

	var decoded string
	for _, authData := range token.AuthorizationData {
		token, err := base64.StdEncoding.DecodeString(aws.ToString(authData.AuthorizationToken))
		if err != nil {
			return "", fmt.Errorf("failed to decode authorization token, %v", err)
		}
		decoded = string(token)
	}
	tokenParts := strings.SplitN(decoded, ":", 2)
	if len(tokenParts) != 2 {
		return "", fmt.Errorf("invalid authorization token format")
	}
	username := tokenParts[0]
	password := tokenParts[1]

	result := fmt.Sprintf(LOGIN_CMD, password, username, "010438477961.dkr.ecr.eu-west-1.amazonaws.com")

	return result, nil
}

func getCredentialProvider(creds *types.Credentials) aws.CredentialsProviderFunc {
	return func(ctx context.Context) (aws.Credentials, error) {
		c := &aws.Credentials{
			AccessKeyID:     aws.ToString(creds.AccessKeyId),
			SecretAccessKey: aws.ToString(creds.SecretAccessKey),
			SessionToken:    aws.ToString(creds.SessionToken),
		}
		return *c, nil
	}
}
