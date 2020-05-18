package capital

import (
	"fmt"
	"github.com/spf13/cobra"
)

var capitalCmd = &cobra.Command{
	Use:   "capital",
	Short: fmt.Sprint("这是capital命令很短的介绍"),
	Long:  fmt.Sprint("这是capital命令很长很长很长的介绍"),
}

func Cmd() *cobra.Command {
	capitalCmd.AddCommand(anyCmd())
	return capitalCmd
}
