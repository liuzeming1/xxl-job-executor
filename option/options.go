package option

import "time"

const (
	defaultAdminAddr = "http://localhost:8080/xxl-job-admin/"
	defaultAppName   = "go-executor"
	defaultPort      = 8081
	defaultTimeout   = 5 * time.Second
	defaultBeatTime  = 20 * time.Second
)

type Option func(*ClientOptions)

type ClientOptions struct {
	IsEnable    bool          `yaml:"isEnable"`
	ServerAddrs []string      `yaml:"serverAddrs"`
	AccessToken string        `yaml:"accessToken"`
	AppName     string        `yaml:"appName"`
	ClientPort  int           `yaml:"clientPort"`
	Timeout     time.Duration `yaml:"timeout"`
	BeatTime    time.Duration `yaml:"beatTime"`
	LogLevel    int           `yaml:"logLevel"`
}

func NewClientOptions(opts ...Option) ClientOptions {
	clientOptions := ClientOptions{}
	for _, opt := range opts {
		opt(&clientOptions)
	}
	return clientOptions
}

func WithDefaultOptions() Option {
	return func(o *ClientOptions) {
		if len(o.ServerAddrs) == 0 {
			o.ServerAddrs = []string{defaultAdminAddr}
		}
		if o.AppName == "" {
			o.AppName = defaultAppName
		}
		if o.ClientPort == 0 {
			o.ClientPort = defaultPort
		}
		if o.Timeout == 0 {
			o.Timeout = defaultTimeout
		}
		if o.BeatTime == 0 {
			o.BeatTime = defaultBeatTime
		}
		if o.LogLevel == 0 {
			o.LogLevel = 1 // default log level
		}
	}
}

func WithServerAddrs(addrs ...string) Option {
	return func(o *ClientOptions) {
		o.ServerAddrs = addrs
	}
}

func WithAccessToken(token string) Option {
	return func(o *ClientOptions) {
		o.AccessToken = token
	}
}

func WithAppName(name string) Option {
	return func(o *ClientOptions) {
		o.AppName = name
	}
}

func WithClientPort(port int) Option {
	return func(o *ClientOptions) {
		o.ClientPort = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *ClientOptions) {
		o.Timeout = timeout
	}
}

func WithBeatTime(beatTime time.Duration) Option {
	return func(o *ClientOptions) {
		o.BeatTime = beatTime
	}
}

func WithLogLevel(level int) Option {
	return func(o *ClientOptions) {
		o.LogLevel = level
	}
}
