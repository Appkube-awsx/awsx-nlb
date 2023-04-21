package loadbalancercmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)
		print(authFlag)
		// authFlag := true
		if authFlag {
			lbArns, _ := cmd.Flags().GetString("lbArns")
			if lbArns != "" {
				getLoadBalancerDetails(region, crossAccountRoleArn, acKey, secKey, lbArns, env, externalId)
			} else {
				log.Fatalln("lbArns not provided. Program exit")
			}
		}
	},
}

func getLoadBalancerDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, lbArns string, env string, externalId string) *elbv2.DescribeLoadBalancersOutput {
	log.Println("Getting load balancer data")
	lbClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &elbv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []*string{
			aws.String(lbArns),
		},
		// LoadBalancerNames: []*string{
		// 	aws.String(lbName),
		// },
	}
	lbResponse, err := lbClient.DescribeLoadBalancers(input)
	log.Println(lbResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return lbResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("lbArns", "t", "", "load balancer Arns")

	if err := GetConfigDataCmd.MarkFlagRequired("lbArns"); err != nil {
		fmt.Println(err)
	}
}
