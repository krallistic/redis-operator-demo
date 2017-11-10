/*
Created by codegen
*/package fake

import (
	v1alpha1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedisDBs implements RedisDBInterface
type FakeRedisDBs struct {
	Fake *FakeKrallisticV1alpha1
	ns   string
}

var redisdbsResource = schema.GroupVersionResource{Group: "krallistic.com", Version: "v1alpha1", Resource: "redisdbs"}

var redisdbsKind = schema.GroupVersionKind{Group: "krallistic.com", Version: "v1alpha1", Kind: "RedisDB"}

// Get takes name of the redisDB, and returns the corresponding redisDB object, and an error if there is any.
func (c *FakeRedisDBs) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisDB, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisdbsResource, c.ns, name), &v1alpha1.RedisDB{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisDB), err
}

// List takes label and field selectors, and returns the list of RedisDBs that match those selectors.
func (c *FakeRedisDBs) List(opts v1.ListOptions) (result *v1alpha1.RedisDBList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisdbsResource, redisdbsKind, c.ns, opts), &v1alpha1.RedisDBList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisDBList{}
	for _, item := range obj.(*v1alpha1.RedisDBList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisDBs.
func (c *FakeRedisDBs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisdbsResource, c.ns, opts))

}

// Create takes the representation of a redisDB and creates it.  Returns the server's representation of the redisDB, and an error, if there is any.
func (c *FakeRedisDBs) Create(redisDB *v1alpha1.RedisDB) (result *v1alpha1.RedisDB, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisdbsResource, c.ns, redisDB), &v1alpha1.RedisDB{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisDB), err
}

// Update takes the representation of a redisDB and updates it. Returns the server's representation of the redisDB, and an error, if there is any.
func (c *FakeRedisDBs) Update(redisDB *v1alpha1.RedisDB) (result *v1alpha1.RedisDB, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisdbsResource, c.ns, redisDB), &v1alpha1.RedisDB{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisDB), err
}

// Delete takes name of the redisDB and deletes it. Returns an error if one occurs.
func (c *FakeRedisDBs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisdbsResource, c.ns, name), &v1alpha1.RedisDB{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisDBs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisdbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisDBList{})
	return err
}

// Patch applies the patch and returns the patched redisDB.
func (c *FakeRedisDBs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisDB, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisdbsResource, c.ns, name, data, subresources...), &v1alpha1.RedisDB{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisDB), err
}
