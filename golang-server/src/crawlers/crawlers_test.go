package crawlers

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestAliBench(t *testing.T) {
	return
	f := NewFetcher("m.qiushibaike.com")
	f.Get("/hot/page")
	data := url.Values{
		"task_from": {"self"},
		"target":    {"http://m.qiushibaike.com/hot/page/"},
		"ac":        {"http"},
	}
	_, body, err := f.PostForm("/new_task.php", data)
	if err != nil {
		t.Fatal(err)
	}
	var a map[string]interface{}
	json.Unmarshal(body, &a)
	if a["code"].(float64) != 0 {
		t.Error("result not match", a)
	}
}

/*
func TestCache(t *testing.T) {
f := NewFetcher("163.com")
f.CacheTime = 60
f.Get("/")
f.RemoveGetCache("/")
f.Get("/")
t.Error()
}*/
