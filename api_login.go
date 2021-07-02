package dd_api

import(
	"context"	
	"github.com/Peanuttown/tzzGoUtil/http"
)


// 企业内部 authCode登陆
// https://developers.dingtalk.com/document/app/get-user-userid-through-login-free-code?spm=ding_open_doc.document.0.0.7e1e47e5Eaivhr#topic-1936806
type ApiCorpInnerLogin struct{
	Code string  
}

func NewApiCorpInnerLogin(authCode string) *ApiCorpInnerLogin{
	return &ApiCorpInnerLogin{
		Code:authCode,
	}
}

type ApiCorpInnerLoginRes struct{
	UserIdEmbed
}

func (this *ApiCorpInnerLogin) ExecBy(ctx context.Context,cli *Client)(*ApiCorpInnerLoginRes,error){
	res := &ApiCorpInnerLoginRes{}
	err := cli.Do(
		ctx,
		"user/getuserinfo",
		http.NewRequestBuilder().MethodGet().PutParamsToUrl(map[string]string{
			"code":this.Code,
		}),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res,nil
}
