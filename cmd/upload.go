package cmd

import (
	"fmt"
	"github.com/feivxs/cbtool/common"
	"github.com/feivxs/cbtool/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var filepath []string

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传图片到 Cnblog",
	Run:   doUpload,
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringArrayVarP(&filepath, "file-path", "f", nil, "文件路径")
	loginCmd.MarkFlagRequired("file-path")
}

func doUpload(cmd *cobra.Command, args []string) {

	if len(filepath) <= 0 {
		os.Stderr.WriteString("No file path specified...\n")
		os.Exit(1)
	}

	for _, fp := range filepath {
		if exist, _ := utils.IsPathExists(fp); !exist {
			os.Stderr.WriteString("图片【" + fp + "】成功！")
			continue
		}

		fileBytes, _ := ioutil.ReadFile(fp)
		args := common.UploadRequest{
			Username: utils.GetConfig().Username,
			Password: utils.GetConfig().Password,
			BlogId:   "14270120",
			File: common.FileInfo{
				Bits: fileBytes,
				Name: "alalala",
				Type: utils.GetImageType(fp),
			},
		}
		utils.Upload(args)
	}

	fmt.Printf("图片【%v】成功！\n", filepath)
}
