// Package cmd implements the command-line interface for chatgo.
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/xavierxcn/chatgo/chatgo"
	"github.com/xavierxcn/chatgo/utils"
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
			fmt.Println("you should set openai token first. example: chatgo set <token>")
			return
		}

		filePath := ""
		if len(args) > 0 {
			filePath = args[0]
		}

		// 初始化一个robot
		robot := chatgo.NewRobot().SetName("chatgo").SetToken(token)
		if filePath != "" {
			fmt.Printf("load messages from %s\n", filePath)
			robot = robot.SetMessagesFromFile(filePath)
			robot.Replay()
		} else {
			fmt.Println("init robot...")
			robot.Init()
			fmt.Println("init robot success.")
		}

		reader := bufio.NewReader(os.Stdin)
		// 循环读取用户输入
		for {
			fmt.Print("> ")
			// 读取一行输入
			sentence, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			if sentence == "\n" {
				continue
			}

			if sentence == "exit\n" {
				// 保存聊天记录
				var savePath string
				defaultSavePath := fmt.Sprintf(chatgo.HistoryPath, robot.Name(), robot.CreateAt.Format(time.RFC3339))
				fmt.Printf("save history to(default: %s): ", defaultSavePath)
				_, _ = fmt.Scan(&savePath)
				if savePath == "" {
					savePath = defaultSavePath
				}
				err := robot.Save(savePath)
				if err != nil {
					panic(err)
				}
				return
			}

			// 机器人回复
			answers, err := robot.TellStream(sentence)
			if err != nil {
				panic(err)
			}
			fmt.Println("chatgo: ")
			for answer := range answers {
				fmt.Print(answer)
			}

			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
