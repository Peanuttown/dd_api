package dd_api

import(
	"context"	
	"github.com/pigfall/gosdk/http"
)

// 获取子部门列表
type ApiDeptGetSubDeptsList struct{
	DeptIdEmbed
}

func NewApiDeptGetSubDeptsList(deptId uint)*ApiDeptGetSubDeptsList{
	return &ApiDeptGetSubDeptsList{
		DeptIdEmbed:DeptIdEmbed{
			DeptId:deptId,
		},
	}
}

type ApiDeptGetSubDeptsListRes struct{
	SubDepts []uint `json:"dept_id_list"`
}

func (this *ApiDeptGetSubDeptsList) ExecBy(ctx context.Context,cli *Client)(subDeptIds []uint,err error){
	res := &ApiDeptGetSubDeptsListRes{}
	err  =cli.Do(
			ctx,
			"topapi/v2/department/listsubid",
			http.NewRequestBuilder().MethodPost().JsonParam(this),
			res,
	)
	if err != nil{
		return nil,err
	}
	return res.SubDepts,nil
}
