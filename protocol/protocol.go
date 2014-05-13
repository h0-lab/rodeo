package protocol

import "net"

// すべてのKVSに対応する
// 文字長、文字列の整形変換などする
type Protocol interface {
	Request(args ...string) Protocol
	Execute(conn net.Conn) Protocol
	WaitFor(conn net.Conn, reciever *chan string)
	ToResult() Result
}
type Result struct {
	Response string
	Error    error
}
