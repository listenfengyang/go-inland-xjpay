package go_inland_xjpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-inland-xjpay/utils"
)

// accesskey=1525364505&gmtrequest=1774351417&randomstr=CdwyauBhzSiZV7gYmc8NYBUQgHObWbIL&key=c87047c344517d5b26d5de992b93ce5b
func (cli *Client) DepositCallback(req InlandXJPayCallbackReq, sign string, processor func(InlandXJPayCallbackReq) error) error {
	params := map[string]string{
		"accesskey":  req.Accesskey,
		"gmtrequest": req.Gmtrequest,
		"randomstr":  req.Randomstr,
	}

	res, err := utils.VerifySignWithSortedParams(params, sign, cli.Params.SecretKey)
	if !res || err != nil {
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("inland-xjpay deposit callback verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}
	return processor(req)
}

func (cli *Client) WithdrawCallback(req InlandXJPayCallbackReq, sign string, processor func(InlandXJPayCallbackReq) error) error {
	params := map[string]string{
		"accesskey":  req.Accesskey,
		"gmtrequest": req.Gmtrequest,
		"randomstr":  req.Randomstr,
	}

	res, err := utils.VerifySignWithSortedParams(params, sign, cli.Params.SecretKey)
	if !res || err != nil {
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("inland-xjpay withdraw callback verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}
	return processor(req)
}
