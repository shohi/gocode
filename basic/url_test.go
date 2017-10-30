package basic

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	fmt.Println(url.PathEscape("lang:>50"))
}
