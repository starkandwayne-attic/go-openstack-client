package nova

import (
    _"fmt"
    "launchpad.net/gocheck"
    "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type NovaTestSuite struct{}

var _ = gocheck.Suite(&NovaTestSuite{})

func (t *NovaTestSuite) Test_Authorization_NoCreds (c *gocheck.C) {
    n := New("http://10.150.0.60:35757","bosh","bosh","bosh")
    n.PrintServerList()
}
