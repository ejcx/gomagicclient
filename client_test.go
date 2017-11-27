package gomagicclient

import (
	"testing"
)

var (
	email  string
	apikey string
)

func init() {
	email = "evan@twiinsen.com"
	apikey = "9cb6604ff7401b2c69af3e65257762bed7019cb3279d8b09"
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

//func TestValidate(t *testing.T) {
//	client := NewEx(email, apikey, nil)
//	b, err := client.Validate("3edb4c1c4cfbe3faa6c17372e4f1744d02d4fbed6c79f3c2")
//	if err != nil {
//		t.Errorf("Could not send: %s", err)
//	}
//	fmt.Println(b)
//}
