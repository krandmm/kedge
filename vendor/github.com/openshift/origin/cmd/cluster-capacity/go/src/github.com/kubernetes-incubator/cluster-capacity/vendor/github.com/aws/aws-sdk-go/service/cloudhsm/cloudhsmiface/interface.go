// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudhsmiface provides an interface to enable mocking the Amazon CloudHSM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudhsmiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
)

// CloudHSMAPI provides an interface to enable mocking the
// cloudhsm.CloudHSM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudHSM.
//    func myFunc(svc cloudhsmiface.CloudHSMAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudhsm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudHSMClient struct {
//        cloudhsmiface.CloudHSMAPI
//    }
//    func (m *mockCloudHSMClient) AddTagsToResource(input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudHSMClient{}
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
type CloudHSMAPI interface {
	AddTagsToResourceRequest(*cloudhsm.AddTagsToResourceInput) (*request.Request, *cloudhsm.AddTagsToResourceOutput)

	AddTagsToResource(*cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)

	CreateHapgRequest(*cloudhsm.CreateHapgInput) (*request.Request, *cloudhsm.CreateHapgOutput)

	CreateHapg(*cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)

	CreateHsmRequest(*cloudhsm.CreateHsmInput) (*request.Request, *cloudhsm.CreateHsmOutput)

	CreateHsm(*cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)

	CreateLunaClientRequest(*cloudhsm.CreateLunaClientInput) (*request.Request, *cloudhsm.CreateLunaClientOutput)

	CreateLunaClient(*cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)

	DeleteHapgRequest(*cloudhsm.DeleteHapgInput) (*request.Request, *cloudhsm.DeleteHapgOutput)

	DeleteHapg(*cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)

	DeleteHsmRequest(*cloudhsm.DeleteHsmInput) (*request.Request, *cloudhsm.DeleteHsmOutput)

	DeleteHsm(*cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)

	DeleteLunaClientRequest(*cloudhsm.DeleteLunaClientInput) (*request.Request, *cloudhsm.DeleteLunaClientOutput)

	DeleteLunaClient(*cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)

	DescribeHapgRequest(*cloudhsm.DescribeHapgInput) (*request.Request, *cloudhsm.DescribeHapgOutput)

	DescribeHapg(*cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)

	DescribeHsmRequest(*cloudhsm.DescribeHsmInput) (*request.Request, *cloudhsm.DescribeHsmOutput)

	DescribeHsm(*cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)

	DescribeLunaClientRequest(*cloudhsm.DescribeLunaClientInput) (*request.Request, *cloudhsm.DescribeLunaClientOutput)

	DescribeLunaClient(*cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)

	GetConfigRequest(*cloudhsm.GetConfigInput) (*request.Request, *cloudhsm.GetConfigOutput)

	GetConfig(*cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)

	ListAvailableZonesRequest(*cloudhsm.ListAvailableZonesInput) (*request.Request, *cloudhsm.ListAvailableZonesOutput)

	ListAvailableZones(*cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)

	ListHapgsRequest(*cloudhsm.ListHapgsInput) (*request.Request, *cloudhsm.ListHapgsOutput)

	ListHapgs(*cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)

	ListHsmsRequest(*cloudhsm.ListHsmsInput) (*request.Request, *cloudhsm.ListHsmsOutput)

	ListHsms(*cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)

	ListLunaClientsRequest(*cloudhsm.ListLunaClientsInput) (*request.Request, *cloudhsm.ListLunaClientsOutput)

	ListLunaClients(*cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)

	ListTagsForResourceRequest(*cloudhsm.ListTagsForResourceInput) (*request.Request, *cloudhsm.ListTagsForResourceOutput)

	ListTagsForResource(*cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)

	ModifyHapgRequest(*cloudhsm.ModifyHapgInput) (*request.Request, *cloudhsm.ModifyHapgOutput)

	ModifyHapg(*cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)

	ModifyHsmRequest(*cloudhsm.ModifyHsmInput) (*request.Request, *cloudhsm.ModifyHsmOutput)

	ModifyHsm(*cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)

	ModifyLunaClientRequest(*cloudhsm.ModifyLunaClientInput) (*request.Request, *cloudhsm.ModifyLunaClientOutput)

	ModifyLunaClient(*cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)

	RemoveTagsFromResourceRequest(*cloudhsm.RemoveTagsFromResourceInput) (*request.Request, *cloudhsm.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
}

var _ CloudHSMAPI = (*cloudhsm.CloudHSM)(nil)