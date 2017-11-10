/*
Created by codegen
*/package v1alpha1

import (
	v1alpha1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1alpha1"
	scheme "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisDBsGetter has a method to return a RedisDBInterface.
// A group's client should implement this interface.
type RedisDBsGetter interface {
	RedisDBs(namespace string) RedisDBInterface
}

// RedisDBInterface has methods to work with RedisDB resources.
type RedisDBInterface interface {
	Create(*v1alpha1.RedisDB) (*v1alpha1.RedisDB, error)
	Update(*v1alpha1.RedisDB) (*v1alpha1.RedisDB, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RedisDB, error)
	List(opts v1.ListOptions) (*v1alpha1.RedisDBList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisDB, err error)
	RedisDBExpansion
}

// redisDBs implements RedisDBInterface
type redisDBs struct {
	client rest.Interface
	ns     string
}

// newRedisDBs returns a RedisDBs
func newRedisDBs(c *KrallisticV1alpha1Client, namespace string) *redisDBs {
	return &redisDBs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisDB, and returns the corresponding redisDB object, and an error if there is any.
func (c *redisDBs) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisDB, err error) {
	result = &v1alpha1.RedisDB{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisDBs that match those selectors.
func (c *redisDBs) List(opts v1.ListOptions) (result *v1alpha1.RedisDBList, err error) {
	result = &v1alpha1.RedisDBList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisDBs.
func (c *redisDBs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a redisDB and creates it.  Returns the server's representation of the redisDB, and an error, if there is any.
func (c *redisDBs) Create(redisDB *v1alpha1.RedisDB) (result *v1alpha1.RedisDB, err error) {
	result = &v1alpha1.RedisDB{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisdbs").
		Body(redisDB).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisDB and updates it. Returns the server's representation of the redisDB, and an error, if there is any.
func (c *redisDBs) Update(redisDB *v1alpha1.RedisDB) (result *v1alpha1.RedisDB, err error) {
	result = &v1alpha1.RedisDB{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(redisDB.Name).
		Body(redisDB).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisDB and deletes it. Returns an error if one occurs.
func (c *redisDBs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisdbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisDBs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisdbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisDB.
func (c *redisDBs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisDB, err error) {
	result = &v1alpha1.RedisDB{}
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
