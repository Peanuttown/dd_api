package dd_api

import (
	"context"
	//	"fmt"
	//"encoding/base64"
"github.com/Peanuttown/tzzGoUtil/http"
)

// 注册回调时间
type ApiContactsCallbackRegUpdate struct {
	ApiContactsCallbackRegister
}

type ApiCallbackRegUpdateRes struct {
	Res
}

func NewApiReqCallbackRegUpdate(
	token string,
	eventTypes []EventType,
	aesKeyBase64 string,
	callbackUrl string,
) *ApiContactsCallbackRegUpdate{
	return &ApiContactsCallbackRegUpdate{
		ApiContactsCallbackRegister: *NewApiContactsCallbackRegister(
			token,
			eventTypes,
			aesKeyBase64,
			callbackUrl,
		),
	}
}


func (this *ApiContactsCallbackRegUpdate) ExecBy(ctx context.Context, cli *Client) error {
	res := &ApiCallbackRegUpdateRes{}
	reqBuilder := http.NewRequestBuilder().MethodPost().JsonParam(this)
	err := cli.Do(
		ctx,
		"call_back/update_call_back",
		reqBuilder,
		res,
	)
	if err != nil {
		return err
	}
	return nil
}
