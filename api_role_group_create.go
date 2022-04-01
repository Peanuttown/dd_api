package dd_api

import(
	"github.com/pigfall/gosdk/http"
	"context"
)

type ApiRoleGroupCreate struct{
	Name string `json:"name"`
}


func NewApiRoleGroupCreate(name string){

}

type ApiRoleGroupCreateRes struct{
	RoleGroupIdEmbed
}

func (this *ApiRoleGroupCreate) ExecBy(ctx context.Context,cli *Client)(*ApiRoleGroupCreateRes,error){
	res := &ApiRoleGroupCreateRes{}
	err := cli.Do(
		ctx,
		"role/add_role_group",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return nil,err
	}
	return res,nil
}
