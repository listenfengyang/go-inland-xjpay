package go_inland_xjpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-inland-xjpay/utils"
)

func (cli *Client) DepositCallback(req InlandXJPayCallbackReq, processor func(InlandXJPayCallbackReq) error) error {
	params := map[string]string{
		"type":       req.Type,
		"orderid":    req.OrderId,
		"pay_status": req.PayStatus,
		"sign":       req.Sign,
	}
	for k, v := range req.Extra {
		params[k] = v
	}
	if !utils.VerifySignWithSortedParams(params, cli.Params.SecretKey) {
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("inland-xjpay deposit callback verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}
	return processor(req)
}

func (cli *Client) WithdrawCallback(req InlandXJPayCallbackReq, processor func(InlandXJPayCallbackReq) error) error {
	params := map[string]string{
		"type":       req.Type,
		"orderid":    req.OrderId,
		"pay_status": req.PayStatus,
		"sign":       req.Sign,
	}
	for k, v := range req.Extra {
		params[k] = v
	}
	if !utils.VerifySignWithSortedParams(params, cli.Params.SecretKey) {
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("inland-xjpay withdraw callback verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}
	return processor(req)
}
