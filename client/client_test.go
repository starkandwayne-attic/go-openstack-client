package client

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
    "go-openstack-client/authhttp/authenticator"
    "go-openstack-client/authhttp/mockauthentication"
    "go-openstack-client/authhttp/mockserver"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type ClientTestSuite struct{}

var _ = gocheck.Suite(&ClientTestSuite{})

func (t *ClientTestSuite) SetUpSuite (c *gocheck.C) {
    authorizers := authenticator.Authenticators{}
    authorizers.Add(mockauthentication.Authenticator{},true)

    mockServer := mockserver.Server{}
    go mockServer.Start(authorizers, "8083")
}

