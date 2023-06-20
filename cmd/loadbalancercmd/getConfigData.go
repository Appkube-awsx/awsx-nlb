package loadbalancercmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-elbv2/authenticator"
	"github.com/Appkube-awsx/awsx-elbv2/controllers"

	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		
		authFlag := authenticator.ChildCommandAuth(cmd)

		lbArns, _ := cmd.Flags().GetString("lbArns")
		
		if authFlag {
			controllers.Elbv2Details(lbArns, authenticator.ClientAuth)
		}
	},
}


func init() {
	GetConfigDataCmd.Flags().StringP("lbArns", "t", "", "load balancer Arns")

	if err := GetConfigDataCmd.MarkFlagRequired("lbArns"); err != nil {
		fmt.Println("--lbArns or -t is required",err)
	}
}
