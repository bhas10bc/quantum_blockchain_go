package core

import (
	"encoding/gob"
	"io"

	"github.com/cloudflare/circl/sign/dilithium"
)

type Encoder[T any] interface {
	Encode(T) error
}

type Decoder[T any] interface {
	Decode(T) error
}

type GobTxEncoder struct {
	w io.Writer
}

func init() {

	mode := dilithium.Mode3

	// Register the public and private key types by creating sample instances
	_, privKey, err := mode.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	pubKey := privKey.Public() // Get public key from private key

	gob.Register(privKey)
	gob.Register(pubKey)
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	return &GobTxEncoder{w: w}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(e.w).Encode(tx)
}

type GobTxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	return &GobTxDecoder{r: r}
}

func (e *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)
}
