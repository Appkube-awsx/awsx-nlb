package controllers

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-elbv2/services"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func Elbv2Details(lbArns string, auth client.Auth) *elbv2.DescribeLoadBalancersOutput {

	// this is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	// elbv2 client
	elbv2Client := client.GetClient(auth, client.ELBV2_CLIENT).(*elbv2.ELBV2)

	elbv2Detail := services.Getelbv2Detail(elbv2Client, lbArns)
	fmt.Println(elbv2Detail)
	return elbv2Detail
}
