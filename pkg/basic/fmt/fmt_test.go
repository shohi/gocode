package fmt_test

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

func TestFormatEnum(t *testing.T) {
	assert := assert.New(t)

	type RaftState uint32
	const (
		Follower uint32 = iota
		Leader
		Candidate
		Shutdown
	)

	s := fmt.Sprintf("%v", Follower)

	// assert.Equal(s, "Follower")
	assert.Equal(s, "0")
}
