package main

import (
	"errors"
	"sync"
)

// ResourcePoolManager manages a pool of reusable Resource objects.
// This version includes concurrency-safe locking.
// ------------------------------------------------------------
type ResourcePoolManager struct {
	emptyResources  []Resource // resources available to use
	useResources    []Resource // resources currently in use
	initialPoolSize int
	maxPoolSize     int

	mu sync.Mutex // protects all fields of ResourcePoolManager
}

// NewResourcePoolManager initializes the pool with 10 resources.
func NewResourcePoolManager() *ResourcePoolManager {
	resources := make([]Resource, 0, 10)
	for i := 0; i < 10; i++ {
		resources = append(resources, NewResource(i))
	}

	return &ResourcePoolManager{
		emptyResources:  resources,
		useResources:    []Resource{},
		initialPoolSize: 10,
		maxPoolSize:     20,
	}
}

// Singleton instance
var (
	rpmInstance *ResourcePoolManager
	once        sync.Once
)

// GetResourcePoolManager returns the global singleton instance.
func GetResourcePoolManager() *ResourcePoolManager {
	once.Do(func() {
		rpmInstance = NewResourcePoolManager()
	})
	return rpmInstance
}

// GetResource retrieves a resource from the pool. Safe for concurrent calls.
func (rpm *ResourcePoolManager) GetResource() (Resource, error) {
	rpm.mu.Lock()
	defer rpm.mu.Unlock()

	// If empty, consider creating new one if pool isn't full
	if rpm.isEmpty() {
		if rpm.isFull() {
			return Resource{}, errors.New("resource pool exhausted")
		}

		// Create a new resource with next available ID
		newID := len(rpm.emptyResources) + len(rpm.useResources)
		rpm.emptyResources = append(rpm.emptyResources, NewResource(newID))
	}

	// Take last empty resource (O(1))
	res := rpm.emptyResources[len(rpm.emptyResources)-1]
	rpm.emptyResources = rpm.emptyResources[:len(rpm.emptyResources)-1]

	// Move resource to the in-use list
	rpm.useResources = append(rpm.useResources, res)

	return res, nil
}

// RemoveResource returns a resource back to the pool.
// Uses swap-delete to remove from useResources in O(1).
func (rpm *ResourcePoolManager) RemoveResource(resource Resource) {
	rpm.mu.Lock()
	defer rpm.mu.Unlock()

	// Find and remove from useResources
	for i, r := range rpm.useResources {
		if r.id == resource.id {
			// Swap with last element, then truncate slice
			rpm.useResources[i] = rpm.useResources[len(rpm.useResources)-1]
			rpm.useResources = rpm.useResources[:len(rpm.useResources)-1]
			break
		}
	}

	// Put it back into empty pool
	rpm.emptyResources = append(rpm.emptyResources, resource)
}

// isEmpty returns true if no resources are available
func (rpm *ResourcePoolManager) isEmpty() bool {
	return len(rpm.emptyResources) == 0
}

// isFull returns true if total resources reached maxPoolSize
func (rpm *ResourcePoolManager) isFull() bool {
	return len(rpm.emptyResources)+len(rpm.useResources) >= rpm.maxPoolSize
}
