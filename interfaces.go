package msign

import "io"

type exporter interface {
	export(io.Writer) error
}

type PrivateKey interface {
	exporter
	Sign(io.Reader) (Signature, error)
}

type PublicKey interface {
	exporter
	Verify(io.Reader, Signature) (bool, error)
}

type Signature interface {
	exporter
}
