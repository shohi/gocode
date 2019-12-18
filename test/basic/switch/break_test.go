package switch_test

import (
	"log"
	"testing"
)

func TestSwitch_Break(t *testing.T) {
	var flag = 10

	switch {
	case flag < 20:
		if flag < 15 {
			log.Printf("flag < 15")
			break
		}

		log.Printf("flag >= 15 && flag < 20")

	case flag < 100:
		log.Printf("flag < 100")

	default:
		log.Printf("flag >= 100")
	}

}
