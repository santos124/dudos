package arguments

import "flag"

type Flag struct {
	Endpoint    string
	RPS         int
	PayloadType string
	Payload     string
	Method      string
	TimeLoad    int
}

func New() *Flag {
	var rpsFlag = flag.Int("rps", 1000, "rps limit")
	var endpointFlag = flag.String("endpoint", "http://localhost:8040", "endpoint")
	var payloadFlag = flag.String("payload", "1kb", "payload")
	var methodFlag = flag.String("method", "GET", "method")
	var timeLoadFlag = flag.Int("load", 1, "load time")
	flag.Parse()

	return &Flag{
		Endpoint:    *endpointFlag,
		RPS:         *rpsFlag,
		PayloadType: *payloadFlag,
		Method:      *methodFlag,
		TimeLoad:    *timeLoadFlag,
	}
}
