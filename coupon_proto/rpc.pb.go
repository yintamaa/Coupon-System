// 声明protobuf版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: rpc.proto

package coupon_proto

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

type GetCouponListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *GetCouponListReq) Reset() {
	*x = GetCouponListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponListReq) ProtoMessage() {}

func (x *GetCouponListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponListReq.ProtoReflect.Descriptor instead.
func (*GetCouponListReq) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetCouponListReq) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type GetCouponListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponList *Coupon `protobuf:"bytes,1,opt,name=CouponList,proto3" json:"CouponList,omitempty"`
}

func (x *GetCouponListResp) Reset() {
	*x = GetCouponListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponListResp) ProtoMessage() {}

func (x *GetCouponListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponListResp.ProtoReflect.Descriptor instead.
func (*GetCouponListResp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetCouponListResp) GetCouponList() *Coupon {
	if x != nil {
		return x.CouponList
	}
	return nil
}

type GetCouponReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponID int32 `protobuf:"varint,1,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
}

func (x *GetCouponReq) Reset() {
	*x = GetCouponReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponReq) ProtoMessage() {}

func (x *GetCouponReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponReq.ProtoReflect.Descriptor instead.
func (*GetCouponReq) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetCouponReq) GetCouponID() int32 {
	if x != nil {
		return x.CouponID
	}
	return 0
}

type GetCouponResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupon *Coupon `protobuf:"bytes,1,opt,name=Coupon,proto3" json:"Coupon,omitempty"`
}

func (x *GetCouponResp) Reset() {
	*x = GetCouponResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponResp) ProtoMessage() {}

func (x *GetCouponResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponResp.ProtoReflect.Descriptor instead.
func (*GetCouponResp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetCouponResp) GetCoupon() *Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x2a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a,
	0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x32, 0x97, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2e, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_proto_goTypes = []interface{}{
	(*GetCouponListReq)(nil),  // 0: pb.GetCouponListReq
	(*GetCouponListResp)(nil), // 1: pb.GetCouponListResp
	(*GetCouponReq)(nil),      // 2: pb.GetCouponReq
	(*GetCouponResp)(nil),     // 3: pb.GetCouponResp
	(*Coupon)(nil),            // 4: pb.Coupon
}
var file_rpc_proto_depIdxs = []int32{
	4, // 0: pb.GetCouponListResp.CouponList:type_name -> pb.Coupon
	4, // 1: pb.GetCouponResp.Coupon:type_name -> pb.Coupon
	0, // 2: pb.CouponsService.GetCouponsByAccountID:input_type -> pb.GetCouponListReq
	2, // 3: pb.CouponsService.GetCouponsByCouponID:input_type -> pb.GetCouponReq
	1, // 4: pb.CouponsService.GetCouponsByAccountID:output_type -> pb.GetCouponListResp
	3, // 5: pb.CouponsService.GetCouponsByCouponID:output_type -> pb.GetCouponResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponListReq); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponListResp); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponReq); i {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponResp); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
