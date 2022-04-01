package dd_api

import(
	"strings"
	"strconv"
	"context"
	"github.com/pigfall/gosdk/http"
)

type ApiRoleAddToUser struct{
	RoleIds string `json:"roleIds"`
	UserIds string `json:"userIds"`
}


func NewApiRoleAddToUser(roleIds []int,userIds []string) *ApiRoleAddToUser{
	roleIdStr := make([]string,0,len(roleIds))
	for _,v := range roleIds{
		roleIdStr = append(roleIdStr,strconv.FormatInt(int64(v),10))
	}
	return &ApiRoleAddToUser{
		RoleIds: strings.Join(roleIdStr,","),
		UserIds: strings.Join(userIds,","),
	}
}

type ApiRoleAddToUserRes struct{}

func (this *ApiRoleAddToUser) ExecBy(ctx context.Context,cli *Client)error{
	res := &ApiRoleAddToUserRes{}
	err := cli.Do(
		ctx,
		"topapi/role/addrolesforemps",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil{
		return err
	}
	return nil

}
