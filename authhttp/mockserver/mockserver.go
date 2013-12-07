package mockserver

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "go-openstack-client/authhttp/authenticator"
  "go-openstack-client/authhttp/handler"
  "github.com/gorilla/mux"
)

type Server struct {
}

func (t Server) ServeHTTP(w http.ResponseWriter, request *http.Request) {
  retval := "<html><body><h1>" + request.Method + " Successful!</h1>"

  body, _ := ioutil.ReadAll(request.Body)
  request.Body.Close()

  if string(body) != "" {
    retval = retval + "<h2>" + string(body) + "</h2>"
  }
  retval = retval + "</body></html>"
  fmt.Fprintf(w, retval)
}

func (t Server) Start(a authenticator.Authenticators, port string) {
  router := mux.NewRouter()
  router.HandleFunc("/", handler.New(a, t).ServeHTTP)
  http.Handle("/", router)
  http.ListenAndServe(":" + port, nil)
}
