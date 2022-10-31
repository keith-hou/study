package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	inTime "study/internal/time"
	"time"
)

var timeCmd = &cobra.Command {
	Use:"now",
	Short:"时间转换",
	Long:"时间转换",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("当前时间： %d", inTime.GetNowTime().Unix())
	},
}

var d string
var currTime string

var calcCmd = &cobra.Command{
	Use:"calc",
	Short:"时间推算",
	Long:"时间推算",
	Run: func(cmd *cobra.Command, args []string) {
		var layout = "2006-01-02 15:04:05"
		tc :=  strings.Count(currTime, " ")
		if  tc == 0 {
			layout = "2006-01-02"
		}

		t, err := time.Parse(layout, currTime)
		if err != nil {
			log.Fatal("时间格式解析错误：", err)
		}
		ct , err := inTime.GetCalculateTime(t, d)
		if err != nil {
			log.Fatal("推算时间错误：", err)
		}
		log.Printf("推算时间戳%d", ct.Unix())
		log.Printf("推算时间 %s", ct.Format(layout))
	},
}

func init() {
	calcCmd.Flags().StringVarP(&currTime, "currTime", "c", "", "当前时间")
	calcCmd.Flags().StringVarP(&d, "d", "d", "", "延迟时间")
}