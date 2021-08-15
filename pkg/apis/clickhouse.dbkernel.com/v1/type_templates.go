// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"github.com/imdario/mergo"
)

// NewChiTemplates
func NewChiTemplates() *ChiTemplates {
	return new(ChiTemplates)
}

// Len
func (templates *ChiTemplates) Len() int {
	if templates == nil {
		return 0
	}

	return 0 +
		len(templates.HostTemplates) +
		len(templates.PodTemplates) +
		len(templates.VolumeClaimTemplates) +
		len(templates.ServiceTemplates)
}

// MergeFrom
func (templates *ChiTemplates) MergeFrom(from *ChiTemplates, _type MergeType) *ChiTemplates {
	if from.Len() == 0 {
		return templates
	}

	if templates == nil {
		templates = NewChiTemplates()
	}

	if len(from.HostTemplates) > 0 {
		// We have templates to copy from
		// Append HostTemplates from `from` to receiver
		if templates.HostTemplates == nil {
			templates.HostTemplates = make([]ChiHostTemplate, 0)
		}
		// Loop over all 'from' templates and copy it in case no such template in receiver
		for fromIndex := range from.HostTemplates {
			fromTemplate := &from.HostTemplates[fromIndex]

			// Try to find entry with the same name among local templates in receiver
			sameNameFound := false
			for toIndex := range templates.HostTemplates {
				toTemplate := &templates.HostTemplates[toIndex]
				if toTemplate.Name == fromTemplate.Name {
					// Receiver already have such a template
					sameNameFound = true
					// Override `to` template with `from` template
					//templates.PodTemplates[toIndex] = *fromTemplate.DeepCopy()
					if err := mergo.Merge(toTemplate, *fromTemplate, mergo.WithOverride); err != nil {
						//errs = append(errs, fmt.Errorf("ERROR merge template(%s): %v", toTemplate.Name, err))
					}
					break
				}
			}

			if !sameNameFound {
				// Receiver does not have template with such a name
				// Append template from `from`
				templates.HostTemplates = append(templates.HostTemplates, *fromTemplate.DeepCopy())
			}
		}
	}

	if len(from.PodTemplates) > 0 {
		// We have templates to copy from
		// Append PodTemplates from `from` to receiver
		if templates.PodTemplates == nil {
			templates.PodTemplates = make([]ChiPodTemplate, 0)
		}
		// Loop over all 'from' templates and copy it in case no such template in receiver
		for fromIndex := range from.PodTemplates {
			fromTemplate := &from.PodTemplates[fromIndex]

			// Try to find entry with the same name among local templates in receiver
			sameNameFound := false
			for toIndex := range templates.PodTemplates {
				toTemplate := &templates.PodTemplates[toIndex]
				if toTemplate.Name == fromTemplate.Name {
					// Receiver already have such a template
					sameNameFound = true
					// Override `to` template with `from` template
					//templates.PodTemplates[toIndex] = *fromTemplate.DeepCopy()
					if err := mergo.Merge(toTemplate, *fromTemplate, mergo.WithOverride); err != nil {
						//errs = append(errs, fmt.Errorf("ERROR merge template(%s): %v", toTemplate.Name, err))
					}
					break
				}
			}

			if !sameNameFound {
				// Receiver does not have template with such a name
				// Append template from `from`
				templates.PodTemplates = append(templates.PodTemplates, *fromTemplate.DeepCopy())
			}
		}
	}

	if len(from.VolumeClaimTemplates) > 0 {
		// We have templates to copy from
		// Append VolumeClaimTemplates from `from` to receiver
		if templates.VolumeClaimTemplates == nil {
			templates.VolumeClaimTemplates = make([]ChiVolumeClaimTemplate, 0)
		}
		// Loop over all 'from' templates and copy it in case no such template in receiver
		for fromIndex := range from.VolumeClaimTemplates {
			fromTemplate := &from.VolumeClaimTemplates[fromIndex]

			// Try to find entry with the same name among local templates in receiver
			sameNameFound := false
			for toIndex := range templates.VolumeClaimTemplates {
				toTemplate := &templates.VolumeClaimTemplates[toIndex]
				if toTemplate.Name == fromTemplate.Name {
					// Receiver already have such a template
					sameNameFound = true
					// Override `to` template with `from` template
					//templates.VolumeClaimTemplates[toIndex] = *fromTemplate.DeepCopy()
					if err := mergo.Merge(toTemplate, *fromTemplate, mergo.WithOverride); err != nil {
						//errs = append(errs, fmt.Errorf("ERROR merge template(%s): %v", toTemplate.Name, err))
					}
					break
				}
			}

			if !sameNameFound {
				// Receiver does not have template with such a name
				// Append template from `from`
				templates.VolumeClaimTemplates = append(templates.VolumeClaimTemplates, *fromTemplate.DeepCopy())
			}
		}
	}

	if len(from.ServiceTemplates) > 0 {
		// We have templates to copy from
		// Append ServiceTemplates from `from` to receiver
		if templates.ServiceTemplates == nil {
			templates.ServiceTemplates = make([]ChiServiceTemplate, 0)
		}
		// Loop over all 'from' templates and copy it in case no such template in receiver
		for fromIndex := range from.ServiceTemplates {
			fromTemplate := &from.ServiceTemplates[fromIndex]

			// Try to find entry with the same name among local templates in receiver
			sameNameFound := false
			for toIndex := range templates.ServiceTemplates {
				toTemplate := &templates.ServiceTemplates[toIndex]
				if toTemplate.Name == fromTemplate.Name {
					// Receiver already have such a template
					sameNameFound = true
					// Override `to` template with `from` template
					//templates.ServiceTemplates[toIndex] = *fromTemplate.DeepCopy()
					if err := mergo.Merge(toTemplate, *fromTemplate, mergo.WithOverride); err != nil {
						//errs = append(errs, fmt.Errorf("ERROR merge template(%s): %v", toTemplate.Name, err))
					}
					break
				}
			}

			if !sameNameFound {
				// Receiver does not have template with such a name
				// Append template from `from`
				templates.ServiceTemplates = append(templates.ServiceTemplates, *fromTemplate.DeepCopy())
			}
		}
	}

	return templates
}

