/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/storage/v1beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // CSIDrivers returns a CSIDriverLister
	CSIDrivers() v1beta1.CSIDriverLister
	// CSINodes returns a CSINodeLister
	CSINodes() v1beta1.CSINodeLister
	// StorageClasses returns a StorageClassLister
	StorageClasses() v1beta1.StorageClassLister
	// VolumeAttachments returns a VolumeAttachmentLister
	VolumeAttachments() v1beta1.VolumeAttachmentLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// CSIDrivers returns a CSIDriverLister.
func (v *version) CSIDrivers() v1beta1.CSIDriverLister {
	return &cSIDriverLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// CSIDrivers returns a CSIDriverLister.
func (v *infromerVersion) CSIDrivers() v1beta1.CSIDriverLister {
	return v.factory.Storage().V1beta1().CSIDrivers().Lister()
}

// CSINodes returns a CSINodeLister.
func (v *version) CSINodes() v1beta1.CSINodeLister {
	return &cSINodeLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// CSINodes returns a CSINodeLister.
func (v *infromerVersion) CSINodes() v1beta1.CSINodeLister {
	return v.factory.Storage().V1beta1().CSINodes().Lister()
}

// StorageClasses returns a StorageClassLister.
func (v *version) StorageClasses() v1beta1.StorageClassLister {
	return &storageClassLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// StorageClasses returns a StorageClassLister.
func (v *infromerVersion) StorageClasses() v1beta1.StorageClassLister {
	return v.factory.Storage().V1beta1().StorageClasses().Lister()
}

// VolumeAttachments returns a VolumeAttachmentLister.
func (v *version) VolumeAttachments() v1beta1.VolumeAttachmentLister {
	return &volumeAttachmentLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// VolumeAttachments returns a VolumeAttachmentLister.
func (v *infromerVersion) VolumeAttachments() v1beta1.VolumeAttachmentLister {
	return v.factory.Storage().V1beta1().VolumeAttachments().Lister()
}