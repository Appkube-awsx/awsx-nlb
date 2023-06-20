package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/Appkube-awsx/awsx-nlb/utils"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func GetLbErrDetail(cloudClient *cloudwatchlogs.CloudWatchLogs, lbArns string) int {
	log.Println("Getting execution number and errors")

	errCount := 0
	executionCount := 0

	logGroupName := fmt.Sprintf("/aws/loadBalancer/%s", lbArns)

	fmt.Println("log group name", logGroupName)
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(logGroupName),
	}

	firstStreamsList, err := cloudClient.DescribeLogStreams(input)
	if err != nil {
		log.Println("no cloud watch log found for this lbArns")
	}

	nextToken := firstStreamsList.NextToken
	executionCount += len(firstStreamsList.LogStreams)

	errCount += errorDetailsInStreamList(firstStreamsList, logGroupName, cloudClient)
	fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)

	for firstStreamsList.NextToken != nil {

		input = &cloudwatchlogs.DescribeLogStreamsInput{
			LogGroupName: aws.String(logGroupName),
			NextToken:    nextToken,
		}

		tokenStreamsList, err := cloudClient.DescribeLogStreams(input)
		if err != nil {
			log.Println("no cloud watch log found for this lbArns")
		}

		executionCount += len(tokenStreamsList.LogStreams)
		errCount += errorDetailsInStreamList(tokenStreamsList, logGroupName, cloudClient)

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)

	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)

	return errCount
}

func errorDetailsInStreamList(firstStreamsList *cloudwatchlogs.DescribeLogStreamsOutput, logGroupName string, cloudClient *cloudwatchlogs.CloudWatchLogs) int {
	errCount := 0
	for _, stream := range firstStreamsList.LogStreams {

		input := &cloudwatchlogs.GetLogEventsInput{
			LogGroupName:  aws.String(logGroupName),
			LogStreamName: aws.String(*stream.LogStreamName),
			StartFromHead: aws.Bool(true),
		}

		resp, err := cloudClient.GetLogEvents(input)

		if err != nil {
			log.Fatalln("Error: in getting event data", err)
		}

		for _, event := range resp.Events {
			//fmt.Println("tracing error in :", *stream.LogStreamName)
			if strings.Contains(*event.Message, "ERROR") {
				fmt.Println()
				fmt.Println()
				fmt.Println("error in ", *stream.LogStreamName, "is::", event)
				errCount++
			}
		}
	}
	return errCount
}

func GetLoadBalancersErrCount(cloudClient *cloudwatchlogs.CloudWatchLogs, elbv2Client *elbv2.ELBV2) int {
	log.Println("Getting execution number and errors for all lbArns")
	allArnsList := GetAllLbList(elbv2Client)

	errCount := 0
	executionCount := 0

	for _, elbv2 := range allArnsList {
		tempErrCount, tempExecutionCount, err := GetElbv2ErrCount(cloudClient, *elbv2.LoadBalancerArn)

		if err != nil {
			if strings.Contains(err.Error(), "ResourceNotFoundException") {
				continue
			}
		}

		errCount += tempErrCount
		executionCount += tempExecutionCount

		fmt.Println("Total executions till now:", executionCount, "and errors are:", errCount)
	}

	fmt.Println("Final execution count is:", executionCount, "errors are:", errCount)
	return errCount
}

func GetElbv2ErrCount(cloudClient *cloudwatchlogs.CloudWatchLogs, lbArns string) (int, int, error) {
	log.Println("Getting execution number and errors")

	layout := "2006-01-02 15:04:05.000000000 -0700 MST"
	//startTime := time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC) // 1 hour ago
	str := "2023-01-01 01:01:01.000000001 +0000 UTC"
	end := "2023-02-01 01:01:01.000000001 +0000 UTC"

	startTime, err := time.Parse(layout, str)
	endTime, err := time.Parse(layout, end)

	if err != nil {
		log.Fatalln(err)
	}

	//endTime1 := time.Now()
	fmt.Println("this is time", endTime)
	//fmt.Println(endTime1)

	query := `filter @message like /(?i)(ERROR)/
									| stats count() as ErrorCount`

	logGroupName := fmt.Sprintf("/aws/elbv2/%s", lbArns)

	fmt.Println("log group name", logGroupName)
	result, err := utils.GetQueryData(cloudClient, logGroupName, startTime, endTime, query)

	if err != nil {
		return 0, 0, err
	}

	fmt.Println("Result json", result)
	fmt.Println("Final execution count is:", *result.Statistics.RecordsScanned, "errors are:", *result.Statistics.RecordsMatched)
	fmt.Println()

	return int(*result.Statistics.RecordsMatched), int(*result.Statistics.RecordsScanned), nil
}



