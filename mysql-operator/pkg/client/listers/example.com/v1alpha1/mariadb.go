// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/tangfeixiong/go-to-kubernetes/mysql-operator/pkg/apis/example.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MariadbLister helps list Mariadbs.
type MariadbLister interface {
	// List lists all Mariadbs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Mariadb, err error)
	// Mariadbs returns an object that can list and get Mariadbs.
	Mariadbs(namespace string) MariadbNamespaceLister
	MariadbListerExpansion
}

// mariadbLister implements the MariadbLister interface.
type mariadbLister struct {
	indexer cache.Indexer
}

// NewMariadbLister returns a new MariadbLister.
func NewMariadbLister(indexer cache.Indexer) MariadbLister {
	return &mariadbLister{indexer: indexer}
}

// List lists all Mariadbs in the indexer.
func (s *mariadbLister) List(selector labels.Selector) (ret []*v1alpha1.Mariadb, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mariadb))
	})
	return ret, err
}

// Mariadbs returns an object that can list and get Mariadbs.
func (s *mariadbLister) Mariadbs(namespace string) MariadbNamespaceLister {
	return mariadbNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MariadbNamespaceLister helps list and get Mariadbs.
type MariadbNamespaceLister interface {
	// List lists all Mariadbs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Mariadb, err error)
	// Get retrieves the Mariadb from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Mariadb, error)
	MariadbNamespaceListerExpansion
}

// mariadbNamespaceLister implements the MariadbNamespaceLister
// interface.
type mariadbNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Mariadbs in the indexer for a given namespace.
func (s mariadbNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Mariadb, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Mariadb))
	})
	return ret, err
}

// Get retrieves the Mariadb from the indexer for a given namespace and name.
func (s mariadbNamespaceLister) Get(name string) (*v1alpha1.Mariadb, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mariadb"), name)
	}
	return obj.(*v1alpha1.Mariadb), nil
}
