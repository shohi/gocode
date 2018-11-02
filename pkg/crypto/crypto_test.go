package crypto_test

import (
	"log"
	"testing"

	"github.com/nats-io/nkeys"
)

// Use nkeys and the public key.
func genID() string {
	kp, _ := nkeys.CreateServer()
	pub, _ := kp.PublicKey()
	return pub
}

func TestEd25519WithHelpOfNATSUtil(t *testing.T) {
	var id string
	count := 10
	for k := 0; k < count; k++ {
		// will generate 32-byte unique id
		// e.g. NDEN5OETJ6FF7QNLGIDFWW2GCYULOZUNX326SOO7YTEAFBZBXZXYHBCP
		id = genID()
		log.Printf("unique id [%d]: %v", k, id)
	}
}
