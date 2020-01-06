package build_test

import (
	"go/build"
	"log"
	"testing"
)

func TestBuildContext(t *testing.T) {
	log.Printf("build context: %+v", build.Default)
}

// refer, https://stackoverflow.com/questions/29087241/finding-imports-and-dependencies-of-a-go-program
func TestBuild_Deps(t *testing.T) {
	pkg, err := build.Default.Import(".", ".", 0)

	log.Printf("packages: %v, error: %v", pkg, err)
	if err == nil {
		log.Printf("imports: %v", pkg.Imports)
	}
}
