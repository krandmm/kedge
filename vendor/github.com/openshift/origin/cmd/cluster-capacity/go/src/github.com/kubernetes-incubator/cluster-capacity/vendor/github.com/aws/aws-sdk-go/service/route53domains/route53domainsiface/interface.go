// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53domainsiface provides an interface to enable mocking the Amazon Route 53 Domains service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53domainsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

// Route53DomainsAPI provides an interface to enable mocking the
// route53domains.Route53Domains service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53 Domains.
//    func myFunc(svc route53domainsiface.Route53DomainsAPI) bool {
//        // Make svc.CheckDomainAvailability request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := route53domains.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53DomainsClient struct {
//        route53domainsiface.Route53DomainsAPI
//    }
//    func (m *mockRoute53DomainsClient) CheckDomainAvailability(input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53DomainsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Route53DomainsAPI interface {
	CheckDomainAvailabilityRequest(*route53domains.CheckDomainAvailabilityInput) (*request.Request, *route53domains.CheckDomainAvailabilityOutput)

	CheckDomainAvailability(*route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)

	DeleteTagsForDomainRequest(*route53domains.DeleteTagsForDomainInput) (*request.Request, *route53domains.DeleteTagsForDomainOutput)

	DeleteTagsForDomain(*route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)

	DisableDomainAutoRenewRequest(*route53domains.DisableDomainAutoRenewInput) (*request.Request, *route53domains.DisableDomainAutoRenewOutput)

	DisableDomainAutoRenew(*route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)

	DisableDomainTransferLockRequest(*route53domains.DisableDomainTransferLockInput) (*request.Request, *route53domains.DisableDomainTransferLockOutput)

	DisableDomainTransferLock(*route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)

	EnableDomainAutoRenewRequest(*route53domains.EnableDomainAutoRenewInput) (*request.Request, *route53domains.EnableDomainAutoRenewOutput)

	EnableDomainAutoRenew(*route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)

	EnableDomainTransferLockRequest(*route53domains.EnableDomainTransferLockInput) (*request.Request, *route53domains.EnableDomainTransferLockOutput)

	EnableDomainTransferLock(*route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)

	GetContactReachabilityStatusRequest(*route53domains.GetContactReachabilityStatusInput) (*request.Request, *route53domains.GetContactReachabilityStatusOutput)

	GetContactReachabilityStatus(*route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error)

	GetDomainDetailRequest(*route53domains.GetDomainDetailInput) (*request.Request, *route53domains.GetDomainDetailOutput)

	GetDomainDetail(*route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)

	GetDomainSuggestionsRequest(*route53domains.GetDomainSuggestionsInput) (*request.Request, *route53domains.GetDomainSuggestionsOutput)

	GetDomainSuggestions(*route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error)

	GetOperationDetailRequest(*route53domains.GetOperationDetailInput) (*request.Request, *route53domains.GetOperationDetailOutput)

	GetOperationDetail(*route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)

	ListDomainsRequest(*route53domains.ListDomainsInput) (*request.Request, *route53domains.ListDomainsOutput)

	ListDomains(*route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)

	ListDomainsPages(*route53domains.ListDomainsInput, func(*route53domains.ListDomainsOutput, bool) bool) error

	ListOperationsRequest(*route53domains.ListOperationsInput) (*request.Request, *route53domains.ListOperationsOutput)

	ListOperations(*route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)

	ListOperationsPages(*route53domains.ListOperationsInput, func(*route53domains.ListOperationsOutput, bool) bool) error

	ListTagsForDomainRequest(*route53domains.ListTagsForDomainInput) (*request.Request, *route53domains.ListTagsForDomainOutput)

	ListTagsForDomain(*route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)

	RegisterDomainRequest(*route53domains.RegisterDomainInput) (*request.Request, *route53domains.RegisterDomainOutput)

	RegisterDomain(*route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)

	RenewDomainRequest(*route53domains.RenewDomainInput) (*request.Request, *route53domains.RenewDomainOutput)

	RenewDomain(*route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error)

	ResendContactReachabilityEmailRequest(*route53domains.ResendContactReachabilityEmailInput) (*request.Request, *route53domains.ResendContactReachabilityEmailOutput)

	ResendContactReachabilityEmail(*route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error)

	RetrieveDomainAuthCodeRequest(*route53domains.RetrieveDomainAuthCodeInput) (*request.Request, *route53domains.RetrieveDomainAuthCodeOutput)

	RetrieveDomainAuthCode(*route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)

	TransferDomainRequest(*route53domains.TransferDomainInput) (*request.Request, *route53domains.TransferDomainOutput)

	TransferDomain(*route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)

	UpdateDomainContactRequest(*route53domains.UpdateDomainContactInput) (*request.Request, *route53domains.UpdateDomainContactOutput)

	UpdateDomainContact(*route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)

	UpdateDomainContactPrivacyRequest(*route53domains.UpdateDomainContactPrivacyInput) (*request.Request, *route53domains.UpdateDomainContactPrivacyOutput)

	UpdateDomainContactPrivacy(*route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)

	UpdateDomainNameserversRequest(*route53domains.UpdateDomainNameserversInput) (*request.Request, *route53domains.UpdateDomainNameserversOutput)

	UpdateDomainNameservers(*route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)

	UpdateTagsForDomainRequest(*route53domains.UpdateTagsForDomainInput) (*request.Request, *route53domains.UpdateTagsForDomainOutput)

	UpdateTagsForDomain(*route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)

	ViewBillingRequest(*route53domains.ViewBillingInput) (*request.Request, *route53domains.ViewBillingOutput)

	ViewBilling(*route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error)
}

var _ Route53DomainsAPI = (*route53domains.Route53Domains)(nil)