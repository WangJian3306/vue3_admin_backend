package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machinedID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	// 默认输出的是一个 64 位的数字，如下设置减少节点和序列位数，从而减少输出的ID的总位数（为了兼容 JavaScript number 数据类型最大值 2 ** 53 - 1）
	sf.NodeBits = 3
	sf.StepBits = 7
	node, err = sf.NewNode(machinedID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2020-07-01", 1); err != nil {
//		fmt.Printf("init failed, err:%v \n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
