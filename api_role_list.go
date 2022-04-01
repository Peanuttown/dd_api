package dd_api

import(
	"github.com/pigfall/gosdk/http"
	"context"
)

type ApiRoleList struct{
	Size int `json:"size"`
	Offset int `json:"offset"`
}


func NewApiRoleList(size,offset int)*ApiRoleList{
	return &ApiRoleList{
		Size: size,
		Offset: offset,
	}
}

type ApiRoleListRes struct{
	HasMore bool `json:"hasMore"`
	RoleGroups  []RoleGroup `json:"list"`
}

type RoleGroup struct{
	Name string `json:"name"`
	GroupId int `json:"group_id"`
	Roles []Role `json:"roles"`
}

type Role struct{
	Name  string `json:"name"`
	Id  int `json:"id"`
}

func (this *ApiRoleList)ExecBy(ctx context.Context,cli *Client)(*ApiRoleListRes,error){
	res := &ApiRoleListRes{}
	err  := cli.Do(
		ctx,
		"topapi/role/list",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res,nil

}
