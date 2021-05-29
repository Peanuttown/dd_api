package client

import(
		"testing"
		"context"
)

func TestDeptTree(t *testing.T){
	cli := NewClient(testAppCfg())
	ctx := context.Background()
	tree,err := BuildDeptTreeByApi(ctx,ROOT_DEPT_ID,cli)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(tree.ToString(ctx))
}