// GetHostTemplatesIndex
func (templates *ChiTemplates) GetHostTemplatesIndex() *HostTemplatesIndex {
	if templates == nil {
		return nil
	}
	return templates.HostTemplatesIndex
}

// EnsureHostTemplatesIndex
func (templates *ChiTemplates) EnsureHostTemplatesIndex() *HostTemplatesIndex {
	if templates == nil {
		return nil
	}
	if templates.HostTemplatesIndex != nil {
		return templates.HostTemplatesIndex
	}
	templates.HostTemplatesIndex = NewHostTemplatesIndex()
	return templates.HostTemplatesIndex
}

// GetPodTemplatesIndex
func (templates *ChiTemplates) GetPodTemplatesIndex() *PodTemplatesIndex {
	if templates == nil {
		return nil
	}
	return templates.PodTemplatesIndex
}

// EnsurePodTemplatesIndex
func (templates *ChiTemplates) EnsurePodTemplatesIndex() *PodTemplatesIndex {
	if templates == nil {
		return nil
	}
	if templates.PodTemplatesIndex != nil {
		return templates.PodTemplatesIndex
	}
	templates.PodTemplatesIndex = NewPodTemplatesIndex()
	return templates.PodTemplatesIndex
}

// GetVolumeClaimTemplatesIndex
func (templates *ChiTemplates) GetVolumeClaimTemplatesIndex() *VolumeClaimTemplatesIndex {
	if templates == nil {
		return nil
	}
	return templates.VolumeClaimTemplatesIndex
}

// EnsureVolumeClaimTemplatesIndex
func (templates *ChiTemplates) EnsureVolumeClaimTemplatesIndex() *VolumeClaimTemplatesIndex {
	if templates == nil {
		return nil
	}
	if templates.VolumeClaimTemplatesIndex != nil {
		return templates.VolumeClaimTemplatesIndex
	}
	templates.VolumeClaimTemplatesIndex = NewVolumeClaimTemplatesIndex()
	return templates.VolumeClaimTemplatesIndex
}

// GetServiceTemplatesIndex
func (templates *ChiTemplates) GetServiceTemplatesIndex() *ServiceTemplatesIndex {
	if templates == nil {
		return nil
	}
	return templates.ServiceTemplatesIndex
}

// EnsureServiceTemplatesIndex
func (templates *ChiTemplates) EnsureServiceTemplatesIndex() *ServiceTemplatesIndex {
	if templates == nil {
		return nil
	}
	if templates.ServiceTemplatesIndex != nil {
		return templates.ServiceTemplatesIndex
	}
	templates.ServiceTemplatesIndex = NewServiceTemplatesIndex()
	return templates.ServiceTemplatesIndex
}

// HostTemplatesIndex
type HostTemplatesIndex struct {
	v map[string]*ChiHostTemplate `json:",omitempty" yaml:",omitempty" testdiff:"ignore"`
}

// NewHostTemplatesIndex
func NewHostTemplatesIndex() *HostTemplatesIndex {
	return &HostTemplatesIndex{
		v: make(map[string]*ChiHostTemplate),
	}
}

