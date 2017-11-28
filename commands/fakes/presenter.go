// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/models"
)

type Presenter struct {
	PresentAvailableProductsStub        func([]models.Product)
	presentAvailableProductsMutex       sync.RWMutex
	presentAvailableProductsArgsForCall []struct {
		arg1 []models.Product
	}
	PresentCertificateAuthoritiesStub        func([]api.CA)
	presentCertificateAuthoritiesMutex       sync.RWMutex
	presentCertificateAuthoritiesArgsForCall []struct {
		arg1 []api.CA
	}
	PresentCredentialReferencesStub        func([]string)
	presentCredentialReferencesMutex       sync.RWMutex
	presentCredentialReferencesArgsForCall []struct {
		arg1 []string
	}
	PresentCredentialsStub        func(map[string]string)
	presentCredentialsMutex       sync.RWMutex
	presentCredentialsArgsForCall []struct {
		arg1 map[string]string
	}
	PresentDeployedProductsStub        func([]api.DiagnosticProduct)
	presentDeployedProductsMutex       sync.RWMutex
	presentDeployedProductsArgsForCall []struct {
		arg1 []api.DiagnosticProduct
	}
	PresentErrandsStub        func([]models.Errand)
	presentErrandsMutex       sync.RWMutex
	presentErrandsArgsForCall []struct {
		arg1 []models.Errand
	}
	PresentCertificateAuthorityStub        func(api.CA)
	presentCertificateAuthorityMutex       sync.RWMutex
	presentCertificateAuthorityArgsForCall []struct {
		arg1 api.CA
	}
	PresentInstallationsStub        func([]models.Installation)
	presentInstallationsMutex       sync.RWMutex
	presentInstallationsArgsForCall []struct {
		arg1 []models.Installation
	}
	PresentPendingChangesStub        func([]api.ProductChange)
	presentPendingChangesMutex       sync.RWMutex
	presentPendingChangesArgsForCall []struct {
		arg1 []api.ProductChange
	}
	PresentStagedProductsStub        func([]api.DiagnosticProduct)
	presentStagedProductsMutex       sync.RWMutex
	presentStagedProductsArgsForCall []struct {
		arg1 []api.DiagnosticProduct
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Presenter) PresentAvailableProducts(arg1 []models.Product) {
	var arg1Copy []models.Product
	if arg1 != nil {
		arg1Copy = make([]models.Product, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentAvailableProductsMutex.Lock()
	fake.presentAvailableProductsArgsForCall = append(fake.presentAvailableProductsArgsForCall, struct {
		arg1 []models.Product
	}{arg1Copy})
	fake.recordInvocation("PresentAvailableProducts", []interface{}{arg1Copy})
	fake.presentAvailableProductsMutex.Unlock()
	if fake.PresentAvailableProductsStub != nil {
		fake.PresentAvailableProductsStub(arg1)
	}
}

func (fake *Presenter) PresentAvailableProductsCallCount() int {
	fake.presentAvailableProductsMutex.RLock()
	defer fake.presentAvailableProductsMutex.RUnlock()
	return len(fake.presentAvailableProductsArgsForCall)
}

func (fake *Presenter) PresentAvailableProductsArgsForCall(i int) []models.Product {
	fake.presentAvailableProductsMutex.RLock()
	defer fake.presentAvailableProductsMutex.RUnlock()
	return fake.presentAvailableProductsArgsForCall[i].arg1
}

func (fake *Presenter) PresentCertificateAuthorities(arg1 []api.CA) {
	var arg1Copy []api.CA
	if arg1 != nil {
		arg1Copy = make([]api.CA, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentCertificateAuthoritiesMutex.Lock()
	fake.presentCertificateAuthoritiesArgsForCall = append(fake.presentCertificateAuthoritiesArgsForCall, struct {
		arg1 []api.CA
	}{arg1Copy})
	fake.recordInvocation("PresentCertificateAuthorities", []interface{}{arg1Copy})
	fake.presentCertificateAuthoritiesMutex.Unlock()
	if fake.PresentCertificateAuthoritiesStub != nil {
		fake.PresentCertificateAuthoritiesStub(arg1)
	}
}

func (fake *Presenter) PresentCertificateAuthoritiesCallCount() int {
	fake.presentCertificateAuthoritiesMutex.RLock()
	defer fake.presentCertificateAuthoritiesMutex.RUnlock()
	return len(fake.presentCertificateAuthoritiesArgsForCall)
}

func (fake *Presenter) PresentCertificateAuthoritiesArgsForCall(i int) []api.CA {
	fake.presentCertificateAuthoritiesMutex.RLock()
	defer fake.presentCertificateAuthoritiesMutex.RUnlock()
	return fake.presentCertificateAuthoritiesArgsForCall[i].arg1
}

func (fake *Presenter) PresentCredentialReferences(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentCredentialReferencesMutex.Lock()
	fake.presentCredentialReferencesArgsForCall = append(fake.presentCredentialReferencesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("PresentCredentialReferences", []interface{}{arg1Copy})
	fake.presentCredentialReferencesMutex.Unlock()
	if fake.PresentCredentialReferencesStub != nil {
		fake.PresentCredentialReferencesStub(arg1)
	}
}

func (fake *Presenter) PresentCredentialReferencesCallCount() int {
	fake.presentCredentialReferencesMutex.RLock()
	defer fake.presentCredentialReferencesMutex.RUnlock()
	return len(fake.presentCredentialReferencesArgsForCall)
}

func (fake *Presenter) PresentCredentialReferencesArgsForCall(i int) []string {
	fake.presentCredentialReferencesMutex.RLock()
	defer fake.presentCredentialReferencesMutex.RUnlock()
	return fake.presentCredentialReferencesArgsForCall[i].arg1
}

func (fake *Presenter) PresentCredentials(arg1 map[string]string) {
	fake.presentCredentialsMutex.Lock()
	fake.presentCredentialsArgsForCall = append(fake.presentCredentialsArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	fake.recordInvocation("PresentCredentials", []interface{}{arg1})
	fake.presentCredentialsMutex.Unlock()
	if fake.PresentCredentialsStub != nil {
		fake.PresentCredentialsStub(arg1)
	}
}

func (fake *Presenter) PresentCredentialsCallCount() int {
	fake.presentCredentialsMutex.RLock()
	defer fake.presentCredentialsMutex.RUnlock()
	return len(fake.presentCredentialsArgsForCall)
}

func (fake *Presenter) PresentCredentialsArgsForCall(i int) map[string]string {
	fake.presentCredentialsMutex.RLock()
	defer fake.presentCredentialsMutex.RUnlock()
	return fake.presentCredentialsArgsForCall[i].arg1
}

func (fake *Presenter) PresentDeployedProducts(arg1 []api.DiagnosticProduct) {
	var arg1Copy []api.DiagnosticProduct
	if arg1 != nil {
		arg1Copy = make([]api.DiagnosticProduct, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentDeployedProductsMutex.Lock()
	fake.presentDeployedProductsArgsForCall = append(fake.presentDeployedProductsArgsForCall, struct {
		arg1 []api.DiagnosticProduct
	}{arg1Copy})
	fake.recordInvocation("PresentDeployedProducts", []interface{}{arg1Copy})
	fake.presentDeployedProductsMutex.Unlock()
	if fake.PresentDeployedProductsStub != nil {
		fake.PresentDeployedProductsStub(arg1)
	}
}

func (fake *Presenter) PresentDeployedProductsCallCount() int {
	fake.presentDeployedProductsMutex.RLock()
	defer fake.presentDeployedProductsMutex.RUnlock()
	return len(fake.presentDeployedProductsArgsForCall)
}

func (fake *Presenter) PresentDeployedProductsArgsForCall(i int) []api.DiagnosticProduct {
	fake.presentDeployedProductsMutex.RLock()
	defer fake.presentDeployedProductsMutex.RUnlock()
	return fake.presentDeployedProductsArgsForCall[i].arg1
}

func (fake *Presenter) PresentErrands(arg1 []models.Errand) {
	var arg1Copy []models.Errand
	if arg1 != nil {
		arg1Copy = make([]models.Errand, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentErrandsMutex.Lock()
	fake.presentErrandsArgsForCall = append(fake.presentErrandsArgsForCall, struct {
		arg1 []models.Errand
	}{arg1Copy})
	fake.recordInvocation("PresentErrands", []interface{}{arg1Copy})
	fake.presentErrandsMutex.Unlock()
	if fake.PresentErrandsStub != nil {
		fake.PresentErrandsStub(arg1)
	}
}

func (fake *Presenter) PresentErrandsCallCount() int {
	fake.presentErrandsMutex.RLock()
	defer fake.presentErrandsMutex.RUnlock()
	return len(fake.presentErrandsArgsForCall)
}

func (fake *Presenter) PresentErrandsArgsForCall(i int) []models.Errand {
	fake.presentErrandsMutex.RLock()
	defer fake.presentErrandsMutex.RUnlock()
	return fake.presentErrandsArgsForCall[i].arg1
}

func (fake *Presenter) PresentCertificateAuthority(arg1 api.CA) {
	fake.presentCertificateAuthorityMutex.Lock()
	fake.presentCertificateAuthorityArgsForCall = append(fake.presentCertificateAuthorityArgsForCall, struct {
		arg1 api.CA
	}{arg1})
	fake.recordInvocation("PresentCertificateAuthority", []interface{}{arg1})
	fake.presentCertificateAuthorityMutex.Unlock()
	if fake.PresentCertificateAuthorityStub != nil {
		fake.PresentCertificateAuthorityStub(arg1)
	}
}

func (fake *Presenter) PresentCertificateAuthorityCallCount() int {
	fake.presentCertificateAuthorityMutex.RLock()
	defer fake.presentCertificateAuthorityMutex.RUnlock()
	return len(fake.presentCertificateAuthorityArgsForCall)
}

func (fake *Presenter) PresentCertificateAuthorityArgsForCall(i int) api.CA {
	fake.presentCertificateAuthorityMutex.RLock()
	defer fake.presentCertificateAuthorityMutex.RUnlock()
	return fake.presentCertificateAuthorityArgsForCall[i].arg1
}

func (fake *Presenter) PresentInstallations(arg1 []models.Installation) {
	var arg1Copy []models.Installation
	if arg1 != nil {
		arg1Copy = make([]models.Installation, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentInstallationsMutex.Lock()
	fake.presentInstallationsArgsForCall = append(fake.presentInstallationsArgsForCall, struct {
		arg1 []models.Installation
	}{arg1Copy})
	fake.recordInvocation("PresentInstallations", []interface{}{arg1Copy})
	fake.presentInstallationsMutex.Unlock()
	if fake.PresentInstallationsStub != nil {
		fake.PresentInstallationsStub(arg1)
	}
}

func (fake *Presenter) PresentInstallationsCallCount() int {
	fake.presentInstallationsMutex.RLock()
	defer fake.presentInstallationsMutex.RUnlock()
	return len(fake.presentInstallationsArgsForCall)
}

func (fake *Presenter) PresentInstallationsArgsForCall(i int) []models.Installation {
	fake.presentInstallationsMutex.RLock()
	defer fake.presentInstallationsMutex.RUnlock()
	return fake.presentInstallationsArgsForCall[i].arg1
}

func (fake *Presenter) PresentPendingChanges(arg1 []api.ProductChange) {
	var arg1Copy []api.ProductChange
	if arg1 != nil {
		arg1Copy = make([]api.ProductChange, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentPendingChangesMutex.Lock()
	fake.presentPendingChangesArgsForCall = append(fake.presentPendingChangesArgsForCall, struct {
		arg1 []api.ProductChange
	}{arg1Copy})
	fake.recordInvocation("PresentPendingChanges", []interface{}{arg1Copy})
	fake.presentPendingChangesMutex.Unlock()
	if fake.PresentPendingChangesStub != nil {
		fake.PresentPendingChangesStub(arg1)
	}
}

func (fake *Presenter) PresentPendingChangesCallCount() int {
	fake.presentPendingChangesMutex.RLock()
	defer fake.presentPendingChangesMutex.RUnlock()
	return len(fake.presentPendingChangesArgsForCall)
}

func (fake *Presenter) PresentPendingChangesArgsForCall(i int) []api.ProductChange {
	fake.presentPendingChangesMutex.RLock()
	defer fake.presentPendingChangesMutex.RUnlock()
	return fake.presentPendingChangesArgsForCall[i].arg1
}

func (fake *Presenter) PresentStagedProducts(arg1 []api.DiagnosticProduct) {
	var arg1Copy []api.DiagnosticProduct
	if arg1 != nil {
		arg1Copy = make([]api.DiagnosticProduct, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.presentStagedProductsMutex.Lock()
	fake.presentStagedProductsArgsForCall = append(fake.presentStagedProductsArgsForCall, struct {
		arg1 []api.DiagnosticProduct
	}{arg1Copy})
	fake.recordInvocation("PresentStagedProducts", []interface{}{arg1Copy})
	fake.presentStagedProductsMutex.Unlock()
	if fake.PresentStagedProductsStub != nil {
		fake.PresentStagedProductsStub(arg1)
	}
}

func (fake *Presenter) PresentStagedProductsCallCount() int {
	fake.presentStagedProductsMutex.RLock()
	defer fake.presentStagedProductsMutex.RUnlock()
	return len(fake.presentStagedProductsArgsForCall)
}

func (fake *Presenter) PresentStagedProductsArgsForCall(i int) []api.DiagnosticProduct {
	fake.presentStagedProductsMutex.RLock()
	defer fake.presentStagedProductsMutex.RUnlock()
	return fake.presentStagedProductsArgsForCall[i].arg1
}

func (fake *Presenter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.presentAvailableProductsMutex.RLock()
	defer fake.presentAvailableProductsMutex.RUnlock()
	fake.presentCertificateAuthoritiesMutex.RLock()
	defer fake.presentCertificateAuthoritiesMutex.RUnlock()
	fake.presentCredentialReferencesMutex.RLock()
	defer fake.presentCredentialReferencesMutex.RUnlock()
	fake.presentCredentialsMutex.RLock()
	defer fake.presentCredentialsMutex.RUnlock()
	fake.presentDeployedProductsMutex.RLock()
	defer fake.presentDeployedProductsMutex.RUnlock()
	fake.presentErrandsMutex.RLock()
	defer fake.presentErrandsMutex.RUnlock()
	fake.presentCertificateAuthorityMutex.RLock()
	defer fake.presentCertificateAuthorityMutex.RUnlock()
	fake.presentInstallationsMutex.RLock()
	defer fake.presentInstallationsMutex.RUnlock()
	fake.presentPendingChangesMutex.RLock()
	defer fake.presentPendingChangesMutex.RUnlock()
	fake.presentStagedProductsMutex.RLock()
	defer fake.presentStagedProductsMutex.RUnlock()
	return fake.invocations
}

func (fake *Presenter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ commands.Presenter = new(Presenter)
