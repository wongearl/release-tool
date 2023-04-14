/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImagerSpec defines the desired state of Imager
type ImagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	RepoUrl string `json:"repoUrl"`
	// +optional
	RepoRemote          string `json:"repoRemote"`
	RepoSecretName      string `json:"repoSecretName"`
	RepoSecretNamespace string `json:"repoSecretNamespace"`
	// +optional
	RepoBranch string `json:"repoBranch"`
	RepoTag    string `json:"repoTag"`
	// +optional
	RepoPr                  string   `json:"repoPr"`
	ImageName               string   `json:"imageName"`
	DockerfilePath          string   `json:"dockerfilePath"`
	Registry                string   `json:"registry"`
	RegistrySecretName      string   `json:"registrySecretName"`
	RegistrySecretNamespace string   `json:"registrySecretNamespace"`
	BuildArgs               []string `json:"buildArgs"`
}

// ImagerStatus defines the observed state of Imager
type ImagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	Conditions     []Condition  `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Imager is the Schema for the imagers API
type Imager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagerSpec   `json:"spec,omitempty"`
	Status ImagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImagerList contains a list of Imager
type ImagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Imager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Imager{}, &ImagerList{})
}
