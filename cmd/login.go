package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/feivxs/cbtool/common"
	"github.com/feivxs/cbtool/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var username string
var password string
var id string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "录入登录信息",
	Run:   doLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "用户名称")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "用户密码")
	loginCmd.Flags().StringVarP(&id, "id", "i", "", "博客园 ID（如：https://www.cnblogs.com/feivxs 的博客id为 feivxs）")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("id")
}

func doLogin(cmd *cobra.Command, args []string) {
	if password == "" {
		fmt.Println("请输入用户密码：")
		fmt.Scanln(&password)
	}
	account := common.CnBlogAccount{
		Id:             id,
		Username:       username,
		Password:       password,
		BlogUrl:        "https://www.cnblogs.com/" + id,
		MetaWebBlogUrl: "https://rpc.cnblogs.com/metaweblog/" + id,
	}
	bytes, err := json.Marshal(account)
	utils.HandleErr(err)
	utils.HandleErr(ioutil.WriteFile(utils.GetDefaultConfigPath(), bytes, 0777))
	fmt.Println("配置设置成功！")
}
