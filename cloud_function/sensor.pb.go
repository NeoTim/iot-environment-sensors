// Copyright 2020 Google LLC
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
// source: sensor.proto

package cloud_function

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

type BMP280Reading struct {
	Temperature          *float32 `protobuf:"fixed32,1,opt,name=temperature" json:"temperature,omitempty"`
	Pressure             *float32 `protobuf:"fixed32,2,opt,name=pressure" json:"pressure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BMP280Reading) Reset()         { *m = BMP280Reading{} }
func (m *BMP280Reading) String() string { return proto.CompactTextString(m) }
func (*BMP280Reading) ProtoMessage()    {}
func (*BMP280Reading) Descriptor() ([]byte, []int) {
	return fileDescriptor_c749425f02243e2d, []int{0}
}

func (m *BMP280Reading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BMP280Reading.Unmarshal(m, b)
}
func (m *BMP280Reading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BMP280Reading.Marshal(b, m, deterministic)
}
func (m *BMP280Reading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BMP280Reading.Merge(m, src)
}
func (m *BMP280Reading) XXX_Size() int {
	return xxx_messageInfo_BMP280Reading.Size(m)
}
func (m *BMP280Reading) XXX_DiscardUnknown() {
	xxx_messageInfo_BMP280Reading.DiscardUnknown(m)
}

var xxx_messageInfo_BMP280Reading proto.InternalMessageInfo

func (m *BMP280Reading) GetTemperature() float32 {
	if m != nil && m.Temperature != nil {
		return *m.Temperature
	}
	return 0
}

func (m *BMP280Reading) GetPressure() float32 {
	if m != nil && m.Pressure != nil {
		return *m.Pressure
	}
	return 0
}

type SCD30Reading struct {
	Temperature          *float32 `protobuf:"fixed32,1,opt,name=temperature" json:"temperature,omitempty"`
	Humidity             *float32 `protobuf:"fixed32,2,opt,name=humidity" json:"humidity,omitempty"`
	Co2                  *float32 `protobuf:"fixed32,3,opt,name=co2" json:"co2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SCD30Reading) Reset()         { *m = SCD30Reading{} }
func (m *SCD30Reading) String() string { return proto.CompactTextString(m) }
func (*SCD30Reading) ProtoMessage()    {}
func (*SCD30Reading) Descriptor() ([]byte, []int) {
	return fileDescriptor_c749425f02243e2d, []int{1}
}

func (m *SCD30Reading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCD30Reading.Unmarshal(m, b)
}
func (m *SCD30Reading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCD30Reading.Marshal(b, m, deterministic)
}
func (m *SCD30Reading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCD30Reading.Merge(m, src)
}
func (m *SCD30Reading) XXX_Size() int {
	return xxx_messageInfo_SCD30Reading.Size(m)
}
func (m *SCD30Reading) XXX_DiscardUnknown() {
	xxx_messageInfo_SCD30Reading.DiscardUnknown(m)
}

var xxx_messageInfo_SCD30Reading proto.InternalMessageInfo

func (m *SCD30Reading) GetTemperature() float32 {
	if m != nil && m.Temperature != nil {
		return *m.Temperature
	}
	return 0
}

func (m *SCD30Reading) GetHumidity() float32 {
	if m != nil && m.Humidity != nil {
		return *m.Humidity
	}
	return 0
}

func (m *SCD30Reading) GetCo2() float32 {
	if m != nil && m.Co2 != nil {
		return *m.Co2
	}
	return 0
}

type SGP30Reading struct {
	Co2                  *uint32  `protobuf:"varint,1,opt,name=co2" json:"co2,omitempty"`
	Tvoc                 *uint32  `protobuf:"varint,2,opt,name=tvoc" json:"tvoc,omitempty"`
	BaselineCo2          *uint32  `protobuf:"varint,3,opt,name=baseline_co2,json=baselineCo2" json:"baseline_co2,omitempty"`
	BaselineTvoc         *uint32  `protobuf:"varint,4,opt,name=baseline_tvoc,json=baselineTvoc" json:"baseline_tvoc,omitempty"`
	H2                   *uint32  `protobuf:"varint,5,opt,name=h2" json:"h2,omitempty"`
	Ethanol              *uint32  `protobuf:"varint,6,opt,name=ethanol" json:"ethanol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SGP30Reading) Reset()         { *m = SGP30Reading{} }
func (m *SGP30Reading) String() string { return proto.CompactTextString(m) }
func (*SGP30Reading) ProtoMessage()    {}
func (*SGP30Reading) Descriptor() ([]byte, []int) {
	return fileDescriptor_c749425f02243e2d, []int{2}
}

func (m *SGP30Reading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SGP30Reading.Unmarshal(m, b)
}
func (m *SGP30Reading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SGP30Reading.Marshal(b, m, deterministic)
}
func (m *SGP30Reading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SGP30Reading.Merge(m, src)
}
func (m *SGP30Reading) XXX_Size() int {
	return xxx_messageInfo_SGP30Reading.Size(m)
}
func (m *SGP30Reading) XXX_DiscardUnknown() {
	xxx_messageInfo_SGP30Reading.DiscardUnknown(m)
}

var xxx_messageInfo_SGP30Reading proto.InternalMessageInfo

func (m *SGP30Reading) GetCo2() uint32 {
	if m != nil && m.Co2 != nil {
		return *m.Co2
	}
	return 0
}

func (m *SGP30Reading) GetTvoc() uint32 {
	if m != nil && m.Tvoc != nil {
		return *m.Tvoc
	}
	return 0
}

func (m *SGP30Reading) GetBaselineCo2() uint32 {
	if m != nil && m.BaselineCo2 != nil {
		return *m.BaselineCo2
	}
	return 0
}

func (m *SGP30Reading) GetBaselineTvoc() uint32 {
	if m != nil && m.BaselineTvoc != nil {
		return *m.BaselineTvoc
	}
	return 0
}

func (m *SGP30Reading) GetH2() uint32 {
	if m != nil && m.H2 != nil {
		return *m.H2
	}
	return 0
}

func (m *SGP30Reading) GetEthanol() uint32 {
	if m != nil && m.Ethanol != nil {
		return *m.Ethanol
	}
	return 0
}

type SMUART04LReading struct {
	Pm10Smoke            *uint32  `protobuf:"varint,1,opt,name=pm10_smoke,json=pm10Smoke" json:"pm10_smoke,omitempty"`
	Pm25Smoke            *uint32  `protobuf:"varint,2,opt,name=pm25_smoke,json=pm25Smoke" json:"pm25_smoke,omitempty"`
	Pm100Smoke           *uint32  `protobuf:"varint,3,opt,name=pm100_smoke,json=pm100Smoke" json:"pm100_smoke,omitempty"`
	Pm10Env              *uint32  `protobuf:"varint,4,opt,name=pm10_env,json=pm10Env" json:"pm10_env,omitempty"`
	Pm25Env              *uint32  `protobuf:"varint,5,opt,name=pm25_env,json=pm25Env" json:"pm25_env,omitempty"`
	Pm100Env             *uint32  `protobuf:"varint,6,opt,name=pm100_env,json=pm100Env" json:"pm100_env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMUART04LReading) Reset()         { *m = SMUART04LReading{} }
func (m *SMUART04LReading) String() string { return proto.CompactTextString(m) }
func (*SMUART04LReading) ProtoMessage()    {}
func (*SMUART04LReading) Descriptor() ([]byte, []int) {
	return fileDescriptor_c749425f02243e2d, []int{3}
}

func (m *SMUART04LReading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMUART04LReading.Unmarshal(m, b)
}
func (m *SMUART04LReading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMUART04LReading.Marshal(b, m, deterministic)
}
func (m *SMUART04LReading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMUART04LReading.Merge(m, src)
}
func (m *SMUART04LReading) XXX_Size() int {
	return xxx_messageInfo_SMUART04LReading.Size(m)
}
func (m *SMUART04LReading) XXX_DiscardUnknown() {
	xxx_messageInfo_SMUART04LReading.DiscardUnknown(m)
}

var xxx_messageInfo_SMUART04LReading proto.InternalMessageInfo

func (m *SMUART04LReading) GetPm10Smoke() uint32 {
	if m != nil && m.Pm10Smoke != nil {
		return *m.Pm10Smoke
	}
	return 0
}

func (m *SMUART04LReading) GetPm25Smoke() uint32 {
	if m != nil && m.Pm25Smoke != nil {
		return *m.Pm25Smoke
	}
	return 0
}

func (m *SMUART04LReading) GetPm100Smoke() uint32 {
	if m != nil && m.Pm100Smoke != nil {
		return *m.Pm100Smoke
	}
	return 0
}

func (m *SMUART04LReading) GetPm10Env() uint32 {
	if m != nil && m.Pm10Env != nil {
		return *m.Pm10Env
	}
	return 0
}

func (m *SMUART04LReading) GetPm25Env() uint32 {
	if m != nil && m.Pm25Env != nil {
		return *m.Pm25Env
	}
	return 0
}

func (m *SMUART04LReading) GetPm100Env() uint32 {
	if m != nil && m.Pm100Env != nil {
		return *m.Pm100Env
	}
	return 0
}

type SensorUpdate struct {
	Timestamp            *uint64           `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Bmp280               *BMP280Reading    `protobuf:"bytes,2,opt,name=bmp280" json:"bmp280,omitempty"`
	Scd30                *SCD30Reading     `protobuf:"bytes,3,opt,name=scd30" json:"scd30,omitempty"`
	Sgp30                *SGP30Reading     `protobuf:"bytes,4,opt,name=sgp30" json:"sgp30,omitempty"`
	Smuart04L            *SMUART04LReading `protobuf:"bytes,5,opt,name=smuart04l" json:"smuart04l,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SensorUpdate) Reset()         { *m = SensorUpdate{} }
func (m *SensorUpdate) String() string { return proto.CompactTextString(m) }
func (*SensorUpdate) ProtoMessage()    {}
func (*SensorUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c749425f02243e2d, []int{4}
}

func (m *SensorUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorUpdate.Unmarshal(m, b)
}
func (m *SensorUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorUpdate.Marshal(b, m, deterministic)
}
func (m *SensorUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorUpdate.Merge(m, src)
}
func (m *SensorUpdate) XXX_Size() int {
	return xxx_messageInfo_SensorUpdate.Size(m)
}
func (m *SensorUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_SensorUpdate proto.InternalMessageInfo

func (m *SensorUpdate) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *SensorUpdate) GetBmp280() *BMP280Reading {
	if m != nil {
		return m.Bmp280
	}
	return nil
}

func (m *SensorUpdate) GetScd30() *SCD30Reading {
	if m != nil {
		return m.Scd30
	}
	return nil
}

func (m *SensorUpdate) GetSgp30() *SGP30Reading {
	if m != nil {
		return m.Sgp30
	}
	return nil
}

func (m *SensorUpdate) GetSmuart04L() *SMUART04LReading {
	if m != nil {
		return m.Smuart04L
	}
	return nil
}

func init() {
	proto.RegisterType((*BMP280Reading)(nil), "gastranslator.BMP280Reading")
	proto.RegisterType((*SCD30Reading)(nil), "gastranslator.SCD30Reading")
	proto.RegisterType((*SGP30Reading)(nil), "gastranslator.SGP30Reading")
	proto.RegisterType((*SMUART04LReading)(nil), "gastranslator.SMUART04LReading")
	proto.RegisterType((*SensorUpdate)(nil), "gastranslator.SensorUpdate")
}

func init() { proto.RegisterFile("sensor.proto", fileDescriptor_c749425f02243e2d) }

var fileDescriptor_c749425f02243e2d = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x61, 0xc5, 0x4e, 0xe2, 0x93, 0x54, 0x04, 0x9c, 0xd4, 0x26, 0x45, 0x52, 0x75, 0xe9,
	0x64, 0x28, 0xb4, 0x0d, 0x64, 0xe9, 0xd0, 0xa6, 0x45, 0x97, 0x1a, 0x08, 0xa8, 0x64, 0x6d, 0xc0,
	0xd8, 0x84, 0x2d, 0x54, 0x12, 0x09, 0x92, 0x36, 0xd0, 0x27, 0xe8, 0x9b, 0xf4, 0x49, 0xfa, 0x60,
	0x05, 0x8f, 0xa2, 0xed, 0x18, 0x5d, 0xb2, 0xf1, 0xee, 0xfb, 0xff, 0xdf, 0x77, 0xd6, 0x41, 0x62,
	0x44, 0x6b, 0xa4, 0x1e, 0x29, 0x2d, 0xad, 0x24, 0xe9, 0x92, 0x1b, 0xab, 0x79, 0x6b, 0x6a, 0x6e,
	0xa5, 0xce, 0x67, 0x90, 0x7e, 0x9e, 0xdd, 0xd1, 0x9b, 0x82, 0x09, 0xbe, 0xa8, 0xda, 0x25, 0xb9,
	0x82, 0xd8, 0x8a, 0x46, 0x09, 0xcd, 0xed, 0x5a, 0x8b, 0xac, 0x77, 0xd5, 0xfb, 0x10, 0xb1, 0xfd,
	0x16, 0x79, 0x03, 0xa7, 0x4a, 0x0b, 0x63, 0x1c, 0x8e, 0x10, 0x6f, 0xeb, 0xfc, 0x07, 0x24, 0xe5,
	0xed, 0x97, 0xf1, 0xcb, 0xd2, 0x56, 0xeb, 0xa6, 0x5a, 0x54, 0xf6, 0x57, 0x48, 0x0b, 0x35, 0x39,
	0x83, 0xa3, 0xb9, 0xa4, 0xd9, 0x11, 0xb6, 0xdd, 0x33, 0xff, 0xd3, 0x83, 0xa4, 0xfc, 0x76, 0xb7,
	0xfb, 0x81, 0x4e, 0xe2, 0x82, 0x53, 0x94, 0x10, 0x02, 0x7d, 0xbb, 0x91, 0x73, 0x0c, 0x4b, 0x19,
	0xbe, 0xc9, 0x3b, 0x48, 0x9e, 0xb8, 0x11, 0x75, 0xd5, 0x8a, 0xc7, 0x90, 0x98, 0xb2, 0x38, 0xf4,
	0x6e, 0x25, 0x25, 0xef, 0x21, 0xdd, 0x4a, 0xd0, 0xdf, 0x47, 0xcd, 0xd6, 0x77, 0xef, 0x72, 0x5e,
	0x41, 0xb4, 0xa2, 0xd9, 0x00, 0x49, 0xb4, 0xa2, 0x24, 0x83, 0x13, 0x61, 0x57, 0xbc, 0x95, 0x75,
	0x76, 0x8c, 0xcd, 0x50, 0xe6, 0x7f, 0x7b, 0x70, 0x56, 0xce, 0x1e, 0x3e, 0xb1, 0xfb, 0x62, 0xf2,
	0x3d, 0x0c, 0xfb, 0x16, 0x40, 0x35, 0xd7, 0xc5, 0xa3, 0x69, 0xe4, 0x4f, 0xd1, 0xcd, 0x3c, 0x74,
	0x9d, 0xd2, 0x35, 0x3c, 0xa6, 0xd3, 0x0e, 0x47, 0x01, 0xd3, 0xa9, 0xc7, 0x97, 0x10, 0x3b, 0x6d,
	0xb0, 0xfb, 0x1d, 0x30, 0xb0, 0xf3, 0xbf, 0x86, 0x53, 0x8c, 0x17, 0xed, 0xa6, 0x9b, 0xfe, 0xc4,
	0xd5, 0x5f, 0xdb, 0x8d, 0x47, 0x74, 0x8a, 0x68, 0x10, 0x10, 0x9d, 0x3a, 0x74, 0x0e, 0x43, 0x1f,
	0xeb, 0x98, 0xdf, 0x02, 0x63, 0x9c, 0x2f, 0xff, 0x1d, 0x41, 0x52, 0xe2, 0xf9, 0x3c, 0xa8, 0x05,
	0xb7, 0x82, 0x5c, 0xc0, 0xd0, 0x56, 0x8d, 0x30, 0x96, 0x37, 0x0a, 0x37, 0xe8, 0xb3, 0x5d, 0x83,
	0x4c, 0xe0, 0xf8, 0xa9, 0x51, 0xf4, 0xa6, 0xc0, 0xe9, 0x63, 0x7a, 0x31, 0x7a, 0x76, 0x6d, 0xa3,
	0x67, 0xa7, 0xc6, 0x3a, 0x2d, 0xb9, 0x86, 0x81, 0x99, 0x2f, 0xc6, 0x05, 0xae, 0x14, 0xd3, 0xf3,
	0x03, 0xd3, 0xfe, 0x41, 0x31, 0xaf, 0x44, 0xcb, 0x52, 0x8d, 0x0b, 0xdc, 0xf3, 0x3f, 0x96, 0xbd,
	0x13, 0x61, 0x5e, 0x49, 0x3e, 0xc2, 0xd0, 0x34, 0x6b, 0xae, 0x6d, 0x31, 0xa9, 0xf1, 0x3f, 0x88,
	0xe9, 0xe5, 0xa1, 0xed, 0xe0, 0x83, 0xb1, 0x9d, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x3b, 0x84, 0xe6, 0x46, 0x03, 0x00, 0x00,
}
