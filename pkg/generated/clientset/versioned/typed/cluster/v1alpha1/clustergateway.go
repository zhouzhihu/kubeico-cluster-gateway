/*
Copyright 2021 The KubeVela Authors.

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
	"time"

	v1alpha1 "github.com/zhouzhihu/kubeico-cluster-gateway/apis/cluster/v1alpha1"
	scheme "github.com/zhouzhihu/kubeico-cluster-gateway/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterGatewaysGetter has a method to return a ClusterGatewayInterface.
// A group's client should implement this interface.
type ClusterGatewaysGetter interface {
	ClusterGateways() ClusterGatewayInterface
}

// ClusterGatewayInterface has methods to work with ClusterGateway resources.
type ClusterGatewayInterface interface {
	Create(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.CreateOptions) (*v1alpha1.ClusterGateway, error)
	Update(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.UpdateOptions) (*v1alpha1.ClusterGateway, error)
	UpdateStatus(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.UpdateOptions) (*v1alpha1.ClusterGateway, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterGateway, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterGatewayList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterGateway, err error)
	ClusterGatewayExpansion
}

// clusterGateways implements ClusterGatewayInterface
type clusterGateways struct {
	client rest.Interface
}

// newClusterGateways returns a ClusterGateways
func newClusterGateways(c *ClusterV1alpha1Client) *clusterGateways {
	return &clusterGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterGateway, and returns the corresponding clusterGateway object, and an error if there is any.
func (c *clusterGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterGateway, err error) {
	result = &v1alpha1.ClusterGateway{}
	err = c.client.Get().
		Resource("clustergateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterGateways that match those selectors.
func (c *clusterGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterGatewayList{}
	err = c.client.Get().
		Resource("clustergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterGateways.
func (c *clusterGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterGateway and creates it.  Returns the server's representation of the clusterGateway, and an error, if there is any.
func (c *clusterGateways) Create(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.CreateOptions) (result *v1alpha1.ClusterGateway, err error) {
	result = &v1alpha1.ClusterGateway{}
	err = c.client.Post().
		Resource("clustergateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGateway).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterGateway and updates it. Returns the server's representation of the clusterGateway, and an error, if there is any.
func (c *clusterGateways) Update(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.UpdateOptions) (result *v1alpha1.ClusterGateway, err error) {
	result = &v1alpha1.ClusterGateway{}
	err = c.client.Put().
		Resource("clustergateways").
		Name(clusterGateway.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGateway).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterGateways) UpdateStatus(ctx context.Context, clusterGateway *v1alpha1.ClusterGateway, opts v1.UpdateOptions) (result *v1alpha1.ClusterGateway, err error) {
	result = &v1alpha1.ClusterGateway{}
	err = c.client.Put().
		Resource("clustergateways").
		Name(clusterGateway.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterGateway).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterGateway and deletes it. Returns an error if one occurs.
func (c *clusterGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustergateways").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustergateways").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterGateway.
func (c *clusterGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterGateway, err error) {
	result = &v1alpha1.ClusterGateway{}
	err = c.client.Patch(pt).
		Resource("clustergateways").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
