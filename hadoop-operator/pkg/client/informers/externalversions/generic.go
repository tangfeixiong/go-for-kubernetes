// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/hadoop-operator/pkg/apis/example.com/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=example.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("hadoops"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().Hadoops().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("hadoophdfses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().HadoopHdfses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("hadoopmapreduces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().HadoopMapreduces().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("hadoopnodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().HadoopNodes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("hadoopyarns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().HadoopYarns().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}