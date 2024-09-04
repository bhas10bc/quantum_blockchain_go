package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	address := pubKey.Address()

	fmt.Println(address)
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privkey := GeneratePrivateKey()
	msg := []byte("Hello")

	sig := privkey.Sign(msg)

	otherPrivateKey := GeneratePrivateKey()
	otherpubKey := otherPrivateKey.PublicKey()
	assert.True(t, sig.Verify(privkey.PublicKey(), msg))

	assert.False(t, sig.Verify(otherpubKey, msg))
	assert.False(t, sig.Verify(otherpubKey, []byte("Whasd")))

}