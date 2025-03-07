// Copyright 2021 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: teleport/lib/teleterm/v1/cluster.proto

package teletermv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Cluster describes cluster fields
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uri is the cluster resource URI
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// name is used throughout the Teleport Connect codebase as the cluster name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// proxy address (only for root clusters)
	ProxyHost string `protobuf:"bytes,3,opt,name=proxy_host,json=proxyHost,proto3" json:"proxy_host,omitempty"`
	// connected indicates if connection to the cluster can be established, that is if we have a
	// cert for the cluster that hasn't expired
	Connected bool `protobuf:"varint,4,opt,name=connected,proto3" json:"connected,omitempty"`
	// leaf indicates if this is a leaf cluster
	Leaf bool `protobuf:"varint,5,opt,name=leaf,proto3" json:"leaf,omitempty"`
	// User is the cluster access control list of the logged-in user
	LoggedInUser *LoggedInUser `protobuf:"bytes,7,opt,name=logged_in_user,json=loggedInUser,proto3" json:"logged_in_user,omitempty"`
	// features describes the auth servers features.
	// Only present when detailed information is queried from the auth server.
	Features *Features `protobuf:"bytes,8,opt,name=features,proto3" json:"features,omitempty"`
	// auth_cluster_id is the unique cluster ID that is set once
	// during the first auth server startup.
	// Only present when detailed information is queried from the auth server.
	AuthClusterId string `protobuf:"bytes,9,opt,name=auth_cluster_id,json=authClusterId,proto3" json:"auth_cluster_id,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetProxyHost() string {
	if x != nil {
		return x.ProxyHost
	}
	return ""
}

func (x *Cluster) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *Cluster) GetLeaf() bool {
	if x != nil {
		return x.Leaf
	}
	return false
}

func (x *Cluster) GetLoggedInUser() *LoggedInUser {
	if x != nil {
		return x.LoggedInUser
	}
	return nil
}

func (x *Cluster) GetFeatures() *Features {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Cluster) GetAuthClusterId() string {
	if x != nil {
		return x.AuthClusterId
	}
	return ""
}

// LoggedInUser describes a logged-in user
type LoggedInUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the user name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// roles is the user roles
	Roles []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	// ssh_logins is the user ssh logins
	SshLogins []string `protobuf:"bytes,3,rep,name=ssh_logins,json=sshLogins,proto3" json:"ssh_logins,omitempty"`
	// acl is the user acl
	Acl *ACL `protobuf:"bytes,4,opt,name=acl,proto3" json:"acl,omitempty"`
	// active_requests is an array of request-id strings of active requests
	ActiveRequests []string `protobuf:"bytes,5,rep,name=active_requests,json=activeRequests,proto3" json:"active_requests,omitempty"`
	// suggested_reviewers for the given user.
	// Only present when detailed information is queried from the auth server.
	SuggestedReviewers []string `protobuf:"bytes,6,rep,name=suggested_reviewers,json=suggestedReviewers,proto3" json:"suggested_reviewers,omitempty"`
	// requestable_roles for the given user.
	// Only present when detailed information is queried from the auth server.
	RequestableRoles []string `protobuf:"bytes,7,rep,name=requestable_roles,json=requestableRoles,proto3" json:"requestable_roles,omitempty"`
}

func (x *LoggedInUser) Reset() {
	*x = LoggedInUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggedInUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggedInUser) ProtoMessage() {}

func (x *LoggedInUser) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggedInUser.ProtoReflect.Descriptor instead.
func (*LoggedInUser) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *LoggedInUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoggedInUser) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *LoggedInUser) GetSshLogins() []string {
	if x != nil {
		return x.SshLogins
	}
	return nil
}

func (x *LoggedInUser) GetAcl() *ACL {
	if x != nil {
		return x.Acl
	}
	return nil
}

func (x *LoggedInUser) GetActiveRequests() []string {
	if x != nil {
		return x.ActiveRequests
	}
	return nil
}

func (x *LoggedInUser) GetSuggestedReviewers() []string {
	if x != nil {
		return x.SuggestedReviewers
	}
	return nil
}

func (x *LoggedInUser) GetRequestableRoles() []string {
	if x != nil {
		return x.RequestableRoles
	}
	return nil
}

// ACL is the access control list of the user
type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// auth_connectors defines access to auth.connectors
	AuthConnectors *ResourceAccess `protobuf:"bytes,2,opt,name=auth_connectors,json=authConnectors,proto3" json:"auth_connectors,omitempty"`
	// Roles defines access to roles
	Roles *ResourceAccess `protobuf:"bytes,3,opt,name=roles,proto3" json:"roles,omitempty"`
	// Users defines access to users.
	Users *ResourceAccess `protobuf:"bytes,4,opt,name=users,proto3" json:"users,omitempty"`
	// trusted_clusters defines access to trusted clusters
	TrustedClusters *ResourceAccess `protobuf:"bytes,5,opt,name=trusted_clusters,json=trustedClusters,proto3" json:"trusted_clusters,omitempty"`
	// Events defines access to audit logs
	Events *ResourceAccess `protobuf:"bytes,6,opt,name=events,proto3" json:"events,omitempty"`
	// Tokens defines access to tokens.
	Tokens *ResourceAccess `protobuf:"bytes,7,opt,name=tokens,proto3" json:"tokens,omitempty"`
	// Servers defines access to servers.
	Servers *ResourceAccess `protobuf:"bytes,8,opt,name=servers,proto3" json:"servers,omitempty"`
	// apps defines access to application servers
	Apps *ResourceAccess `protobuf:"bytes,9,opt,name=apps,proto3" json:"apps,omitempty"`
	// dbs defines access to database servers.
	Dbs *ResourceAccess `protobuf:"bytes,10,opt,name=dbs,proto3" json:"dbs,omitempty"`
	// kubeservers defines access to kubernetes servers.
	Kubeservers *ResourceAccess `protobuf:"bytes,11,opt,name=kubeservers,proto3" json:"kubeservers,omitempty"`
	// access_requests defines access to access requests
	AccessRequests *ResourceAccess `protobuf:"bytes,12,opt,name=access_requests,json=accessRequests,proto3" json:"access_requests,omitempty"`
	// recorded_sessions defines access to recorded sessions.
	RecordedSessions *ResourceAccess `protobuf:"bytes,13,opt,name=recorded_sessions,json=recordedSessions,proto3" json:"recorded_sessions,omitempty"`
	// active_sessions defines access to active sessions.
	ActiveSessions *ResourceAccess `protobuf:"bytes,14,opt,name=active_sessions,json=activeSessions,proto3" json:"active_sessions,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ACL) GetAuthConnectors() *ResourceAccess {
	if x != nil {
		return x.AuthConnectors
	}
	return nil
}

