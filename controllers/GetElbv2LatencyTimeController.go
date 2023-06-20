package controllers

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-nlb/authenticator"
	"github.com/Appkube-awsx/awsx-nlb/services"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetElbv2LatencyTimeController(lbArns string, startTime string, endTime string, auth client.Auth) string {

	// this is Api auth and compulsory for every controller
	authenticator.ApiAuth(auth)

	cloudClient := client.GetClient(auth, client.CLOUDWATCH_LOG).(*cloudwatchlogs.CloudWatchLogs)

	if lbArns != "" {
		result, _, _ := services.GetElbv2LatencyTime(cloudClient, lbArns, startTime, endTime)
		return result
	}

	return "Please send lbArns"
}
