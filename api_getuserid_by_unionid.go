package dd_api

import(
	"context"
	"fmt"
	"github.com/pigfall/gosdk/http"
)

// 根据unionid获取用户userid
// ref: https://developers.dingtalk.com/document/app/query-a-user-by-the-union-id?spm=ding_open_doc.document.0.0.4b57479192VwD0#topic-1960045
type ApiGetUserIdByUnionId struct{
	UnionId string `json:"unionid"`
}

func NewApiGetUserIdByUnionId(unionId string)(*ApiGetUserIdByUnionId){
	return &ApiGetUserIdByUnionId{
		UnionId:unionId,
	}
}

type ApiGetUserIdByUnionIdRes struct{
	ContactType int `json:"contact_type"`
	UserId string `json:"userid"`
}

func (this *ApiGetUserIdByUnionIdRes) IsInnerCorpUser() bool{
	return this.ContactType == 0
}

func (this *ApiGetUserIdByUnionId) ExecBy(ctx context.Context,cli *Client)(*ApiGetUserIdByUnionIdRes,error){
	res := &ApiGetUserIdByUnionIdRes{
		ContactType :-1,
	}
	reqBuilder := http.NewRequestBuilder().JsonParam(this).MethodPost()
	err := cli.Do(ctx,"topapi/user/getbyunionid",reqBuilder,res)
	if err != nil{
		return nil,err
	}
	if len(res.UserId) == 0{
		return nil, fmt.Errorf("res code is ok, but userId is nil")
	}
	return res,nil
}


