/*
Created by codegen
*/package fake

import (
	krallistic_com_v1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedisdbs implements RedisdbInterface
type FakeRedisdbs struct {
	Fake *FakeKrallisticV1
	ns   string
}

var redisdbsResource = schema.GroupVersionResource{Group: "krallistic.com", Version: "v1", Resource: "redisdbs"}

var redisdbsKind = schema.GroupVersionKind{Group: "krallistic.com", Version: "v1", Kind: "Redisdb"}

// Get takes name of the redisdb, and returns the corresponding redisdb object, and an error if there is any.
func (c *FakeRedisdbs) Get(name string, options v1.GetOptions) (result *krallistic_com_v1.Redisdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisdbsResource, c.ns, name), &krallistic_com_v1.Redisdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*krallistic_com_v1.Redisdb), err
}

// List takes label and field selectors, and returns the list of Redisdbs that match those selectors.
func (c *FakeRedisdbs) List(opts v1.ListOptions) (result *krallistic_com_v1.RedisdbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisdbsResource, redisdbsKind, c.ns, opts), &krallistic_com_v1.RedisdbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &krallistic_com_v1.RedisdbList{}
	for _, item := range obj.(*krallistic_com_v1.RedisdbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisdbs.
func (c *FakeRedisdbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisdbsResource, c.ns, opts))

}

// Create takes the representation of a redisdb and creates it.  Returns the server's representation of the redisdb, and an error, if there is any.
func (c *FakeRedisdbs) Create(redisdb *krallistic_com_v1.Redisdb) (result *krallistic_com_v1.Redisdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisdbsResource, c.ns, redisdb), &krallistic_com_v1.Redisdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*krallistic_com_v1.Redisdb), err
}

// Update takes the representation of a redisdb and updates it. Returns the server's representation of the redisdb, and an error, if there is any.
func (c *FakeRedisdbs) Update(redisdb *krallistic_com_v1.Redisdb) (result *krallistic_com_v1.Redisdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisdbsResource, c.ns, redisdb), &krallistic_com_v1.Redisdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*krallistic_com_v1.Redisdb), err
}

// Delete takes name of the redisdb and deletes it. Returns an error if one occurs.
func (c *FakeRedisdbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisdbsResource, c.ns, name), &krallistic_com_v1.Redisdb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisdbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisdbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &krallistic_com_v1.RedisdbList{})
	return err
}

// Patch applies the patch and returns the patched redisdb.
func (c *FakeRedisdbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *krallistic_com_v1.Redisdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisdbsResource, c.ns, name, data, subresources...), &krallistic_com_v1.Redisdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*krallistic_com_v1.Redisdb), err
}
