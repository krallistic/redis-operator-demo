/*
Created by codegen
*/package fake

import (
	v1alpha1 "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/typed/krallistic/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKrallisticV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKrallisticV1alpha1) RedisDBs(namespace string) v1alpha1.RedisDBInterface {
	return &FakeRedisDBs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKrallisticV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
