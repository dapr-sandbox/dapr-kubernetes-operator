/*
Copyright 2023.

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

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/dapr-sandbox/dapr-kubernetes-operator/api/operator/v1alpha1"
	operatorv1alpha1 "github.com/dapr-sandbox/dapr-kubernetes-operator/pkg/client/applyconfiguration/operator/v1alpha1"
	scheme "github.com/dapr-sandbox/dapr-kubernetes-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DaprInstancesGetter has a method to return a DaprInstanceInterface.
// A group's client should implement this interface.
type DaprInstancesGetter interface {
	DaprInstances(namespace string) DaprInstanceInterface
}

// DaprInstanceInterface has methods to work with DaprInstance resources.
type DaprInstanceInterface interface {
	Create(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.CreateOptions) (*v1alpha1.DaprInstance, error)
	Update(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.UpdateOptions) (*v1alpha1.DaprInstance, error)
	UpdateStatus(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.UpdateOptions) (*v1alpha1.DaprInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DaprInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DaprInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DaprInstance, err error)
	Apply(ctx context.Context, daprInstance *operatorv1alpha1.DaprInstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaprInstance, err error)
	ApplyStatus(ctx context.Context, daprInstance *operatorv1alpha1.DaprInstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaprInstance, err error)
	DaprInstanceExpansion
}

// daprInstances implements DaprInstanceInterface
type daprInstances struct {
	client rest.Interface
	ns     string
}

// newDaprInstances returns a DaprInstances
func newDaprInstances(c *OperatorV1alpha1Client, namespace string) *daprInstances {
	return &daprInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the daprInstance, and returns the corresponding daprInstance object, and an error if there is any.
func (c *daprInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DaprInstance, err error) {
	result = &v1alpha1.DaprInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daprinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DaprInstances that match those selectors.
func (c *daprInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DaprInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DaprInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daprinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested daprInstances.
func (c *daprInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("daprinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a daprInstance and creates it.  Returns the server's representation of the daprInstance, and an error, if there is any.
func (c *daprInstances) Create(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.CreateOptions) (result *v1alpha1.DaprInstance, err error) {
	result = &v1alpha1.DaprInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("daprinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daprInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a daprInstance and updates it. Returns the server's representation of the daprInstance, and an error, if there is any.
func (c *daprInstances) Update(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.UpdateOptions) (result *v1alpha1.DaprInstance, err error) {
	result = &v1alpha1.DaprInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daprinstances").
		Name(daprInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daprInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *daprInstances) UpdateStatus(ctx context.Context, daprInstance *v1alpha1.DaprInstance, opts v1.UpdateOptions) (result *v1alpha1.DaprInstance, err error) {
	result = &v1alpha1.DaprInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("daprinstances").
		Name(daprInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(daprInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the daprInstance and deletes it. Returns an error if one occurs.
func (c *daprInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daprinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *daprInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("daprinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched daprInstance.
func (c *daprInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DaprInstance, err error) {
	result = &v1alpha1.DaprInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("daprinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied daprInstance.
func (c *daprInstances) Apply(ctx context.Context, daprInstance *operatorv1alpha1.DaprInstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaprInstance, err error) {
	if daprInstance == nil {
		return nil, fmt.Errorf("daprInstance provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(daprInstance)
	if err != nil {
		return nil, err
	}
	name := daprInstance.Name
	if name == nil {
		return nil, fmt.Errorf("daprInstance.Name must be provided to Apply")
	}
	result = &v1alpha1.DaprInstance{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("daprinstances").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *daprInstances) ApplyStatus(ctx context.Context, daprInstance *operatorv1alpha1.DaprInstanceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaprInstance, err error) {
	if daprInstance == nil {
		return nil, fmt.Errorf("daprInstance provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(daprInstance)
	if err != nil {
		return nil, err
	}

	name := daprInstance.Name
	if name == nil {
		return nil, fmt.Errorf("daprInstance.Name must be provided to Apply")
	}

	result = &v1alpha1.DaprInstance{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("daprinstances").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}