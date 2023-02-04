package msign

import "io"

type exporter interface {
	export(io.Writer) error
}

type KeyId []byte
type PrivateKey interface {
	exporter
	Id() KeyId
	Public() PublicKey
	Sign(io.Reader) (Signature, error)
}

type PublicKey interface {
	exporter
	Id() KeyId
	Verify(io.Reader, Signature) (bool, error)
}

type Signature interface {
	exporter
	KeyId() KeyId
}
