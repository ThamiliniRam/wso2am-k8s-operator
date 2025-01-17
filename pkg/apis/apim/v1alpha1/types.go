/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Apimanager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApimanagerSpec   `json:"spec"`
	Status ApimanagerStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type ApimanagerSpec struct {
	DeploymentName string             `json:"deploymentName"`
	Replicas       *int32             `json:"replicas"`
	ServiceName    string             `json:"serviceName"`
	ServType       corev1.ServiceType `json:"servType"`
	ExternalIps    []string           `json:"externalIps"`
	//NodePort       int32              `json:"nodePort"`
}

// FooStatus is the status for a Foo resource
type ApimanagerStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApimanagerList is a list of Apimanager resources
type ApimanagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Apimanager `json:"items"`
}
