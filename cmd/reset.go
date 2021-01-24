package cmd

import (
	"fmt"
	"github.com/feivxs/cbtool/utils"
	"github.com/spf13/cobra"
	"os"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "删除登录信息",
	Run:   doReset,
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func doReset(cmd *cobra.Command, args []string) {
	os.Remove(utils.GetDefaultConfigPath())
	fmt.Println("删除配置成功！")
}
