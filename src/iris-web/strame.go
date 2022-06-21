/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-09 20:42:37
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-19 14:16:46
 */
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt" //只是一个可选的助手
	"io"
	"net/http"
	"strings"
	"time" //展示延迟

	iris "github.com/kataras/iris/v12"
)

func NewInts() []int {
	ints := []int{}
	for i := 0; i < 200; i++ {
		ints = append(ints, i)
	}
	return ints
}

type messageNumber struct {
	Number int `json:"number"`
}

var errDone = errors.New("done")

func HtmlHandler(ctx iris.Context) {
	// ctx.ContentType("text/html")
	// ctx.Header("Transfer-Encoding", "chunked")
	// i := 0
	// ints := NewInts()
	// // Send the response in chunks and
	// // wait for half a second between each chunk,
	// // until connection closed.
	// err := ctx.StreamWriter(func(w io.Writer) error {
	// 	ctx.Writef("Message number %d<br>", ints[i])
	// 	time.Sleep(500 * time.Millisecond) // simulate delay.
	// 	if i == len(ints)-1 {
	// 		return errDone // ends the loop.
	// 	}
	// 	i++
	// 	return nil // continue write
	// })
	// if err != errDone {
	// 	// Test it by canceling the request before the stream ends:
	// 	// [ERRO] $DATETIME stream: context canceled.
	// 	ctx.Application().Logger().Errorf("stream: %v", err)
	// }
}

func JsonHandler(ctx iris.Context) {
	ctx.ContentType("application/json")
	ctx.Header("Transfer-Encoding", "chunked")
	i := 0
	ints := NewInts()
	notifyClose := ctx.Request().Context().Done()
	fmt.Println("request body:", ctx.Request().Body)
	for {
		select {
		case <-notifyClose:
			// err := ctx.Request().Context().Err()
			ctx.Application().Logger().Infof("Connection closed, loop end.")
			return
		default:
			ctx.JSON(messageNumber{Number: ints[i]})
			ctx.WriteString("\n")
			// time.Sleep(500 * time.Millisecond) // simulate delay.
			time.Sleep(time.Second * 2)
			if i == len(ints)-1 {
				ctx.Application().Logger().Infof("Loop end.")
				return
			}
			i++
			ctx.ResponseWriter().Flush()
		}
	}
}

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.ContentType("text/html")
		ctx.Header("Transfer-Encoding", "chunked")
		i := 0
		ints := NewInts()
		//以块的形式发送响应，并在每个块之间等待半秒钟
		ctx.StreamWriter(func(w io.Writer) bool {
			fmt.Fprintf(w, "Message number %d<br>", ints[i])
			// time.Sleep(500 * time.Millisecond) // simulate delay.
			time.Sleep(time.Second * 2)
			if i == len(ints)-1 {
				return false //关闭并刷新
			}
			i++
			return true //继续写入数据
		})
	})

	app.Any("/json", JsonHandler)
	app.Get("/html", HtmlHandler)
	app.Get("/foo", foo)
	app.Get("/bar", bar)
	app.Run(iris.Addr(":8080"))
}

func foo(ctx iris.Context) {
	s := strings.NewReader("test a\r\nb test\r\nc test\r\n")
	r := bufio.NewReader(s)
	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
			ctx.WriteString(fmt.Sprintf("Token (ReadLine): %q\n", token))
		}
		if err != nil {
			break
		}
	}
	s.Seek(0, io.SeekStart)
	r.Reset(s)
	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		ctx.WriteString(fmt.Sprintf("Token (ReadBytes): %q\n", token))
		if err != nil {
			break
		}
	}
	s.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
		ctx.WriteString(fmt.Sprintf("Token (Scanner): %q\n", scanner.Text()))
	}
}

func bar(ctx iris.Context) {
	// request body is a stream of bytes.
	// we need to read all bytes from the request body.
	// and write them to the response body.
	api := "http://localhost:8080/json"
	ctx.ContentType("application/json")
	ctx.Header("Transfer-Encoding", "chunked")
	reqData := []byte(`{"name":"wf-test"}`)
	stranBytes, err := json.Marshal(reqData)
	fmt.Println("stranBytes:", stranBytes)
	if err != nil {
		fmt.Println("failed to request HTTP, url: %s, err: %v", api, err)
		return
	}
	resp, err := http.Post(api, "application/json", bytes.NewBuffer(stranBytes))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	// s := strings.NewReader(resp.Body)
	r := bufio.NewReader(resp.Body)
	for {
		// token, _, err := r.ReadLine()
		// if len(token) > 0 {
		// 	fmt.Printf("Token (ReadLine): %q\n", token)
		// 	ctx.WriteString(fmt.Sprintf("Token (ReadLine): %q\n", token))
		// }
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %s", string(token))
		if err != nil {
			break
		}
	}

}
