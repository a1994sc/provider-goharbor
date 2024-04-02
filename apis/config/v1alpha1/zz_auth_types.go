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

type AuthInitParameters struct {

	// (String) Harbor authentication mode. Can be "oidc_auth", "db_auth" or "ldap_auth". (Default: "db_auth")
	AuthMode *string `json:"authMode,omitempty" tf:"auth_mode,omitempty"`

	// (String) A user's DN who has the permission to search the LDAP/AD server.
	LdapBaseDn *string `json:"ldapBaseDn,omitempty" tf:"ldap_base_dn,omitempty"`

	// (String) ldap filters
	LdapFilter *string `json:"ldapFilter,omitempty" tf:"ldap_filter,omitempty"`

	// (String) Specify an LDAP group DN. All LDAP user in this group will have harbor admin privilege
	LdapGroupAdminDn *string `json:"ldapGroupAdminDn,omitempty" tf:"ldap_group_admin_dn,omitempty"`

	// (String) The base DN from which to look up a group in LDAP/AD
	LdapGroupBaseDn *string `json:"ldapGroupBaseDn,omitempty" tf:"ldap_group_base_dn,omitempty"`

	// (String) The filter to look up an LDAP/AD group
	LdapGroupFilter *string `json:"ldapGroupFilter,omitempty" tf:"ldap_group_filter,omitempty"`

	// (String) The attribute used in a search to match a group
	LdapGroupGID *string `json:"ldapGroupGid,omitempty" tf:"ldap_group_gid,omitempty"`

	// (String) The attribute indicates the membership of LDAP group
	LdapGroupMembership *string `json:"ldapGroupMembership,omitempty" tf:"ldap_group_membership,omitempty"`

	// (String) The scope to search for groups
	LdapGroupScope *string `json:"ldapGroupScope,omitempty" tf:"ldap_group_scope,omitempty"`

	// (String) The attribute used in a search to match a user
	LdapGroupUID *string `json:"ldapGroupUid,omitempty" tf:"ldap_group_uid,omitempty"`

	// (String) LDAP Group Scope
	LdapScope *string `json:"ldapScope,omitempty" tf:"ldap_scope,omitempty"`

	// (String) The base DN from which to look up a user in LDAP/AD.
	LdapSearchDn *string `json:"ldapSearchDn,omitempty" tf:"ldap_search_dn,omitempty"`

	// (String) The attribute used in a search to match a user. It could be uid, cn, email, sAMAccountName or other attributes depending on your LDAP/AD.
	LdapUID *string `json:"ldapUid,omitempty" tf:"ldap_uid,omitempty"`

	// (String) The ldap server. Required when auth_mode is set to ldap.
	LdapURL *string `json:"ldapUrl,omitempty" tf:"ldap_url,omitempty"`

	// (Boolean) Verify Cert from LDAP Server.
	LdapVerifyCert *bool `json:"ldapVerifyCert,omitempty" tf:"ldap_verify_cert,omitempty"`

	// (String) All members of this group get Harbor admin permissions.
	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	// (Boolean) Default is "false", set to "true" if you want to skip the user onboarding screen, so user cannot change its username
	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	// (String) The client id for the oidc server.
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// complaint server.
	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	// (String) The OIDC group filter to filter which groups could be onboarded to Harbor.
	OidcGroupFilter *string `json:"oidcGroupFilter,omitempty" tf:"oidc_group_filter,omitempty"`

	// (String) The name of the claim in the token whose values is the list of group names.
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// (String) The name of the oidc provider name.
	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	// (String) The scope sent to OIDC server during authentication. It has to contain “openid”.
	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	// (String) Default is name
	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	// signed certificate.
	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	// (Boolean) Default is false, set to true if you want to use the OIDC or LDAP mode as the primary auth mode.
	PrimaryAuthMode *bool `json:"primaryAuthMode,omitempty" tf:"primary_auth_mode,omitempty"`
}

