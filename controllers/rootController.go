package controllers

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/services"
	"github.com/aws/aws-sdk-go/service/elbv2"
	// "github.com/aws/aws-sdk-go/service/elbv2"
)

func AllElbv2ListController(auth client.Auth) []*elbv2.LoadBalancer{

	authenticator.ApiAuth(auth)

	elbv2Client := client.GetClient(auth, client.ELBV2_CLIENT).(*elbv2.ELBV2)

	loadBalancersList := services.GetAllLbList(elbv2Client)

	fmt.Println("List of all load balancer", loadBalancersList)

	return loadBalancersList
}

// elbv2ListController for pagination elbv2 list
func Elbv2ListController(marker string, auth client.Auth) *elbv2.DescribeLoadBalancersOutput {
 
	// This is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	elbv2Client := client.GetClient(auth, client.ELBV2_CLIENT).(*elbv2.ELBV2)

	loadBalancersList := services.GetElbv2List(elbv2Client, marker)

	fmt.Println("List of elbv2 functions", loadBalancersList)
	return loadBalancersList
}

