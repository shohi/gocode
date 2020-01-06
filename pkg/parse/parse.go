package parse

import (
	"log"
	"strings"
)

// Label is a key/value pair info.
type Label struct {
	Key   string
	Value string
}

// Labels is a collection of Label.
type Labels []Label

// ParseLabels parses labels string into []Labels.
// labels string is in the format - "k1:v1,k2:v2,k3".
// Key-only labels and value with colon are also supported.
func ParseLabels(input string) Labels {
	kvs := strings.Split(input, ",")
	if len(kvs) == 0 {
		return nil
	}

	labels := make(Labels, 0, len(kvs))
	for _, kv := range kvs {
		tkv := strings.TrimSpace(kv)
		tokens := strings.Split(tkv, ":")

		key := strings.TrimSpace(tokens[0])
		if key == "" {
			log.Printf("invalid label - %s", kv)
			continue

		}
		labels = append(labels, Label{
			Key:   key,
			Value: strings.Join(tokens[1:], ":"),
		})

	}
	if len(labels) > 0 {
		return labels
	}

	return nil
}
