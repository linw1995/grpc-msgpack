package grpc_msgpack

import (
	"google.golang.org/grpc/encoding"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	encoding.RegisterCodec(MsgPack{})
}

type MsgPack struct {
}

func (_ MsgPack) Name() string {
	return "msgpack"
}

func (j MsgPack) Marshal(v interface{}) (out []byte, err error) {
	return msgpack.Marshal(v)
}

func (j MsgPack) Unmarshal(data []byte, v interface{}) (err error) {
	return msgpack.Unmarshal(data, v)
}
