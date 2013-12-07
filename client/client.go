package client

type Client struct {
    
}

func (c Client) Authenticate() {
    
}

func (c Client) UnAuthenticate() {
    
}

func (c Client) GetServiceList() {
    
}

func (c Client) Get() {
    
}

func (c Client) Post() {

}

func (c Client) Put() {
    
}

func (c Client) Delete() {
    
}

func New () Client {
    newClient := Client{}
    newClient.Connect()
    return newClient
}
