package gomagicclient

import (
	"testing"
)

var (
	email  string
	apikey string
)

func init() {
	email = "e@ejj.io"
	apikey = "c2fbc139886b860cac261e0bc00a7e767b9e011299b289f7"
}

func TestSend(t *testing.T) {
	client := NewEx(email, apikey, nil)
	success, err := client.Send("e@ejj.io", "poopfeast", "https://hexauth.com/callback")
	if err != nil {
		t.Errorf("Could not send: %s", err)
	}
	if !success {
		t.Errorf("Could not send. False success")
	}
}

func TestValidate(t *testing.T) {
	client := NewEx(email, apikey, nil)
	_, err := client.Validate("3edb4c1c4cfbe3faa6c17372e4f1744d02d4fbed6c79f3c2")
	if err != nil {
		t.Errorf("Could not send: %s", err)
	}

}
