package db

import (
	tcp_model "github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	tls_model "github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	proto_tcp "github.com/k8spacket/k8spacket/internal/proto/nodegraph/model"
	proto_tls "github.com/k8spacket/k8spacket/internal/proto/tlsparser/model"
	"google.golang.org/protobuf/proto"
)

// Converter functions for TLSDetails
func tlsDetailsToProto(in *tls_model.TLSDetails) *proto_tls.TLSDetails {
	_ = "STUB: not implemented"
	return nil
}

func tlsDetailsFromProto(in *proto_tls.TLSDetails) *tls_model.TLSDetails {
	_ = "STUB: not implemented"
	return nil
}

// Converter functions for TLSConnection
func tlsConnectionToProto(in *tls_model.TLSConnection) *proto_tls.TLSConnection {
	_ = "STUB: not implemented"
	return nil
}

func tlsConnectionFromProto(in *proto_tls.TLSConnection) *tls_model.TLSConnection {
	_ = "STUB: not implemented"
	return nil
}

// Converter functions for ConnectionItem
func connectionItemToProto(in *tcp_model.ConnectionItem) *proto_tcp.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}

func connectionItemFromProto(in *proto_tcp.ConnectionItem) *tcp_model.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}

// marshalProto marshals a domain model to protobuf
func marshalProto(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	// Handle basic types that bolthold might try to encode
	return nil, nil
}

// unmarshalProto unmarshals protobuf to domain model
func unmarshalProto(data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	// Handle basic types that bolthold might try to decode
	return nil
}

// marshalMessage is a helper to marshal protobuf messages
func marshalMessage(msg proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// unmarshalMessage is a helper to unmarshal protobuf messages
		nil
}

func unmarshalMessage(data []byte, msg proto.Message) error { _ = "STUB: not implemented"; return nil }
