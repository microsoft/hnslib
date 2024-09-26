//go:build windows && integration
// +build windows,integration

package hcn

import (
	"os"
	"testing"

	"github.com/Microsoft/hnslib"
)

const (
	NatTestNetworkName     string = "GoTestNat"
	NatTestEndpointName    string = "GoTestNatEndpoint"
	OverlayTestNetworkName string = "GoTestOverlay"
	BridgeTestNetworkName  string = "GoTestL2Bridge"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func CreateTestNetwork() (*hnslib.HNSNetwork, error) {
	network := &hnslib.HNSNetwork{
		Type: "NAT",
		Name: NatTestNetworkName,
		Subnets: []hnslib.Subnet{
			{
				AddressPrefix:  "192.168.100.0/24",
				GatewayAddress: "192.168.100.1",
			},
		},
	}

	return network.Create()
}

func TestEndpoint(t *testing.T) {
	network, err := CreateTestNetwork()
	if err != nil {
		t.Fatal(err)
	}

	Endpoint := &hnslib.HNSEndpoint{
		Name: NatTestEndpointName,
	}

	Endpoint, err = network.CreateEndpoint(Endpoint)
	if err != nil {
		t.Fatal(err)
	}

	err = Endpoint.HostAttach(1)
	if err != nil {
		t.Fatal(err)
	}

	err = Endpoint.HostDetach()
	if err != nil {
		t.Fatal(err)
	}

	_, err = Endpoint.Delete()
	if err != nil {
		t.Fatal(err)
	}

	_, err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEndpointGetAll(t *testing.T) {
	_, err := hnslib.HNSListEndpointRequest()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEndpointStatsAll(t *testing.T) {
	network, err := CreateTestNetwork()
	if err != nil {
		t.Fatal(err)
	}

	Endpoint := &hnslib.HNSEndpoint{
		Name: NatTestEndpointName,
	}

	_, err = network.CreateEndpoint(Endpoint)
	if err != nil {
		t.Fatal(err)
	}

	epList, err := hnslib.HNSListEndpointRequest()
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range epList {
		_, err := hnslib.GetHNSEndpointStats(e.Id)
		if err != nil {
			t.Fatal(err)
		}
	}

	_, err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkGetAll(t *testing.T) {
	_, err := hnslib.HNSListNetworkRequest("GET", "", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetwork(t *testing.T) {
	network, err := CreateTestNetwork()
	if err != nil {
		t.Fatal(err)
	}
	_, err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
