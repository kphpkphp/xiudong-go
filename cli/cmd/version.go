//所属的包
package cmd

import (
	"fmt"
	"os"
	//这是个用于构建命令行项目的库
	"github.com/spf13/cobra"
)

var (
	version = "unknown"
	commit  = "?"
	date    = "?"
)

var versionCmd = &cobra.Command{
	//使用方式
	Use:   "version",
	//描述
	Short: "check version",
	//详细描述
	Long:  `check showstart cli version`,
	//当命令被执行时调用的方法
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, `git commit hash: %s
build timestamp: %s
version: %s
`, commit, date, version)
	},
}

//将versionCmd命令添加到rootCmd命令中，一般rootCmd是整个应用程序的入口点
//所有引入包中的init()函数在程序启动时都会被执行，也仅会执行这一次
func init() {
	rootCmd.AddCommand(versionCmd)
}
