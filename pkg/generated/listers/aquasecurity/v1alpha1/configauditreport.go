// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigAuditReportLister helps list ConfigAuditReports.
// All objects returned here must be treated as read-only.
type ConfigAuditReportLister interface {
	// List lists all ConfigAuditReports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigAuditReport, err error)
	// ConfigAuditReports returns an object that can list and get ConfigAuditReports.
	ConfigAuditReports(namespace string) ConfigAuditReportNamespaceLister
	ConfigAuditReportListerExpansion
}

// configAuditReportLister implements the ConfigAuditReportLister interface.
type configAuditReportLister struct {
	indexer cache.Indexer
}

// NewConfigAuditReportLister returns a new ConfigAuditReportLister.
func NewConfigAuditReportLister(indexer cache.Indexer) ConfigAuditReportLister {
	return &configAuditReportLister{indexer: indexer}
}

// List lists all ConfigAuditReports in the indexer.
func (s *configAuditReportLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigAuditReport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigAuditReport))
	})
	return ret, err
}

// ConfigAuditReports returns an object that can list and get ConfigAuditReports.
func (s *configAuditReportLister) ConfigAuditReports(namespace string) ConfigAuditReportNamespaceLister {
	return configAuditReportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigAuditReportNamespaceLister helps list and get ConfigAuditReports.
// All objects returned here must be treated as read-only.
type ConfigAuditReportNamespaceLister interface {
	// List lists all ConfigAuditReports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigAuditReport, err error)
	// Get retrieves the ConfigAuditReport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConfigAuditReport, error)
	ConfigAuditReportNamespaceListerExpansion
}

// configAuditReportNamespaceLister implements the ConfigAuditReportNamespaceLister
// interface.
type configAuditReportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigAuditReports in the indexer for a given namespace.
func (s configAuditReportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigAuditReport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigAuditReport))
	})
	return ret, err
}

// Get retrieves the ConfigAuditReport from the indexer for a given namespace and name.
func (s configAuditReportNamespaceLister) Get(name string) (*v1alpha1.ConfigAuditReport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configauditreport"), name)
	}
	return obj.(*v1alpha1.ConfigAuditReport), nil
}
