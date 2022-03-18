package dd_api

import(
	"context"
	"github.com/pigfall/gosdk/http"
)


// 获取部门用户userid列表
type ApiDeptGetUserIds struct{
	DeptIdEmbed
}

func NewApiDeptGetUserIds(deptId uint)* ApiDeptGetUserIds{
	return  &ApiDeptGetUserIds{
		DeptIdEmbed:DeptIdEmbed{
			DeptId:deptId,
		},
	}
}

type ApiDeptGetUserIdsRes struct{
	UserIds []string `json:"userid_list"`
}

func (this *ApiDeptGetUserIds) ExecBy(ctx context.Context,cli *Client)(userIds []string,err error){
	res := &ApiDeptGetUserIdsRes{}
	err  = cli.Do(
		ctx,
		"topapi/user/listid",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res.UserIds,nil
}
