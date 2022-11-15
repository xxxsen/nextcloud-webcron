package crontask

type config struct {
	url        string
	cronexp    string
	runOnStart bool
}

type Option func(c *config)

func WithURL(url string) Option {
	return func(c *config) {
		c.url = url
	}
}

func WithCronExpression(exp string) Option {
	return func(c *config) {
		c.cronexp = exp
	}
}

func WithRunOnStart(v bool) Option {
	return func(c *config) {
		c.runOnStart = v
	}
}
