package controllers

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-nlb/authenticator"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-nlb/services"
)

func Elbv2DetailsErrorController(lbArns string, auth client.Auth) {

	// this is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	// Cloud client
	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	detail := services.GetLbErrDetail(cloudClient, lbArns)
	fmt.Println(detail)
}
