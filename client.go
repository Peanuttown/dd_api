package dd_api

import(
	"time"
	"fmt"
	"encoding/json"
	"log"
	"context"
	"github.com/Peanuttown/tzzGoUtil/http"
	stdhttp "net/http"
	"net/http/httputil"
)

type Client struct{
	cfg *Cfg
	accessToken string
	accessTokenExpiredAt time.Time
}

func NewClient(cfg *Cfg) *Client{
	return &Client{
		cfg:cfg,
	}
}

type Res struct{
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	Result interface{} `json:"result"`
}

type ResI interface{
	Err() error
	Error() string
}

func (this *Res) Err()error{
	if this.ErrMsg ==  "ok"{
		return nil
	}
	return this
}

func (this *Res) Error()string{
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %v",this.ErrCode,this.ErrMsg)
}

func(this *Client) buildApiURL(apiPath string)(string){
return fmt.Sprintf("https://%s/%s",this.cfg.ApiHost,apiPath)
}

func (this *Client) GetAppKey()string{
	return this.cfg.AppKey
}

func (this *Client) GetAppSecret() string{
	return this.cfg.AppSecret
}


func (this *Client) updateTokenIfExpired(ctx context.Context)error{
	now := time.Now()
	if this.accessTokenExpiredAt.After(now){
		return nil
	}
	// update token
	req,err := http.NewRequestBuilder().MethodGet().PutParamsToUrl(
		map[string]string{
			"appkey":this.cfg.AppKey,
			"appsecret":this.cfg.AppSecret,
		},
	).URL(this.buildApiURL("gettoken")).Build(ctx)
	type TokenRes struct{
		Res
		Token string `json:"access_token"`
		ExpiresInSecond int `json:"expires_in"`
	}
	if err != nil{
		return err
	}
	res,err := stdhttp.DefaultClient.Do(req)
	if err != nil{
		return err
	}
	bodyBytes,err := http.ReadResBody(res)
	if err != nil{
		return err
	}
	var tokenRes = &TokenRes{}
	err = this.handleResEntity(bodyBytes,tokenRes)
	if err != nil{
		return err
	}
	this.accessToken = tokenRes.Token
	this.accessTokenExpiredAt = time.Now().Add(time.Duration(tokenRes.ExpiresInSecond)*time.Second / 2)
	return nil
}
func (this *Client) DoNoNeedToken(ctx context.Context,apiPath string,reqBuilder *http.RequestBuilder,resEntity interface{})(error){
	return this.do(ctx,false,apiPath,reqBuilder,resEntity)
}

func (this *Client) do(ctx context.Context,appendToken bool,apiPath string,reqBuilder *http.RequestBuilder,resEntity interface{})(error){
	err := this.updateTokenIfExpired(ctx)
	if err != nil{
		return err
	}
	url := fmt.Sprintf("https://%s/%s",this.cfg.ApiHost,apiPath)
	reqBuilder.URL(url)
	var httpReq *stdhttp.Request
	if appendToken{
		httpReq,err=reqBuilder.PutParamsToUrl(map[string]string{
			"access_token":this.accessToken,
		}).Build(ctx)
	}else{
		httpReq,err =reqBuilder.Build(ctx)
	}
		if err != nil{
			return err
		}
	reqDump,err := httputil.DumpRequest(httpReq,true)
	if err != nil{
		panic(err)
	}
	log.Println(string(reqDump))
	res,err := stdhttp.DefaultClient.Do(httpReq)
	if err != nil{
		return err
	}
	resBytes,err := http.ReadResBody(res)
	if err != nil{
		return err
	}

	return this.handleWrapResEntity(resBytes,resEntity)

}


func (this *Client) Do(ctx context.Context,apiPath string,reqBuilder *http.RequestBuilder,resEntity interface{})(error){
	return this.do(ctx,true,apiPath,reqBuilder,resEntity)
}

func (this *Client) handleResEntity(bytes []byte,res ResI)error{
	err := json.Unmarshal(bytes,res)
	if err != nil{
		return err
	}
	err=res.Err()
	if err != nil{
		return err
	}
	return nil
}


func (this *Client) handleWrapResEntity(bytes []byte,resEntity interface{})(error){
	resI,ok := resEntity.(ResI)
	if ok{
		return this.handleResEntity(bytes,resI)
	}else{
		resWrapper := &Res{
			ErrCode: -1,
			Result:resEntity,
		}
		return this.handleResEntity(bytes,resWrapper)
	}
}
