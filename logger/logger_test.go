package logger_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/yujiahaol68/rossy/logger"
)

var (
	errSample = errors.New("Sample")
	//errTest   = errors.New("Test")
)

func Test_ErrMsgFactory(t *testing.T) {
	m := logger.NewErrMsg(errSample)

	if m.Level != "ERROR" {
		t.Fatalf("Err message should be <Error> level, got <%s> level", m.Level)
	}

	if m.Msg != "Sample" {
		t.Fatalf("Message should exactly the same with error, but got: %s", m.Msg)
	}
}

func Test_logInfo(t *testing.T) {
	mi := logger.Message{Level: "info", Msg: "test"}
	md := logger.Message{Level: "debug", Msg: "testing"}

	fmt.Println("Info level msg will not show its level")
	mi.ShowInCmd()
	fmt.Println("Other level msg should show its level")
	md.ShowInCmd()
}

func Test_msgGroup(t *testing.T) {
	_, err := logger.NewMsgGroup("messy")
	if err == nil {
		t.Fatalf("NewMsgGroup need at least 2 parameter, err should not be nil")
	}

	mg, err := logger.NewMsgGroup("debug", "msg1", "msg2", "msg3")
	if err != nil {
		t.Fatal(err)
	}

	for i, m := range mg {
		if m.Msg != fmt.Sprintf("msg%d", i+1) {
			t.Fatalf("expect msg group contains: msg1, msg2, msg3\nBut got: %s", m.Msg)
		}
	}
}
