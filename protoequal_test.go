package gomockgrpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FakeProtoMessage struct {
}

func (f *FakeProtoMessage) ProtoReflect() protoreflect.Message {
	return nil
}

func TestProtoEqual(t *testing.T) {
	wantMessage := &FakeProtoMessage{}
	got := ProtoEqual(wantMessage)
	var pe *protoEqual
	require.IsType(t, got, pe)
	pe = got.(*protoEqual)
	assert.Equal(t, wantMessage, pe.expected)
}

func Test_protoEqual_Matches(t *testing.T) {

	wantMessage := helloworld.HelloRequest{
		Name: "Given Name",
	}

	wantMessageData, err := proto.Marshal(&wantMessage)
	require.NoError(t, err, "failed marshaling data")

	var encodedMessage helloworld.HelloRequest
	err = proto.Unmarshal(wantMessageData, &encodedMessage)
	require.NoError(t, err, "failed unmarshaling data")

	tests := []struct {
		name  string
		given interface{}
		want  bool
	}{
		{"should return false when given match is not a proto.Message", "not a proto.Message", false},
		{"should match", &wantMessage, true},
		{"should not match due to different values", &helloworld.HelloRequest{
			Name: "Not Given Name",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &protoEqual{expected: &encodedMessage}

			got := p.Matches(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_protoEqual_String(t *testing.T) {
	wantString := `is equal to name:  "Given Name"
`
	gotString := ProtoEqual(&helloworld.HelloRequest{
		Name: "Given Name",
	}).String()
	assert.Equal(t, wantString, gotString)
}

func TestGomockMatcher(t *testing.T) {
	//
}
