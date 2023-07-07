package ojclient

const (
	ServerRegisterPath = "/server/register"
)

// Config of client
type Config struct {
	ServerAddr string
}

func NewConfig() *Config {
	return &Config{}
}

// Client of openjob-server
type Client struct {
	cfg *Config
}

// New client instance
func New() *Client {
	return &Client{}
}

// BuildAPI of request server
func (c *Client) BuildAPI(path string) string {
	return c.cfg.ServerAddr + path
}

func (c *Client) Start(cfg *Config) error {
	c.cfg = cfg

	// register
	err := c.Register()
	if err != nil {
		return err
	}

	// fetch topic
	err = c.FetchTopic()
	if err != nil {
		return err
	}

	// start task queue

	return nil
}

func Start(cfg *Config) error {
	return New().Start(cfg)
}
