package sqwiggle

type Request struct {
	ApiKey string
	Method string
	Args   map[string]string
}

