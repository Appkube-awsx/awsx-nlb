/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package loadbalancercmd

import (
	"github.com/Appkube-awsx/awsx-nlb/authenticator"
	"github.com/Appkube-awsx/awsx-nlb/controllers"
	"github.com/spf13/cobra"
)

// GetTotalNumberOfLambdaCmd represents the number command
var GetTotalNumberOfElbv2Cmd = &cobra.Command{
	Use:   "totalCount",
	Short: "gets total number of lambdas present in aws account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticator.ChildCommandAuth(cmd)

		if authFlag {
			controllers.Elbv2TotalNumberOfLb(authenticator.ClientAuth)
		}
	},
}

func init() {
	
}
