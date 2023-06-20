/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package loadbalancercmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/controllers"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetDetailOfErrorCmd = &cobra.Command{
	Use:   "errorDetail",
	Short: "to get details of error in elbv2",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticator.ChildCommandAuth(cmd)
		lbArns, _ := cmd.Flags().GetString("lbArns")

		if authFlag {
			controllers.Elbv2DetailsErrorController(lbArns, authenticator.ClientAuth)

		}

	},
}

func init() {

	GetDetailOfErrorCmd.Flags().StringP("lbArns", "f", "", "LOAD BALANCER  lbArns ")

	if err := GetDetailOfErrorCmd.MarkFlagRequired("lbArns"); err != nil {
		fmt.Println("--lbArns is required", err)
	}

}
