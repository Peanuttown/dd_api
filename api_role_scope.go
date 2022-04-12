package dd_api

import (
	"context"
	"strconv"
	"strings"

	"github.com/pigfall/gosdk/http"
)

//ref: https://open.dingtalk.com/document/orgapp-server/update-role-member-management-department-scope
type ApiRoleScopeUpdate struct {
	UserIdEmbed
	RoleId  int    `json:"role_id"`
	DeptIds string `json:"dept_ids"`
}

func NewApiRoleScopeUpdate(userId string, roleId int, deptIds []uint) *ApiRoleScopeUpdate {
	deptIdsStr := make([]string, 0, len(deptIds))
	for _, v := range deptIds {
		deptIdsStr = append(deptIdsStr, strconv.FormatUint(uint64(v), 10))
	}
	return &ApiRoleScopeUpdate{
		DeptIds:     strings.Join(deptIdsStr, ","),
		RoleId:      roleId,
		UserIdEmbed: UserIdEmbed{UserId: userId},
	}
}

func (this *ApiRoleScopeUpdate) ExecBy(ctx context.Context, cli *Client) error {
	err := cli.Do(
		ctx,
		"topapi/role/scope/update",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		nil,
	)
	if err != nil {
		return err
	}
	return nil

}
