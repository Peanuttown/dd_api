package dd_api

import (
	"context"
	"fmt"
	"github.com/pigfall/gosdk/http"
	//	"fmt"
	//"encoding/base64"
)



// 注册回调时间
type ApiContactsCallbackRegister struct {
	EventTypes   []EventType `json:"call_back_tag"`
	Token        string         `json:"token"`
	AESKeyBase64 string         `json:"aes_key"`
	CallbackUrl  string         `json:"url"`
}


func NewApiContactsCallbackRegister(
	token string,
	eventTypes []EventType,
	aesKeyBase64 string,
	callbackUrl string,
) *ApiContactsCallbackRegister{
	fmt.Println(callbackUrl)

	return &ApiContactsCallbackRegister{
		EventTypes:   eventTypes,
		Token:        token,
		AESKeyBase64: aesKeyBase64,
		CallbackUrl:  callbackUrl,
	}
}


type ApiContactsCallbackRegisterRes struct{
	Res
}

func (this *ApiContactsCallbackRegister) ExecBy(ctx context.Context, cli *Client) error {
	reqBuilder := http.NewRequestBuilder().MethodPost().JsonParam(this)
	res := &ApiContactsCallbackRegisterRes{}
	err := cli.Do(ctx,"call_back/register_call_back",reqBuilder,res)
	if err != nil {
		return err
	}
	return nil
}
