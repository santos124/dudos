package CETb

import (
	"dudos/arguments"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

type CETb struct {
	client *resty.Client
	cntOK  int
	cntKO  int
}

type task struct {
	headerIn  map[string]string
	headerOut map[string]string
	bodyIn    map[string]string
	bodyOut   map[string]string
}

func New() *CETb {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	return &CETb{client: client}
}

func (c *CETb) StartLoad(flags *arguments.Flag) {
	numWorkers := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	taskChan := make(chan task)
	resultChan := make(chan task)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.NewWorker()
		}()
	}
	for
	wg.Wait()
	logrus.Infof(":")
}

func (c *CETb) NewWorker(in, out chan task, flags arguments.Flag) {
	for taska := range in {
		c.doRequest(flags)
	}
}

func (c *CETb) doRequest(flags arguments.Flag) (*resty.Response, error) {
	if flags.Method == "POST" {
		return c.doPost(flags), nil
	}
	if flags.Method == "GET" {
		return c.doGet(flags), nil
	}
}

func (c *CETb) doPost(arguments arguments.Flag) (*resty.Response, error) {
	response, err := c.client.R().
		SetBody(arguments.Payload).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).Post(arguments.Endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *CETb) doGet(arguments arguments.Flag) (*resty.Response, error) {
	response, err := c.client.R().
		SetBody(arguments.Payload).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).Get(arguments.Endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}
