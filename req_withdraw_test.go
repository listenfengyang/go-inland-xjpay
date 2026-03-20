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
		OrderId:    "OUT202508091146507666",
		Amount:     "1000",
		Diqu:       "1",
		PayeeName:  "张三",
		BankCardNo: "6222021234567890",
		BankName:   "ICBC",
		BankBranch: "Shenzhen",
		Remark:     "test",
	}
}
