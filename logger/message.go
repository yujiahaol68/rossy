package logger

import (
	"errors"
	"fmt"
)

// TODO: format word Map

const (
	defaultLevel = "ERROR"
)

type Message struct {
	Level string
	Msg   string
}

func (m Message) ShowInCmd() {
	if m.Level != "info" {
		fmt.Printf("%s: %s\n", m.Level, m.Msg)
		return
	}

	fmt.Println(m.Msg)
}

func NewErrMsg(e error) *Message {
	return &Message{
		Level: defaultLevel,
		Msg:   e.Error(),
	}
}

// NewMsgGroup will log info msg to term. Args[0] is level and others are many msgs
func NewMsgGroup(args ...string) ([]*Message, error) {
	if len(args) < 2 {
		//panic("Cmd Msg need at least 2 parameters")
		return nil, errors.New("Cmd Msg need at least 2 parameters")
	}
	s := make([]*Message, len(args)-1)

	level := args[0]

	for i := 1; i < len(args); i++ {
		s[i-1] = &Message{level, args[i]}
	}
	return s, nil
}
