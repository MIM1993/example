package node

import (
	"fmt"
	"github.com/spf13/cobra"
)

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "这是node start命令很短的介绍",
	Long:  "这是node start命令很长很长很长的介绍",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}

func startCmd() *cobra.Command {
	return nodeStartCmd
}

func serve(args []string) error {
	fmt.Println("测试使用node start命令", args)
	return nil
}
