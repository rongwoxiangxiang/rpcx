package util

import (
	. "bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"math/rand"
	"reflect"
	"rpc/common"
	"strconv"
	"strings"
)

func IsEmpty(str string) bool {
	if str == "" {
		return true
	} else if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}

func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		i = 0
	}
	return i
}

func BoolenToInt8(bool2 bool) int8 {
	if bool2 {
		return common.YES_VALUE
	}
	return common.NO_VALUE
}

func Int8ToBoolen(i int8) bool {
	if i == common.YES_VALUE {
		return true
	} else {
		return false
	}
}

func HmacSHA1Base64Encrypt(encryptText, seed string) string {
	mac := hmac.New(sha1.New, []byte(seed))
	mac.Write([]byte(encryptText))
	bytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(bytes)
}

func StringJoin(strings ...string) string {
	var buffer Buffer
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func Get6RandomNumber() string {
	return strconv.FormatInt(rand.Int63n(899999)+100000, 10)
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

func GetFromMap(mp map[string]interface{}, key string) interface{} {
	if val, ok := mp[key]; ok {
		return val
	}
	return nil
}
