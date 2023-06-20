/*
Copyright © 2023 AFREEN KHAN <afreen.khan@synectiks.com>
*/
package loadbalancercmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/controllers"
	"github.com/spf13/cobra"
)

// GetLatencyCmd represents the GetLatencyCmd		 command
var GetLatencyCmd = &cobra.Command{
	Use:   "latency",
	Short: "to get latency of elbv2 lbArns",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticator.ChildCommandAuth(cmd)
		lbArns, _ := cmd.Flags().GetString("lbArns")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")

		if authFlag {
			controllers.GetElbv2LatencyTimeController(lbArns, startTime, endTime, authenticator.ClientAuth)
		}

	},
}

func init() {

	GetLatencyCmd.Flags().StringP("lbArns", "f", "", "elbv2 lbArns")
	GetLatencyCmd.Flags().StringP("startTime", "s", "", "elbv2 start Time")
	GetLatencyCmd.Flags().StringP("endTime", "e", "", "elbv2 endtime")

	if err := GetLatencyCmd.MarkFlagRequired("lbArns"); err != nil {
		fmt.Println("--lbArns is required", err)
	}

}
