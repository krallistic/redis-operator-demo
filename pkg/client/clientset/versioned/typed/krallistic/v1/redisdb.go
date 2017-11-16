/*
Created by codegen
*/package v1

import (
	v1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1"
	scheme "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisdbsGetter has a method to return a RedisdbInterface.
// A group's client should implement this interface.
type RedisdbsGetter interface {
	Redisdbs(namespace string) RedisdbInterface
}

// RedisdbInterface has methods to work with Redisdb resources.
type RedisdbInterface interface {
	Create(*v1.Redisdb) (*v1.Redisdb, error)
	Update(*v1.Redisdb) (*v1.Redisdb, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Redisdb, error)
	List(opts meta_v1.ListOptions) (*v1.RedisdbList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Redisdb, err error)
	RedisdbExpansion
}

// redisdbs implements RedisdbInterface
type redisdbs struct {
	client rest.Interface
	ns     string
}

// newRedisdbs returns a Redisdbs
func newRedisdbs(c *KrallisticV1Client, namespace string) *redisdbs {
	return &redisdbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisdb, and returns the corresponding redisdb object, and an error if there is any.
func (c *redisdbs) Get(name string, options meta_v1.GetOptions) (result *v1.Redisdb, err error) {
	result = &v1.Redisdb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Redisdbs that match those selectors.
func (c *redisdbs) List(opts meta_v1.ListOptions) (result *v1.RedisdbList, err error) {
	result = &v1.RedisdbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisdbs.
func (c *redisdbs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a redisdb and creates it.  Returns the server's representation of the redisdb, and an error, if there is any.
func (c *redisdbs) Create(redisdb *v1.Redisdb) (result *v1.Redisdb, err error) {
	result = &v1.Redisdb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisdbs").
		Body(redisdb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisdb and updates it. Returns the server's representation of the redisdb, and an error, if there is any.
func (c *redisdbs) Update(redisdb *v1.Redisdb) (result *v1.Redisdb, err error) {
	result = &v1.Redisdb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(redisdb.Name).
		Body(redisdb).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisdb and deletes it. Returns an error if one occurs.
func (c *redisdbs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisdbs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisdb.
func (c *redisdbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Redisdb, err error) {
	result = &v1.Redisdb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisdbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
