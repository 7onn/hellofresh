package logger

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type customOutput struct {
	Msg string
}

//Err is fot printing at stderr
func Err(Msg string) {
	e := customOutput{
		Msg,
	}
	o, _ := json.Marshal(e)
	io.Copy(os.Stderr, strings.NewReader(string(o)+"\n"))
}

//Out is for printing at stdout
func Out(Msg string) {
	m := customOutput{
		Msg,
	}
	o, _ := json.Marshal(m)
	io.Copy(os.Stderr, strings.NewReader(string(o)+"\n"))
}
