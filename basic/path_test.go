package basic

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPathJoin(t *testing.T) {
	fmt.Println(filepath.Join("", "a"))
	fmt.Println(filepath.Join("", "a", ".dat"))
}
