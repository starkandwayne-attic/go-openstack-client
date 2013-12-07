package handler

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type AuthHandlerTestSuite struct{}

var _ = gocheck.Suite(&AuthHandlerTestSuite{})

func (t *AuthHandlerTestSuite) Test_ServeHTTP (c *gocheck.C) {
    //Test not needed - function is used in testserver.go, and exercised in
    //client_test.go.
}
