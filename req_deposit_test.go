package go_inland_xjpay

import (
	"fmt"
	"testing"
)

type VLog struct{}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	cli := NewClient(vLog, &InlandXJPayInitParams{
		AccessKey:      ACCESS_KEY,
		SecretKey:      SECRET_KEY,
		DepositUrl:     DEPOSIT_URL,
		DepositBackUrl: DEPOSIT_BACK_URL,
	})

	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() InlandXJPayDepositReq {
	return InlandXJPayDepositReq{
		OrderId:   "202508091146507665",
		Amount:    "3600",
		Diqu:      "1",
		YxTime:    "900",
		PayerName: "李四",
	}
}
