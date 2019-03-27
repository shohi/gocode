package cpu_test

import (
	"log"
	"testing"

	"github.com/klauspost/cpuid"
)

func TestCPU(t *testing.T) {
	log.Printf("cpu info: %+v", cpuid.CPU)
}