func (x *ACL) GetRoles() *ResourceAccess {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *ACL) GetUsers() *ResourceAccess {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ACL) GetTrustedClusters() *ResourceAccess {
	if x != nil {
		return x.TrustedClusters
	}
	return nil
}

func (x *ACL) GetEvents() *ResourceAccess {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ACL) GetTokens() *ResourceAccess {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *ACL) GetServers() *ResourceAccess {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *ACL) GetApps() *ResourceAccess {
	if x != nil {
		return x.Apps
	}
	return nil
}

func (x *ACL) GetDbs() *ResourceAccess {
	if x != nil {
		return x.Dbs
	}
	return nil
}

func (x *ACL) GetKubeservers() *ResourceAccess {
	if x != nil {
		return x.Kubeservers
	}
	return nil
}

func (x *ACL) GetAccessRequests() *ResourceAccess {
	if x != nil {
		return x.AccessRequests
	}
	return nil
}

func (x *ACL) GetRecordedSessions() *ResourceAccess {
	if x != nil {
		return x.RecordedSessions
	}
	return nil
}

func (x *ACL) GetActiveSessions() *ResourceAccess {
	if x != nil {
		return x.ActiveSessions
	}
	return nil
}

// ResourceAccess describes access verbs
type ResourceAccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list determines "list" access
	List bool `protobuf:"varint,1,opt,name=list,proto3" json:"list,omitempty"`
	// read determines "read" access
	Read bool `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
	// edit determines "edit" access
	Edit bool `protobuf:"varint,3,opt,name=edit,proto3" json:"edit,omitempty"`
	// create determines "create" access
	Create bool `protobuf:"varint,4,opt,name=create,proto3" json:"create,omitempty"`
	// delete determines "delete" access
	Delete bool `protobuf:"varint,5,opt,name=delete,proto3" json:"delete,omitempty"`
	// use determines "use" access
	Use bool `protobuf:"varint,6,opt,name=use,proto3" json:"use,omitempty"`
}

func (x *ResourceAccess) Reset() {
	*x = ResourceAccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAccess) ProtoMessage() {}

func (x *ResourceAccess) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAccess.ProtoReflect.Descriptor instead.
func (*ResourceAccess) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceAccess) GetList() bool {
	if x != nil {
		return x.List
	}
	return false
}

func (x *ResourceAccess) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *ResourceAccess) GetEdit() bool {
	if x != nil {
		return x.Edit
	}
	return false
}

func (x *ResourceAccess) GetCreate() bool {
	if x != nil {
		return x.Create
	}
	return false
}

func (x *ResourceAccess) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *ResourceAccess) GetUse() bool {
	if x != nil {
		return x.Use
	}
	return false
}

// Features describes the auth servers features
type Features struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// advanced_access_workflows enables search-based access requests
	AdvancedAccessWorkflows bool `protobuf:"varint,1,opt,name=advanced_access_workflows,json=advancedAccessWorkflows,proto3" json:"advanced_access_workflows,omitempty"`
	// is_usage_based_billing determines if the cloud user subscription is usage-based (pay-as-you-go).
	IsUsageBasedBilling bool `protobuf:"varint,2,opt,name=is_usage_based_billing,json=isUsageBasedBilling,proto3" json:"is_usage_based_billing,omitempty"`
}

func (x *Features) Reset() {
	*x = Features{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Features) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Features) ProtoMessage() {}

func (x *Features) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Features.ProtoReflect.Descriptor instead.
func (*Features) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *Features) GetAdvancedAccessWorkflows() bool {
	if x != nil {
		return x.AdvancedAccessWorkflows
	}
	return false
}

func (x *Features) GetIsUsageBasedBilling() bool {
	if x != nil {
		return x.IsUsageBasedBilling
	}
	return false
}

var File_teleport_lib_teleterm_v1_cluster_proto protoreflect.FileDescriptor

var file_teleport_lib_teleterm_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x22, 0xb6, 0x02, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6c, 0x65, 0x61, 0x66, 0x12, 0x4c, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8f, 0x02, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x73, 0x68, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x73, 0x68, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x03, 0x61, 0x63, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69,
	0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43,
	0x4c, 0x52, 0x03, 0x61, 0x63, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xc8, 0x07,
	0x0a, 0x03, 0x41, 0x43, 0x4c, 0x12, 0x51, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x10, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69,
	0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0f, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x40, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x42, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69,
	0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c,
	0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x04, 0x61,
	0x70, 0x70, 0x73, 0x12, 0x3a, 0x0a, 0x03, 0x64, 0x62, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x03, 0x64, 0x62, 0x73, 0x12,
	0x4a, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0b,
	0x6b, 0x75, 0x62, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x55,
	0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x08,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x65, 0x64, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x75, 0x73, 0x65, 0x22, 0x7b, 0x0a, 0x08, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x69, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_lib_teleterm_v1_cluster_proto_rawDescOnce sync.Once
	file_teleport_lib_teleterm_v1_cluster_proto_rawDescData = file_teleport_lib_teleterm_v1_cluster_proto_rawDesc
)

func file_teleport_lib_teleterm_v1_cluster_proto_rawDescGZIP() []byte {
	file_teleport_lib_teleterm_v1_cluster_proto_rawDescOnce.Do(func() {
		file_teleport_lib_teleterm_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_lib_teleterm_v1_cluster_proto_rawDescData)
	})
	return file_teleport_lib_teleterm_v1_cluster_proto_rawDescData
}

var file_teleport_lib_teleterm_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_lib_teleterm_v1_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil),        // 0: teleport.lib.teleterm.v1.Cluster
	(*LoggedInUser)(nil),   // 1: teleport.lib.teleterm.v1.LoggedInUser
	(*ACL)(nil),            // 2: teleport.lib.teleterm.v1.ACL
	(*ResourceAccess)(nil), // 3: teleport.lib.teleterm.v1.ResourceAccess
	(*Features)(nil),       // 4: teleport.lib.teleterm.v1.Features
}
var file_teleport_lib_teleterm_v1_cluster_proto_depIdxs = []int32{
	1,  // 0: teleport.lib.teleterm.v1.Cluster.logged_in_user:type_name -> teleport.lib.teleterm.v1.LoggedInUser
	4,  // 1: teleport.lib.teleterm.v1.Cluster.features:type_name -> teleport.lib.teleterm.v1.Features
	2,  // 2: teleport.lib.teleterm.v1.LoggedInUser.acl:type_name -> teleport.lib.teleterm.v1.ACL
	3,  // 3: teleport.lib.teleterm.v1.ACL.auth_connectors:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 4: teleport.lib.teleterm.v1.ACL.roles:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 5: teleport.lib.teleterm.v1.ACL.users:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 6: teleport.lib.teleterm.v1.ACL.trusted_clusters:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 7: teleport.lib.teleterm.v1.ACL.events:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 8: teleport.lib.teleterm.v1.ACL.tokens:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 9: teleport.lib.teleterm.v1.ACL.servers:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 10: teleport.lib.teleterm.v1.ACL.apps:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 11: teleport.lib.teleterm.v1.ACL.dbs:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 12: teleport.lib.teleterm.v1.ACL.kubeservers:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 13: teleport.lib.teleterm.v1.ACL.access_requests:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 14: teleport.lib.teleterm.v1.ACL.recorded_sessions:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	3,  // 15: teleport.lib.teleterm.v1.ACL.active_sessions:type_name -> teleport.lib.teleterm.v1.ResourceAccess
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_teleport_lib_teleterm_v1_cluster_proto_init() }
func file_teleport_lib_teleterm_v1_cluster_proto_init() {
	if File_teleport_lib_teleterm_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggedInUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceAccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_lib_teleterm_v1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Features); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_lib_teleterm_v1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_lib_teleterm_v1_cluster_proto_goTypes,
		DependencyIndexes: file_teleport_lib_teleterm_v1_cluster_proto_depIdxs,
		MessageInfos:      file_teleport_lib_teleterm_v1_cluster_proto_msgTypes,
	}.Build()
	File_teleport_lib_teleterm_v1_cluster_proto = out.File
	file_teleport_lib_teleterm_v1_cluster_proto_rawDesc = nil
	file_teleport_lib_teleterm_v1_cluster_proto_goTypes = nil
	file_teleport_lib_teleterm_v1_cluster_proto_depIdxs = nil
}
