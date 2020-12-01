// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: student.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// 学生数据结构
type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sno  uint64 `protobuf:"varint,1,opt,name=sno,proto3" json:"sno,omitempty" bson:"sno"`   // 学生编号
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`  // 学生名称
	Age  uint32 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty" bson:"age"`   // 学生年龄
	Male bool   `protobuf:"varint,4,opt,name=male,proto3" json:"male,omitempty" bson:"male"` // 性别为男
	Desc string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc"`  // 学生介绍
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetSno() uint64 {
	if x != nil {
		return x.Sno
	}
	return 0
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Student) GetMale() bool {
	if x != nil {
		return x.Male
	}
	return false
}

func (x *Student) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// 搜索请求
type AllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   int32  `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty" bson:"from"`    // 分页开始下标(从0开始)
	Size   int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty" bson:"size"`    // 分页大小
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty" bson:"search"` // 模糊查询的输入内容
	Field  string `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty" bson:"field"`   // 查询结果排序字段
	Desc   bool   `protobuf:"varint,5,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc"`    // 查询结果排序是否DESC
}

func (x *AllReq) Reset() {
	*x = AllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReq) ProtoMessage() {}

func (x *AllReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReq.ProtoReflect.Descriptor instead.
func (*AllReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *AllReq) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *AllReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AllReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *AllReq) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *AllReq) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

// 搜索响应
type AllRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty" bson:"total"` //查询结果总数
	Data  []*Student `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" bson:"data"`    // 查询结果数据
}

func (x *AllRsp) Reset() {
	*x = AllRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRsp) ProtoMessage() {}

func (x *AllRsp) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRsp.ProtoReflect.Descriptor instead.
func (*AllRsp) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *AllRsp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AllRsp) GetData() []*Student {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x69, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x6e,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0x72, 0x0a, 0x06, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x22, 0x40, 0x0a, 0x06, 0x41, 0x6c, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xbd, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x44,
	0x65, 0x6c, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x03, 0x55, 0x70, 0x64, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x6c, 0x6c, 0x52, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_student_proto_goTypes = []interface{}{
	(*Student)(nil), // 0: api.Student
	(*AllReq)(nil),  // 1: api.AllReq
	(*AllRsp)(nil),  // 2: api.AllRsp
}
var file_student_proto_depIdxs = []int32{
	0, // 0: api.AllRsp.data:type_name -> api.Student
	0, // 1: api.StudentService.Add:input_type -> api.Student
	0, // 2: api.StudentService.Del:input_type -> api.Student
	0, // 3: api.StudentService.Upd:input_type -> api.Student
	0, // 4: api.StudentService.Get:input_type -> api.Student
	1, // 5: api.StudentService.All:input_type -> api.AllReq
	0, // 6: api.StudentService.Add:output_type -> api.Student
	0, // 7: api.StudentService.Del:output_type -> api.Student
	0, // 8: api.StudentService.Upd:output_type -> api.Student
	0, // 9: api.StudentService.Get:output_type -> api.Student
	2, // 10: api.StudentService.All:output_type -> api.AllRsp
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReq); i {
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
		file_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRsp); i {
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
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}