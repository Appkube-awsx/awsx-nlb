package services

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Getelbv2List -> get elbv2 list with pagination
func GetElbv2List(elbv2Client *elbv2.ELBV2, marker string) *elbv2.DescribeLoadBalancersOutput {
	log.Println("Getting elbv2 list summary")
	input := &elbv2.DescribeLoadBalancersInput{}

	if marker != "" {
		input = &elbv2.DescribeLoadBalancersInput{
			Marker: &marker,
			
		}
	}

	loadBalancersList, err := elbv2Client.DescribeLoadBalancers(input)

	if err != nil {
		log.Fatalln("Error: in getting elbv2 list", err)
	}

	return loadBalancersList

}

func GetAllLbList(elbv2Client *elbv2.ELBV2) []*elbv2.LoadBalancer{
	log.Println("Getting all elbv2 list summary")
	loadBalancersList := GetElbv2List(elbv2Client, "")

	allArnsList := loadBalancersList.LoadBalancers
	marker := loadBalancersList.NextMarker

	// loop for getting all values
	for marker != nil{
		loadBalancersList = GetElbv2List(elbv2Client, *loadBalancersList.NextMarker)
		allArnsList = append(allArnsList, loadBalancersList.LoadBalancers...)
		marker = loadBalancersList.NextMarker
		fmt.Println("load balancers got till now:: ", len(allArnsList))
	}
	return allArnsList

}


