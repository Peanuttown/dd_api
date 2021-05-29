package client

import(
		"context"
		"testing"
)

func TestApiDeptCreate(t *testing.T){
	cli :=NewClient(testAppCfg())
	ctx := context.Background()
	deptId,err := NewApiDeptCreate("deptTest2",ROOT_DEPT_ID).ExecBy(ctx,cli)
	if err != nil{
		t.Fatal(err)
	}
	t.Log("< Create dept success > depId:", deptId)

	subDepts,err := NewApiDeptGetSubDeptsList(ROOT_DEPT_ID).ExecBy(ctx,cli)
	if err != nil{
		t.Fatal(err)
	}
	var subDeptsExist bool
	for _,v := range subDepts{
		if v == deptId{
			subDeptsExist = true
			break
		}
	}
	if !subDeptsExist{
		t.Fatal("Not found deptid in subDepts list")
	}

	err =	NewApiDeptDelete(deptId).ExecBy(ctx,cli)
	if err != nil{
		t.Fatal(err)
	}
}
