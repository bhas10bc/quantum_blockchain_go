package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"y/types"

	"github.com/cloudflare/circl/sign/dilithium"
)

type PrivateKey struct {
	Key dilithium.PrivateKey
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	mode := dilithium.Mode3
	sign := mode.Sign(k.Key, data)
	

	return &Signature{
		Signature: sign,
	}, nil
}

func NewPrivateKeyFromReader(r io.Reader) PrivateKey {
	mode := dilithium.Mode3
	_, privKey, err := mode.GenerateKey(r)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		Key: privKey,
	}
}

func GeneratePrivateKey() PrivateKey {
	return NewPrivateKeyFromReader(rand.Reader)
}

func (k PrivateKey) PublicKey() PublicKey {
	if k.Key == nil {
		panic("private key is not initialized")
	}
	return PublicKey{
		Key: k.Key.Public().(dilithium.PublicKey),
	}
}

type PublicKey struct {
	Key dilithium.PublicKey
}

func (k PublicKey) String() string {
	
	if k.Key == nil {
		panic("public key is not initialized---11")
	}
	return hex.EncodeToString(k.Key.Bytes())
}

func (k PublicKey) Address() types.Address {
		if k.Key == nil {
		panic("public key is not initialized---22")
	}
	h := sha256.Sum256(k.Key.Bytes())
	return types.AddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	Signature []byte
}

func (sig Signature) String() string {
	return hex.EncodeToString(sig.Signature)
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool {
	mode := dilithium.Mode3
	return mode.Verify(pubKey.Key, data, sig.Signature)
}
