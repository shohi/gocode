package regexp_test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRe_MatchString(t *testing.T) {
	re := regexp.MustCompile(".*" + "github.comcast.com" + ".*")

	result := re.MatchString("github.com/prometheus/common")

	fmt.Println(result)
}
