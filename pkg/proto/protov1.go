/*
	name : proto v1
	包总长	4字节
	消息类型 1字节
	序列号	4字节
	type 1 消息包4个字节  前2个是主版本 后2字节是次版本号

*/
package proto

import (
	"fmt"
	"time"
	"unsafe"
)

type nodeProto struct {
	data   []byte
	length int
}
type protoV1 struct {
	TotalLen   int
	MsgType    int
	SequenceId int
	data       string
}

func newV1() *nodeProto {
	return &nodeProto{
		data:   make([]byte, 0),
		length: 0,
	}
}

// tcp粘包
func (n *nodeProto) Stickpacket(buf []byte) int {

	n.data = append(n.data, buf...)
	n.length = len(n.data)
	head := *(*protoV1)(unsafe.Pointer(&n.data[0]))
	for {
		if n.length != 0 && n.length >= head.TotalLen {
			completePacket := n.data[:head.TotalLen]
			n.PacketProcess(completePacket)
			n.data = n.data[head.TotalLen:]
			n.length -= head.TotalLen
			if n.length > 0 {
				head = *(*protoV1)(unsafe.Pointer(&n.data[0]))
			}
		} else {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return 0
}

// 处理一个完整的包
func (n *nodeProto) PacketProcess(buf []byte) error {
	head := *(*protoV1)(unsafe.Pointer(&n.data[0]))
	switch head.MsgType {
	case 1: //
	case 2: //
	case 3: // 写文件或输出
		fmt.Printf("%v\n", head)
	}
	return nil
}
