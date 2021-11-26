package kubernetes

import (
	"k8s.io/client-go/tools/clientcmd"
)

func InitK8S()  {
	config, err := clientcmd.BuildConfigFromFlags("","./config/kubeconfig")
	if err != nil {

		return
	}

}