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

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	return &GobTxEncoder{
		w: w,
	}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(e.w).Encode(tx)
}

type GobTxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	return &GobTxDecoder{
		r: r,
	}
}

func (e *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)
}

type GobBlockEncoder struct {
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	return &GobBlockEncoder{
		w: w,
	}
}

func (enc *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(enc.w).Encode(b)
}

type GobBlockDecoder struct {
	r io.Reader
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {
	return &GobBlockDecoder{
		r: r,
	}
}

func (dec *GobBlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(dec.r).Decode(b)
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