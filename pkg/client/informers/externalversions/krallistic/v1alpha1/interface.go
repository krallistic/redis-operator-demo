/*
Created by codegen
*/
// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/krallistic/redis-operator-demo/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// RedisDBs returns a RedisDBInformer.
	RedisDBs() RedisDBInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// RedisDBs returns a RedisDBInformer.
func (v *version) RedisDBs() RedisDBInformer {
	return &redisDBInformer{factory: v.SharedInformerFactory}
}
