/*
Created by codegen
*/package fake

import (
	v1 "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/typed/krallistic/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKrallisticV1 struct {
	*testing.Fake
}

func (c *FakeKrallisticV1) Redisdbs(namespace string) v1.RedisdbInterface {
	return &FakeRedisdbs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKrallisticV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
