package dd_api

import(
	"testing"
	"context"
)


func TestApiDeptGetDetail(t *testing.T){
	cli := NewClient(testAppCfg())
	ctx := context.Background()
	res,err := NewApiDeptGetDetail(1).ExecBy(ctx,cli)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("%+v",res)


}
