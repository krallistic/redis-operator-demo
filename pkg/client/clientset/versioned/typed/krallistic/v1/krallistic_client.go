/*
Created by codegen
*/package v1

import (
	v1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1"
	"github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type KrallisticV1Interface interface {
	RESTClient() rest.Interface
	RedisdbsGetter
}

// KrallisticV1Client is used to interact with features provided by the krallistic.com group.
type KrallisticV1Client struct {
	restClient rest.Interface
}

func (c *KrallisticV1Client) Redisdbs(namespace string) RedisdbInterface {
	return newRedisdbs(c, namespace)
}

// NewForConfig creates a new KrallisticV1Client for the given config.
func NewForConfig(c *rest.Config) (*KrallisticV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KrallisticV1Client{client}, nil
}

// NewForConfigOrDie creates a new KrallisticV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KrallisticV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KrallisticV1Client for the given RESTClient.
func New(c rest.Interface) *KrallisticV1Client {
	return &KrallisticV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KrallisticV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
