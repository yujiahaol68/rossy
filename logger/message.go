package logger

import (
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
	fmt.Printf("%s: %s\n", m.Level, m.Msg)
}

func NewErrMsg(e error) *Message {
	return &Message{
		Level: defaultLevel,
		Msg:   e.Error(),
	}
}

// func NewMsg(args []string, e error) *Message {
// 	if e != nil {
// 		defaultLevel := "ERROR"

// 		switch len(args) {
// 		case 1:

// 		}
// 	}
// }
