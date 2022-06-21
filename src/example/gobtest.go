/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-24 09:45:27
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-24 10:45:44
 */
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// A struct with a mix of fields, used for the GOB example.
type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
	E Addr
}

type Addr struct {
	Comment string
}

func main() {

	testStruct := complexData{
		N: 23,
		S: "string data",
		M: map[string]int{"one": 1, "two": 2, "three": 3},
		P: []byte("abc"),
		C: &complexData{
			N: 256,
			S: "Recursive structs? Piece of cake!",
			M: map[string]int{"01": 1, "10": 2, "11": 3},
			E: Addr{
				Comment: "InnerTest123123123123",
			},
		},
		E: Addr{
			Comment: "Test123123123",
		},
	}

	fmt.Println("Outer complexData struct: ", testStruct)
	fmt.Println("Inner complexData struct: ", testStruct.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)
	fmt.Println("===========================")

	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(testStruct)
	if err != nil {
		fmt.Println(err)
	}

	var data complexData
	err = gob.NewDecoder(&b).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding GOB data:", err)
		return
	}

	fmt.Println("Outer complexData struct: ", data)
	fmt.Println("Inner complexData struct: ", data.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)
	foo()
}

type A struct {
	Comment string
}

func foo() {
	var a A
	var err error
	ch := bar()
	for {
		select {
		case buf, ok := <-ch:
			if !ok {
				fmt.Println("close")
				return
			}
			// err = gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(&a)
			err = gob.NewDecoder(buf).Decode(&a)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("----->", a)
		case <-time.After(time.Second * 10):
			fmt.Println("timeout")
			return
		}

	}
}

func bar() chan *bytes.Buffer {
	var a A

	ch := make(chan *bytes.Buffer, 10)
	go func() {
		for i := 0; i < 10; i++ {
			a.Comment = fmt.Sprintf("%d", i)
			buf := bytes.NewBuffer([]byte{})
			err := gob.NewEncoder(buf).Encode(a)
			if err != nil {
				fmt.Println(err)
				return
			}
			ch <- buf
		}
		close(ch)
	}()
	return ch

}
