package dd_api

import(
	"encoding/base64"
//	"net/url"
	"fmt"
	"time"
		"context"
	"github.com/Peanuttown/tzzGoUtil/http"
	"github.com/Peanuttown/tzzGoUtil/crypto/hmac"
)

// 根据sns临时授权码获取用户信息
// ref: https://developers.dingtalk.com/document/app/obtain-the-user-information-based-on-the-sns-temporary-authorization?spm=ding_open_doc.document.0.0.4b574791rcsrts#topic-1995619
type ApiGetUserInfoByteCode struct{
	AccessKey string
	AuthCode string
	AccessSecret string
}

func NewApiGetUserInfoByteCode(accessKey string,accessSecret string, authCode string)*ApiGetUserInfoByteCode{
	return &ApiGetUserInfoByteCode{
		AccessKey:accessKey,
		AuthCode:authCode,
		AccessSecret:accessSecret,
	}
}


type ApiGetUserInfoByteCodeRes struct{
	Res
	UserInfo *ApiGetUserInfoByteCodeResUserInfo  `json:"user_info"`
}

type  ApiGetUserInfoByteCodeResUserInfo struct{
	UnionId string `json:"unionid"`
	OpenId string `json:"openid"`
}

func (this *ApiGetUserInfoByteCode) ExecBy(ctx context.Context,cli *Client)(*ApiGetUserInfoByteCodeResUserInfo,error){
	nowInMillSecond := fmt.Sprintf("%d",time.Now().Unix() * 1000)
	signatureBytes,err := hmac.EnCrypt([]byte(this.AccessSecret),[]byte(nowInMillSecond))
	if err != nil{
		return nil, fmt.Errorf("计算签名出错: %w",err)
	}
	signature:=base64.StdEncoding.EncodeToString(signatureBytes)
	// signature= url.QueryEscape(signature)
	res := &ApiGetUserInfoByteCodeRes{}
	reqBuilder:= http.NewRequestBuilder().
	MethodPost().
	PutParamsToUrl(map[string]string{
		"accessKey":this.AccessKey,
		"timestamp":nowInMillSecond,
		"signature":signature,
	}).
	JsonParam(map[string]interface{}{"tmp_auth_code":this.AuthCode})
	err =  cli.DoNoNeedToken(ctx,"sns/getuserinfo_bycode",reqBuilder,res)
	if err != nil{
		return nil,err
	}
	if res.UserInfo == nil{
		return nil, fmt.Errorf("response is ok, but response body after unmarshaled is nil")
	}
	return res.UserInfo,nil
}
