package CETb

import (
	"dudos/arguments"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
	"time"
)

type CETb struct {
	client *resty.Client
	cntOK  int
	cntKO  int
}

type task struct {
	headerIn   map[string]string
	headerOut  map[string]string
	bodyIn     map[string]string
	bodyOut    map[string]string
	statusCode int
}

func New() *CETb {
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	return &CETb{client: client}
}

func (c *CETb) StartLoad(flags *arguments.Flag) {
	numWorkers := runtime.NumCPU() * 10
	wg := &sync.WaitGroup{}
	taskChan := make(chan task, numWorkers)
	resultChan := make(chan task, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.NewWorker(taskChan, resultChan, *flags)
		}()
	}

	delay := time.Second * 5
	timer := time.NewTimer(delay)

	go func() {
		for {
			select {
			case <-timer.C:
				close(taskChan)
				close(resultChan)

				break
			default:
				taskChan <- task{
					headerIn:   map[string]string{},
					headerOut:  map[string]string{},
					bodyIn:     map[string]string{},
					bodyOut:    map[string]string{},
					statusCode: 0,
				}
			}
		}
	}()

	arrResult := []task{}

	timerResult := time.NewTimer(delay)
	go func() {
		for {
			select {
			case <-timerResult.C:
				close(taskChan)
				close(resultChan)

				break
			case result := <-resultChan:
				arrResult = append(arrResult, result)
			default:
			}
		}
	}()
	wg.Wait()
	rps := len(arrResult) / int(delay.Seconds())
	logrus.Infof("rps:%v", rps)
}

func (c *CETb) NewWorker(in, out chan task, flags arguments.Flag) {
	for taska := range in {
		resp, err := c.doRequest(flags)
		if err != nil {
			logrus.Error(err)
			continue
		}

		taska2 := task{
			headerIn:   taska.headerIn,
			headerOut:  taska.headerOut,
			bodyIn:     nil,
			bodyOut:    nil,
			statusCode: resp.StatusCode(),
		}

		out <- taska2
	}
}

func (c *CETb) doRequest(flags arguments.Flag) (*resty.Response, error) {
	if flags.Method == "POST" {
		return c.doPost(flags)
	}
	if flags.Method == "GET" {
		return c.doGet(flags)
	}
	return nil, errors.New("unknow method")
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
