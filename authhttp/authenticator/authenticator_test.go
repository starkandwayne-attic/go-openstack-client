package authenticator

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
    "net/http"
    "go_openstack_client/authhttp/mockauthentication"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type AuthenticatorTestSuite struct{}

var _ = gocheck.Suite(&AuthenticatorTestSuite{})

func (t *AuthenticatorTestSuite) SetUpSuite (c *gocheck.C) {
}

func (t *AuthenticatorTestSuite) Test_AuthenticateRequest (c *gocheck.C) {
    authorizers := Authenticators{}
    authorizers.Add(mockauthentication.Authenticator{},true)

    req, _ := http.NewRequest("GET", "http://localhost", nil)

    authorized, err := authorizers.AuthenticateRequest(req)

    c.Assert(authorized, gocheck.Equals, false)
    c.Assert(err.Error(), gocheck.Equals, "Unauthorized")

    req.Header.Add("Authorization", "Pickles")

    authorized, err = authorizers.AuthenticateRequest(req)

    c.Assert(authorized, gocheck.Equals, false)
    c.Assert(err.Error(), gocheck.Equals, "Unsupported authentication provider requested.")

    req.Header.Del("Authorization")
    req.Header.Add("Authorization", "TestAUTH")

    authorized, err = authorizers.AuthenticateRequest(req)

    c.Assert(authorized, gocheck.Equals, false)
    c.Assert(err.Error(), gocheck.Equals, "Unauthorized")

    req.Header.Add("Authorization","Jeremy:PicklesRGR8")

    authorized, err = authorizers.AuthenticateRequest(req)

    c.Assert(authorized, gocheck.Equals, true)
    c.Assert(err, gocheck.Equals, nil)
}
