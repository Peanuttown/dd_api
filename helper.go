package dd_api

import(
		"context"
		"fmt"
)

type ApiHelper struct{
	cli *Client
}

func ApiHelperNew(cli *Client)*ApiHelper{
	return &ApiHelper{
		cli:cli,
	}
}


func (this *ApiHelper)ForEachDepsUsersByDeptFirst(ctx context.Context,do func(ctx context.Context,deptId uint ,userIds []string)error)error{
	return  this.forEachDeptUsersByDeptFisrt(ctx,ROOT_DEPT_ID,do)
}



func (this *ApiHelper) forEachDeptUsersByDeptFisrt(ctx context.Context,deptId uint,do func(ctx context.Context,deptId uint,userIds []string)error)error{
	var apiCli = this.cli
	subDepts,err := NewApiDeptGetSubDeptsList(deptId).ExecBy(ctx,apiCli)
	if err != nil{
		return fmt.Errorf("Get dept %v sub depts failed %w",deptId,err)
	}
	deptUserIds,err  :=NewApiDeptGetUserIds(deptId).ExecBy(ctx,this.cli)
	if err != nil{
		return fmt.Errorf("Get dept %v users failed %w",deptId,err)
	}
	err =do(ctx,deptId,deptUserIds)
	if err != nil{
		return err
	}

	for _,subDeptId := range subDepts{
		err = this.forEachDeptUsersByDeptFisrt(ctx,subDeptId,do)
		if err != nil{
			return err
		}
	}

	return nil
}
