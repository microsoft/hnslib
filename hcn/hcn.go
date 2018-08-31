// Package hcn is a shim for the Host Compute Networking (HCN) service, which manages networking for Windows Server
// containers and Hyper-V containers. Previous to RS5, HCN was referred to as Host Networking Service (HNS).
package hcn

import (
	"fmt"
	"syscall"

	"github.com/Microsoft/hcsshim/internal/hcserror"
	"github.com/Microsoft/hcsshim/internal/interop"
)

// hns flag imports "github.com/Microsoft/hcsshim/internal/guid"
//go:generate go run ../mksyscall_windows.go -output zsyscall_windows.go -hns hcn.go

/// HNS V1 API

//sys SetCurrentThreadCompartmentId(compartmentId uint32) (hr error) = iphlpapi.SetCurrentThreadCompartmentId
//sys _hnsCall(method string, path string, object string, response **uint16) (hr error) = vmcompute.HNSCall?

/// HCN V2 API

// Network
//sys hcnEnumerateNetworks(query string, networks **uint16, result **uint16) (hr error) = computenetwork.HcnEnumerateNetworks?
//sys hcnCreateNetwork(id *guid.GUID, settings string, network *hcnNetwork, result **uint16) (hr error) = computenetwork.HcnCreateNetwork?
//sys hcnOpenNetwork(id *guid.GUID, network *hcnNetwork, result **uint16) (hr error) = computenetwork.HcnOpenNetwork?
//sys hcnModifyNetwork(network hcnNetwork, settings string, result **uint16) (hr error) = computenetwork.HcnModifyNetwork?
//sys hcnQueryNetworkProperties(network hcnNetwork, query string, properties **uint16, result **uint16) (hr error) = computenetwork.HcnQueryNetworkProperties?
//sys hcnDeleteNetwork(id *guid.GUID, result **uint16) (hr error) = computenetwork.HcnDeleteNetwork?
//sys hcnCloseNetwork(network hcnNetwork) (hr error) = computenetwork.HcnCloseNetwork?

// Endpoint
//sys hcnEnumerateEndpoints(query string, endpoints **uint16, result **uint16) (hr error) = computenetwork.HcnEnumerateEndpoints?
//sys hcnCreateEndpoint(network hcnNetwork, id *guid.GUID, settings string, endpoint *hcnEndpoint, result **uint16) (hr error) = computenetwork.HcnCreateEndpoint?
//sys hcnOpenEndpoint(id *guid.GUID, endpoint *hcnEndpoint, result **uint16) (hr error) = computenetwork.HcnOpenEndpoint?
//sys hcnModifyEndpoint(endpoint hcnEndpoint, settings string, result **uint16) (hr error) = computenetwork.HcnModifyEndpoint?
//sys hcnQueryEndpointProperties(endpoint hcnEndpoint, query string, properties **uint16, result **uint16) (hr error) = computenetwork.HcnQueryEndpointProperties?
//sys hcnDeleteEndpoint(id *guid.GUID, result **uint16) (hr error) = computenetwork.HcnDeleteEndpoint?
//sys hcnCloseEndpoint(endpoint hcnEndpoint) (hr error) = computenetwork.HcnCloseEndpoint?

// Namespace
//sys hcnEnumerateNamespaces(query string, namespaces **uint16, result **uint16) (hr error) = computenetwork.HcnEnumerateNamespaces?
//sys hcnCreateNamespace(id *guid.GUID, settings string, namespace *hcnNamespace, result **uint16) (hr error) = computenetwork.HcnCreateNamespace?
//sys hcnOpenNamespace(id *guid.GUID, namespace *hcnNamespace, result **uint16) (hr error) = computenetwork.HcnOpenNamespace?
//sys hcnModifyNamespace(namespace hcnNamespace, settings string, result **uint16) (hr error) = computenetwork.HcnModifyNamespace?
//sys hcnQueryNamespaceProperties(namespace hcnNamespace, query string, properties **uint16, result **uint16) (hr error) = computenetwork.HcnQueryNamespaceProperties?
//sys hcnDeleteNamespace(id *guid.GUID, result **uint16) (hr error) = computenetwork.HcnDeleteNamespace?
//sys hcnCloseNamespace(namespace hcnNamespace) (hr error) = computenetwork.HcnCloseNamespace?

