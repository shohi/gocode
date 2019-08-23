package serde

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration_Marshall(t *testing.T) {
	assert := assert.New(t)

	d := Duration(10 * time.Second)
	b, err := json.Marshal(d)

	assert.Nil(err)
	assert.Equal(`"10s"`, string(b))

	assert.Equal("10s", d.String())
}

func TestDuration_Unmarshall(t *testing.T) {
	tests := []struct {
		name string
		// input
		val string

		// output
		expNilErr bool
		expVal    Duration
	}{
		{"success", `"10s"`, true, Duration(10 * time.Second)},
		{"failure-no-quote", `10s`, false, Duration(0)},
		{"failure-int", `10`, false, Duration(0)},
		{"failure-other-string", `"abc"`, false, Duration(0)},
		{"corner-zero", `"0"`, true, Duration(0)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var d Duration
			err := json.Unmarshal([]byte(test.val), &d)

			assert.Equal(t, test.expNilErr, err == nil)
			if err == nil {
				assert.Equal(t, test.expVal, d)
			}
		})
	}
}

func TestDuration_Embed(t *testing.T) {
	type myStruct struct {
		Key string   `json:"k"`
		Dur Duration `json:"d"`
	}

	tests := []struct {
		name    string
		skipped bool

		// input
		val string

		// output
		expNilErr bool
		expVal    Duration
	}{
		{"success", false,
			`{"d": "10s"}`, true, Duration(10 * time.Second)},
		{"failure-no-quote", false,
			`{"d":10s}`, false, Duration(0)},
		{"failure-int", false,
			`{"d":10}`, false, Duration(0)},
		{"failure-other-string", false,
			`{"d":"abc"}`, false, Duration(0)},
		{"corner-zero", false,
			`{"d":"0"}`, true, Duration(0)},
		{"corner-no-value", false,
			`{"d":}`, false, Duration(0)},
		{"corner-no-entry", false,
			`{"key": "key"}`, true, Duration(0)},
		{"corner-empty", false,
			`{}`, true, Duration(0)},
	}
	for _, test := range tests {
		if test.skipped {
			continue
		}

		t.Run(test.name, func(t *testing.T) {
			var m myStruct
			err := json.Unmarshal([]byte(test.val), &m)

			assert.Equal(t, test.expNilErr, err == nil)
			if err == nil {
				assert.Equal(t, test.expVal, m.Dur)
			}
		})
	}
}
