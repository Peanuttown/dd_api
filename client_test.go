package client

import(
	"os"
	"path"
	"encoding/json"
		"testing"
	"github.com/Peanuttown/tzzGoUtil/encoding"
		"context"
)

func testAppCfg() *Cfg{
	wd,err := os.Getwd()
	if err != nil{
		panic(err)
	}
	cfg := &Cfg{}
	err = encoding.UnMarshalByFile(path.Join(wd,"debug","test_app_cfg.json"),cfg,json.Unmarshal)
	if err != nil{
		panic(err)
	}
	return cfg

}

func TestClient(t *testing.T){
	cli := NewClient(testAppCfg())
	ctx :=context.Background()
	err := cli.updateTokenIfExpired(ctx)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(cli.accessToken)
	t.Log(cli.accessTokenExpiredAt.String())
}
