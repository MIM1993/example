package node

import (
	"fmt"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: fmt.Sprint("这是node命令很短的介绍"),
	Long:  fmt.Sprint("这是node命令很长很长很长的介绍"),
}

func Cmd() *cobra.Command {
	nodeCmd.AddCommand(startCmd())
	return nodeCmd
}
