package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myit.fun/kubernetes/admission-webhook/utils"
	"net/http"

	"k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

var (
	cert = "/opt/certs/tls.crt"
	key  = "/opt/certs/tls.key"
)

func serve(w http.ResponseWriter, r *http.Request, admit utils.AdmitFunc) {
	klog.Info(r.RequestURI)
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}
	klog.Info(fmt.Sprintf("handling request: %s", body))

	reqAdmissionReview := v1.AdmissionReview{}                          // 请求
	resAdmissionReview := v1.AdmissionReview{TypeMeta: metav1.TypeMeta{ // 响应
		Kind:       "AdmissionReview",
		APIVersion: "admission.k8s.io/v1",
	}}

	deserializer := utils.Codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(body, nil, &reqAdmissionReview); err != nil {
		klog.Error(err)
		resAdmissionReview.Response = utils.ToAdmissionResponse(err)
	} else {
		// pass to admitFunc
		resAdmissionReview.Response = admit(reqAdmissionReview) // 业务逻辑
	}

	// 以下是固定写法
	resAdmissionReview.Response.UID = reqAdmissionReview.Request.UID

	klog.V(2).Info(fmt.Sprintf("sending response: %v", resAdmissionReview.Response))

	respBytes, err := json.Marshal(resAdmissionReview)
	if err != nil {
		klog.Error(err)
	}
	if _, err := w.Write(respBytes); err != nil {
		klog.Error(err)
	}
}

func servePods(w http.ResponseWriter, r *http.Request) {
	serve(w, r, utils.AdmitPods)
}

func main() {

	http.HandleFunc("/", servePods)

	//http.ListenAndServe(":8080", nil)

	err := http.ListenAndServeTLS(":8080", cert, key, nil)
	if err != nil {
		klog.Fatal(err)
		return
	}

}
