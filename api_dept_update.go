package dd_api

import (
	"context"

	"github.com/pigfall/gosdk/http"
)

// 获取部门用户userid列表
type ApiDeptUpdate struct {
	DeptIdEmbed
	Name      string `json:"name,omitempty"`
	ParentId  uint   `json:"parent_id,omitempty"`
}

func NewApiDeptUpdate(deptId uint, name string, parentDeptId uint) *ApiDeptUpdate {
	return &ApiDeptUpdate{
		DeptIdEmbed: DeptIdEmbed{
			DeptId: deptId,
		},
	}
}

func (this *ApiDeptUpdate) ExecBy(ctx context.Context, cli *Client) (err error) {
	err = cli.Do(
		ctx,
		"topapi/v2/department/update",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}
