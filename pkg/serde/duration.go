package serde

import (
	"encoding/json"
	"fmt"
	"time"
)

// Duration is an alias of time.Duration, which is convenient for marshalling from/to json.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s", time.Duration(d)))
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	tmp, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*d = Duration(tmp)
	return nil
}

func (d Duration) String() string {
	return time.Duration(d).String()
}
