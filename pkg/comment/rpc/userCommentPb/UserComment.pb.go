// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: UserComment.proto

package userCommentPb

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

// 用户对视频评论
type UpdateCommentStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId    int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`       //videoId
	UserId     int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`         //userId
	CommentId  int64 `protobuf:"varint,3,opt,name=commentId,proto3" json:"commentId,omitempty"`   //commentId
	ActionType int64 `protobuf:"varint,4,opt,name=actionType,proto3" json:"actionType,omitempty"` //评论 取消评论
}

func (x *UpdateCommentStatusReq) Reset() {
	*x = UpdateCommentStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserComment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentStatusReq) ProtoMessage() {}

func (x *UpdateCommentStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_UserComment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateCommentStatusReq) Descriptor() ([]byte, []int) {
	return file_UserComment_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCommentStatusReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *UpdateCommentStatusReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateCommentStatusReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *UpdateCommentStatusReq) GetActionType() int64 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type UpdateCommentStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentStatusResp) Reset() {
	*x = UpdateCommentStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserComment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentStatusResp) ProtoMessage() {}

func (x *UpdateCommentStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_UserComment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentStatusResp.ProtoReflect.Descriptor instead.
func (*UpdateCommentStatusResp) Descriptor() ([]byte, []int) {
	return file_UserComment_proto_rawDescGZIP(), []int{1}
}

// 获得视频的评论列表
type GetVideoCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"` //video_id
}

func (x *GetVideoCommentReq) Reset() {
	*x = GetVideoCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserComment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoCommentReq) ProtoMessage() {}

func (x *GetVideoCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_UserComment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoCommentReq.ProtoReflect.Descriptor instead.
func (*GetVideoCommentReq) Descriptor() ([]byte, []int) {
	return file_UserComment_proto_rawDescGZIP(), []int{2}
}

func (x *GetVideoCommentReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId  int64  `protobuf:"varint,1,opt,name=commentId,proto3" json:"commentId,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate string `protobuf:"bytes,4,opt,name=createDate,proto3" json:"createDate,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserComment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_UserComment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_UserComment_proto_rawDescGZIP(), []int{3}
}

func (x *Comment) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *Comment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type GetVideoCommentReqResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentList []*Comment `protobuf:"bytes,1,rep,name=commentList,proto3" json:"commentList,omitempty"` //comment_list
}

func (x *GetVideoCommentReqResp) Reset() {
	*x = GetVideoCommentReqResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UserComment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoCommentReqResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoCommentReqResp) ProtoMessage() {}

func (x *GetVideoCommentReqResp) ProtoReflect() protoreflect.Message {
	mi := &file_UserComment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoCommentReqResp.ProtoReflect.Descriptor instead.
func (*GetVideoCommentReqResp) Descriptor() ([]byte, []int) {
	return file_UserComment_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideoCommentReqResp) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

var File_UserComment_proto protoreflect.FileDescriptor

var file_UserComment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2e, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x79, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x47, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xa4, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_UserComment_proto_rawDescOnce sync.Once
	file_UserComment_proto_rawDescData = file_UserComment_proto_rawDesc
)

func file_UserComment_proto_rawDescGZIP() []byte {
	file_UserComment_proto_rawDescOnce.Do(func() {
		file_UserComment_proto_rawDescData = protoimpl.X.CompressGZIP(file_UserComment_proto_rawDescData)
	})
	return file_UserComment_proto_rawDescData
}

var file_UserComment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_UserComment_proto_goTypes = []interface{}{
	(*UpdateCommentStatusReq)(nil),  // 0: pb.UpdateCommentStatusReq
	(*UpdateCommentStatusResp)(nil), // 1: pb.UpdateCommentStatusResp
	(*GetVideoCommentReq)(nil),      // 2: pb.GetVideoCommentReq
	(*Comment)(nil),                 // 3: pb.Comment
	(*GetVideoCommentReqResp)(nil),  // 4: pb.GetVideoCommentReqResp
}
var file_UserComment_proto_depIdxs = []int32{
	3, // 0: pb.GetVideoCommentReqResp.commentList:type_name -> pb.Comment
	0, // 1: pb.UserComment.UpdateCommentStatus:input_type -> pb.UpdateCommentStatusReq
	2, // 2: pb.UserComment.GetVideoComment:input_type -> pb.GetVideoCommentReq
	1, // 3: pb.UserComment.UpdateCommentStatus:output_type -> pb.UpdateCommentStatusResp
	4, // 4: pb.UserComment.GetVideoComment:output_type -> pb.GetVideoCommentReqResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UserComment_proto_init() }
func file_UserComment_proto_init() {
	if File_UserComment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UserComment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentStatusReq); i {
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
		file_UserComment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentStatusResp); i {
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
		file_UserComment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoCommentReq); i {
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
		file_UserComment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_UserComment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoCommentReqResp); i {
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
			RawDescriptor: file_UserComment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_UserComment_proto_goTypes,
		DependencyIndexes: file_UserComment_proto_depIdxs,
		MessageInfos:      file_UserComment_proto_msgTypes,
	}.Build()
	File_UserComment_proto = out.File
	file_UserComment_proto_rawDesc = nil
	file_UserComment_proto_goTypes = nil
	file_UserComment_proto_depIdxs = nil
}
