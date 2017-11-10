package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisDB describes a RedisDB.
type RedisDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RedisDBSpec `json:"spec"`
}

// RedisDBSpec is the spec for a Foo resource
type RedisDBSpec struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// RedisDBList is a list of RedisDB resources
type RedisDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RedisDB `json:"items"`
}
