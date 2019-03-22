package build_test

import (
	"go/build"
	"log"
	"testing"
)

func TestBuildContext(t *testing.T) {
	log.Printf("build context: %+v", build.Default)
}
