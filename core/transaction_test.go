package core

import (
	"testing"
	"y/crypto"

	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("test"),
	}
	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestTransactionVerify(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("test"),
	}
	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())

}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("test"),
	}

	assert.Nil(t, tx.Sign(privKey))

	return tx
}