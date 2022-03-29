package dd_api

import(
	"fmt"
	"context"
		
	dt "github.com/pigfall/gosdk/datastruct"
)

type DeptTree struct{
	*dt.Tree
}


type DeptNodeValue struct{
	DeptId uint
	Name string
}

func (this *DeptNodeValue) DeptIdStr()string{
	return fmt.Sprintf("%d",this.DeptId)
}


type deptTreeBuilderByApi struct{
	deptId  uint
	client *Client
}

func newDeptTreeBuilderByApi(deptId uint,cli *Client) *deptTreeBuilderByApi{
	return &deptTreeBuilderByApi{
		deptId:deptId,
		client:cli,
	}
}

func (this *deptTreeBuilderByApi) GetValue(ctx context.Context)(interface{},error){
	res,err := NewApiDeptGetDetail(this.deptId).ExecBy(ctx,this.client)
	if err != nil{
		return nil,err
	}
	return &DeptNodeValue{
		DeptId:res.DeptId,
		Name:res.Name,
	},nil
}

func (this *deptTreeBuilderByApi) GetChildren(ctx context.Context)([]dt.TreeBuilderI,error){
	subDepts,err := NewApiDeptGetSubDeptsList(this.deptId).ExecBy(ctx,this.client)
	if err != nil{
		return nil,err
	}
	builders := make([]dt.TreeBuilderI,0,len(subDepts))
	for _,id := range subDepts{
		builders = append(builders,newDeptTreeBuilderByApi(id,this.client))
	}
	return builders,nil
}



func BuildDeptTreeByApi(ctx context.Context,deptId uint,cli *Client)(*DeptTree,error){
	tree,err := dt.BuildTree(ctx,newDeptTreeBuilderByApi(deptId,cli))
	if err != nil{
		return nil,err
	}
	return &DeptTree{
		Tree:tree,
	},nil
}

