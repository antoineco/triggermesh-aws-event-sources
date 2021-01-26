/*
Copyright (c) 2020-2021 TriggerMesh Inc.

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

package awssnssource

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/system"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	"github.com/triggermesh/aws-event-sources/pkg/apis/sources/v1alpha1"
	"github.com/triggermesh/aws-event-sources/pkg/reconciler/common"
	"github.com/triggermesh/aws-event-sources/pkg/reconciler/common/resource"
)

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"gcr.io/triggermesh/awssnssource"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

// Verify that Reconciler implements common.AdapterServiceBuilder.
var _ common.AdapterServiceBuilder = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterServiceBuilder.
func (r *Reconciler) BuildAdapter(src v1alpha1.EventSource, sinkURI *apis.URL) *servingv1.Service {
	typedSrc := src.(*v1alpha1.AWSSNSSource)

	return common.NewAdapterKnService(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVar(common.EnvARN, typedSrc.Spec.ARN.String()),
		resource.EnvVars(common.MakeSecurityCredentialsEnvVars(typedSrc.Spec.Credentials)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),

		// required by multi-tenant adapters to enable HA
		resource.EnvVar(system.NamespaceEnvKey, src.GetNamespace()),
	)
}

// RBACOwners implements common.AdapterServiceBuilder.
func (r *Reconciler) RBACOwners(namespace string) ([]kmeta.OwnerRefable, error) {
	srcs, err := r.srcLister(namespace).List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("listing objects from cache: %w", err)
	}

	ownerRefables := make([]kmeta.OwnerRefable, len(srcs))
	for i := range srcs {
		ownerRefables[i] = srcs[i]
	}

	return ownerRefables, nil
}
