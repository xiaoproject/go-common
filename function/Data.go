package function

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

/**
格式化数据处理
@data object
@defaultStr string
*/
func DefaultHandleStr(data interface{}, defaultStr string) string {

	if data != nil {
		return data.(string)
	} else {
		return defaultStr
	}
}


func FunMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// string to int
// @param source 字符串
func Str2Int(source string) int {

	i, e := strconv.Atoi(source)
	if e != nil{
		return 0
	}
	return i
}


