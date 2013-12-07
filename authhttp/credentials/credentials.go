package credentials

import (
    _"fmt"
    "net/http"
)

type Credentials interface {
    SignRequest(request *http.Request) *http.Request
    SignClient(client *http.Client) *http.Client
    Name() string
}
