package go_inland_xjpay

type InlandXJPayInitParams struct {
	AccessKey       string `json:"accessKey" mapstructure:"accessKey" config:"accessKey" yaml:"accessKey"`
	SecretKey       string `json:"secretKey" mapstructure:"secretKey" config:"secretKey" yaml:"secretKey"`
	DepositUrl      string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl" yaml:"depositUrl"`
	WithdrawUrl     string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl" yaml:"withdrawUrl"`
	DepositBackUrl  string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl" yaml:"depositBackUrl"`
	WithdrawBackUrl string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl" yaml:"withdrawBackUrl"`
}

type InlandXJPayDepositReq struct {
	OrderId   string            `json:"orderid" form:"orderid" mapstructure:"orderid"`
	Amount    string            `json:"amount" form:"amount" mapstructure:"amount"`
	Diqu      string            `json:"diqu" form:"diqu" mapstructure:"diqu"`
	BackUrl   string            `json:"backurl" form:"backurl" mapstructure:"backurl"`
	YxTime    string            `json:"yx_time" form:"yx_time" mapstructure:"yx_time"`
	PayerName string            `json:"payername" form:"payername" mapstructure:"payername"`
	Extra     map[string]string `json:"extra" mapstructure:"extra"`
}

type InlandXJPayWithdrawReq struct {
	OrderId    string            `json:"orderid" form:"orderid" mapstructure:"orderid"`
	Amount     string            `json:"amount" form:"amount" mapstructure:"amount"`
	Diqu       string            `json:"diqu" form:"diqu" mapstructure:"diqu"`
	WebhookUrl string            `json:"webhookurl" form:"webhookurl" mapstructure:"webhookurl"`
	PayeeName  string            `json:"payername" form:"payername" mapstructure:"payername"`
	BankCardNo string            `json:"bankcard" form:"bankcard" mapstructure:"bankcard"`
	BankName   string            `json:"bankname" form:"bankname" mapstructure:"bankname"`
	BankBranch string            `json:"bankbranch" form:"bankbranch" mapstructure:"bankbranch"`
	Remark     string            `json:"remark" form:"remark" mapstructure:"remark"`
	Extra      map[string]string `json:"extra" mapstructure:"extra"`
}

type InlandXJPayCommonRsp struct {
	HttpStatusCode int               `json:"httpStatusCode" mapstructure:"httpStatusCode"`
	ResponseBody   string            `json:"responseBody" mapstructure:"responseBody"`
	Headers        map[string]string `json:"headers" mapstructure:"headers"`
}

type InlandXJPayCallbackReq struct {
	OrderId string            `json:"orderid" form:"orderid" mapstructure:"orderid"`
	Status  string            `json:"status" form:"status" mapstructure:"status"`
	Amount  string            `json:"amount" form:"amount" mapstructure:"amount"`
	Sign    string            `json:"sign" form:"sign" mapstructure:"sign"`
	Extra   map[string]string `json:"extra" mapstructure:"extra"`
}
