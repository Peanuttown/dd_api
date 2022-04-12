package dd_api

import (
	"context"
	"fmt"
)

type ApiHelper struct {
	cli *Client
}

func ApiHelperNew(cli *Client) *ApiHelper {
	return &ApiHelper{
		cli: cli,
	}
}

func (this *ApiHelper) ForEachDepsUsersByDeptFirst(ctx context.Context, do func(ctx context.Context, deptId uint, userIds []string) error) error {
	return this.forEachDeptUsersByDeptFisrt(ctx, ROOT_DEPT_ID, do)
}

func (this *ApiHelper) forEachDeptUsersByDeptFisrt(ctx context.Context, deptId uint, do func(ctx context.Context, deptId uint, userIds []string) error) error {
	var apiCli = this.cli
	subDepts, err := NewApiDeptGetSubDeptsList(deptId).ExecBy(ctx, apiCli)
	if err != nil {
		return fmt.Errorf("Get dept %v sub depts failed %w", deptId, err)
	}
	deptUserIds, err := NewApiDeptGetUserIds(deptId).ExecBy(ctx, this.cli)
	if err != nil {
		return fmt.Errorf("Get dept %v users failed %w", deptId, err)
	}
	err = do(ctx, deptId, deptUserIds)
	if err != nil {
		return err
	}

	for _, subDeptId := range subDepts {
		err = this.forEachDeptUsersByDeptFisrt(ctx, subDeptId, do)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *ApiHelper) FindRoleScopeOfUser(ctx context.Context, userId string, roleId uint) (scopes []ManageScope, findUserInTheRole bool, err error) {
	err = this.ForEachUserOfRole(ctx, roleId, func(ctx context.Context, user *ApiUsersOfRoleResData) (stop bool, err error) {
		if user.UserId == user.UserId {
			findUserInTheRole = true
			scopes = user.ManageScopes
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, false, err
	}
	return
}

func (this *ApiHelper) ForEachUserOfRole(ctx context.Context, roleId uint, do func(ctx context.Context, user *ApiUsersOfRoleResData) (stop bool, err error)) error {
	var pageIndex = 0
	var pageSize = 20
	for {
		res, err := NewApiUsersOfRole(roleId, pageIndex, pageSize).ExecBy(ctx, this.cli)
		if err != nil {
			return err
		}
		for _, v := range res.List {
			stop, err := do(ctx, &v)
			if err != nil {
				return err
			}
			if stop {
				return nil
			}
		}
		if res.HasMore {
			pageIndex = res.NextCursor
		} else {
			break
		}
	}
	return nil
}

func (this *ApiHelper) FindRoleGroupByName(ctx context.Context,roleGroupName string)(*RoleGroup,error){
	var toFind *RoleGroup
	err := this.ForEachRoleGroup(ctx, func(ctx context.Context,rg *RoleGroup)(bool,error){
		if rg.Name == roleGroupName{
			toFind =rg
			return true,nil
		}
		return false,nil
	})
	if err != nil{
		return nil,err
	}
	return toFind,nil
}

func (this *ApiHelper) ForEachRoleGroup(ctx context.Context,f func(ctx context.Context,roleGroup *RoleGroup)(stop bool,err error))(error){
	var pageIndex = 0
	var pageSize = 20
	for {
		res, err := NewApiRoleList(pageSize,pageIndex).ExecBy(ctx, this.cli)
		if err != nil {
			return err
		}
		for _, v := range res.RoleGroups{
			stop, err := f(ctx, &v)
			if err != nil {
				return err
			}
			if stop {
				return nil
			}
		}
		if res.HasMore {
			pageIndex++
		} else {
			break
		}
	}
	return nil
}
