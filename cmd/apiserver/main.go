/*
Copyright 2021 The KubeVela Authors.

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

package main

import (
	clusterv1alpha1 "github.com/zhouzhihu/kubeico-cluster-gateway/apis/cluster/v1alpha1"
	"github.com/zhouzhihu/kubeico-cluster-gateway/pkg/metrics"
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

func main() {

	// registering metrics
	metrics.Register()

	err := builder.APIServer.
		// +kubebuilder:scaffold:resource-register
		WithResource(&clusterv1alpha1.ClusterGateway{}).
		WithLocalDebugExtension().
		ExposeLoopbackClientConfig().
		ExposeLoopbackAuthorizer().
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
