package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"study/internal/word"
)

const (
	ModeUpper  =  iota + 1
	ModeLower
	ModeUnderlineToUpperCamelCase
	ModeUnderlineToLowerCamelCase
	ModeCamelCaseToUnderline
)

var desc = strings.Join([]string{
	"该命令用于转换单词模式:",
	"1 全部转大写",
	"2 全部转小写",
	"3 下划线转大写驼峰",
	"4 下划线转小写驼峰",
	"5 驼峰转下划线",
}, ",")
var mode int8
var s string
var wordCmd = &cobra.Command{
	Use:"word",
	Short:"单词格式转换",
	Long:desc,
	Run:func(cmd *cobra.Command, args []string){
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(s)
		case ModeLower:
			content = word.ToLower(s)
		case ModeUnderlineToUpperCamelCase:
			content = word.UnderlineToUpperCamelCase(s)
		case ModeUnderlineToLowerCamelCase:
			content = word.UnderlineToLowerCamelCase(s)
		case ModeCamelCaseToUnderline:
			content = word.CamelCaseToUnderline(s)
		default:
			log.Fatalf("暂不支持该转换模式")
		}
		log.Printf("输出结果： %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&s, "s", "s", "", "请输入内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入内容")
}