// Has
func (i *HostTemplatesIndex) Has(name string) bool {
	if i == nil {
		return false
	}
	if i.v == nil {
		return false
	}
	_, ok := i.v[name]
	return ok
}

// Get
func (i *HostTemplatesIndex) Get(name string) *ChiHostTemplate {
	if !i.Has(name) {
		return nil
	}
	return i.v[name]
}

// Set
func (i *HostTemplatesIndex) Set(name string, entry *ChiHostTemplate) {
	if i == nil {
		return
	}
	if i.v == nil {
		return
	}
	i.v[name] = entry
}

// Walk
func (i *HostTemplatesIndex) Walk(f func(template *ChiHostTemplate)) {
	if i == nil {
		return
	}
	for _, entry := range i.v {
		f(entry)
	}
}

// PodTemplatesIndex
type PodTemplatesIndex struct {
	v map[string]*ChiPodTemplate `json:",omitempty" yaml:",omitempty" testdiff:"ignore"`
}

// NewPodTemplatesIndex
func NewPodTemplatesIndex() *PodTemplatesIndex {
	return &PodTemplatesIndex{
		v: make(map[string]*ChiPodTemplate),
	}
}

// Has
func (i *PodTemplatesIndex) Has(name string) bool {
	if i == nil {
		return false
	}
	if i.v == nil {
		return false
	}
	_, ok := i.v[name]
	return ok
}

// Get
func (i *PodTemplatesIndex) Get(name string) *ChiPodTemplate {
	if !i.Has(name) {
		return nil
	}
	return i.v[name]
}

// Set
func (i *PodTemplatesIndex) Set(name string, entry *ChiPodTemplate) {
	if i == nil {
		return
	}
	if i.v == nil {
		return
	}
	i.v[name] = entry
}

// Walk
func (i *PodTemplatesIndex) Walk(f func(template *ChiPodTemplate)) {
	if i == nil {
		return
	}
	for _, entry := range i.v {
		f(entry)
	}
}

// VolumeClaimTemplatesIndex
type VolumeClaimTemplatesIndex struct {
	v map[string]*ChiVolumeClaimTemplate `json:",omitempty" yaml:",omitempty" testdiff:"ignore"`
}

// NewVolumeClaimTemplatesIndex
func NewVolumeClaimTemplatesIndex() *VolumeClaimTemplatesIndex {
	return &VolumeClaimTemplatesIndex{
		v: make(map[string]*ChiVolumeClaimTemplate),
	}
}

// Has
func (i *VolumeClaimTemplatesIndex) Has(name string) bool {
	if i == nil {
		return false
	}
	if i.v == nil {
		return false
	}
	_, ok := i.v[name]
	return ok
}

// Get
func (i *VolumeClaimTemplatesIndex) Get(name string) *ChiVolumeClaimTemplate {
	if !i.Has(name) {
		return nil
	}
	return i.v[name]
}

// Set
func (i *VolumeClaimTemplatesIndex) Set(name string, entry *ChiVolumeClaimTemplate) {
	if i == nil {
		return
	}
	if i.v == nil {
		return
	}
	i.v[name] = entry
}

// Walk
func (i *VolumeClaimTemplatesIndex) Walk(f func(template *ChiVolumeClaimTemplate)) {
	if i == nil {
		return
	}
	for _, entry := range i.v {
		f(entry)
	}
}

// ServiceTemplatesIndex
type ServiceTemplatesIndex struct {
	v map[string]*ChiServiceTemplate `json:",omitempty" yaml:",omitempty" testdiff:"ignore"`
}

// NewServiceTemplatesIndex
func NewServiceTemplatesIndex() *ServiceTemplatesIndex {
	return &ServiceTemplatesIndex{
		v: make(map[string]*ChiServiceTemplate),
	}
}

// Has
func (i *ServiceTemplatesIndex) Has(name string) bool {
	if i == nil {
		return false
	}
	if i.v == nil {
		return false
	}
	_, ok := i.v[name]
	return ok
}

// Get
func (i *ServiceTemplatesIndex) Get(name string) *ChiServiceTemplate {
	if !i.Has(name) {
		return nil
	}
	return i.v[name]
}

// Set
func (i *ServiceTemplatesIndex) Set(name string, entry *ChiServiceTemplate) {
	if i == nil {
		return
	}
	if i.v == nil {
		return
	}
	i.v[name] = entry
}

// Walk
func (i *ServiceTemplatesIndex) Walk(f func(template *ChiServiceTemplate)) {
	if i == nil {
		return
	}
	for _, entry := range i.v {
		f(entry)
	}
}