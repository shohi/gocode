package basic

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatVsLog(t *testing.T) {
	fmt.Println("fmt first")
	log.Println("log first")
}

func TestFormatSprint(t *testing.T) {
	assert := assert.New(t)
	arr := []string{"hello", "world"}
	assert.NotNil(fmt.Sprint(arr))
	log.Println(fmt.Sprint(arr))
}
