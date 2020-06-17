package goroseadapter

import (
	"fmt"
	"testing"

	"github.com/dafanshu/simplejson"
	"github.com/stretchr/testify/assert"
)

var rule = &CasbinRule{
	Id:     0,
	P_type:  "p",
	V0:     "admin1",
	V1:     "",
	V2:     "",
	V3:     "",
	V4:     "",
	V5:     "",
	Legion: "",
}

func TestClient_Retrieve(t *testing.T) {
	client := New("1254789", "http://192.168.1.108:8088/function/xmysql-rbac", "test")
	retrieve, err := client.Retrieve(rule)
	if err != nil {
		panic(err)
	}
	json, err := simplejson.NewJson(retrieve)
	if err != nil {
		panic(err)
	}
	index := json.GetIndex(0)
	s, _ := index.Get("id").Int()
	fmt.Println(s)
	marshalJSON, err := index.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshalJSON))
	fmt.Println(string(retrieve))
}

func TestLoadParamsWhere(t *testing.T) {
	where := LoadParamsWhere(rule)
	assert.Equal(t, "(id,eq,0)~and(p_type,eq,p)~and(v0,eq,admin)", where)
}

func TestLoadParams(t *testing.T) {
	params := LoadParams(rule)
	assert.Equal(t, "id=0&p_type=p&v0=admin", params)
}
