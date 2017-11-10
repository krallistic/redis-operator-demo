/*
Created by codegen
*/package v1alpha1

import (
	v1alpha1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1alpha1"
	"github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type KrallisticV1alpha1Interface interface {
	RESTClient() rest.Interface
	RedisDBsGetter
}

// KrallisticV1alpha1Client is used to interact with features provided by the krallistic.com group.
type KrallisticV1alpha1Client struct {
	restClient rest.Interface
}

func (c *KrallisticV1alpha1Client) RedisDBs(namespace string) RedisDBInterface {
	return newRedisDBs(c, namespace)
}

// NewForConfig creates a new KrallisticV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*KrallisticV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KrallisticV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new KrallisticV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KrallisticV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KrallisticV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *KrallisticV1alpha1Client {
	return &KrallisticV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
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
func (c *KrallisticV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
