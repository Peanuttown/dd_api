package client

import(
//	"fmt"
		"context"
	"github.com/Peanuttown/tzzGoUtil/http"
)

type DeptIdEmbed struct{
	DeptId uint `json:"dept_id"`
}

type DeptParentIdEmbed struct{
	ParentId uint `json:"parent_id"`
}

// 获取部门详情
type ApiDeptGetDetail struct{
	DeptIdEmbed
}

func NewApiDeptGetDetail(deptId uint)*ApiDeptGetDetail{
	return &ApiDeptGetDetail{
		DeptIdEmbed:DeptIdEmbed{
			DeptId:deptId,
		},
	}
}

type ApiDeptGetDetailRes struct{
	DeptIdEmbed
	DeptParentIdEmbed
}



func (this *ApiDeptGetDetail) ExecBy(ctx context.Context,cli *Client)(*ApiDeptGetDetailRes,error){
	res := &ApiDeptGetDetailRes{}
	
	err :=cli.Do(
		ctx,"topapi/v2/department/get",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res,nil
}
