package capital

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var capitalAnyCmd = &cobra.Command{
	Use:   "any",
	Short: "这是capital any命令很短的介绍",
	Long:  "这是capital any命令很长很长很长的介绍",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}

func anyCmd() *cobra.Command {
	return capitalAnyCmd
}

func serve(args []string) error {
	var temStrArr []string
	for _, v := range args {
		s := strings.ToUpper(v)
		temStrArr = append(temStrArr, s)
	}
	fmt.Println(temStrArr)
	return nil
}
