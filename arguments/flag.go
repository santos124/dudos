package arguments

import "flag"

type Flag struct {
	Endpoint string
	RPS      int
	Payload  string
	Method   string
}

func New() *Flag {
	var rpsFlag = flag.Int("rps", 1000, "rps limit")
	var endpointFlag = flag.String("endpoint", "", "endpoint")
	var payloadFlag = flag.String("payload", "", "payload")
	flag.Parse()

	return &Flag{
		Endpoint: *endpointFlag,
		RPS:      *rpsFlag,
	}
}
