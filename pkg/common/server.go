package common

import (
	"fmt"
	"net"
	"sync"

	"k8s.io/klog"
	"sdtp.io/pkg/proto"
)

type NonBlockingServer interface {
	Start(endpoint, protoType string)
	Wait()
	// Stop()
	// ForceStop()
}

func NewNonBlockingServer() NonBlockingServer {
	return &nonBlockingServer{}
}

type nonBlockingServer struct {
	wg         sync.WaitGroup
	protoNodes proto.ProtoNode
}

func (s *nonBlockingServer) Start(endpoint, protoType string) {
	err := s.Init(protoType)
	if err != nil {
		klog.Fatal(err)
		return
	}
	s.serve(endpoint)
}

func (s *nonBlockingServer) Wait() {
	s.wg.Wait()
}

func (s *nonBlockingServer) Init(protoType string) error {
	var err error
	s.protoNodes, err = proto.NewProtoNode(protoType)
	if err != nil {
		return err
	}
	fmt.Println("proto type is ", protoType)
	return nil
}

func (s *nonBlockingServer) serve(endpoint string) {
	proto, addr, err := parseEndpoint(endpoint)
	if err != nil {
		klog.Fatal(err.Error())
	}

	listener, err := net.Listen(proto, addr)
	if err != nil {
		klog.Fatal("Failed to listen: %v", err)
	}
	for {
		rawConn, err := listener.Accept()
		if err != nil {
			// TODO
			return
		}
		fmt.Println("client connect ...")
		s.wg.Add(1)
		go func() {
			s.handleRawConn(rawConn)
			s.wg.Done()
		}()
	}
}

func (s *nonBlockingServer) handleRawConn(c net.Conn) {
	var buf []byte = make([]byte, 100)

	for {
		n, err := c.Read(buf)
		if err != nil {
			// TODO
			return
		}
		defer c.Close()
		// 接收到的数据 粘包
		s.protoNodes.Stickpacket(buf[:n])
	}
}
