package es

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

var esUrl string

// HitsTwoData struct
type HitsTwoData struct {
	Source json.RawMessage `json:"_source"`
}

// ReqSearchData struct
type ReqSearchData struct {
	Hits HitsData `json:"hits"`
}

// TotalData struct
type TotalData struct {
	Value    int
	Relation string
}

type HitsData struct {
	Total TotalData     `json:"total"`
	Hits  []HitsTwoData `json:"hits"`
}

// 初始化
func init() {
	esUrl = "http://127.0.0.1:9200/"
}

func EsSearch(indexName string, query map[string]interface{}, from int, size int, sort []map[string]string) HitsData {
	searchQuery := map[string]interface{}{
		"query": query,
		"from":  from,
		"size":  size,
		"sort":  sort,
	}

	req := httplib.Get(esUrl + indexName + "/_search")
	req.JSONBody(searchQuery)
	str, err := req.String()

	if err != nil {
		fmt.Println(err)
	}
	var stb ReqSearchData
	err = json.Unmarshal([]byte(str), &stb)

	return stb.Hits
}

// 添加
func EsAdd(indexName string, id string, body map[string]interface{}) bool {
	req := httplib.Post(esUrl + indexName + "/_doc/" + id)
	req.JSONBody(body)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}

// EsEdit modify es database
func EsEdit(indexName string, id string, body map[string]interface{}) bool {
	bodyData := map[string]interface{}{
		"doc": body,
	}

	req := httplib.Post(esUrl + indexName + "/_doc/" + id + "/_update")
	req.JSONBody(bodyData)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}

// EsDelete delete es database
func EsDelete(indexName string, id string) bool {
	req := httplib.Delete(esUrl + indexName + "/_doc/" + id)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}
