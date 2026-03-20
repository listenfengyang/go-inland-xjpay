package go_inland_xjpay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	cli := NewClient(vLog, &InlandXJPayInitParams{
		AccessKey:       ACCESS_KEY,
		SecretKey:       SECRET_KEY,
		WithdrawUrl:     WITHDRAW_URL,
		WithdrawBackUrl: WITHDRAW_BACK_URL,
	})

	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() InlandXJPayWithdrawReq {
	return InlandXJPayWithdrawReq{
		RealName:       "王五",
		CardNumber:     "665111111111",
		BankName:       "建设银行",
		BankBranchName: "建设银行长江路支行",
		PayType:        "bank",
		PayAccount:     "wxpay",
		Usdt:           "200",
		OrderId:        "123456789",
	}
}
