/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/rook/rook/pkg/apis/cassandra.rook.io/v1alpha1"
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	cockroachdbrookiov1alpha1 "github.com/rook/rook/pkg/apis/cockroachdb.rook.io/v1alpha1"
	v1beta1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1beta1"
	miniorookiov1alpha1 "github.com/rook/rook/pkg/apis/minio.rook.io/v1alpha1"
	nfsrookiov1alpha1 "github.com/rook/rook/pkg/apis/nfs.rook.io/v1alpha1"
	v1alpha2 "github.com/rook/rook/pkg/apis/rook.io/v1alpha2"
	yugabytedbrookiov1alpha1 "github.com/rook/rook/pkg/apis/yugabytedb.rook.io/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=cassandra.rook.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cassandra().V1alpha1().Clusters().Informer()}, nil

		// Group=ceph.rook.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("cephblockpools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephBlockPools().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cephclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephClusters().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cephfilesystems"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephFilesystems().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cephnfses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephNFSes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cephobjectstores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephObjectStores().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cephobjectstoreusers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ceph().V1().CephObjectStoreUsers().Informer()}, nil

		// Group=cockroachdb.rook.io, Version=v1alpha1
	case cockroachdbrookiov1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cockroachdb().V1alpha1().Clusters().Informer()}, nil

		// Group=edgefs.rook.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().Clusters().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("iscsis"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().ISCSIs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("isgws"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().ISGWs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("nfss"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().NFSs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("s3s"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().S3s().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("s3xs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().S3Xs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("swifts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Edgefs().V1beta1().SWIFTs().Informer()}, nil

		// Group=minio.rook.io, Version=v1alpha1
	case miniorookiov1alpha1.SchemeGroupVersion.WithResource("objectstores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Minio().V1alpha1().ObjectStores().Informer()}, nil

		// Group=nfs.rook.io, Version=v1alpha1
	case nfsrookiov1alpha1.SchemeGroupVersion.WithResource("nfsservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nfs().V1alpha1().NFSServers().Informer()}, nil

		// Group=rook.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("volumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rook().V1alpha2().Volumes().Informer()}, nil

		// Group=yugabytedb.rook.io, Version=v1alpha1
	case yugabytedbrookiov1alpha1.SchemeGroupVersion.WithResource("ybclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Yugabytedb().V1alpha1().YBClusters().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
