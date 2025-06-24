package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// Init 初始化
func Init(startTime string, machineID int64) (err error) {
	// 解析时间字符串
	var st time.Time
	// 2006-01-02 15:04:05 是Go语言中时间格式化的模板，用于将时间对象转换为字符串。
	st, err = time.Parse("2006-01-02", startTime)
	// 如果解析失败，输出错误信息并返回
	if err != nil {
		fmt.Println("time.Parse err:", err)
		return
	}
	// 设置snowflake的初始时间
	snowflake.Epoch = st.UnixNano() / 1000000
	// 创建一个节点
	node, err = snowflake.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
