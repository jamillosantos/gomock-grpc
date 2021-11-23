package gomockgrpc

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

type protoEqual struct {
	expected proto.Message
}

// ProtoEqual is a gomock.Matcher that implements gomock.Eq but for GRPC messages. With this matcher you can compare
// messages directly without needing to worry about the internal fields.
func ProtoEqual(m proto.Message) gomock.Matcher {
	return &protoEqual{m}
}

func (p *protoEqual) Matches(x interface{}) bool {
	m, ok := x.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(p.expected, m)
}

func (p protoEqual) String() string {
	return fmt.Sprintf("is equal to %s", prototext.Format(p.expected))
}
