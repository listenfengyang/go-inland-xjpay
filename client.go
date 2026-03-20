package go_inland_xjpay

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-inland-xjpay/utils"
)

type Client struct {
	Params *InlandXJPayInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *InlandXJPayInitParams) *Client {
	return &Client{
		Params:   params,
		ryClient: resty.New(),
		logger:   logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}

