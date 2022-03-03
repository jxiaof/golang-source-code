package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func handleResponse(resp *http.Response) (*http.Response, error) {
	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func newHttpRequest(url string, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http request err, err: %v", err)
	}
	req.Header.Add("token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request err, err: %v", err)
	}
	return handleResponse(resp)
}

// 向kas发送get请求，获取用户space信息
func GetUserSpace(url, token string) ([]byte, error) {
	res, err := newHttpRequest(url, token)
	if err != nil {
		return nil, fmt.Errorf("get user space err: %v, token: %s", err, token)
	}
	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read user space api response err: %v", err)
	}
	// data := string(rawBody)
	return rawBody, nil
}

type Json struct {
	data interface{}
}

func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

func NewJson(body []byte) (*Json, error) {
	j := new(Json)
	err := j.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// func parserByte(){
// 	json.Unmarshal(body,&res)
// }

type SUserSpaceData struct {
	Data []map[string]interface{} `json:"data"`
}
type SUserSpaceRes struct {
	Code   int            `json:"code"`
	Msg    string         `json:"msg"`
	Result SUserSpaceData `json:"result"`
}

// func main() {
// 	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJleHAiOjE2NDIwNjQ4MzcsImlzcyI6Imthcy1iYWNrZW5kIn0.0cq6gC6snQpaBugpfYd8dXeWMPf6wqKC4cBo9WlcrIQ"
// 	url := "http://120.92.93.100:30808/api/space/list"
// 	// url := "http://120.92.93.100:30808/api/fetch_user"
// 	data, err := GetUserSpace(url, token)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var res SUserSpaceRes
// 	json.Unmarshal(data, &res)
// 	// fmt.Println(res.Result.Data...)
// 	for _, v := range res.Result.Data {
// 		for k, vv := range v {
// 			if k == "id" {
// 				fmt.Println(vv)
// 			}
// 		}
// 	}

// 	// fmt.Println(userInfoJs.data.(map[string]interface{})["result"])
// }

var pool *redis.Pool //创建redis连接池

func init() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("Set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)
	pool.Close() //关闭连接池
}
