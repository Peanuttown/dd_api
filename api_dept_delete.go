package dd_api

import(
		"context"
	"github.com/Peanuttown/tzzGoUtil/http"
)

// 删除部门
type ApiDeptDelete struct{
	DeptIdEmbed
}

func NewApiDeptDelete(deptId uint)*ApiDeptDelete{
	return &ApiDeptDelete{
		DeptIdEmbed:DeptIdEmbed{
			DeptId:deptId,
		},
	}
}


func (this *ApiDeptDelete) ExecBy(ctx context.Context,cli *Client)error{
return 	cli.Do(
		ctx,
		"topapi/v2/department/delete",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		nil,
	)
}
