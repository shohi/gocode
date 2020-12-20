package time_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// https://github.com/polaris1119/golangweekly/blob/master/docs/issue-074.md
func TestTime_Unmarshal(t *testing.T) {
	s := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	// TODO: Why?
	// Out: 2020-12-20T00:00:00Z
	m, _ := json.Marshal(s)
	fmt.Printf("%s\n", m)
}
