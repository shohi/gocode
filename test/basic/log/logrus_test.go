package log_test

import (
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

type MyData struct {
	Key   string
	Value string
}

func TestLogrus_Formatter(t *testing.T) {

	fr := new(logrus.TextFormatter)
	fr.TimestampFormat = "02-01-2006 15:04:05"
	fr.FullTimestamp = true

	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetFormatter(fr)
	logrus.Info("Some \n info")
	logrus.Warnf("Some warning")

	str := "some \n warn"
	logrus.Warnf(str)

	d := MyData{
		Key:   "season",
		Value: "winter",
	}
	dj, _ := json.Marshal(&d)

	logrus.Warnf("values: %v", string(dj))

	// s, err := strconv.Unquote(string(dj))
	// log.Printf("values: %v, err: %v", s, err)
}

func TestLogrus_Default(t *testing.T) {
	logrus.Info("hello")
}
