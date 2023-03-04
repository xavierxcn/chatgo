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

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "say something to chatgo",
	Long: `say something to chatgo, for example:
chatgo say
you should set openai token first.
For example:
chatgo set <token>`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		// 判断是否存在 ~/.chatgo/token
		if !utils.IsFileExist(chatgo.TokenPath) {
			err = utils.CreateFile(chatgo.TokenPath)
			if err != nil {
				panic(err)
			}
		}

		// 读取 ~/.chatgo/token
		token, err := utils.ReadFile(chatgo.TokenPath)
		if err != nil {
			panic(err)
		}

		if token == "" {
			fmt.Println("you should set openai token first.")
		}

		// 初始化一个robot
		robot := chatgo.NewRobot().SetName("chatgo").SetToken(token)

		// 循环读取用户输入
		for {
			var sentence string
			fmt.Print("> ")
			fmt.Scanln(&sentence)

			// 机器人回复
			fmt.Printf("Robot: %s\n", robot.Tell(sentence))
		}

	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
