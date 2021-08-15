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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/DBKernel/clickhouse-operator/pkg/apis/clickhouse.dbkernel.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClickHouseOperatorConfigurationLister helps list ClickHouseOperatorConfigurations.
// All objects returned here must be treated as read-only.
type ClickHouseOperatorConfigurationLister interface {
	// List lists all ClickHouseOperatorConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClickHouseOperatorConfiguration, err error)
	// ClickHouseOperatorConfigurations returns an object that can list and get ClickHouseOperatorConfigurations.
	ClickHouseOperatorConfigurations(namespace string) ClickHouseOperatorConfigurationNamespaceLister
	ClickHouseOperatorConfigurationListerExpansion
}

// clickHouseOperatorConfigurationLister implements the ClickHouseOperatorConfigurationLister interface.
type clickHouseOperatorConfigurationLister struct {
	indexer cache.Indexer
}

// NewClickHouseOperatorConfigurationLister returns a new ClickHouseOperatorConfigurationLister.
func NewClickHouseOperatorConfigurationLister(indexer cache.Indexer) ClickHouseOperatorConfigurationLister {
	return &clickHouseOperatorConfigurationLister{indexer: indexer}
}

// List lists all ClickHouseOperatorConfigurations in the indexer.
func (s *clickHouseOperatorConfigurationLister) List(selector labels.Selector) (ret []*v1.ClickHouseOperatorConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClickHouseOperatorConfiguration))
	})
	return ret, err
}

// ClickHouseOperatorConfigurations returns an object that can list and get ClickHouseOperatorConfigurations.
func (s *clickHouseOperatorConfigurationLister) ClickHouseOperatorConfigurations(namespace string) ClickHouseOperatorConfigurationNamespaceLister {
	return clickHouseOperatorConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClickHouseOperatorConfigurationNamespaceLister helps list and get ClickHouseOperatorConfigurations.
// All objects returned here must be treated as read-only.
type ClickHouseOperatorConfigurationNamespaceLister interface {
	// List lists all ClickHouseOperatorConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClickHouseOperatorConfiguration, err error)
	// Get retrieves the ClickHouseOperatorConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClickHouseOperatorConfiguration, error)
	ClickHouseOperatorConfigurationNamespaceListerExpansion
}

// clickHouseOperatorConfigurationNamespaceLister implements the ClickHouseOperatorConfigurationNamespaceLister
// interface.
type clickHouseOperatorConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClickHouseOperatorConfigurations in the indexer for a given namespace.
func (s clickHouseOperatorConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1.ClickHouseOperatorConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClickHouseOperatorConfiguration))
	})
	return ret, err
}

// Get retrieves the ClickHouseOperatorConfiguration from the indexer for a given namespace and name.
func (s clickHouseOperatorConfigurationNamespaceLister) Get(name string) (*v1.ClickHouseOperatorConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clickhouseoperatorconfiguration"), name)
	}
	return obj.(*v1.ClickHouseOperatorConfiguration), nil
}
