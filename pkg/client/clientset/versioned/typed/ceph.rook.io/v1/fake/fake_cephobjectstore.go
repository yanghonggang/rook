/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	cephrookiov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephObjectStores implements CephObjectStoreInterface
type FakeCephObjectStores struct {
	Fake *FakeCephV1
	ns   string
}

var cephobjectstoresResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephobjectstores"}

var cephobjectstoresKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1", Kind: "CephObjectStore"}

// Get takes name of the cephObjectStore, and returns the corresponding cephObjectStore object, and an error if there is any.
func (c *FakeCephObjectStores) Get(name string, options v1.GetOptions) (result *cephrookiov1.CephObjectStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephobjectstoresResource, c.ns, name), &cephrookiov1.CephObjectStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephObjectStore), err
}

// List takes label and field selectors, and returns the list of CephObjectStores that match those selectors.
func (c *FakeCephObjectStores) List(opts v1.ListOptions) (result *cephrookiov1.CephObjectStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephobjectstoresResource, cephobjectstoresKind, c.ns, opts), &cephrookiov1.CephObjectStoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cephrookiov1.CephObjectStoreList{ListMeta: obj.(*cephrookiov1.CephObjectStoreList).ListMeta}
	for _, item := range obj.(*cephrookiov1.CephObjectStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephObjectStores.
func (c *FakeCephObjectStores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephobjectstoresResource, c.ns, opts))

}

// Create takes the representation of a cephObjectStore and creates it.  Returns the server's representation of the cephObjectStore, and an error, if there is any.
func (c *FakeCephObjectStores) Create(cephObjectStore *cephrookiov1.CephObjectStore) (result *cephrookiov1.CephObjectStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephobjectstoresResource, c.ns, cephObjectStore), &cephrookiov1.CephObjectStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephObjectStore), err
}

// Update takes the representation of a cephObjectStore and updates it. Returns the server's representation of the cephObjectStore, and an error, if there is any.
func (c *FakeCephObjectStores) Update(cephObjectStore *cephrookiov1.CephObjectStore) (result *cephrookiov1.CephObjectStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephobjectstoresResource, c.ns, cephObjectStore), &cephrookiov1.CephObjectStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephObjectStore), err
}

// Delete takes name of the cephObjectStore and deletes it. Returns an error if one occurs.
func (c *FakeCephObjectStores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cephobjectstoresResource, c.ns, name), &cephrookiov1.CephObjectStore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephObjectStores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephobjectstoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cephrookiov1.CephObjectStoreList{})
	return err
}

// Patch applies the patch and returns the patched cephObjectStore.
func (c *FakeCephObjectStores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cephrookiov1.CephObjectStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephobjectstoresResource, c.ns, name, data, subresources...), &cephrookiov1.CephObjectStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephObjectStore), err
}