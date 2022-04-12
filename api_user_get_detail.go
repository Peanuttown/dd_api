package dd_api

import (
	"context"

	"github.com/pigfall/gosdk/http"
)

type UserIdEmbed struct {
	UserId string `json:"userid"`
}

// 根据userid获取用户详情
type ApiUserGetDetail struct {
	UserIdEmbed
}

func NewApiUserGetDetail(userId string) *ApiUserGetDetail {
	return &ApiUserGetDetail{
		UserIdEmbed: UserIdEmbed{
			UserId: userId,
		},
	}
}

type ApiUserGetDetailRes struct {
	UserId          string                             `json:"-"`
	Name            string                             `json:"name"`
	Mobile          string                             `json:"mobile"`
	Title           string                             `json:"title"`
	IsLeaderInDepts []ApiUserGetDetailRes_LeaderInDept `json:"leader_in_dept"`
	Avatar          string                             `json:"avatar"`
	RoleList        []RoleListData                     `json:"role_list"`
}

type RoleListData struct {
	RoleId    uint   `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"group_name"`
}

type ApiUserGetDetailRes_LeaderInDept struct {
	DeptIdEmbed
	Leader bool `json:"leader"`
}

func (this *ApiUserGetDetail) ExecBy(ctx context.Context, cli *Client) (*ApiUserGetDetailRes, error) {
	res := &ApiUserGetDetailRes{
		UserId: this.UserId,
	}
	err := cli.Do(
		ctx,
		"topapi/v2/user/get",
		http.NewRequestBuilder().MethodPost().JsonParam(this),
		res,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
