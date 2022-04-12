package dd_api

import (
	"context"
	"github.com/pigfall/gosdk/http"
)

type ApiUsersOfRole struct {
	PageIndex int  `json:"offset"`
	PageSize  int  `json:"size"`
	RoleId    uint `json:"role_id"`
}

func NewApiUsersOfRole(roleId uint, pageIndex int, pageSize int) *ApiUsersOfRole {
	return &ApiUsersOfRole{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		RoleId:    roleId,
	}
}

type ApiUsersOfRoleRes struct {
	HasMore    bool                    `json:"hasMore"`
	NextCursor int                     `json:"nextCursor"`
	List       []ApiUsersOfRoleResData `json:"list"`
}

type ApiUsersOfRoleResData struct {
	UserId       string        `json:"userid"`
	Name         string        `json:"mame"`
	ManageScopes []ManageScope `json:"manageScopes"`
}

type ManageScope struct {
	DeptId uint   `json:"dept_id"`
	Name   string `json:"name"`
}

func (this *ApiUsersOfRole) ExecBy(ctx context.Context, cli *Client)(*ApiUsersOfRoleRes,error){
	var res = &ApiUsersOfRoleRes{}
	err := cli.Do(
		ctx,
		"topapi/role/simplelist",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil {
		return nil, err
	}
	return res, nil

}
