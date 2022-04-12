package dd_api

import(
		
	"github.com/pigfall/gosdk/http"
	"context"
)

type ApiRoleCreate struct{
	RoleGroupId uint `json:"groupId"`
	RoleName string `json:"roleName"`
}

func NewApiRoleCreate(groupId uint,roleName string)*ApiRoleCreate{
	return &ApiRoleCreate{
		RoleGroupId:groupId,
		RoleName:roleName,
	}
}

type ApiRoleCreateRes struct{
	RoleId uint `json:"roleId"`
}

func (this *ApiRoleCreate) ExecBy(ctx context.Context,cli *Client)(*ApiRoleCreateRes,error){
	if len(this.RoleName) == 0{
		panic("unexpect")
	}
	if this.RoleGroupId ==0{
		panic("invalid param")
	}
	res := &ApiRoleCreateRes{}
	err  := cli.Do(
		ctx,
		"role/add_role",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res,nil
}
