package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"y/core"
	"y/crypto"
	"y/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	// trRemoteA := network.NewLocalTransport("REMOTE_A")
	// trRemoteB := network.NewLocalTransport("REMOTE_B")
	// trRemoteC := network.NewLocalTransport("REMOTE_C")

	// trLocal.Connect(trRemoteA)
	// trRemoteA.Connect(trRemoteB)
	// trRemoteB.Connect(trRemoteC)
	// trRemoteA.Connect(trLocal)

	// initRemoteServers([] network.Transport{trRemoteA,trRemoteB,trRemoteC})
	// go func() {
	// 	for {
	// 		if err := sendTransaction(trRemoteA, trLocal.Addr()); err != nil {
	// 			logrus.Error(err)
	// 		}
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	privKey := crypto.GeneratePrivateKey()
 localServer := makeServer("LOCAL", trLocal, &privKey)
 localServer.Start()
}

func initRemoteServers(trs []network.Transport){
	for i :=0; i < len(trs); i++ {
		id := fmt.Sprintf("Remote_%d",i)
		s := makeServer(id,trs[i],nil)
		go s.Start()
	}
}

func makeServer (id string, tr network.Transport, pk *crypto.PrivateKey) *network.Server{
	opts := network.ServerOpts{
		PrivateKey: pk,
		ID:         id,
		Transports: []network.Transport{tr},
	}
	s, err := network.NewServer(opts)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000000000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())
}