// LoadBalancer
//sys hcnEnumerateLoadBalancers(query string, loadBalancers **uint16, result **uint16) (hr error) = computenetwork.HcnEnumerateLoadBalancers?
//sys hcnCreateLoadBalancer(id *guid.GUID, settings string, loadBalancer *hcnLoadBalancer, result **uint16) (hr error) = computenetwork.HcnCreateLoadBalancer?
//sys hcnOpenLoadBalancer(id *guid.GUID, loadBalancer *hcnLoadBalancer, result **uint16) (hr error) = computenetwork.HcnOpenLoadBalancer?
//sys hcnModifyLoadBalancer(loadBalancer hcnLoadBalancer, settings string, result **uint16) (hr error) = computenetwork.HcnModifyLoadBalancer?
//sys hcnQueryLoadBalancerProperties(loadBalancer hcnLoadBalancer, query string, properties **uint16, result **uint16) (hr error) = computenetwork.HcnQueryLoadBalancerProperties?
//sys hcnDeleteLoadBalancer(id *guid.GUID, result **uint16) (hr error) = computenetwork.HcnDeleteLoadBalancer?
//sys hcnCloseLoadBalancer(loadBalancer hcnLoadBalancer) (hr error) = computenetwork.HcnCloseLoadBalancer?

// Service
//sys hcnOpenService(service *hcnService, result **uint16) (hr error) = computenetwork.HcnOpenService?
//sys hcnRegisterServiceCallback(service hcnService, callback int32, context int32, callbackHandle *hcnCallbackHandle) (hr error) = computenetwork.HcnRegisterServiceCallback?
//sys hcnUnregisterServiceCallback(callbackHandle hcnCallbackHandle) (hr error) = computenetwork.HcnUnregisterServiceCallback?
//sys hcnCloseService(service hcnService) (hr error) = computenetwork.HcnCloseService?

type hcnNetwork syscall.Handle
type hcnEndpoint syscall.Handle
type hcnNamespace syscall.Handle
type hcnLoadBalancer syscall.Handle
type hcnService syscall.Handle
type hcnCallbackHandle syscall.Handle

func checkForErrors(methodName string, hr error, resultBuffer *uint16) error {
	errorFound := false

	if hr != nil {
		errorFound = true
	}

	result := ""
	if resultBuffer != nil {
		result = interop.ConvertAndFreeCoTaskMemString(resultBuffer)
		if result != "" {
			errorFound = true
		}
	}

	if errorFound {
		return hcserror.New(hr, methodName, result)
	}

	return nil
}

// SchemaVersion for HCN Objects/Queries.
type SchemaVersion struct {
	Major uint32 `json:",omitempty"`
	Minor uint32 `json:",omitempty"`
}

// MarshalJSON formats the SchemaVersion.
func (s *SchemaVersion) MarshalJSON() ([]byte, error) {
	jsonString := fmt.Sprintf("{\"Major\":%d,\"Minor\":%d}", s.Major, s.Minor)
	return []byte(jsonString), nil
}

// HostComputeQuery is the format for HCN queries.
type HostComputeQuery struct {
	SchemaVersion SchemaVersion `json:""`
	Flags         uint32        `json:",omitempty"` // 0: None, 1: Detailed
	Filter        string        `json:",omitempty"`
}

// QuerySchema generates HCN Query.
// Passed into get/enumerate calls to filter results.
func QuerySchema(version int) HostComputeQuery {
	var query HostComputeQuery
	switch version {
	case 2:
		query = HostComputeQuery{
			SchemaVersion: SchemaVersion{
				Major: 2,
				Minor: 0,
			},
		}
	}
	return query
}

// PlatformDoesNotSupportError happens when users are attempting to use a newer shim on an older OS
func platformDoesNotSupportError(featureName string) error {
	return fmt.Errorf("Platform does not support feature %s", featureName)
}

// V2ApiSupported returns an error if the HCN version does not support the V2 Apis.
func V2ApiSupported() error {
	supported := GetSupportedFeatures()
	if supported.Api.V2 {
		return nil
	}
	return platformDoesNotSupportError("V2 Api/Schema")
}
