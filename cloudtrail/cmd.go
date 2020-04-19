package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/urfave/cli"
)

func main() {
	var region string

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Ivan Katliarchuk",
			Email: "ivan.katliarchuk@gmail.com",
		},
	}
	app.Name = "cloudtrail"
	app.Usage = "Operations on cloudtrail"
	app.Description = "Should help to create IAM role from given CloudTrail"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "describe, d",
			Usage: "Describe Trails input",
		},
		cli.StringFlag{
			Name:        "region, r",
			Value:       "region",
			Destination: &region,
			Usage:       "AWS Region to run against",
			EnvVar:      "AWS_REGION",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"d"},
			Usage:   "provide a list of trails for given region",
			Action: func(c *cli.Context) error {
				fmt.Println("describe cloud trail for: ", region)
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String(region)},
				)
				if err != nil {
					fmt.Println("Not Able to obtain session > ")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				result := describe(region, sess)
				for _, trail := range result {
					fmt.Println("Trail name:  " + trail.Name)
					fmt.Println("Bucket name: " + trail.Bucket)
					fmt.Println("")
				}
				events("cloudtrail-us-west-2", sess)
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
