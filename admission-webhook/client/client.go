package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"strings"
)

var str = `{
  "apiVersion": "admission.k8s.io/v1",
  "kind": "AdmissionReview",
  "request": {
    "uid": "705ab4f5-6393-11e8-b7cc-42010a800002",
    "kind": {"group":"","version":"v1","kind":"pods"},
    "resource": {"group":"","version":"v1","resource":"pods"},
    "name": "test",
    "namespace": "default",
    "operation": "CREATE",
    "userInfo": {
      "username": "admin",
      "uid": "014fbff9a07c",
      "groups": ["system:authenticated","my-admin-group"],
      "extra": {
        "some-key":["some-value1", "some-value2"]
      }
    },
    "object": {"apiVersion":"v1","kind":"Pod","metadata":{"name":"123","namespace":"default"}},
    "dryRun": false
  }
}`

func main() {
	cfg := rest.Config{Host: "http://localhost:8080"}
	clientset, err := kubernetes.NewForConfig(&cfg)
	if err != nil {
		klog.Error(err)
		return
	}
	res := clientset.AdmissionregistrationV1().RESTClient().Post().Body(strings.NewReader(str)).Do(context.Background())
	raw, err := res.Raw()
	if err != nil {
		klog.Error(err)
		return
	}
	//klog.V(7).Infoln(raw)
	fmt.Println(string(raw))
}
