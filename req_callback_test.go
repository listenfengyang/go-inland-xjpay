package go_inland_xjpay

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/listenfengyang/go-inland-xjpay/utils"
)

func TestCallback(t *testing.T) {
	vLog := VLog{}
	cli := NewClient(vLog, &InlandXJPayInitParams{
		AccessKey:   ACCESS_KEY,
		SecretKey:   SECRET_KEY,
		DepositUrl:  DEPOSIT_URL,
		WithdrawUrl: WITHDRAW_URL,
	})

	req := GenCallbackRequestDemo()
	var backReq InlandXJPayCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	err = cli.DepositCallback(backReq, func(InlandXJPayCallbackReq) error {
		return nil
	})
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

// {\"type\":\"recharge\",\"orderid\":\"202508091146507663\",\"pay_status\":3}
func GenCallbackRequestDemo() string {
	params := map[string]string{
		"orderid":    "202508091146507666",
		"pay_status": "3",
		"type":       "recharge",
	}
	signSource := utils.BuildSortedSignSource(params, SECRET_KEY)
	sign := utils.Md5Hex(signSource)
	return fmt.Sprintf(`{"orderid":"202508091146507666","pay_status":"3","type":"recharge","sign":"%s"}`, sign)
}
