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

// resp:&{HttpStatusCode:200 ResponseBody:{"code":1,"msg":"success","time":"1774317615","data":"https:\/\/bingocn.wobeis.com\/cash\/#\/?orderid=2025080911465070&access_key=1525364505"} Headers:map[accesskey:1525364505 gmtrequest:1774317614 randomstr:47069f407698bd5c2c1108c43c0b5225 signature:D8536C1E3647460C5CEBE3D7C82B58F7]}
func GenDepositRequestDemo() InlandXJPayDepositReq {
	return InlandXJPayDepositReq{
		OrderId:   "2025080911465072",
		Amount:    "3600",
		Diqu:      "1",
		YxTime:    "900",
		PayerName: "李四",
	}
}
