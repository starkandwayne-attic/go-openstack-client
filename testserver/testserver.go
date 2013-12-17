package testserver

import (
  "fmt"
  "math/rand"
  "net"
  "net/http"
  "io/ioutil"
  "strconv"
  "strings"
  "encoding/json"
  "git.smf.sh/jrbudnack/go_openstack_client/authhttp/authenticator"
  "git.smf.sh/jrbudnack/go_openstack_client/authhttp/handler"
  "github.com/gorilla/mux"
)

type TestServer struct {
    index []map[string]string
    router *mux.Router
    Port string
}

func New(a authenticator.Authenticators, port string, testFilePath string) TestServer{
    ts := TestServer{}
    ts.router = mux.NewRouter()
    if port == "" {
        ts.PickRandomLocalPort()
    } else {
        ts.Port = port
    }
    ts.LoadTestResponses(a,testFilePath)
    return ts
}

func (t *TestServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
    queryParameters := make(map[string]string)
    queryParameters["method"] = request.Method
    queryParameters["path"] = request.URL.Path

    body, _ := ioutil.ReadAll(request.Body)
    request.Body.Close()

    queryParameters["body"] = string(body)

    retval := t.GetContent(queryParameters)

    fmt.Fprintf(w, retval)
}

func (t *TestServer) Start() {
    http.ListenAndServe(":" + t.Port, nil)
}

func (t *TestServer) PickRandomLocalPort() {
    portIsOpen := false
    portToTry := -1

    for portIsOpen == false {
        portToTry = rand.Intn(55535) + 1000
        conn, err := net.Dial("tcp", "127.0.0.1:" + strconv.Itoa(portToTry))
        if err != nil {
            portIsOpen = true
        } else {
            conn.Close()
        }
    }
    t.Port = strconv.Itoa(portToTry)
}

func (t *TestServer) LoadTestResponses(a authenticator.Authenticators, path string) {
    indexFile, _ := ioutil.ReadFile(path + "/index.json")
    var index []map[string]string
    registeredEndpoints := make(map[string]string)
    json.Unmarshal(indexFile,&index)

    for _, resp := range index {
        contentFile, _ := ioutil.ReadFile(path + "/" + resp["file"])
        resp["content"] = strings.Replace(strings.TrimRight(string(contentFile),"\n"),"PORT_NUM",t.Port,-1)
        _, isRegistered := registeredEndpoints[resp["path"]]
        if isRegistered == false {
            t.router.HandleFunc(resp["path"], handler.New(a, t).ServeHTTP)
            http.Handle(resp["path"], t.router)
            registeredEndpoints[resp["path"]] = resp["path"]
        }
    }
    t.index = index
}

func (t *TestServer) GetContent(queryParameters map[string]string) string {
    //Query Parameters:
    //  method (GET,PUT,...)
    //  path
    for _, respMap := range t.index {
        meetsAll := true
        //respMap := resp.(map[string]string)
        for k,v := range queryParameters {
            _, keyExists := respMap[k]
            if keyExists && respMap[k] != v {
                meetsAll = false
                break
            }
        }
        if meetsAll {
            _, hasBody := queryParameters["body"]
            if hasBody {
                return strings.Replace(string(respMap["content"]),"POSTED_BODY",queryParameters["body"],-1)
            } else {
                return respMap["content"]
            }
        }
    }
    return ""
}
