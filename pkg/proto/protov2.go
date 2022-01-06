package proto

import (
	"fmt"
	"time"
	"unsafe"
)

type nodeProtoV2 struct {
	data   []byte
	length int
}

type protoV2Head struct {
	TotalLen   int
	MsgType    int
	SequenceId int
}

func newV2() *nodeProtoV2 {
	return &nodeProtoV2{
		data: make([]byte, 0),
	}
}

func (n *nodeProtoV2) Stickpacket(buf []byte) int {
	n.data = append(n.data, buf...)
	n.length = len(n.data)
	head := (*protoV2Head)(unsafe.Pointer(&n.data[0]))
	for {
		if n.length != 0 && n.length >= head.TotalLen {
			completePacket := n.data[:head.TotalLen]
			n.PacketProcess(completePacket)
			n.data = n.data[head.TotalLen:]
			n.length -= head.TotalLen
			if n.length > 0 {
				head = (*protoV2Head)(unsafe.Pointer(&n.data[0]))
			}
		} else {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return 0
}

// 处理一个完整的包
func (n *nodeProtoV2) PacketProcess(buf []byte) error {
	head := *(*protoV2Head)(unsafe.Pointer(&n.data[0]))
	str := unsafe.Pointer(uintptr(unsafe.Pointer(&n.data[0])) + uintptr(unsafe.Sizeof(head)))
	switch head.MsgType {
	case 1: //
	case 2: //
	case 3: // 写文件或输出
		fmt.Printf("%v， %v\n", head, *(*string)(str))
	}
	return nil
}
