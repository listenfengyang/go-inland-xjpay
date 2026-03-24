package go_inland_xjpay

type InlandXJPayInitParams struct {
	AccessKey       string `json:"accessKey" mapstructure:"accessKey" config:"accessKey" yaml:"accessKey"`                         // accesskey 请求头鉴权账号
	SecretKey       string `json:"secretKey" mapstructure:"secretKey" config:"secretKey" yaml:"secretKey"`                         // key 请求头签名密钥
	DepositUrl      string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl" yaml:"depositUrl"`                     // 入金接口地址
	WithdrawUrl     string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl" yaml:"withdrawUrl"`                 // 出金接口地址
	DepositBackUrl  string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl" yaml:"depositBackUrl"`     // 入金回调地址
	WithdrawBackUrl string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl" yaml:"withdrawBackUrl"` // 出金回调地址
}

type InlandXJPayDepositReq struct {
	OrderId   string            `json:"orderid" form:"orderid" mapstructure:"orderid"`       // 商户订单号
	Amount    string            `json:"amount" form:"amount" mapstructure:"amount"`          // 金额
	Diqu      string            `json:"diqu" form:"diqu" mapstructure:"diqu"`                // 区域: 1大陆 2江城 3海外
	BackUrl   string            `json:"backurl" form:"backurl" mapstructure:"backurl"`       // 入金回调地址
	YxTime    string            `json:"yx_time" form:"yx_time" mapstructure:"yx_time"`       // 订单有效时间(秒)
	PayerName string            `json:"payername" form:"payername" mapstructure:"payername"` // 付款人姓名
	Extra     map[string]string `json:"extra" mapstructure:"extra"`                          // 扩展参数
}

type InlandXJPayWithdrawReq struct {
	RealName       string            `json:"realName" form:"realName" mapstructure:"realName"`                   // 收款人姓名
	CardNumber     string            `json:"cardNumber" form:"cardNumber" mapstructure:"cardNumber"`             // 收款卡号
	BankName       string            `json:"bankName" form:"bankName" mapstructure:"bankName"`                   // 银行名称
	BankBranchName string            `json:"bankBranchName" form:"bankBranchName" mapstructure:"bankBranchName"` // 开户支行
	PayType        string            `json:"pay_type" form:"pay_type" mapstructure:"pay_type"`                   // 支付类型
	PayAccount     string            `json:"pay_account" form:"pay_account" mapstructure:"pay_account"`          // 付款账户类型
	Usdt           string            `json:"usdt" form:"usdt" mapstructure:"usdt"`                               // 出金USDT数量
	OrderId        string            `json:"orderid" form:"orderid" mapstructure:"orderid"`                      // 商户订单号
	WebhookUrl     string            `json:"webhookUrl" form:"webhookUrl" mapstructure:"webhookUrl"`             // 出金回调地址
	Extra          map[string]string `json:"extra" mapstructure:"extra"`                                         // 扩展参数
}

type ResponseBody struct {
	Code int64  `json:"code" mapstructure:"code"`
	Msg  string `json:"msg" mapstructure:"msg"`
	Time string `json:"time" mapstructure:"time"`
	Data string `json:"data" mapstructure:"data"`
}

type InlandXJPayCommonRsp struct {
	HttpStatusCode int               `json:"httpStatusCode" mapstructure:"httpStatusCode"` // HTTP状态码
	ResponseBody   string            `json:"responseBody" mapstructure:"responseBody"`     // 原始响应体
	Headers        map[string]string `json:"headers" mapstructure:"headers"`               // 请求头(含签名参数)
	BodyData       ResponseBody      `json:"bodyData" mapstructure:"bodyData"`
}

type InlandXJPayCallbackReq struct {
	Type       string            `json:"type" form:"type" mapstructure:"type"`                   // 回调类型: recharge/withdraw
	OrderId    string            `json:"orderid" form:"orderid" mapstructure:"orderid"`          // 商户订单号
	PayStatus  int64             `json:"pay_status" form:"pay_status" mapstructure:"pay_status"` // 支付状态
	Sign       string            `json:"sign" form:"sign" mapstructure:"sign"`                   // 回调签名
	Extra      map[string]string `json:"extra" mapstructure:"extra"`                             // 扩展参数
	Accesskey  string            `json:"accesskey" mapstructure:"accesskey"`                     // header中的值
	Gmtrequest string            `json:"gmtrequest" mapstructure:"gmtrequest"`                   // header中的值
	Randomstr  string            `json:"randomstr" mapstructure:"randomstr"`                     // header中的值
}
