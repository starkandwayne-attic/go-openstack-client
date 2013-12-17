package apitestharness

import (
    _"fmt"
    "os"
    "go_openstack_client/apiconnection"
    "go_openstack_client/authhttp/authenticator"
    "go_openstack_client/authhttp/none"
    "go_openstack_client/testserver"
)

type ApiTestHarness struct{
    TestServer testserver.TestServer
    ApiConnection apiconnection.ApiConnection
    ApiServerType string
    Host string
    Port string
    Url string
    Username string
    Password string
    Tenant string
}

func New(apiservertype string, https bool) ApiTestHarness {
    a := ApiTestHarness{}
    a.ApiServerType = apiservertype
    a.ParseEnvironmentVariables()
    if a.Host == ""  {
        a.RunTestApiServer(apiservertype)
        a.Host = "127.0.0.1"
        a.Port = a.TestServer.Port
    }
    a.Url = "http://" + a.Host + ":" + a.Port
    a.ApiConnection = apiconnection.New(a.Url,
                                        apiservertype,
                                        a.Username,
                                        a.Password,
                                        a.Tenant)
    return a
}

func (t *ApiTestHarness) RunTestApiServer(apiservertype string) {
    authenticators := authenticator.Authenticators{}
    // We use Authentication = none because we aren't writing the 
    // authentication server.  We are only writing the client.
    // Therefore, we only need the API server to simulate responses,
    // not do actual authentication.
    authenticators.Add(none.Authenticator{},true)

    workingDir, _ := os.Getwd()
    t.TestServer = testserver.New(authenticators, "", workingDir + "/" + apiservertype + "_testfiles")
    go t.TestServer.Start()
}

func (t *ApiTestHarness) ParseEnvironmentVariables() {
    //If environment variables pointing to an existing  are set,
    //parse them into internal variables so we can point our ApiConnection
    //to the proper environment.
    t.Host = os.Getenv("GO_OPENSTACK_CLIENT_KEYSTONE_HOST")
    t.Port = os.Getenv("GO_OPENSTACK_CLIENT_KEYSTONE_PORT")
    t.Username = os.Getenv("GO_OPENSTACK_CLIENT_USERNAME")
    t.Password = os.Getenv("GO_OPENSTACK_CLIENT_PASSWORD")
    t.Tenant = os.Getenv("GO_OPENSTACK_CLIENT_TENANT")
}
