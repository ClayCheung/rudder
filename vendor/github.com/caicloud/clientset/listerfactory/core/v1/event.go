/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
)

var _ v1.EventLister = &eventLister{}

var _ v1.EventNamespaceLister = &eventNamespaceLister{}

// eventLister implements the EventLister interface.
type eventLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEventLister returns a new EventLister.
func NewEventLister(client kubernetes.Interface) v1.EventLister {
	return NewFilteredEventLister(client, nil)
}

func NewFilteredEventLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.EventLister {
	return &eventLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all Events in the indexer.
func (s *eventLister) List(selector labels.Selector) (ret []*corev1.Event, err error) {
	listopt := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoreV1().Events(metav1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Events returns an object that can list and get Events.
func (s *eventLister) Events(namespace string) v1.EventNamespaceLister {
	return eventNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// eventNamespaceLister implements the EventNamespaceLister
// interface.
type eventNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all Events in the indexer for a given namespace.
func (s eventNamespaceLister) List(selector labels.Selector) (ret []*corev1.Event, err error) {
	listopt := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoreV1().Events(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the Event from the indexer for a given namespace and name.
func (s eventNamespaceLister) Get(name string) (*corev1.Event, error) {
	return s.client.CoreV1().Events(s.namespace).Get(name, metav1.GetOptions{})
}