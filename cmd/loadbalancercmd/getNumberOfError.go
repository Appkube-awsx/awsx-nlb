/*
Copyright © 2023 Afreen Khan afreen.khan@synectiks.com
*/
package loadbalancercmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-nlb/authenticator"
	"github.com/Appkube-awsx/awsx-nlb/controllers"
	"github.com/spf13/cobra"
)

// GetNumberOfErrorCmd represents the GetNumberOfErrorCmd command
var GetNumberOfErrorCmd = &cobra.Command{
	Use:   "errorCount",
	Short: "to total number of errors",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticator.ChildCommandAuth(cmd)
		lbArns, _ := cmd.Flags().GetString("lbArns")

		if authFlag {
			controllers.Elbv2GetNumberOfErrorController(lbArns, authenticator.ClientAuth)

		}

	},
}

func init() {
	GetNumberOfErrorCmd.Flags().StringP("lbArns", "f", "", "elbv2 lbArns ")

	if err := GetNumberOfErrorCmd.MarkFlagRequired("lbArns"); err != nil {
		fmt.Println("--lbArns is required", err)
	}
}
