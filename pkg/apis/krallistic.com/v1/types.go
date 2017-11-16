package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=redisdb

// RedisDB describes a RedisDB.
type Redisdb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RedisdbSpec `json:"spec"`
}

// RedisDBSpec is the spec for a Foo resource
type RedisdbSpec struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Replicas *int32 `json:"replicas"`
	Name     string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=redisdbs

// RedisDBList is a list of RedisDB resources
type RedisdbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Redisdb `json:"items"`
}
