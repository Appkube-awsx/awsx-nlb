package controllers

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/services"
)

func Elbv2TotalNumberOfLb(auth client.Auth) int {

	// this is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	// ELBV2  client
	elbv2Client := client.GetClient(auth, client.ELBV2_CLIENT).(*elbv2.ELBV2)

	totalNumber := len(services.GetAllLbList(elbv2Client))

	fmt.Println("total number of load balancer present in aws account in", authenticator.ClientAuth.Region, "is:", totalNumber)
	return totalNumber
}
