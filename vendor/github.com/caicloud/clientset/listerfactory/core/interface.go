/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package core

import (
	v1 "github.com/caicloud/clientset/listerfactory/core/v1"
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to listers for resources in V1.
	V1() v1.Interface
}

type group struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type informerGroup struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &informerGroup{factory: factory}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.client, g.tweakListOptions)
}

func (g *informerGroup) V1() v1.Interface {
	return v1.NewFrom(g.factory)
}