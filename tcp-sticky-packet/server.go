package main

import (
	"bytes"
	"encoding/binary"
	"net"
)


var deviceName ="{1121A8E9-251B-4059-BF5C-225B605E43E0}"
var serverSenderIp=net.ParseIP("10.33.0.103")

func main() {

}


// 利用数据的前4个字节代表长度
func Encode(message string) ([]byte, error) {
	// 读取消息的长度
	var length  = int32(len(message))
	var pkg  = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}



