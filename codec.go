package grpc_msgpack

import (
	"bytes"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/grpc/encoding"
)

func init() {
	codec := MsgPack{OmitEmpty: true, CustomStructTag: "json", UseArrayEncodedStructs: true}
	encoding.RegisterCodec(codec)
}

type MsgPack struct {
	OmitEmpty              bool
	CustomStructTag        string
	UseArrayEncodedStructs bool
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
	encoder.SetCustomStructTag(j.CustomStructTag)
	encoder.UseArrayEncodedStructs(j.UseArrayEncodedStructs)
	if err = encoder.Encode(v); err != nil {
		return
	}
	out = b.Bytes()
	return
}

func (j MsgPack) Unmarshal(data []byte, v interface{}) (err error) {
	return msgpack.Unmarshal(data, v)
}
