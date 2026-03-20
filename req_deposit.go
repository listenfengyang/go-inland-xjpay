package go_inland_xjpay

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-inland-xjpay/utils"
)

func (cli *Client) Deposit(req InlandXJPayDepositReq) (*InlandXJPayCommonRsp, error) {

	req.BackUrl = cli.Params.DepositBackUrl
	params := mapFromDepositReq(req)
	headers := cli.buildAuthHeaders()

	resp2, err := cli.ryClient.R().
		SetHeaders(headers).
		SetBody(params).
		SetFormData(params).
		SetDebug(cli.debugMode).
		Post(cli.Params.DepositUrl)
	if err != nil {
		return nil, err
	}

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#inland-xjpay#deposit->%s", string(restLog))

	if resp2.StatusCode() >= 400 {
		return nil, fmt.Errorf("status code: %d, body:%s", resp2.StatusCode(), resp2.String())
	}

	return &InlandXJPayCommonRsp{
		HttpStatusCode: resp2.StatusCode(),
		ResponseBody:   resp2.String(),
		Headers:        headers,
	}, nil
}
