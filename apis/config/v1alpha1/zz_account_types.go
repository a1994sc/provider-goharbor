// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessInitParameters struct {

	// (String) Eg. push, pull, read, etc. Check available actions.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Either allow or deny. Defaults to allow.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// (String) Eg. repository, labels, etc. Check available resources.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type AccessObservation struct {

	// (String) Eg. push, pull, read, etc. Check available actions.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Either allow or deny. Defaults to allow.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// (String) Eg. repository, labels, etc. Check available resources.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`
}

type AccessParameters struct {

	// (String) Eg. push, pull, read, etc. Check available actions.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// (String) Either allow or deny. Defaults to allow.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// (String) Eg. repository, labels, etc. Check available resources.
	// +kubebuilder:validation:Optional
	Resource *string `json:"resource" tf:"resource,omitempty"`
}

type AccountInitParameters struct {

	// (String) The description of the robot account will be displayed in harbor.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Disables the robot account when set to true.
	Disable *bool `json:"disable,omitempty" tf:"disable,omitempty"`

	// (Number) By default, the robot account will not expire. Set it to the amount of days until the account should expire.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) Level of the robot account, currently either system or project.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The name of the project that will be created in harbor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Permissions []PermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type AccountObservation struct {

	// (String) The description of the robot account will be displayed in harbor.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Disables the robot account when set to true.
	Disable *bool `json:"disable,omitempty" tf:"disable,omitempty"`

	// (Number) By default, the robot account will not expire. Set it to the amount of days until the account should expire.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String)
	FullName *string `json:"fullName,omitempty" tf:"full_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Level of the robot account, currently either system or project.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The name of the project that will be created in harbor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Permissions []PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// (String)
	RobotID *string `json:"robotId,omitempty" tf:"robot_id,omitempty"`
}

type AccountParameters struct {

	// (String) The description of the robot account will be displayed in harbor.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Disables the robot account when set to true.
	// +kubebuilder:validation:Optional
	Disable *bool `json:"disable,omitempty" tf:"disable,omitempty"`

	// (Number) By default, the robot account will not expire. Set it to the amount of days until the account should expire.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) Level of the robot account, currently either system or project.
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The name of the project that will be created in harbor.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// (String, Sensitive) The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`
}

type PermissionsInitParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Access []AccessInitParameters `json:"access,omitempty" tf:"access,omitempty"`

	// (String) Either system or project.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) namespace is the name of your project. For kind system permissions, always use / as namespace. Use * to match all projects.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type PermissionsObservation struct {

	// (Block Set, Min: 1) (see below for nested schema)
	Access []AccessObservation `json:"access,omitempty" tf:"access,omitempty"`

	// (String) Either system or project.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) namespace is the name of your project. For kind system permissions, always use / as namespace. Use * to match all projects.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type PermissionsParameters struct {

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Access []AccessParameters `json:"access" tf:"access,omitempty"`

	// (String) Either system or project.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// (String) namespace is the name of your project. For kind system permissions, always use / as namespace. Use * to match all projects.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccountInitParameters `json:"initProvider,omitempty"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,goharbor}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.level) || (has(self.initProvider) && has(self.initProvider.level))",message="spec.forProvider.level is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
