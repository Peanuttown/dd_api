package client

import(
	"context"	
	"github.com/Peanuttown/tzzGoUtil/http"
)


// 创建部门
type ApiDeptCreate struct{
	Name string `json:"name"`
	DeptParentIdEmbed
}

func NewApiDeptCreate(name string,parentDeptId uint) *ApiDeptCreate{
	return &ApiDeptCreate{
		Name:name,
		DeptParentIdEmbed:DeptParentIdEmbed{
			ParentId:parentDeptId,
		},
	}
}

type ApiDeptCreateRes struct{
	DeptIdEmbed
}

func (this *ApiDeptCreate) ExecBy(ctx context.Context,cli *Client)(deptId uint,err error){
	res := &ApiDeptCreateRes{}
	err = cli.Do(
		ctx,
		"topapi/v2/department/create",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return 0,err
	}
	return res.DeptId,nil
}
