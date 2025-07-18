//go:build windows && integration
// +build windows,integration

package hcn

import (
	"testing"
	"time"
)

func TestIntegration_CreateAndDeleteNetwork(t *testing.T) {
	networkName := "integration-test-network"
	network := &HostComputeNetwork{
		Type: NAT,
		Name: networkName,
		Ipams: []Ipam{{
			Type: "Static",
			Subnets: []Subnet{{
				IpAddressPrefix: "192.168.250.0/24",
				Routes: []Route{{
					NextHop: "192.168.250.1",
					DestinationPrefix: "0.0.0.0/0",
				}},
			}},
		}},
		SchemaVersion: SchemaVersion{Major: 2, Minor: 0},
	}

	// Create network
	created, err := network.Create()
	if err != nil {
		t.Fatalf("Failed to create network: %v", err)
	}
	if created == nil {
		t.Fatal("Create() returned nil network")
	}

	// Verify network exists
	found, err := GetNetworkByID(created.Id)
	if err != nil {
		t.Fatalf("Failed to get network by ID: %v", err)
	}
	if found == nil || found.Name != networkName {
		t.Fatalf("Network not found after creation")
	}

	// Delete network
	if err := created.Delete(); err != nil {
		t.Fatalf("Failed to delete network: %v", err)
	}

	// Give HNS a moment to update
	time.Sleep(2 * time.Second)

	// Verify network is gone
	_, err = GetNetworkByID(created.Id)
	if err == nil {
		t.Fatalf("Network still exists after deletion")
	}
}
