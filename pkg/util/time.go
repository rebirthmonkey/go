package util

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func Now() *metav1.Time {
	now := metav1.Now()
	return &now
}
