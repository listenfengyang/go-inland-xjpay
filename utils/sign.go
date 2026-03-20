package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
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

func VerifySignWithSortedParams(params map[string]string, secretKey string) bool {
	sign := ""
	for k, v := range params {
		if strings.EqualFold(k, "sign") || strings.EqualFold(k, "signature") {
			sign = v
			break
		}
	}
	if sign == "" {
		return false
	}
	expected := Md5Hex(BuildSortedSignSource(params, secretKey))
	return strings.EqualFold(sign, expected)
}
