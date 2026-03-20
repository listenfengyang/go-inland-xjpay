package go_inland_xjpay

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/listenfengyang/go-inland-xjpay/utils"
)

func (cli *Client) buildAuthHeaders() map[string]string {
	gmtRequest := strconv.FormatInt(time.Now().Unix(), 10)
	randomStr := randomHex32()
	signature := generateHeaderSignature(cli.Params.AccessKey, gmtRequest, randomStr, cli.Params.SecretKey)
	return map[string]string{
		"accesskey":  cli.Params.AccessKey,
		"gmtrequest": gmtRequest,
		"randomstr":  randomStr,
		"signature":  signature,
	}
}

func randomHex32() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return hex.EncodeToString(buf)
}

func generateHeaderSignature(accessKey, gmtRequest, randomStr, secretKey string) string {
	signSource := fmt.Sprintf("accesskey=%s&gmtrequest=%s&randomstr=%s&key=%s", accessKey, gmtRequest, randomStr, secretKey)
	return utils.Md5Hex(signSource)
}

func buildBodyCanonical(body map[string]string) string {
	keys := make([]string, 0, len(body))
	for k, v := range body {
		if v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sortStrings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, body[k]))
	}
	return strings.Join(parts, "&")
}

func sortStrings(values []string) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[i] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func mapFromDepositReq(req InlandXJPayDepositReq) map[string]string {
	params := map[string]string{
		"orderid":   req.OrderId,
		"amount":    req.Amount,
		"diqu":      req.Diqu,
		"backurl":   req.BackUrl,
		"yx_time":   req.YxTime,
		"payername": req.PayerName,
	}
	for k, v := range req.Extra {
		if v != "" {
			params[k] = v
		}
	}
	return params
}

func mapFromWithdrawReq(req InlandXJPayWithdrawReq) map[string]string {
	params := map[string]string{
		"orderid":    req.OrderId,
		"amount":     req.Amount,
		"diqu":       req.Diqu,
		"webhookurl": req.WebhookUrl,
		"payername":  req.PayeeName,
		"bankcard":   req.BankCardNo,
		"bankname":   req.BankName,
		"bankbranch": req.BankBranch,
		"remark":     req.Remark,
	}
	for k, v := range req.Extra {
		if v != "" {
			params[k] = v
		}
	}
	return params
}
