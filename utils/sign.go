package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Md5Hex(input string) string {
	sum := md5.Sum([]byte(input))
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}

func BuildSortedSignSource(params map[string]string, secretKey string) string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v == "" {
			continue
		}
		if strings.EqualFold(k, "sign") || strings.EqualFold(k, "signature") {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys)+1)
	for _, k := range keys {
		parts = append(parts, k+"="+params[k])
	}
	parts = append(parts, "key="+secretKey)
	return strings.Join(parts, "&")
}

func VerifySignWithSortedParams(params map[string]string, sign, secretKey string) (bool, error) {
	// 1. 依照 ASCII 顺序由小到大做排序
	//  key1=value1&key2=value2...的方式组出字串，最后再加上&secret_key={密钥}
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" && value != "" { // && value != ""
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}
	signStr := sb.String()
	signStr += fmt.Sprintf("key=%s", secretKey)
	signStr, err := url.QueryUnescape(signStr)
	if err != nil {
		fmt.Println("QueryUnescape error:", err)
		return false, err
	}

	fmt.Printf("[rawString]%s\n", signStr)

	// 第2步骤产生签名字串做 md5 加签得到sign
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("MD5签名str: %s\n\n", strings.ToUpper(signResult))
	return strings.ToUpper(signResult) == sign, nil
}
