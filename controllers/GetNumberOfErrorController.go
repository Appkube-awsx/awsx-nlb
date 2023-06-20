package controllers

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-nlb/authenticator"
	"github.com/Appkube-awsx/awsx-nlb/services"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func Elbv2GetNumberOfErrorController(lbArns string, auth client.Auth) int {

	// this is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	services.GetElbv2ErrCount(cloudClient, lbArns)
	return -1
}
