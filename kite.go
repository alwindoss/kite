package kite

const (
	rootAPIEndpoint = "https://api.kite.trade"
)

// Client provides the functionality required for the consumer to access all of the Kite Connect Endpoints
type Client interface {
	Login() string
}

type kite struct {
}

func (k kite) Login() string {
	return "ACCESS TOKEN"
}

// NewClient creates a Kite Connect Client that can be used by the consumer to consume the Kite Connect Endpoints
func NewClient(apiKey string) Client {
	var client = kite{}

	return client
}
