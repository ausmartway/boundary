// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: controller/api/resources/sessions/v1/session.proto

package sessions

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	scopes "github.com/hashicorp/boundary/internal/gen/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/internal/gen/controller/api/resources/targets"
	_ "github.com/hashicorp/boundary/internal/gen/controller/protooptions"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type WorkerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the worker
	Address string `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *WorkerInfo) Reset() {
	*x = WorkerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerInfo) ProtoMessage() {}

func (x *WorkerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerInfo.ProtoReflect.Descriptor instead.
func (*WorkerInfo) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_sessions_v1_session_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type SessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the status of the session.  For example "pending", "active",
	// "canceling", "terminated".
	Status string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	// The time the session entered this state.
	// Output only.
	StartTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=start_time,proto3" json:"start_time,omitempty"`
	// The time the session stopped being in this state.
	// Output only.
	EndTime *timestamp.Timestamp `protobuf:"bytes,30,opt,name=end_time,proto3" json:"end_time,omitempty"`
}

func (x *SessionState) Reset() {
	*x = SessionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionState) ProtoMessage() {}

func (x *SessionState) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionState.ProtoReflect.Descriptor instead.
func (*SessionState) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_sessions_v1_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionState) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SessionState) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SessionState) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Session contains all fields related to a Session resource
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the resource
	// Output only.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the parent of this resource.  This must be defined for creation
	// of this resource, but is otherwise read only.
	TargetId string `protobuf:"bytes,20,opt,name=target_id,proto3" json:"target_id,omitempty"`
	// Scope information for this resource
	// Output only.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// The time this resource was created
	// Output only.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,40,opt,name=created_time,proto3" json:"created_time,omitempty"`
	// The time this resource was last updated.
	// Output only.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,50,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	// After this time the connection will be expired, e.g. forcefully terminated
	ExpirationTime *timestamp.Timestamp `protobuf:"bytes,60,opt,name=expiration_time,proto3" json:"expiration_time,omitempty"`
	// The version can be used in subsequent write requests to ensure this
	// resource has not changed and to fail the write if it has.
	Version uint32 `protobuf:"varint,70,opt,name=version,proto3" json:"version,omitempty"`
	// Type of the session (e.g. tcp, ssh, etc.)
	Type string `protobuf:"bytes,80,opt,name=type,proto3" json:"type,omitempty"`
	// The id of the token used to authenticate.
	AuthTokenId string `protobuf:"bytes,90,opt,name=auth_token_id,proto3" json:"auth_token_id,omitempty"`
	// The ID of the user that requested the session
	UserId    string `protobuf:"bytes,100,opt,name=user_id,proto3" json:"user_id,omitempty"`
	HostSetId string `protobuf:"bytes,110,opt,name=host_set_id,proto3" json:"host_set_id,omitempty"`
	HostId    string `protobuf:"bytes,120,opt,name=host_id,proto3" json:"host_id,omitempty"`
	ScopeId   string `protobuf:"bytes,130,opt,name=scope_id,proto3" json:"scope_id,omitempty"`
	Endpoint  string `protobuf:"bytes,140,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// The states of this session in descending from the current state to the
	// first.
	States []*SessionState `protobuf:"bytes,150,rep,name=states,proto3" json:"states,omitempty"`
	// The name of the current status of this session.
	Status string `protobuf:"bytes,160,opt,name=status,proto3" json:"status,omitempty"`
	// Worker information. The first worker in the slice should be prioritized.
	WorkerInfo []*WorkerInfo `protobuf:"bytes,170,rep,name=worker_info,proto3" json:"worker_info,omitempty"`
	// Deprecated, to be moved to another proto by another PR:
	// The certificate to use when connecting (or if using custom certs, to
	// serve as the "login"). Raw DER bytes.
	Certificate []byte `protobuf:"bytes,220,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// The private key to use when connecting (or if using custom certs, to pass
	// as the "password").
	PrivateKey []byte `protobuf:"bytes,230,opt,name=private_key,proto3" json:"private_key,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_sessions_v1_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_sessions_v1_session_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Session) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Session) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Session) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Session) GetExpirationTime() *timestamp.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *Session) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Session) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Session) GetAuthTokenId() string {
	if x != nil {
		return x.AuthTokenId
	}
	return ""
}

func (x *Session) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Session) GetHostSetId() string {
	if x != nil {
		return x.HostSetId
	}
	return ""
}

func (x *Session) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Session) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Session) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Session) GetStates() []*SessionState {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *Session) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Session) GetWorkerInfo() []*WorkerInfo {
	if x != nil {
		return x.WorkerInfo
	}
	return nil
}

func (x *Session) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Session) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

var File_controller_api_resources_sessions_v1_session_proto protoreflect.FileDescriptor

var file_controller_api_resources_sessions_v1_session_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x9a, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa7, 0x06, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x46, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x6e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x96, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xa0, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xaa, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x21, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0xdc, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0xe6, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_sessions_v1_session_proto_rawDescOnce sync.Once
	file_controller_api_resources_sessions_v1_session_proto_rawDescData = file_controller_api_resources_sessions_v1_session_proto_rawDesc
)

func file_controller_api_resources_sessions_v1_session_proto_rawDescGZIP() []byte {
	file_controller_api_resources_sessions_v1_session_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_sessions_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_sessions_v1_session_proto_rawDescData)
	})
	return file_controller_api_resources_sessions_v1_session_proto_rawDescData
}

var file_controller_api_resources_sessions_v1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_api_resources_sessions_v1_session_proto_goTypes = []interface{}{
	(*WorkerInfo)(nil),          // 0: controller.api.resources.sessions.v1.WorkerInfo
	(*SessionState)(nil),        // 1: controller.api.resources.sessions.v1.SessionState
	(*Session)(nil),             // 2: controller.api.resources.sessions.v1.Session
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*scopes.ScopeInfo)(nil),    // 4: controller.api.resources.scopes.v1.ScopeInfo
}
var file_controller_api_resources_sessions_v1_session_proto_depIdxs = []int32{
	3, // 0: controller.api.resources.sessions.v1.SessionState.start_time:type_name -> google.protobuf.Timestamp
	3, // 1: controller.api.resources.sessions.v1.SessionState.end_time:type_name -> google.protobuf.Timestamp
	4, // 2: controller.api.resources.sessions.v1.Session.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	3, // 3: controller.api.resources.sessions.v1.Session.created_time:type_name -> google.protobuf.Timestamp
	3, // 4: controller.api.resources.sessions.v1.Session.updated_time:type_name -> google.protobuf.Timestamp
	3, // 5: controller.api.resources.sessions.v1.Session.expiration_time:type_name -> google.protobuf.Timestamp
	1, // 6: controller.api.resources.sessions.v1.Session.states:type_name -> controller.api.resources.sessions.v1.SessionState
	0, // 7: controller.api.resources.sessions.v1.Session.worker_info:type_name -> controller.api.resources.sessions.v1.WorkerInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_controller_api_resources_sessions_v1_session_proto_init() }
func file_controller_api_resources_sessions_v1_session_proto_init() {
	if File_controller_api_resources_sessions_v1_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_sessions_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerInfo); i {
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
		file_controller_api_resources_sessions_v1_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionState); i {
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
		file_controller_api_resources_sessions_v1_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_controller_api_resources_sessions_v1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_sessions_v1_session_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_sessions_v1_session_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_sessions_v1_session_proto_msgTypes,
	}.Build()
	File_controller_api_resources_sessions_v1_session_proto = out.File
	file_controller_api_resources_sessions_v1_session_proto_rawDesc = nil
	file_controller_api_resources_sessions_v1_session_proto_goTypes = nil
	file_controller_api_resources_sessions_v1_session_proto_depIdxs = nil
}
