package dd_api

import(
		"errors"
)

type ApiErrCode int

const (
	err_not_found_user     ApiErrCode = 60121
	err_callback_url_exist ApiErrCode = 71006
)

func ErrIsCallbackUrlExist(err error) bool {
	var res *Res
	if errors.As(err, &res) {
		return res.ErrCode == int(err_callback_url_exist)
	}
	return false
}
