package errortest

import (
	"os"
	"testing"
)

type Conf struct{
	Name string `json:"name"`
}

func TestReadConfig(t *testing.T){
	os.WriteFile("test.json",[]byte(`{"name":"go"}`),6604)
	defer os.Remove("test.json")

	var cfg Conf 

	if err := Readconfig("test.conf",&cfg);err != nil{
		t.Fatal("expected no error , got ",err)
	}
	if cfg.Name != "go"{
		t.Errorf("Name = %q ; want %q",cfg.Name,"Go")
	}
}


func TestReadConfig_NotFound(t *testing.T) {
    var cfg Conf
    err := Readconfig("nope.json", &cfg)
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}