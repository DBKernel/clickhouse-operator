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

	clickhousedbkernelcomv1 "github.com/DBKernel/clickhouse-operator/pkg/apis/clickhouse.dbkernel.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClickHouseOperatorConfigurations implements ClickHouseOperatorConfigurationInterface
type FakeClickHouseOperatorConfigurations struct {
	Fake *FakeClickhouseV1
	ns   string
}

var clickhouseoperatorconfigurationsResource = schema.GroupVersionResource{Group: "clickhouse.dbkernel.com", Version: "v1", Resource: "clickhouseoperatorconfigurations"}

var clickhouseoperatorconfigurationsKind = schema.GroupVersionKind{Group: "clickhouse.dbkernel.com", Version: "v1", Kind: "ClickHouseOperatorConfiguration"}

// Get takes name of the clickHouseOperatorConfiguration, and returns the corresponding clickHouseOperatorConfiguration object, and an error if there is any.
func (c *FakeClickHouseOperatorConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clickhouseoperatorconfigurationsResource, c.ns, name), &clickhousedbkernelcomv1.ClickHouseOperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfiguration), err
}

// List takes label and field selectors, and returns the list of ClickHouseOperatorConfigurations that match those selectors.
func (c *FakeClickHouseOperatorConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clickhouseoperatorconfigurationsResource, clickhouseoperatorconfigurationsKind, c.ns, opts), &clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList{ListMeta: obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList).ListMeta}
	for _, item := range obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clickHouseOperatorConfigurations.
func (c *FakeClickHouseOperatorConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clickhouseoperatorconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a clickHouseOperatorConfiguration and creates it.  Returns the server's representation of the clickHouseOperatorConfiguration, and an error, if there is any.
func (c *FakeClickHouseOperatorConfigurations) Create(ctx context.Context, clickHouseOperatorConfiguration *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, opts v1.CreateOptions) (result *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clickhouseoperatorconfigurationsResource, c.ns, clickHouseOperatorConfiguration), &clickhousedbkernelcomv1.ClickHouseOperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfiguration), err
}

// Update takes the representation of a clickHouseOperatorConfiguration and updates it. Returns the server's representation of the clickHouseOperatorConfiguration, and an error, if there is any.
func (c *FakeClickHouseOperatorConfigurations) Update(ctx context.Context, clickHouseOperatorConfiguration *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, opts v1.UpdateOptions) (result *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clickhouseoperatorconfigurationsResource, c.ns, clickHouseOperatorConfiguration), &clickhousedbkernelcomv1.ClickHouseOperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfiguration), err
}

// Delete takes name of the clickHouseOperatorConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeClickHouseOperatorConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clickhouseoperatorconfigurationsResource, c.ns, name), &clickhousedbkernelcomv1.ClickHouseOperatorConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClickHouseOperatorConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clickhouseoperatorconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clickhousedbkernelcomv1.ClickHouseOperatorConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched clickHouseOperatorConfiguration.
func (c *FakeClickHouseOperatorConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clickhousedbkernelcomv1.ClickHouseOperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clickhouseoperatorconfigurationsResource, c.ns, name, pt, data, subresources...), &clickhousedbkernelcomv1.ClickHouseOperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhousedbkernelcomv1.ClickHouseOperatorConfiguration), err
}
