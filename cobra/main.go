/*
参考文档：go插件cobra命令用法
http://blog.csdn.net/qq_27809391/article/details/54089774

编译：go build -o command
运行：
./command
./command -h
./command --help
./command node
./command node -h
./command node --help
./command node start
./command node start -h
./command node start --help
./command node start 测试
*/

package main

import (
	"fmt"
	"github.com/fengchunjian/goexamples/cobra/capital"
	"github.com/fengchunjian/goexamples/cobra/node"
	"github.com/spf13/cobra"
	"os"
)

var versionFlag bool

var mainCmd = &cobra.Command{
	Use: "command",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("目前版本为1.0.0")
		} else {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

func main() {
	mainFlags := mainCmd.PersistentFlags()
	mainFlags.BoolVarP(&versionFlag, "version", "v", false, "打印系统版本")

	mainCmd.AddCommand(node.Cmd())
	mainCmd.AddCommand(capital.Cmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
