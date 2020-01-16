package function

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
