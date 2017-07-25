# gomagicclient
gomagicclient is the client for using (hexauth)[https://hexauth.com]!

## Usage

*Installing*
```
go get github.com/ejcx/gomagicclient
```

### Sending a link

```
c := gomagicclient.NewEx("youremail@youremail.com", "deadbeefdeadbeef", nil)
c.Send("signmein@gmail.com", "MyBusiness", "https://callback")

```

### Validating a link

```
c := gomagicclient.NewEx("youremail@youremail.com", "deadbeefdeadbeef", nil)
signin, err := c.Validate("signmein@gmail.com", "MyBusiness", "https://callback")
if !signin.Verified {
  // Unsuccessful login
  return
}

```

## Loading your email and apikey
You can create a client in two ways. First is explicitly
and second is by using environment variables to set your
email address and apikey.

Any easier way to create a client, instead of the above
examples is to set the environment variables that contain
your email address and apikey.

```
MAGIC_EMAIL=youremail@youremail.com
MAGIC_APIKEY=deadbeefdeadbeef
```

The code to create a client will then be much simpler.

```
c := gomagicclient.New()
```
