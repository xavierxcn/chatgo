package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xavierxcn/chatgo/web"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start web page for chatgo",
	Long:  `start web page for chatgo`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Boot()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
