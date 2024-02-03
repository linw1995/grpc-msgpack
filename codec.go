package grpc_msgpack

import (
	"bytes"

	"google.golang.org/grpc/encoding"
	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	codec := MsgPack{OmitEmpty: true}
	encoding.RegisterCodec(codec)
}

type MsgPack struct {
	OmitEmpty bool
}

func (_ MsgPack) Name() string {
	return "msgpack"
}

func (j MsgPack) Marshal(v interface{}) (out []byte, err error) {
	var (
		b       = bytes.NewBuffer(nil)
		encoder = msgpack.NewEncoder(b)
	)
	encoder.SetOmitEmpty(j.OmitEmpty)
	if err = encoder.Encode(v); err != nil {
		return
	}
	out = b.Bytes()
	return
}

func (j MsgPack) Unmarshal(data []byte, v interface{}) (err error) {
	return msgpack.Unmarshal(data, v)
}
