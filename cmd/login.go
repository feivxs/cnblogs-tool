package cmd

import "github.com/spf13/cobra"

var username string
var password string
var id string

var loginCmd = &cobra.Command{
	Use:   "login xxx",
	Short: "录入登录信息",
	Run:   doLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "用户名称")
	loginCmd.Flags().StringVarP(&username, "password", "p", "", "用户密码")
	loginCmd.Flags().StringVarP(&username, "id", "i", "", "博客园 ID（如：https://www.cnblogs.com/feivxs 的博客id为 feivxs）")
}

func doLogin(cmd *cobra.Command, args []string) {

}
