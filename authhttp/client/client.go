package client

import (
    "bytes"
    _"fmt"
    "io/ioutil"
    "net/http"
    "go_openstack_client/authhttp"
    "go_openstack_client/authhttp/credentials"
)

type Client struct {
    credentials credentials.Credentials
    url string
}

func New(c credentials.Credentials, url string) Client {
    return Client{c, url}
}

func (c Client) GetAndParseBody(uri string) ([]byte, error) {
    res, err := c.Get(uri)

    if err != nil {
        return nil, err
    }
    return c.ParseBody(res)
}

func (c Client) PostAndParseBody(uri string, body string) ([]byte, error) {
    res, err := c.Post(uri, body)

    if err != nil {
        return nil, err
    }
    return c.ParseBody(res)
}

func (c Client) ParseBody(resp *http.Response) ([]byte, error) {
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    return body, err
}

func (c Client) Get(uri string) (*http.Response, error) {
    client := &http.Client{}

    fullURL := c.url + uri

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Date", authentication.GetCurrentGMTTime())
    req.Header.Add("Content-Type", GetContentType())
    req.Header.Add("rack.input", "")

    req = c.credentials.SignRequest(req)
    client = c.credentials.SignClient(client)

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (c Client) Post(uri string, body string) (*http.Response, error) {
    client := &http.Client{}
    b := bytes.NewBufferString(body)

    fullUrl := c.url + uri

    req, err := http.NewRequest("POST", fullUrl, b)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Date", authentication.GetCurrentGMTTime())
    req.Header.Add("Content-Type", GetContentType())
    req.Header.Add("Content-Length", "1")
    req.Header.Add("Connection", "keep-alive")

    req = c.credentials.SignRequest(req)
    client = c.credentials.SignClient(client)

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

func (c Client) Delete (uri string) (*http.Response, error) {
    client := &http.Client{}

    fullURL := c.url + uri

    req, err := http.NewRequest("DELETE", fullURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Date", authentication.GetCurrentGMTTime())
    req.Header.Add("Content-Type", GetContentType())
    req.Header.Add("rack.input", "")

    req = c.credentials.SignRequest(req)
    client = c.credentials.SignClient(client)

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

//Returns Content Type
func GetContentType() string {
    return "application/json"
    /* return "application/x-www-form-urlencoded" */
}
