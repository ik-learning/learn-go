package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/goccy/go-yaml"
)

func main() {
	read()
}

func read() {
	path, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/aws/aws-zones-tags/zones.yaml", path))
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var zones HostedZonesR
	if err := decoder.Decode(&zones); err != nil {
		log.Fatalf("failed to decode zones from YAML, %v", err)
	}
	fmt.Println(len(zones.Zones))
	// for _, zone := range zones.Zones {
	// 	log.Printf("Hosted zone %s has ID %s", zone.Name, zone.ID)
	// 	// for _, tag := range zone.Tags {
	// 	// 	log.Printf("Tag %s has value %s", *tag.Key, *tag.Value)
	// 	// }
	// }

	zoneTags := map[string][]*Tag{}

	for _, zone := range zones.Zones {
		zoneTags[zone.ID] = zone.Tags
	}

	fmt.Println(zoneTags)
}

func write() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	client := route53.NewFromConfig(cfg)
	paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{})
	z := HostedZones{
		Zones: []*HostedZone{},
	}
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.Background())
		if err != nil {
			log.Fatalf("failed to list hosted zones, %v", err)
		}
		for _, hostedZone := range output.HostedZones {
			log.Printf("Hosted zone %s has ID %s", *hostedZone.Name, *hostedZone.Id)
			zone := &HostedZone{
				Name: *hostedZone.Name,
				ID:   *hostedZone.Id,
			}
			tags := fetchTags(client, hostedZone.Id)
			zone.Tags = tags
			z.Zones = append(z.Zones, zone)
		}
	}
	data, err := yaml.Marshal(z)
	if err != nil {
		log.Fatalf("failed to marshal zones to YAML, %v", err)
	}
	path, _ := os.Getwd()

	file, err := os.Create(fmt.Sprintf("%s/aws/aws-zones-tags/zones.yaml", path))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()
	//
	// // Write bytes to the file
	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}
}

func fetchTags(client *route53.Client, id *string) []*Tag {
	res, err := client.ListTagsForResource(context.TODO(), &route53.ListTagsForResourceInput{
		ResourceType: types.TagResourceTypeHostedzone,
		ResourceId:   aws.String(cleanZoneID(*id)),
	})
	if err != nil {
		log.Fatalf("failed to list tags for hosted zone, %v", err)
	}
	var result []*Tag
	for _, tag := range res.ResourceTagSet.Tags {
		key := tag.Key
		if strings.Contains(*key, "issue") || strings.Contains(*key, "vpcname") {
			// skip this tag
			continue
		}
		value := tag.Value
		if strings.Contains(*key, "vpcid") {
			value = aws.String("vpc-123456")
		}
		if strings.Contains(*key, "owner") {
			value = aws.String("ext-dns")
		}
		result = append(result, &Tag{
			Key:   key,
			Value: value,
		})

	}
	return result
}

func cleanZoneID(id string) string {
	return strings.TrimPrefix(id, "/hostedzone/")
}

type HostedZonesR struct {
	Zones []*HostedZoneR `yaml:"zones"`
}
type HostedZoneR struct {
	Name string
	ID   string
	Tags []*Tag `yaml:"tags"`
}

type HostedZones struct {
	Zones []*HostedZone `yaml:"zones"`
}

type HostedZone struct {
	Name string
	ID   string
	Tags []*Tag `yaml:"tags"`
}

type Tag struct {
	Key   *string
	Value *string
}
