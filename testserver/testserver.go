package testserver

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "go-openstack-client/authhttp/authenticator"
  "go-openstack-client/authhttp/handler"
  "github.com/gorilla/mux"
)

type TestServer struct {
    index []map[string]string
    router *mux.Router
}

func (t *TestServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
    queryParameters := make(map[string]string)
    queryParameters["method"] = request.Method
    queryParameters["path"] = request.URL.Path

    retval := t.GetContent(queryParameters)

    //retval := "<html><body><h1>" + request.Method + " Successful!</h1>"

    //body, _ := ioutil.ReadAll(request.Body)
    //request.Body.Close()

    //if string(body) != "" {
    //    retval = retval + "<h2>" + string(body) + "</h2>"
    //}
    //retval = retval + "</body></html>"
    fmt.Fprintf(w, retval)
}

func (t *TestServer) Start(a authenticator.Authenticators, port string, testFilePath string) {
    t.router = mux.NewRouter()
    t.LoadTestResponses(a,testFilePath)
    //router.HandleFunc("/", handler.New(a, t).ServeHTTP)
    //http.Handle("/", router)
    http.ListenAndServe(":" + port, nil)
}

func (t *TestServer) LoadTestResponses(a authenticator.Authenticators, path string) {
    indexFile, _ := ioutil.ReadFile(path + "/index.json")
    var index []map[string]string
    json.Unmarshal(indexFile,&index)

    for _, resp := range index {
        contentFile, _ := ioutil.ReadFile(path + "/" + resp["file"])
        resp["content"] = string(contentFile)
        t.router.HandleFunc(resp["path"], handler.New(a, t).ServeHTTP)
        http.Handle(resp["path"], t.router)
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
            return respMap["content"]
        }
    }
    return ""
}
