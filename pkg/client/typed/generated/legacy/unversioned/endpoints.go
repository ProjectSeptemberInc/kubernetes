/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	api "k8s.io/kubernetes/pkg/api"
	watch "k8s.io/kubernetes/pkg/watch"
)

// EndpointsNamespacer has methods to work with Endpoints resources in a namespace
type EndpointsNamespacer interface {
	Endpoints(namespace string) EndpointsInterface
}

// EndpointsInterface has methods to work with Endpoints resources.
type EndpointsInterface interface {
	Create(*api.Endpoints) (*api.Endpoints, error)
	Update(*api.Endpoints) (*api.Endpoints, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*api.Endpoints, error)
	List(opts api.ListOptions) (*api.EndpointsList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
}

// endpoints implements EndpointsInterface
type endpoints struct {
	client *LegacyClient
	ns     string
}

// newEndpoints returns a Endpoints
func newEndpoints(c *LegacyClient, namespace string) *endpoints {
	return &endpoints{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Create(endpoints *api.Endpoints) (result *api.Endpoints, err error) {
	result = &api.Endpoints{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpoints").
		Body(endpoints).
		Do().
		Into(result)
	return
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *endpoints) Update(endpoints *api.Endpoints) (result *api.Endpoints, err error) {
	result = &api.Endpoints{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpoints").
		Name(endpoints.Name).
		Body(endpoints).
		Do().
		Into(result)
	return
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *endpoints) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpoints) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		NamespaceIfScoped(c.ns, len(c.ns) > 0).
		Resource("endpoints").
		VersionedParams(&listOptions, api.Scheme).
		Body(options).
		Do().
		Error()
}

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *endpoints) Get(name string) (result *api.Endpoints, err error) {
	result = &api.Endpoints{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *endpoints) List(opts api.ListOptions) (result *api.EndpointsList, err error) {
	result = &api.EndpointsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, api.Scheme).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *endpoints) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, api.Scheme).
		Watch()
}
