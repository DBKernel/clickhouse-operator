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

package v1

import (
	"context"
	"time"

	v1 "github.com/DBKernel/clickhouse-operator/pkg/apis/clickhouse.dbkernel.com/v1"
	scheme "github.com/DBKernel/clickhouse-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClickHouseInstallationsGetter has a method to return a ClickHouseInstallationInterface.
// A group's client should implement this interface.
type ClickHouseInstallationsGetter interface {
	ClickHouseInstallations(namespace string) ClickHouseInstallationInterface
}

// ClickHouseInstallationInterface has methods to work with ClickHouseInstallation resources.
type ClickHouseInstallationInterface interface {
	Create(ctx context.Context, clickHouseInstallation *v1.ClickHouseInstallation, opts metav1.CreateOptions) (*v1.ClickHouseInstallation, error)
	Update(ctx context.Context, clickHouseInstallation *v1.ClickHouseInstallation, opts metav1.UpdateOptions) (*v1.ClickHouseInstallation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClickHouseInstallation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClickHouseInstallationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClickHouseInstallation, err error)
	ClickHouseInstallationExpansion
}

// clickHouseInstallations implements ClickHouseInstallationInterface
type clickHouseInstallations struct {
	client rest.Interface
	ns     string
}

// newClickHouseInstallations returns a ClickHouseInstallations
func newClickHouseInstallations(c *ClickhouseV1Client, namespace string) *clickHouseInstallations {
	return &clickHouseInstallations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clickHouseInstallation, and returns the corresponding clickHouseInstallation object, and an error if there is any.
func (c *clickHouseInstallations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClickHouseInstallations that match those selectors.
func (c *clickHouseInstallations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClickHouseInstallationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClickHouseInstallationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clickHouseInstallations.
func (c *clickHouseInstallations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clickHouseInstallation and creates it.  Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *clickHouseInstallations) Create(ctx context.Context, clickHouseInstallation *v1.ClickHouseInstallation, opts metav1.CreateOptions) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clickHouseInstallation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clickHouseInstallation and updates it. Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *clickHouseInstallations) Update(ctx context.Context, clickHouseInstallation *v1.ClickHouseInstallation, opts metav1.UpdateOptions) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(clickHouseInstallation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clickHouseInstallation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clickHouseInstallation and deletes it. Returns an error if one occurs.
func (c *clickHouseInstallations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clickHouseInstallations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clickHouseInstallation.
func (c *clickHouseInstallations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
