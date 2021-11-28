package utils

import (
	"fmt"
	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func AdmitPods(ar v1.AdmissionReview) *v1.AdmissionResponse {
	klog.Info("admitting pods")
	podResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"} // pod
	if ar.Request.Resource != podResource {
		err := fmt.Errorf("expect resource to be %s", podResource)
		klog.Error(err)
		return ToAdmissionResponse(err)
	}

	raw := ar.Request.Object.Raw
	pod := corev1.Pod{}
	deserializer := Codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
		klog.Error(err)
		return ToAdmissionResponse(err)
	}
	klog.Info(pod)

	reviewResponse := v1.AdmissionResponse{}
	// 定义准入规则
	if pod.ObjectMeta.Name != "hsx" {
		reviewResponse.Allowed = false
		reviewResponse.Result = &metav1.Status{Code: 403, Message: "pod name must be hsx"}
	} else {
		reviewResponse.Allowed = true
	}

	return &reviewResponse
}
