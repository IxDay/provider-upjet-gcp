/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/official-providers/provider-gcp/apis/appengine/v1beta1"
	v1beta1cloudfunctions "github.com/upbound/official-providers/provider-gcp/apis/cloudfunctions/v1beta1"
	v1beta1cloudplatform "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1"
	v1beta1cloudrun "github.com/upbound/official-providers/provider-gcp/apis/cloudrun/v1beta1"
	v1beta1cloudscheduler "github.com/upbound/official-providers/provider-gcp/apis/cloudscheduler/v1beta1"
	v1beta1cloudtasks "github.com/upbound/official-providers/provider-gcp/apis/cloudtasks/v1beta1"
	v1beta1composer "github.com/upbound/official-providers/provider-gcp/apis/composer/v1beta1"
	v1beta1compute "github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1"
	v1beta1container "github.com/upbound/official-providers/provider-gcp/apis/container/v1beta1"
	v1beta1containeranalysis "github.com/upbound/official-providers/provider-gcp/apis/containeranalysis/v1beta1"
	v1beta1datacatalog "github.com/upbound/official-providers/provider-gcp/apis/datacatalog/v1beta1"
	v1beta1dns "github.com/upbound/official-providers/provider-gcp/apis/dns/v1beta1"
	v1beta1gkehub "github.com/upbound/official-providers/provider-gcp/apis/gkehub/v1beta1"
	v1beta1iap "github.com/upbound/official-providers/provider-gcp/apis/iap/v1beta1"
	v1beta1identityplatform "github.com/upbound/official-providers/provider-gcp/apis/identityplatform/v1beta1"
	v1beta1kms "github.com/upbound/official-providers/provider-gcp/apis/kms/v1beta1"
	v1beta1monitoring "github.com/upbound/official-providers/provider-gcp/apis/monitoring/v1beta1"
	v1beta1notebooks "github.com/upbound/official-providers/provider-gcp/apis/notebooks/v1beta1"
	v1beta1osconfig "github.com/upbound/official-providers/provider-gcp/apis/osconfig/v1beta1"
	v1beta1oslogin "github.com/upbound/official-providers/provider-gcp/apis/oslogin/v1beta1"
	v1beta1privateca "github.com/upbound/official-providers/provider-gcp/apis/privateca/v1beta1"
	v1beta1pubsub "github.com/upbound/official-providers/provider-gcp/apis/pubsub/v1beta1"
	v1beta1redis "github.com/upbound/official-providers/provider-gcp/apis/redis/v1beta1"
	v1beta1secretmanager "github.com/upbound/official-providers/provider-gcp/apis/secretmanager/v1beta1"
	v1beta1servicenetworking "github.com/upbound/official-providers/provider-gcp/apis/servicenetworking/v1beta1"
	v1beta1sourcerepo "github.com/upbound/official-providers/provider-gcp/apis/sourcerepo/v1beta1"
	v1beta1spanner "github.com/upbound/official-providers/provider-gcp/apis/spanner/v1beta1"
	v1beta1sql "github.com/upbound/official-providers/provider-gcp/apis/sql/v1beta1"
	v1beta1storage "github.com/upbound/official-providers/provider-gcp/apis/storage/v1beta1"
	v1alpha1 "github.com/upbound/official-providers/provider-gcp/apis/v1alpha1"
	v1beta1apis "github.com/upbound/official-providers/provider-gcp/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1cloudfunctions.SchemeBuilder.AddToScheme,
		v1beta1cloudplatform.SchemeBuilder.AddToScheme,
		v1beta1cloudrun.SchemeBuilder.AddToScheme,
		v1beta1cloudscheduler.SchemeBuilder.AddToScheme,
		v1beta1cloudtasks.SchemeBuilder.AddToScheme,
		v1beta1composer.SchemeBuilder.AddToScheme,
		v1beta1compute.SchemeBuilder.AddToScheme,
		v1beta1container.SchemeBuilder.AddToScheme,
		v1beta1containeranalysis.SchemeBuilder.AddToScheme,
		v1beta1datacatalog.SchemeBuilder.AddToScheme,
		v1beta1dns.SchemeBuilder.AddToScheme,
		v1beta1gkehub.SchemeBuilder.AddToScheme,
		v1beta1iap.SchemeBuilder.AddToScheme,
		v1beta1identityplatform.SchemeBuilder.AddToScheme,
		v1beta1kms.SchemeBuilder.AddToScheme,
		v1beta1monitoring.SchemeBuilder.AddToScheme,
		v1beta1notebooks.SchemeBuilder.AddToScheme,
		v1beta1osconfig.SchemeBuilder.AddToScheme,
		v1beta1oslogin.SchemeBuilder.AddToScheme,
		v1beta1privateca.SchemeBuilder.AddToScheme,
		v1beta1pubsub.SchemeBuilder.AddToScheme,
		v1beta1redis.SchemeBuilder.AddToScheme,
		v1beta1secretmanager.SchemeBuilder.AddToScheme,
		v1beta1servicenetworking.SchemeBuilder.AddToScheme,
		v1beta1sourcerepo.SchemeBuilder.AddToScheme,
		v1beta1spanner.SchemeBuilder.AddToScheme,
		v1beta1sql.SchemeBuilder.AddToScheme,
		v1beta1storage.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
