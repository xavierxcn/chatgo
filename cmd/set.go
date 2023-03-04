/*
Copyright © 2023 Xavier X <xavier@xavierx.cn>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xavierxcn/chatgo/chatgo"
	"github.com/xavierxcn/chatgo/utils"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set openai token",
	Long:  `Set openai token to ~/.chatgo/token`,
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		var err error

		// 判断是否存在 ~/.chatgo/token
		if !utils.IsFileExist(chatgo.TokenPath) {
			err = utils.CreateFile(chatgo.TokenPath)
			if err != nil {
				panic(err)
			}
		}

		// 将token写入到 ~/.chatgo/token
		err = utils.WriteFile(chatgo.TokenPath, token)
		if err != nil {
			panic(err)
		}

		fmt.Println("set openai token success.")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
