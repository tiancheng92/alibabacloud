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
	"context"
	v1 "github.com/tiancheng92/alibabacloud/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlbConfigs implements AlbConfigInterface
type FakeAlbConfigs struct {
	Fake *FakeAlibabacloudV1
	ns   string
}

var albconfigsResource = v1.SchemeGroupVersion.WithResource("albconfigs")

var albconfigsKind = v1.SchemeGroupVersion.WithKind("AlbConfig")

// Get takes name of the albConfig, and returns the corresponding albConfig object, and an error if there is any.
func (c *FakeAlbConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AlbConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(albconfigsResource, c.ns, name), &v1.AlbConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AlbConfig), err
}

// List takes label and field selectors, and returns the list of AlbConfigs that match those selectors.
func (c *FakeAlbConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AlbConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(albconfigsResource, albconfigsKind, c.ns, opts), &v1.AlbConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.AlbConfigList{ListMeta: obj.(*v1.AlbConfigList).ListMeta}
	for _, item := range obj.(*v1.AlbConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested albConfigs.
func (c *FakeAlbConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(albconfigsResource, c.ns, opts))

}

// Create takes the representation of a albConfig and creates it.  Returns the server's representation of the albConfig, and an error, if there is any.
func (c *FakeAlbConfigs) Create(ctx context.Context, albConfig *v1.AlbConfig, opts metav1.CreateOptions) (result *v1.AlbConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(albconfigsResource, c.ns, albConfig), &v1.AlbConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AlbConfig), err
}

// Update takes the representation of a albConfig and updates it. Returns the server's representation of the albConfig, and an error, if there is any.
func (c *FakeAlbConfigs) Update(ctx context.Context, albConfig *v1.AlbConfig, opts metav1.UpdateOptions) (result *v1.AlbConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(albconfigsResource, c.ns, albConfig), &v1.AlbConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AlbConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlbConfigs) UpdateStatus(ctx context.Context, albConfig *v1.AlbConfig, opts metav1.UpdateOptions) (*v1.AlbConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(albconfigsResource, "status", c.ns, albConfig), &v1.AlbConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AlbConfig), err
}

// Delete takes name of the albConfig and deletes it. Returns an error if one occurs.
func (c *FakeAlbConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(albconfigsResource, c.ns, name, opts), &v1.AlbConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlbConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(albconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.AlbConfigList{})
	return err
}

// Patch applies the patch and returns the patched albConfig.
func (c *FakeAlbConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AlbConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(albconfigsResource, c.ns, name, pt, data, subresources...), &v1.AlbConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.AlbConfig), err
}
