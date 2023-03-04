/*
Copyright © 2023 Xavier X <xavier@xavierx.cn>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xavierxcn/chatgo/chatgo"
	"github.com/xavierxcn/chatgo/utils"
	"os"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "chat",
	Short: "say something to chatgo",
	Long: `say something to chatgo, for example:
chatgo chat
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

		reader := bufio.NewReader(os.Stdin)
		// 循环读取用户输入
		for {
			fmt.Print("> ")
			// 读取一行输入
			sentence, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			// 机器人回复
			fmt.Printf("Robot: %s\n", robot.Tell(sentence))
		}

	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
