package floating_ips

import (
	"fmt"
	"git.smf.sh/jrbudnack/go_openstack_client/apitestharness"
	"launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type FloatingIpsTestSuite struct {
	ApiTestHarness apitestharness.ApiTestHarness
}

var _ = gocheck.Suite(&FloatingIpsTestSuite{})

func (t *FloatingIpsTestSuite) SetUpSuite(c *gocheck.C) {
	t.ApiTestHarness = apitestharness.New("compute", false)
}

func (t *FloatingIpsTestSuite) Test_List(c *gocheck.C) {
	floatingIps := New(t.ApiTestHarness.ApiConnection)
	list, _ := floatingIps.List()
	fmt.Println(list)
}

func (t *FloatingIpsTestSuite) Test_GetById(c *gocheck.C) {
	floatingIps := New(t.ApiTestHarness.ApiConnection)
	ips, _ := floatingIps.List()
	fmt.Println(floatingIps.GetById(ips[0].Id))
}

func (t *FloatingIpsTestSuite) Test_CreateAndDelete(c *gocheck.C) {
	floatingIps := New(t.ApiTestHarness.ApiConnection)
	floatingIp, _ := floatingIps.Create()
	fmt.Printf("Floating IP Created %v", floatingIp)
	floatingIps.Delete(floatingIp.Id)

}
