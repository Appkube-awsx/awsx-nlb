/*
Copyright © 2023 Afreen Khan afreen.khan@synectiks.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/cmd/loadbalancercmd"
	"github.com/Appkube-awsx/awsx-elbv2/controllers"
	"github.com/spf13/cobra"
)

var AwsxLoadBalancerCmd = &cobra.Command{
	Use:   "get List of loadbalancers",
	Short: "get List of loadbalancers command gets resource counts",
	Long:  `get List of loadbalancers command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command elbv2 started")

		authFlag := authenticator.RootCommandAuth(cmd)

        marker := cmd.Flags().Lookup("marker").Value.String()
		all, _ := cmd.Flags().GetBool("all")

		if authFlag{
			
			if all{
				controllers.AllElbv2ListController(authenticator.ClientAuth)
			}else{
				controllers.Elbv2ListController(marker, authenticator.ClientAuth)
			}
		}

	},
}


func Execute() {
	err := AwsxLoadBalancerCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {

	AwsxLoadBalancerCmd.AddCommand(loadbalancercmd.GetConfigDataCmd)
	AwsxLoadBalancerCmd.AddCommand(loadbalancercmd.GetLatencyCmd)
	AwsxLoadBalancerCmd.AddCommand(loadbalancercmd.GetNumberOfErrorCmd)
	AwsxLoadBalancerCmd.AddCommand(loadbalancercmd.GetDetailOfErrorCmd)
	AwsxLoadBalancerCmd.AddCommand(loadbalancercmd.GetTotalNumberOfElbv2Cmd)


	AwsxLoadBalancerCmd.Flags().String("marker", "", "marker for next list")
	AwsxLoadBalancerCmd.Flags().Bool("all", false, "to get all lbArns at once")
	

	AwsxLoadBalancerCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLoadBalancerCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxLoadBalancerCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLoadBalancerCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLoadBalancerCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLoadBalancerCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxLoadBalancerCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
