package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func Getelbv2Detail(elbv2Client *elbv2.ELBV2, lbArns string) *elbv2.DescribeLoadBalancersOutput  {
	log.Println("Getting elbv2  data")

	input := &elbv2.DescribeLoadBalancersInput {
		LoadBalancerArns: []*string{
			aws.String(lbArns),
		},
	}

	elbv2Data, err := elbv2Client.DescribeLoadBalancers(input)
	if err != nil {
		log.Fatalln("Error: in getting elbv2 data", err)
	}

	return elbv2Data
}
