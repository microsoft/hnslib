//go:build windows && integration && featureaccelnetsupport
// +build windows,integration,featureaccelnetsupport

package hcn

import (
	"os"
	"testing"

	"github.com/Microsoft/hnslib"
)

func TestAccelnetNnvManagementMacAddresses(t *testing.T) {
	network, err := CreateTestNetwork()
	if err != nil {
		t.Fatal(err)
	}

	macList := []string{"00-15-5D-0A-B7-C6", "00-15-5D-38-01-00"}
	newMacList, err := hnslib.SetNnvManagementMacAddresses(macList)

	if err != nil {
		t.Fatal(err)
	}

	if len(newMacList.MacAddressList) != 2 {
		t.Errorf("After Create: Expected macaddress count %d, got %d", 2, len(newMacList.MacAddressList))
	}

	newMacList, err = hnslib.GetNnvManagementMacAddresses()
	if err != nil {
		t.Fatal(err)
	}

	if len(newMacList.MacAddressList) != 2 {
		t.Errorf("Get After Create: Expected macaddress count %d, got %d", 2, len(newMacList.MacAddressList))
	}

	newMacList, err = hnslib.DeleteNnvManagementMacAddresses()
	if err != nil {
		t.Fatal(err)
	}

	if len(newMacList.MacAddressList) != 0 {
		t.Errorf("After Delete: Expected macaddress count %d, got %d", 0, len(newMacList.MacAddressList))
	}

	newMacList, err = hnslib.GetNnvManagementMacAddresses()
	if err != nil {
		t.Fatal(err)
	}

	if len(newMacList.MacAddressList) != 0 {
		t.Errorf("Get After Delete: Expected macaddress count %d, got %d", 0, len(newMacList.MacAddressList))
	}

	_, err = network.Delete()
	if err != nil {
		t.Fatal(err)
	}
}