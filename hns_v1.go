//go:build windows

package hnslib

import (
	"github.com/Microsoft/hnslib/internal/hns"
)

// HNSEndpoint represents a network endpoint in HNS
type HNSEndpoint = hns.HNSEndpoint

// HNSEndpointStats represent the stats for an networkendpoint in HNS
type HNSEndpointStats = hns.EndpointStats

// Namespace represents a Compartment.
type Namespace = hns.Namespace

// GetHNSEndpointStats gets the endpoint stats by ID
func GetHNSEndpointStats(endpointName string) (*HNSEndpointStats, error) {
	return hns.GetHNSEndpointStats(endpointName)
}