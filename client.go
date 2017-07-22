package magicclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	MagicEmv = "MAGIC_EMAIL"
	MagicEnv = "MAGIC_APIKEY"
)

type MagicClient struct {
	e string
	k string
	c *http.Client
}

type MagicSignin struct {
	ToEmail     string
	CompanyName string
	Verified    bool

	SentTime     time.Time
	VerifiedTime time.Time
}

type MagicSend struct {
	CallbackURL string
	CompanyName string
	ToEmail     string
}

func magicRequest(method, route, email, apikey string, r io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "http://localhost:8000"+route, r)
	if err != nil {
		return nil, fmt.Errorf("Could not create magic request: %s", err)
	}
	req.SetBasicAuth(email, apikey)

	return req, err
}

func (m *MagicClient) Validate(callback string) (*MagicSignin, error) {
	req, err := magicRequest("GET", "/validate/"+callback, m.e, m.k, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.c.Do(req)
	if err != nil {
		return nil, err
	}
	cb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(cb))
	var s MagicSignin
	err = json.Unmarshal(cb, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (m *MagicClient) Send(ToEmail string, CompanyName string, CallbackURL string) (bool, error) {
	ms := &MagicSend{
		ToEmail:     ToEmail,
		CallbackURL: CallbackURL,
		CompanyName: CompanyName,
	}
	buf, err := json.Marshal(ms)
	if err != nil {
		return false, err
	}
	req, err := magicRequest("PUT", "/send", m.e, m.k, bytes.NewReader(buf))
	if err != nil {
		return false, err
	}
	resp, err := m.c.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode >= 200 || resp.StatusCode <= 299 {
		return true, nil
	}
	return false, fmt.Errorf("Did not successfully send magic link.")
}

func getEmv() string {
	return os.Getenv(MagicEmv)
}

func getEnv() string {
	return os.Getenv(MagicEnv)
}

func n(k string, e string, c *http.Client) *MagicClient {
	if c == nil {
		c = &http.Client{}
	}
	return &MagicClient{k: k, e: e, c: &http.Client{}}
}

// New generates a new MagicClient by loading
// it from the environment variable.
func New() *MagicClient {
	return n(getEnv(), getEmv(), nil)
}

// New generates a new MagicClient by loading
// explicitly allowing it as a parameter. If
// the environment variable is set, it will
// instead return a New object containing the
// environment variable.
func NewEx(email, key string, c *http.Client) *MagicClient {
	k := getEnv()
	if k != "" {
		key = k
	}
	e := getEmv()
	if e != "" {
		email = e
	}
	return n(key, email, c)
}
