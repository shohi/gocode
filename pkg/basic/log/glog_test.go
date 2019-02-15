package log_test

import (
	"flag"
	"testing"

	"github.com/golang/glog"
)

func TestGLog_Basic(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	flag.Parse()
	glog.Infof("hello")
	glog.Warningf("hello")
	glog.V(0).Infof("LINE: %d", 10)
	glog.Error("hello")
	glog.Flush()

}
