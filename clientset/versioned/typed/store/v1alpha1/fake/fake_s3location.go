/*
Copyright 2021 The Kubesphere Authors.

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
	"context"

	v1alpha1 "github.com/kubesphere/cloud-client/apis/store/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeS3Locations implements S3LocationInterface
type FakeS3Locations struct {
	Fake *FakeStoreV1alpha1
	ns   string
}

var s3locationsResource = schema.GroupVersionResource{Group: "store.kubesphere.cloud", Version: "v1alpha1", Resource: "s3locations"}

var s3locationsKind = schema.GroupVersionKind{Group: "store.kubesphere.cloud", Version: "v1alpha1", Kind: "S3Location"}

// Get takes name of the s3Location, and returns the corresponding s3Location object, and an error if there is any.
func (c *FakeS3Locations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.S3Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(s3locationsResource, c.ns, name), &v1alpha1.S3Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3Location), err
}

// List takes label and field selectors, and returns the list of S3Locations that match those selectors.
func (c *FakeS3Locations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.S3LocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(s3locationsResource, s3locationsKind, c.ns, opts), &v1alpha1.S3LocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.S3LocationList{ListMeta: obj.(*v1alpha1.S3LocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.S3LocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested s3Locations.
func (c *FakeS3Locations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(s3locationsResource, c.ns, opts))

}

// Create takes the representation of a s3Location and creates it.  Returns the server's representation of the s3Location, and an error, if there is any.
func (c *FakeS3Locations) Create(ctx context.Context, s3Location *v1alpha1.S3Location, opts v1.CreateOptions) (result *v1alpha1.S3Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(s3locationsResource, c.ns, s3Location), &v1alpha1.S3Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3Location), err
}

// Update takes the representation of a s3Location and updates it. Returns the server's representation of the s3Location, and an error, if there is any.
func (c *FakeS3Locations) Update(ctx context.Context, s3Location *v1alpha1.S3Location, opts v1.UpdateOptions) (result *v1alpha1.S3Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(s3locationsResource, c.ns, s3Location), &v1alpha1.S3Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3Location), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeS3Locations) UpdateStatus(ctx context.Context, s3Location *v1alpha1.S3Location, opts v1.UpdateOptions) (*v1alpha1.S3Location, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(s3locationsResource, "status", c.ns, s3Location), &v1alpha1.S3Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3Location), err
}

// Delete takes name of the s3Location and deletes it. Returns an error if one occurs.
func (c *FakeS3Locations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(s3locationsResource, c.ns, name), &v1alpha1.S3Location{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeS3Locations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(s3locationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.S3LocationList{})
	return err
}

// Patch applies the patch and returns the patched s3Location.
func (c *FakeS3Locations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.S3Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(s3locationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.S3Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3Location), err
}
