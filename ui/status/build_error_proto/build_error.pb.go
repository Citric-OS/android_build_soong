// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build_error.proto

package soong_build_error_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuildError struct {
	// List of error messages of the overall build. The error messages
	// are not associated with a build action.
	ErrorMessages []string `protobuf:"bytes,1,rep,name=error_messages,json=errorMessages" json:"error_messages,omitempty"`
	// List of build action errors.
	ActionErrors         []*BuildActionError `protobuf:"bytes,2,rep,name=action_errors,json=actionErrors" json:"action_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BuildError) Reset()         { *m = BuildError{} }
func (m *BuildError) String() string { return proto.CompactTextString(m) }
func (*BuildError) ProtoMessage()    {}
func (*BuildError) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e15b05802a5501, []int{0}
}

func (m *BuildError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildError.Unmarshal(m, b)
}
func (m *BuildError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildError.Marshal(b, m, deterministic)
}
func (m *BuildError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildError.Merge(m, src)
}
func (m *BuildError) XXX_Size() int {
	return xxx_messageInfo_BuildError.Size(m)
}
func (m *BuildError) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildError.DiscardUnknown(m)
}

var xxx_messageInfo_BuildError proto.InternalMessageInfo

func (m *BuildError) GetErrorMessages() []string {
	if m != nil {
		return m.ErrorMessages
	}
	return nil
}

func (m *BuildError) GetActionErrors() []*BuildActionError {
	if m != nil {
		return m.ActionErrors
	}
	return nil
}

// Build is composed of a list of build action. There can be a set of build
// actions that can failed.
type BuildActionError struct {
	// Description of the command.
	Description *string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// The command name that raised the error.
	Command *string `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	// The command output stream.
	Output *string `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	// List of artifacts (i.e. files) that was produced by the command.
	Artifacts []string `protobuf:"bytes,4,rep,name=artifacts" json:"artifacts,omitempty"`
	// The error string produced by the build action.
	Error                *string  `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildActionError) Reset()         { *m = BuildActionError{} }
func (m *BuildActionError) String() string { return proto.CompactTextString(m) }
func (*BuildActionError) ProtoMessage()    {}
func (*BuildActionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e15b05802a5501, []int{1}
}

func (m *BuildActionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildActionError.Unmarshal(m, b)
}
func (m *BuildActionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildActionError.Marshal(b, m, deterministic)
}
func (m *BuildActionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildActionError.Merge(m, src)
}
func (m *BuildActionError) XXX_Size() int {
	return xxx_messageInfo_BuildActionError.Size(m)
}
func (m *BuildActionError) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildActionError.DiscardUnknown(m)
}

var xxx_messageInfo_BuildActionError proto.InternalMessageInfo

func (m *BuildActionError) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *BuildActionError) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

func (m *BuildActionError) GetOutput() string {
	if m != nil && m.Output != nil {
		return *m.Output
	}
	return ""
}

func (m *BuildActionError) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *BuildActionError) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildError)(nil), "soong_build_error.BuildError")
	proto.RegisterType((*BuildActionError)(nil), "soong_build_error.BuildActionError")
}

func init() { proto.RegisterFile("build_error.proto", fileDescriptor_a2e15b05802a5501) }

var fileDescriptor_a2e15b05802a5501 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x63, 0x95, 0x4c, 0xad, 0xd8, 0x41, 0x74, 0x04, 0x0f, 0xa1, 0x22, 0xe4, 0x94,
	0x83, 0x6f, 0x60, 0x41, 0xf0, 0xe2, 0x25, 0x47, 0x2f, 0x61, 0xdd, 0xac, 0x65, 0xc1, 0x64, 0xc2,
	0xce, 0xe6, 0xe8, 0x8b, 0xf8, 0xb4, 0x92, 0x69, 0xa5, 0xa5, 0x39, 0x7e, 0xdf, 0x3f, 0xfb, 0xef,
	0xce, 0xc2, 0xea, 0x73, 0xf0, 0xdf, 0x4d, 0xed, 0x42, 0xe0, 0x50, 0xf6, 0x81, 0x23, 0xe3, 0x4a,
	0x98, 0xbb, 0x6d, 0x7d, 0x14, 0xac, 0x7f, 0x00, 0x36, 0x23, 0xbe, 0x8e, 0x84, 0x4f, 0x70, 0xa5,
	0xba, 0x6e, 0x9d, 0x88, 0xd9, 0x3a, 0xa1, 0x24, 0x4f, 0x8b, 0xac, 0x5a, 0xaa, 0x7d, 0xdf, 0x4b,
	0x7c, 0x83, 0xa5, 0xb1, 0xd1, 0x73, 0xb7, 0x2b, 0x11, 0x9a, 0xe5, 0x69, 0xb1, 0x78, 0x7e, 0x2c,
	0x27, 0xfd, 0xa5, 0x96, 0xbf, 0xe8, 0xb0, 0x5e, 0x51, 0x5d, 0x9a, 0x03, 0xc8, 0xfa, 0x37, 0x81,
	0xeb, 0xd3, 0x11, 0xcc, 0x61, 0xd1, 0x38, 0xb1, 0xc1, 0xf7, 0xa3, 0xa3, 0x24, 0x4f, 0x8a, 0xac,
	0x3a, 0x56, 0x48, 0x70, 0x61, 0xb9, 0x6d, 0x4d, 0xd7, 0xd0, 0x4c, 0xd3, 0x7f, 0xc4, 0x5b, 0x38,
	0xe7, 0x21, 0xf6, 0x43, 0xa4, 0x54, 0x83, 0x3d, 0xe1, 0x03, 0x64, 0x26, 0x44, 0xff, 0x65, 0x6c,
	0x14, 0x3a, 0xd3, 0xa5, 0x0e, 0x02, 0x6f, 0x60, 0xae, 0xcf, 0xa5, 0xb9, 0x1e, 0xda, 0xc1, 0xe6,
	0xfe, 0xe3, 0x6e, 0xb2, 0x50, 0xad, 0x3f, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x18, 0x9e,
	0x17, 0x5d, 0x01, 0x00, 0x00,
}