type AuthObservation struct {

	// (String) Harbor authentication mode. Can be "oidc_auth", "db_auth" or "ldap_auth". (Default: "db_auth")
	AuthMode *string `json:"authMode,omitempty" tf:"auth_mode,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) A user's DN who has the permission to search the LDAP/AD server.
	LdapBaseDn *string `json:"ldapBaseDn,omitempty" tf:"ldap_base_dn,omitempty"`

	// (String) ldap filters
	LdapFilter *string `json:"ldapFilter,omitempty" tf:"ldap_filter,omitempty"`

	// (String) Specify an LDAP group DN. All LDAP user in this group will have harbor admin privilege
	LdapGroupAdminDn *string `json:"ldapGroupAdminDn,omitempty" tf:"ldap_group_admin_dn,omitempty"`

	// (String) The base DN from which to look up a group in LDAP/AD
	LdapGroupBaseDn *string `json:"ldapGroupBaseDn,omitempty" tf:"ldap_group_base_dn,omitempty"`

	// (String) The filter to look up an LDAP/AD group
	LdapGroupFilter *string `json:"ldapGroupFilter,omitempty" tf:"ldap_group_filter,omitempty"`

	// (String) The attribute used in a search to match a group
	LdapGroupGID *string `json:"ldapGroupGid,omitempty" tf:"ldap_group_gid,omitempty"`

	// (String) The attribute indicates the membership of LDAP group
	LdapGroupMembership *string `json:"ldapGroupMembership,omitempty" tf:"ldap_group_membership,omitempty"`

	// (String) The scope to search for groups
	LdapGroupScope *string `json:"ldapGroupScope,omitempty" tf:"ldap_group_scope,omitempty"`

	// (String) The attribute used in a search to match a user
	LdapGroupUID *string `json:"ldapGroupUid,omitempty" tf:"ldap_group_uid,omitempty"`

	// (String) LDAP Group Scope
	LdapScope *string `json:"ldapScope,omitempty" tf:"ldap_scope,omitempty"`

	// (String) The base DN from which to look up a user in LDAP/AD.
	LdapSearchDn *string `json:"ldapSearchDn,omitempty" tf:"ldap_search_dn,omitempty"`

	// (String) The attribute used in a search to match a user. It could be uid, cn, email, sAMAccountName or other attributes depending on your LDAP/AD.
	LdapUID *string `json:"ldapUid,omitempty" tf:"ldap_uid,omitempty"`

	// (String) The ldap server. Required when auth_mode is set to ldap.
	LdapURL *string `json:"ldapUrl,omitempty" tf:"ldap_url,omitempty"`

	// (Boolean) Verify Cert from LDAP Server.
	LdapVerifyCert *bool `json:"ldapVerifyCert,omitempty" tf:"ldap_verify_cert,omitempty"`

	// (String) All members of this group get Harbor admin permissions.
	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	// (Boolean) Default is "false", set to "true" if you want to skip the user onboarding screen, so user cannot change its username
	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	// (String) The client id for the oidc server.
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// complaint server.
	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	// (String) The OIDC group filter to filter which groups could be onboarded to Harbor.
	OidcGroupFilter *string `json:"oidcGroupFilter,omitempty" tf:"oidc_group_filter,omitempty"`

	// (String) The name of the claim in the token whose values is the list of group names.
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// (String) The name of the oidc provider name.
	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	// (String) The scope sent to OIDC server during authentication. It has to contain “openid”.
	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	// (String) Default is name
	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	// signed certificate.
	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	// (Boolean) Default is false, set to true if you want to use the OIDC or LDAP mode as the primary auth mode.
	PrimaryAuthMode *bool `json:"primaryAuthMode,omitempty" tf:"primary_auth_mode,omitempty"`
}

type AuthParameters struct {

	// (String) Harbor authentication mode. Can be "oidc_auth", "db_auth" or "ldap_auth". (Default: "db_auth")
	// +kubebuilder:validation:Optional
	AuthMode *string `json:"authMode,omitempty" tf:"auth_mode,omitempty"`

	// (String) A user's DN who has the permission to search the LDAP/AD server.
	// +kubebuilder:validation:Optional
	LdapBaseDn *string `json:"ldapBaseDn,omitempty" tf:"ldap_base_dn,omitempty"`

	// (String) ldap filters
	// +kubebuilder:validation:Optional
	LdapFilter *string `json:"ldapFilter,omitempty" tf:"ldap_filter,omitempty"`

	// (String) Specify an LDAP group DN. All LDAP user in this group will have harbor admin privilege
	// +kubebuilder:validation:Optional
	LdapGroupAdminDn *string `json:"ldapGroupAdminDn,omitempty" tf:"ldap_group_admin_dn,omitempty"`

	// (String) The base DN from which to look up a group in LDAP/AD
	// +kubebuilder:validation:Optional
	LdapGroupBaseDn *string `json:"ldapGroupBaseDn,omitempty" tf:"ldap_group_base_dn,omitempty"`

	// (String) The filter to look up an LDAP/AD group
	// +kubebuilder:validation:Optional
	LdapGroupFilter *string `json:"ldapGroupFilter,omitempty" tf:"ldap_group_filter,omitempty"`

	// (String) The attribute used in a search to match a group
	// +kubebuilder:validation:Optional
	LdapGroupGID *string `json:"ldapGroupGid,omitempty" tf:"ldap_group_gid,omitempty"`

	// (String) The attribute indicates the membership of LDAP group
	// +kubebuilder:validation:Optional
	LdapGroupMembership *string `json:"ldapGroupMembership,omitempty" tf:"ldap_group_membership,omitempty"`

	// (String) The scope to search for groups
	// +kubebuilder:validation:Optional
	LdapGroupScope *string `json:"ldapGroupScope,omitempty" tf:"ldap_group_scope,omitempty"`

	// (String) The attribute used in a search to match a user
	// +kubebuilder:validation:Optional
	LdapGroupUID *string `json:"ldapGroupUid,omitempty" tf:"ldap_group_uid,omitempty"`

	// (String) LDAP Group Scope
	// +kubebuilder:validation:Optional
	LdapScope *string `json:"ldapScope,omitempty" tf:"ldap_scope,omitempty"`

	// (String) The base DN from which to look up a user in LDAP/AD.
	// +kubebuilder:validation:Optional
	LdapSearchDn *string `json:"ldapSearchDn,omitempty" tf:"ldap_search_dn,omitempty"`

	// (String, Sensitive) The password for the user that will perform the LDAP search
	// +kubebuilder:validation:Optional
	LdapSearchPasswordSecretRef *v1.SecretKeySelector `json:"ldapSearchPasswordSecretRef,omitempty" tf:"-"`

	// (String) The attribute used in a search to match a user. It could be uid, cn, email, sAMAccountName or other attributes depending on your LDAP/AD.
	// +kubebuilder:validation:Optional
	LdapUID *string `json:"ldapUid,omitempty" tf:"ldap_uid,omitempty"`

	// (String) The ldap server. Required when auth_mode is set to ldap.
	// +kubebuilder:validation:Optional
	LdapURL *string `json:"ldapUrl,omitempty" tf:"ldap_url,omitempty"`

	// (Boolean) Verify Cert from LDAP Server.
	// +kubebuilder:validation:Optional
	LdapVerifyCert *bool `json:"ldapVerifyCert,omitempty" tf:"ldap_verify_cert,omitempty"`

	// (String) All members of this group get Harbor admin permissions.
	// +kubebuilder:validation:Optional
	OidcAdminGroup *string `json:"oidcAdminGroup,omitempty" tf:"oidc_admin_group,omitempty"`

	// (Boolean) Default is "false", set to "true" if you want to skip the user onboarding screen, so user cannot change its username
	// +kubebuilder:validation:Optional
	OidcAutoOnboard *bool `json:"oidcAutoOnboard,omitempty" tf:"oidc_auto_onboard,omitempty"`

	// (String) The client id for the oidc server.
	// +kubebuilder:validation:Optional
	OidcClientID *string `json:"oidcClientId,omitempty" tf:"oidc_client_id,omitempty"`

	// (String, Sensitive) The client secert for the oidc server.
	// +kubebuilder:validation:Optional
	OidcClientSecretSecretRef *v1.SecretKeySelector `json:"oidcClientSecretSecretRef,omitempty" tf:"-"`

	// complaint server.
	// +kubebuilder:validation:Optional
	OidcEndpoint *string `json:"oidcEndpoint,omitempty" tf:"oidc_endpoint,omitempty"`

	// (String) The OIDC group filter to filter which groups could be onboarded to Harbor.
	// +kubebuilder:validation:Optional
	OidcGroupFilter *string `json:"oidcGroupFilter,omitempty" tf:"oidc_group_filter,omitempty"`

	// (String) The name of the claim in the token whose values is the list of group names.
	// +kubebuilder:validation:Optional
	OidcGroupsClaim *string `json:"oidcGroupsClaim,omitempty" tf:"oidc_groups_claim,omitempty"`

	// (String) The name of the oidc provider name.
	// +kubebuilder:validation:Optional
	OidcName *string `json:"oidcName,omitempty" tf:"oidc_name,omitempty"`

	// (String) The scope sent to OIDC server during authentication. It has to contain “openid”.
	// +kubebuilder:validation:Optional
	OidcScope *string `json:"oidcScope,omitempty" tf:"oidc_scope,omitempty"`

	// (String) Default is name
	// +kubebuilder:validation:Optional
	OidcUserClaim *string `json:"oidcUserClaim,omitempty" tf:"oidc_user_claim,omitempty"`

	// signed certificate.
	// +kubebuilder:validation:Optional
	OidcVerifyCert *bool `json:"oidcVerifyCert,omitempty" tf:"oidc_verify_cert,omitempty"`

	// (Boolean) Default is false, set to true if you want to use the OIDC or LDAP mode as the primary auth mode.
	// +kubebuilder:validation:Optional
	PrimaryAuthMode *bool `json:"primaryAuthMode,omitempty" tf:"primary_auth_mode,omitempty"`
}

// AuthSpec defines the desired state of Auth
type AuthSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthParameters `json:"forProvider"`
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
	InitProvider AuthInitParameters `json:"initProvider,omitempty"`
}

// AuthStatus defines the observed state of Auth.
type AuthStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Auth is the Schema for the Auths API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,goharbor}
type Auth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authMode) || (has(self.initProvider) && has(self.initProvider.authMode))",message="spec.forProvider.authMode is a required parameter"
	Spec   AuthSpec   `json:"spec"`
	Status AuthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthList contains a list of Auths
type AuthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Auth `json:"items"`
}

// Repository type metadata.
var (
	Auth_Kind             = "Auth"
	Auth_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Auth_Kind}.String()
	Auth_KindAPIVersion   = Auth_Kind + "." + CRDGroupVersion.String()
	Auth_GroupVersionKind = CRDGroupVersion.WithKind(Auth_Kind)
)

func init() {
	SchemeBuilder.Register(&Auth{}, &AuthList{})
}
