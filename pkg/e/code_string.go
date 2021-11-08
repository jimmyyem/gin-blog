// Code generated by "stringer -type ErrCode -linecomment -output code_string.go"; DO NOT EDIT.

package e

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUCCESS-0]
	_ = x[ERROR-1]
	_ = x[INVALID_PARAMS-2]
	_ = x[ERROR_EXIST_TAG-3]
	_ = x[ERROR_NOT_EXIST_TAG-4]
	_ = x[ERROR_NOT_EXIST_ARTICLE-5]
	_ = x[ERROR_AUTH_CHECK_TOKEN_FAIL-6]
	_ = x[ERROR_AUTH_CHECK_TOKEN_TIMEOUT-7]
	_ = x[ERROR_AUTH_TOKEN-8]
	_ = x[ERROR_AUTH-9]
}

const _ErrCode_name = "okfail请求参数错误已存在该标签名称该标签不存在该文章不存在Token鉴权失败Token已超时Token生成失败Token错误"

var _ErrCode_index = [...]uint8{0, 2, 6, 24, 48, 66, 84, 101, 115, 132, 143}

func (i ErrCode) String() string {
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}
