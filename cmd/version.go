package cmd

import (
	"fmt"
	"github.com/xavierxcn/chatgo/chatgo"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  `show version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("chatgo version: %s\n", chatgo.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
