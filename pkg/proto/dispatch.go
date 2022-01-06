package proto

import (
	"fmt"
	"strings"
)

type ProtoNode interface {
	Stickpacket(buf []byte) int                //粘包
	PacketProcess(completePacket []byte) error //完整包处理
}

func NewProtoNode(protoType string) (ProtoNode, error) {
	switch strings.ToLower(protoType) {
	case "v1":
		return newV1(), nil
	case "v2":
		return newV2(), nil
	default:
		return nil, fmt.Errorf("unknown proto type: %s", protoType)
	}
}
