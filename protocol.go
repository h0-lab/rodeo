package rodeo

// すべてのKVSに対応する
// 文字長、文字列の整形変換などする
type Protocol interface {
	Request(args ...interface{}) Protocol
	Execute() Protocol
	ToResult() Result
}
type Result struct{}
