package main

import (
	"bytes"
	"flag"
	"fmt"

	"testing"

	log "github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
)

func main() {
	fmt.Println("Hello World")
}

func logSomething() {
	klog.Info("test this")
}

func LogsToBuffer(level log.Level, t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetLevel(level)
	klog.SetOutput(buf)
	flags := &flag.FlagSet{}
	klog.InitFlags(flags)
	// make sure klog doesn't write to stderr by default in tests
	_ = flags.Set("logtostderr", "false")
	_ = flags.Set("alsologtostderr", "false")
	_ = flags.Set("stderrthreshold", "4")
	return buf
}
