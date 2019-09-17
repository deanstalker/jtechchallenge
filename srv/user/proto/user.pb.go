// Code generated by protoc-gen-go.
// source: srv/user/proto/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	srv/user/proto/user.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	BatchCreateRequest
	BatchCreateResponse
	GetRequest
	GetResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	UserItem
*/
package user

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type CreateRequest struct {
	User *UserItem `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}

func (m *CreateRequest) GetUser() *UserItem {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}

type BatchCreateRequest struct {
	Users []*UserItem `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *BatchCreateRequest) Reset()         { *m = BatchCreateRequest{} }
func (m *BatchCreateRequest) String() string { return proto.CompactTextString(m) }
func (*BatchCreateRequest) ProtoMessage()    {}

func (m *BatchCreateRequest) GetUsers() []*UserItem {
	if m != nil {
		return m.Users
	}
	return nil
}

type BatchCreateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *BatchCreateResponse) Reset()         { *m = BatchCreateResponse{} }
func (m *BatchCreateResponse) String() string { return proto.CompactTextString(m) }
func (*BatchCreateResponse) ProtoMessage()    {}

type GetRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}

type GetResponse struct {
	User *UserItem `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}

func (m *GetResponse) GetUser() *UserItem {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateRequest struct {
	User *UserItem `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}

func (m *UpdateRequest) GetUser() *UserItem {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}

type DeleteRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}

type DeleteResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}

type UserItem struct {
	Id         int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=last_name" json:"last_name,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Password   []byte `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Phone      string `protobuf:"bytes,7,opt,name=phone" json:"phone,omitempty"`
	UserStatus int64  `protobuf:"varint,8,opt,name=userStatus" json:"userStatus,omitempty"`
	Token      string `protobuf:"bytes,9,opt,name=token" json:"token,omitempty"`
}

func (m *UserItem) Reset()         { *m = UserItem{} }
func (m *UserItem) String() string { return proto.CompactTextString(m) }
func (*UserItem) ProtoMessage()    {}

func init() {
}
