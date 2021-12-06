package pinata

const (
	DefaultNode = "https://api.pinata.cloud"
	ApiPinFile = "/pinning/pinFileToIPFS"
)

type Client struct {
	Node string
	JWT string
	ApiKey string
	SecretApiKey string
}

func New(node string, jwt string, apiKey string, apiSecret string) *Client {
	if jwt != "" || (apiKey != "" && apiSecret != "") {
		return &Client{
			Node: node,
			JWT: jwt,
			ApiKey: apiKey,
			SecretApiKey: apiSecret,
		}
	}

	return nil
}