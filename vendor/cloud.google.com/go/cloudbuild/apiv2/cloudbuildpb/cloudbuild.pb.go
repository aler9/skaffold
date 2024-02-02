// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/devtools/cloudbuild/v2/cloudbuild.proto

package cloudbuildpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time the operation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the operation finished running.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Server-defined resource path for the target of the operation.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Output only. Name of the verb executed by the operation.
	Verb string `protobuf:"bytes,4,opt,name=verb,proto3" json:"verb,omitempty"`
	// Output only. Human-readable status of the operation, if any.
	StatusMessage string `protobuf:"bytes,5,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	// Output only. Identifies whether the user has requested cancellation
	// of the operation. Operations that have successfully been cancelled
	// have [Operation.error][] value with a
	// [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding to
	// `Code.CANCELLED`.
	RequestedCancellation bool `protobuf:"varint,6,opt,name=requested_cancellation,json=requestedCancellation,proto3" json:"requested_cancellation,omitempty"`
	// Output only. API version used to start the operation.
	ApiVersion string `protobuf:"bytes,7,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *OperationMetadata) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadata) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *OperationMetadata) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *OperationMetadata) GetRequestedCancellation() bool {
	if x != nil {
		return x.RequestedCancellation
	}
	return false
}

func (x *OperationMetadata) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

// Represents the custom metadata of the RunWorkflow long-running operation.
type RunWorkflowCustomOperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time the operation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time the operation finished running.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Output only. Name of the verb executed by the operation.
	Verb string `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	// Output only. Identifies whether the user has requested cancellation
	// of the operation. Operations that have successfully been cancelled
	// have [Operation.error][] value with a
	// [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding to
	// `Code.CANCELLED`.
	RequestedCancellation bool `protobuf:"varint,4,opt,name=requested_cancellation,json=requestedCancellation,proto3" json:"requested_cancellation,omitempty"`
	// Output only. API version used to start the operation.
	ApiVersion string `protobuf:"bytes,5,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Output only. Server-defined resource path for the target of the operation.
	Target string `protobuf:"bytes,6,opt,name=target,proto3" json:"target,omitempty"`
	// Output only. ID of the pipeline run created by RunWorkflow.
	PipelineRunId string `protobuf:"bytes,7,opt,name=pipeline_run_id,json=pipelineRunId,proto3" json:"pipeline_run_id,omitempty"`
}

func (x *RunWorkflowCustomOperationMetadata) Reset() {
	*x = RunWorkflowCustomOperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunWorkflowCustomOperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunWorkflowCustomOperationMetadata) ProtoMessage() {}

func (x *RunWorkflowCustomOperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunWorkflowCustomOperationMetadata.ProtoReflect.Descriptor instead.
func (*RunWorkflowCustomOperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescGZIP(), []int{1}
}

func (x *RunWorkflowCustomOperationMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RunWorkflowCustomOperationMetadata) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RunWorkflowCustomOperationMetadata) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *RunWorkflowCustomOperationMetadata) GetRequestedCancellation() bool {
	if x != nil {
		return x.RequestedCancellation
	}
	return false
}

func (x *RunWorkflowCustomOperationMetadata) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *RunWorkflowCustomOperationMetadata) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *RunWorkflowCustomOperationMetadata) GetPipelineRunId() string {
	if x != nil {
		return x.PipelineRunId
	}
	return ""
}

var File_google_devtools_cloudbuild_v2_cloudbuild_proto protoreflect.FileDescriptor

var file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x76, 0x32, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x2a, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x16, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe7, 0x02, 0x0a, 0x22,
	0x52, 0x75, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x3a, 0x0a, 0x16, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x15,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x42, 0xa3, 0x05, 0xea, 0x41, 0x4e, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2c, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f,
	0x7b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x7d, 0xea, 0x41, 0x59, 0x0a, 0x21, 0x69, 0x61,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x7d, 0xea, 0x41, 0x4a, 0x0a, 0x23, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x23, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x7d, 0xea, 0x41, 0x64, 0x0a, 0x2a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x7d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0xea, 0x41, 0x7c, 0x0a, 0x30, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x76, 0x32, 0x42, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x47, 0x43, 0x42, 0xaa, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x5c, 0x56, 0x32,
	0xea, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescOnce sync.Once
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescData = file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDesc
)

func file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescGZIP() []byte {
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescOnce.Do(func() {
		file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescData)
	})
	return file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDescData
}

var file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_devtools_cloudbuild_v2_cloudbuild_proto_goTypes = []interface{}{
	(*OperationMetadata)(nil),                  // 0: google.devtools.cloudbuild.v2.OperationMetadata
	(*RunWorkflowCustomOperationMetadata)(nil), // 1: google.devtools.cloudbuild.v2.RunWorkflowCustomOperationMetadata
	(*timestamppb.Timestamp)(nil),              // 2: google.protobuf.Timestamp
}
var file_google_devtools_cloudbuild_v2_cloudbuild_proto_depIdxs = []int32{
	2, // 0: google.devtools.cloudbuild.v2.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.devtools.cloudbuild.v2.OperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.devtools.cloudbuild.v2.RunWorkflowCustomOperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	2, // 3: google.devtools.cloudbuild.v2.RunWorkflowCustomOperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_devtools_cloudbuild_v2_cloudbuild_proto_init() }
func file_google_devtools_cloudbuild_v2_cloudbuild_proto_init() {
	if File_google_devtools_cloudbuild_v2_cloudbuild_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
		file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunWorkflowCustomOperationMetadata); i {
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
			RawDescriptor: file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_cloudbuild_v2_cloudbuild_proto_goTypes,
		DependencyIndexes: file_google_devtools_cloudbuild_v2_cloudbuild_proto_depIdxs,
		MessageInfos:      file_google_devtools_cloudbuild_v2_cloudbuild_proto_msgTypes,
	}.Build()
	File_google_devtools_cloudbuild_v2_cloudbuild_proto = out.File
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_rawDesc = nil
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_goTypes = nil
	file_google_devtools_cloudbuild_v2_cloudbuild_proto_depIdxs = nil
}
