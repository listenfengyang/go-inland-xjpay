package go_inland_xjpay

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-inland-xjpay/utils"
)

func (cli *Client) Withdraw(req InlandXJPayWithdrawReq) (*InlandXJPayCommonRsp, error) {
	if cli.Params.WithdrawUrl == "" {
		return nil, fmt.Errorf("withdrawUrl is empty")
	}
	if cli.Params.AccessKey == "" {
		return nil, fmt.Errorf("accessKey is empty")
	}
	if cli.Params.SecretKey == "" {
		return nil, fmt.Errorf("secretKey is empty")
	}
	if req.OrderId == "" {
		return nil, fmt.Errorf("orderid is empty")
	}
	if req.RealName == "" {
		return nil, fmt.Errorf("realName is empty")
	}
	if req.CardNumber == "" {
		return nil, fmt.Errorf("cardNumber is empty")
	}
	if req.BankName == "" {
		return nil, fmt.Errorf("bankName is empty")
	}
	if req.BankBranchName == "" {
		return nil, fmt.Errorf("bankBranchName is empty")
	}
	if req.PayType == "" {
		return nil, fmt.Errorf("pay_type is empty")
	}
	if req.PayAccount == "" {
		return nil, fmt.Errorf("pay_account is empty")
	}
	if req.Usdt == "" {
		return nil, fmt.Errorf("usdt is empty")
	}

	if req.WebhookUrl == "" {
		req.WebhookUrl = cli.Params.WithdrawBackUrl
	}
	if req.WebhookUrl == "" {
		return nil, fmt.Errorf("webhookUrl is empty")
	}
	params := mapFromWithdrawReq(req)
	headers := cli.buildAuthHeaders()

	resp2, err := cli.ryClient.R().
		SetHeaders(headers).
		SetFormData(params).
		SetDebug(cli.debugMode).
		Post(cli.Params.WithdrawUrl)
	if err != nil {
		return nil, err
	}

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#inland-xjpay#withdraw->%s", string(restLog))

	if resp2.StatusCode() >= 400 {
		return nil, fmt.Errorf("status code: %d, body:%s", resp2.StatusCode(), resp2.String())
	}

	return &InlandXJPayCommonRsp{
		HttpStatusCode: resp2.StatusCode(),
		ResponseBody:   resp2.String(),
		Headers:        headers,
	}, nil
}
