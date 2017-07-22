package magicclient

import (
	"testing"
)

var (
	email  string
	apikey string
)

func init() {
	email = "e@ejj.io"
	apikey = "8a87ff6c3ca402705d913bf13a0aec42fa4981835d901ea3"
}

func TestSend(t *testing.T) {
	client := NewEx(email, apikey, nil)
	success, err := client.Send("e@ejj.io", "poopfeast", "http://localhost:8080/callback")
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
