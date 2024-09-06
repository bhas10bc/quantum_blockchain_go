package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	// address := pubKey.Address()

	fmt.Println(pubKey)
}

func TestKeypairSignVerifyFail(t *testing.T) {
	privkey := GeneratePrivateKey()
	msg := []byte("Hello")

	sig,err := privkey.Sign(msg)
	assert.Nil(t,err)

	otherPrivateKey := GeneratePrivateKey()
	otherpubKey := otherPrivateKey.PublicKey()
	assert.True(t, sig.Verify(privkey.PublicKey(), msg))

	assert.False(t, sig.Verify(otherpubKey, msg))
	assert.False(t, sig.Verify(otherpubKey, []byte("Whasd")))

}