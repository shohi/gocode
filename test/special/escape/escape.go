package main

import "log"

type session struct {
	id     string
	fields sessionFields
}

type field [2]interface{}
type sessionFields [2]field

func (s *session) getFields() sessionFields {
	return s.fields
}

func (s *session) setFields(f sessionFields) {
	s.fields = f
}

func emit(s session, fs []field) {
	res := append(fs, s.fields[:]...)

	log.Printf("field: %v", res)
}

func main() {
	s := session{id: "1"}
	s.setFields(sessionFields{
		{"k1", 1},
		{"k2", "v2"},
	})
	emit(s, []field{{"k3", "v3"}})

}
