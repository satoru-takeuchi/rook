/*
Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"net"

	cephver "github.com/rook/rook/pkg/operator/ceph/version"
	"github.com/rook/rook/pkg/operator/k8sutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// OwnerInfo is used by both controller and non-controller. If used from controller,
// you should only set owner, controller, and scheme. Otherwise, you should only set
// ownerRef. controlelr will be ignored.
type OwnerInfo struct {
	owner      metav1.Object
	ownerRef   metav1.OwnerReference
	controller bool
	scheme     *runtime.Scheme
}

// ClusterInfo is a collection of information about a particular Ceph cluster. Rook uses information
// about the cluster to configure daemons to connect to the desired cluster.
type ClusterInfo struct {
	FSID          string
	MonitorSecret string
	CephCred      CephCred
	Monitors      map[string]*MonInfo
	CephVersion   cephver.CephVersion
	Namespace     string
	OwnerInfo     OwnerInfo
	// Hide the name of the cluster since in 99% of uses we want to use the cluster namespace.
	// If the CR name is needed, access it through the NamespacedName() method.
	name string
}

// MonInfo is a collection of information about a Ceph mon.
type MonInfo struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}

// CephCred represents the Ceph cluster username and key used by the operator.
// For converged clusters it will be the admin key, but external clusters will have a
// lower-privileged key.
type CephCred struct {
	Username string `json:"name"`
	Secret   string `json:"secret"`
}

// NewOwnerInfo creates a new owner info used by controller
func NewOwnerInfo(owner metav1.Object, controller bool, scheme *runtime.Scheme) *OwnerInfo {
	return &OwnerInfo{owner: owner, scheme: scheme}
}

// NewOwnerInfoWithRef creates a new owner info used by non-controller
func NewOwnerInfoWithRef(ownerRef metav1.OwnerReference) *OwnerInfo {
	return &OwnerInfo{ownerRef: ownerRef}
}

// NewClusterInfo creates a new cluster info
func NewClusterInfo(namespace, name string) *ClusterInfo {
	return &ClusterInfo{Namespace: namespace, name: name}
}

// SetName sets the cluster name of a cluster
func (c *ClusterInfo) SetName(name string) {
	c.name = name
}

// NamespacedName gets a namespaced name of a cluster
func (c *ClusterInfo) NamespacedName() types.NamespacedName {
	if c.name == "" {
		panic("name is not set on the clusterInfo")
	}
	return types.NamespacedName{Namespace: c.Namespace, Name: c.name}
}

// AdminClusterInfo creates a ClusterInfo with the basic info to access the cluster
// as an admin. Only the namespace and the ceph username fields are set in the struct,
// so this clusterInfo cannot be used to generate the mon config or request the
// namespacedName. A full cluster info must be populated for those operations.
func AdminClusterInfo(namespace string) *ClusterInfo {
	return &ClusterInfo{
		Namespace: namespace,
		CephCred: CephCred{
			Username: AdminUsername,
		},
	}
}

// IsInitialized returns true if the critical information in the ClusterInfo struct has been filled
// in. This method exists less out of necessity than the desire to be explicit about the lifecycle
// of the ClusterInfo struct during startup, specifically that it is expected to exist after the
// Rook operator has started up or connected to the first components of the Ceph cluster.
func (c *ClusterInfo) IsInitialized(logError bool) bool {
	var isInitialized bool

	if c == nil {
		if logError {
			logger.Error("clusterInfo is nil")
		}
	} else if c.FSID == "" {
		if logError {
			logger.Error("cluster fsid is empty")
		}
	} else if c.MonitorSecret == "" {
		if logError {
			logger.Error("monitor secret is empty")
		}
	} else if c.CephCred.Username == "" {
		if logError {
			logger.Error("ceph username is empty")
		}
	} else if c.CephCred.Secret == "" {
		if logError {
			logger.Error("ceph secret is empty")
		}
	} else {
		isInitialized = true
	}

	return isInitialized
}

// NewMonInfo returns a new Ceph mon info struct from the given inputs.
func NewMonInfo(name, ip string, port int32) *MonInfo {
	return &MonInfo{Name: name, Endpoint: net.JoinHostPort(ip, fmt.Sprintf("%d", port))}
}

// SetOwnerReference add the owner reference to the owner
func (o *OwnerInfo) SetOwnerReference(object metav1.Object) error {
	if o.scheme != nil {
		if o.controller {
			return controllerutil.SetControllerReference(o.owner, object, o.scheme)
		} else {
			return controllerutil.SetOwnerReference(o.owner, object, o.scheme)
		}
	} else {
		k8sutil.SetOwnerRef(object, &o.ownerRef)
		return nil
	}
}

// GetUID gets the UID of the owner
func (o *OwnerInfo) GetUID() types.UID {
	if o.scheme != nil {
		return o.owner.GetUID()
	} else {
		return o.ownerRef.UID
	}
}
