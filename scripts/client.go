package scripts

import (
	"fmt"
	"net"
	"os"
	"unsafe"

	"k8s.io/klog"
)

type head struct {
	totalLen   int
	msgType    int
	sequenceId int
}
type testData struct {
	head
	data string
}
type SliceTestData struct {
	addr uintptr
	len  int
	cap  int
}

func Client() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		klog.Fatal(err)
	}
	defer conn.Close()

	var data testData
	data.data = "hello world!"
	header := head{
		totalLen:   int(unsafe.Sizeof(data)),
		msgType:    3,
		sequenceId: 1,
	}

	data.head = header
	len := unsafe.Sizeof(data)

	byteData := SliceTestData{
		addr: uintptr(unsafe.Pointer(&data)),
		len:  int(len),
		cap:  int(len),
	}
	structToBytes := *(*[]byte)(unsafe.Pointer(&byteData))
	fmt.Println("structToBytes ", structToBytes)
	bytesToStruct := *(*testData)(unsafe.Pointer(&structToBytes[0]))
	fmt.Println("bytesToStruct ", bytesToStruct)
	structToBytes = append(structToBytes, structToBytes...)

	for {
		_, err := conn.Write(structToBytes)
		if err != nil {
			klog.Fatal(err)
			break
		}
		conn.Close()
		os.Exit(0)
	}
}
