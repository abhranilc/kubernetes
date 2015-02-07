/*
Copyright 2015 Google Inc. All rights reserved.

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
	kubeapiserver "github.com/GoogleCloudPlatform/kubernetes/cmd/kube-apiserver/app"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/hyperkube"
)

// NewKubeAPIServer creates a new hyperkube Server object that includes the
// description and flags.
func NewKubeAPIServer() *hyperkube.Server {
	s := kubeapiserver.NewAPIServer()

	hks := hyperkube.Server{
		SimpleUsage: "apiserver",
		Long:        "The main API entrypoint and interface to the storage system.  The API server is also the focal point for all authorization decisions.",
		Run: func(_ *hyperkube.Server, args []string) error {
			return s.Run(args)
		},
	}
	s.AddFlags(hks.Flags())
	return &hks
}
