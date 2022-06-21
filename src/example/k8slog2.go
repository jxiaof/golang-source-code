/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-11 11:17:45
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-11 11:17:57
 */

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", Kubeconfig-path)
	if err != nil {
		log.Error(err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err.Error())
	}

	podLogOpts := v1.PodLogOptions{}
	req := client.CoreV1().Pods("default").GetLogs("pod-name", &podLogOpts)
	podLogs, err := req.Stream()
	if err != nil {
		log.Error(err.Error())
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		log.Error(err.Error())
	}
	str := buf.String()
	log.Infof("check pod logs:%s", str)
}