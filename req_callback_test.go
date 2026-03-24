package go_inland_xjpay

import (
	"testing"
)

func TestCallback(t *testing.T) {
	vLog := VLog{}
	cli := NewClient(vLog, &InlandXJPayInitParams{
		AccessKey:   ACCESS_KEY,
		SecretKey:   SECRET_KEY,
		DepositUrl:  DEPOSIT_URL,
		WithdrawUrl: WITHDRAW_URL,
	})

	sign := "A2C40DEC9D33C26DDCDF9E5F47806F79"
	err := cli.DepositCallback(GenCallbackRequestDemo(), sign, func(InlandXJPayCallbackReq) error {
		return nil
	})
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
}

// {\"type\":\"recharge\",\"orderid\":\"202508091146507663\",\"pay_status\":3}
func GenCallbackRequestDemo() InlandXJPayCallbackReq {
	return InlandXJPayCallbackReq{
		OrderId:    "202603241421520700",
		PayStatus:  3,
		Type:       "recharge",
		Accesskey:  "1525364505",
		Gmtrequest: "1774351417",
		Randomstr:  "CdwyauBhzSiZV7gYmc8NYBUQgHObWbIL",
	}
}
