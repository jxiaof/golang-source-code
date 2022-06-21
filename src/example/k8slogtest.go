/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-11 11:12:51
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-11 11:13:26
 */
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func log() error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	opts := &v1.PodLogOptions{
		Follow: true, // 对应kubectl logs -f参数
	}
	request := clientset.CoreV1().Pods("default").GetLogs("your pod name", opts)
	readCloser, err := request.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer readCloser.Close()

	r := bufio.NewReader(readCloser)
	for {
		bytes, err := r.ReadBytes('\n')
		fmt.Println(string(bytes))
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
	return nil
}

func main() {
	log()
}
