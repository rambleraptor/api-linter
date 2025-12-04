// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package desc provides wrapper types around protoreflect descriptors
// to maintain compatibility with the existing codebase while migrating
// from jhump/protoreflect to google.golang.org/protobuf.
package desc

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor is the common interface for all descriptor types.
type Descriptor interface {
	GetName() string
	GetFullyQualifiedName() string
	GetFile() *FileDescriptor
	GetSourceInfo() *descriptorpb.SourceCodeInfo_Location
	GetParent() Descriptor
	AsProto() proto.Message
}

// FileDescriptor wraps a protoreflect.FileDescriptor.
type FileDescriptor struct {
	desc protoreflect.FileDescriptor
}

// WrapFile wraps a protoreflect.FileDescriptor.
func WrapFile(fd protoreflect.FileDescriptor) *FileDescriptor {
	if fd == nil {
		return nil
	}
	return &FileDescriptor{desc: fd}
}

// Unwrap returns the underlying protoreflect.FileDescriptor.
func (f *FileDescriptor) Unwrap() protoreflect.FileDescriptor {
	return f.desc
}

func (f *FileDescriptor) GetName() string {
	return f.desc.Path()
}

func (f *FileDescriptor) GetFullyQualifiedName() string {
	return string(f.desc.FullName())
}

func (f *FileDescriptor) GetPackage() string {
	return string(f.desc.Package())
}

func (f *FileDescriptor) GetFile() *FileDescriptor {
	return f
}

func (f *FileDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	// File descriptors don't have a source location themselves
	return nil
}

func (f *FileDescriptor) GetParent() Descriptor {
	// File descriptors have no parent
	return nil
}

func (f *FileDescriptor) AsProto() proto.Message {
	return protodesc.ToFileDescriptorProto(f.desc)
}

func (f *FileDescriptor) GetFileOptions() *descriptorpb.FileOptions {
	return protodesc.ToFileDescriptorProto(f.desc).Options
}

func (f *FileDescriptor) AsFileDescriptorProto() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(f.desc)
}

func (f *FileDescriptor) IsProto3() bool {
	return f.desc.Syntax() == protoreflect.Proto3
}

func (f *FileDescriptor) Edition() descriptorpb.Edition {
	fdp := protodesc.ToFileDescriptorProto(f.desc)
	if fdp.Edition != nil {
		return *fdp.Edition
	}
	return descriptorpb.Edition_EDITION_UNKNOWN
}

func (f *FileDescriptor) GetDependencies() []*FileDescriptor {
	imports := f.desc.Imports()
	result := make([]*FileDescriptor, imports.Len())
	for i := 0; i < imports.Len(); i++ {
		result[i] = WrapFile(imports.Get(i).FileDescriptor)
	}
	return result
}

func (f *FileDescriptor) FindMessage(name string) *MessageDescriptor {
	msg := f.desc.Messages().ByName(protoreflect.Name(name))
	return WrapMessage(msg)
}

func (f *FileDescriptor) GetMessageTypes() []*MessageDescriptor {
	messages := f.desc.Messages()
	result := make([]*MessageDescriptor, messages.Len())
	for i := 0; i < messages.Len(); i++ {
		result[i] = WrapMessage(messages.Get(i))
	}
	return result
}

func (f *FileDescriptor) GetEnumTypes() []*EnumDescriptor {
	enums := f.desc.Enums()
	result := make([]*EnumDescriptor, enums.Len())
	for i := 0; i < enums.Len(); i++ {
		result[i] = WrapEnum(enums.Get(i))
	}
	return result
}

func (f *FileDescriptor) GetServices() []*ServiceDescriptor {
	services := f.desc.Services()
	result := make([]*ServiceDescriptor, services.Len())
	for i := 0; i < services.Len(); i++ {
		result[i] = WrapService(services.Get(i))
	}
	return result
}

func (f *FileDescriptor) GetExtensions() []*FieldDescriptor {
	extensions := f.desc.Extensions()
	result := make([]*FieldDescriptor, extensions.Len())
	for i := 0; i < extensions.Len(); i++ {
		result[i] = WrapField(extensions.Get(i))
	}
	return result
}

func (f *FileDescriptor) FindSymbol(symbol string) Descriptor {
	name := protoreflect.FullName(symbol)
	if msg := f.desc.Messages().ByName(name.Name()); msg != nil {
		return WrapMessage(msg)
	}
	if enum := f.desc.Enums().ByName(name.Name()); enum != nil {
		return WrapEnum(enum)
	}
	if svc := f.desc.Services().ByName(name.Name()); svc != nil {
		return WrapService(svc)
	}
	return nil
}

// MessageDescriptor wraps a protoreflect.MessageDescriptor.
type MessageDescriptor struct {
	desc protoreflect.MessageDescriptor
}

// WrapMessage wraps a protoreflect.MessageDescriptor.
func WrapMessage(md protoreflect.MessageDescriptor) *MessageDescriptor {
	if md == nil {
		return nil
	}
	return &MessageDescriptor{desc: md}
}

// Unwrap returns the underlying protoreflect.MessageDescriptor.
func (m *MessageDescriptor) Unwrap() protoreflect.MessageDescriptor {
	return m.desc
}

func (m *MessageDescriptor) GetName() string {
	return string(m.desc.Name())
}

func (m *MessageDescriptor) GetFullyQualifiedName() string {
	return string(m.desc.FullName())
}

func (m *MessageDescriptor) GetFile() *FileDescriptor {
	return WrapFile(m.desc.ParentFile())
}

func (m *MessageDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(m.desc)
}

func (m *MessageDescriptor) GetParent() Descriptor {
	parent := m.desc.Parent()
	switch p := parent.(type) {
	case protoreflect.MessageDescriptor:
		return WrapMessage(p)
	case protoreflect.FileDescriptor:
		return WrapFile(p)
	default:
		return nil
	}
}

func (m *MessageDescriptor) AsProto() proto.Message {
	return protodesc.ToDescriptorProto(m.desc)
}

func (m *MessageDescriptor) GetMessageOptions() *descriptorpb.MessageOptions {
	return protodesc.ToDescriptorProto(m.desc).Options
}

func (m *MessageDescriptor) GetFields() []*FieldDescriptor {
	fields := m.desc.Fields()
	result := make([]*FieldDescriptor, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		result[i] = WrapField(fields.Get(i))
	}
	return result
}

func (m *MessageDescriptor) GetNestedMessageTypes() []*MessageDescriptor {
	messages := m.desc.Messages()
	result := make([]*MessageDescriptor, messages.Len())
	for i := 0; i < messages.Len(); i++ {
		result[i] = WrapMessage(messages.Get(i))
	}
	return result
}

func (m *MessageDescriptor) GetNestedEnumTypes() []*EnumDescriptor {
	enums := m.desc.Enums()
	result := make([]*EnumDescriptor, enums.Len())
	for i := 0; i < enums.Len(); i++ {
		result[i] = WrapEnum(enums.Get(i))
	}
	return result
}

func (m *MessageDescriptor) GetNestedExtensions() []*FieldDescriptor {
	extensions := m.desc.Extensions()
	result := make([]*FieldDescriptor, extensions.Len())
	for i := 0; i < extensions.Len(); i++ {
		result[i] = WrapField(extensions.Get(i))
	}
	return result
}

func (m *MessageDescriptor) GetOneofs() []*OneofDescriptor {
	oneofs := m.desc.Oneofs()
	result := make([]*OneofDescriptor, oneofs.Len())
	for i := 0; i < oneofs.Len(); i++ {
		result[i] = WrapOneof(oneofs.Get(i))
	}
	return result
}

func (m *MessageDescriptor) FindFieldByName(name string) *FieldDescriptor {
	return WrapField(m.desc.Fields().ByName(protoreflect.Name(name)))
}

func (m *MessageDescriptor) FindFieldByNumber(num int) *FieldDescriptor {
	return WrapField(m.desc.Fields().ByNumber(protoreflect.FieldNumber(num)))
}

func (m *MessageDescriptor) IsMapEntry() bool {
	return m.desc.IsMapEntry()
}

// FieldDescriptor wraps a protoreflect.FieldDescriptor.
type FieldDescriptor struct {
	desc protoreflect.FieldDescriptor
}

// WrapField wraps a protoreflect.FieldDescriptor.
func WrapField(fd protoreflect.FieldDescriptor) *FieldDescriptor {
	if fd == nil {
		return nil
	}
	return &FieldDescriptor{desc: fd}
}

// Unwrap returns the underlying protoreflect.FieldDescriptor.
func (f *FieldDescriptor) Unwrap() protoreflect.FieldDescriptor {
	return f.desc
}

func (f *FieldDescriptor) GetName() string {
	return string(f.desc.Name())
}

func (f *FieldDescriptor) GetFullyQualifiedName() string {
	return string(f.desc.FullName())
}

func (f *FieldDescriptor) GetFile() *FileDescriptor {
	return WrapFile(f.desc.ParentFile())
}

func (f *FieldDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(f.desc)
}

func (f *FieldDescriptor) GetParent() Descriptor {
	parent := f.desc.Parent()
	switch p := parent.(type) {
	case protoreflect.MessageDescriptor:
		return WrapMessage(p)
	case protoreflect.FileDescriptor:
		return WrapFile(p)
	default:
		return nil
	}
}

func (f *FieldDescriptor) AsProto() proto.Message {
	return protodesc.ToFieldDescriptorProto(f.desc)
}

func (f *FieldDescriptor) GetFieldOptions() *descriptorpb.FieldOptions {
	return protodesc.ToFieldDescriptorProto(f.desc).Options
}

func (f *FieldDescriptor) GetNumber() int32 {
	return int32(f.desc.Number())
}

func (f *FieldDescriptor) GetType() descriptorpb.FieldDescriptorProto_Type {
	// Convert protoreflect.Kind to descriptorpb.FieldDescriptorProto_Type
	return protodesc.ToFieldDescriptorProto(f.desc).GetType()
}

func (f *FieldDescriptor) GetTypeName() string {
	if f.desc.Message() != nil {
		return string(f.desc.Message().FullName())
	}
	if f.desc.Enum() != nil {
		return string(f.desc.Enum().FullName())
	}
	return ""
}

func (f *FieldDescriptor) GetMessageType() *MessageDescriptor {
	return WrapMessage(f.desc.Message())
}

func (f *FieldDescriptor) GetEnumType() *EnumDescriptor {
	return WrapEnum(f.desc.Enum())
}

func (f *FieldDescriptor) GetOwner() *MessageDescriptor {
	parent := f.desc.Parent()
	if msg, ok := parent.(protoreflect.MessageDescriptor); ok {
		return WrapMessage(msg)
	}
	return nil
}

func (f *FieldDescriptor) GetOneof() *OneofDescriptor {
	return WrapOneof(f.desc.ContainingOneof())
}

func (f *FieldDescriptor) GetOneOf() *OneofDescriptor {
	return f.GetOneof()
}

func (f *FieldDescriptor) IsRepeated() bool {
	return f.desc.Cardinality() == protoreflect.Repeated
}

func (f *FieldDescriptor) IsRequired() bool {
	return f.desc.Cardinality() == protoreflect.Required
}

func (f *FieldDescriptor) IsMap() bool {
	return f.desc.IsMap()
}

func (f *FieldDescriptor) IsExtension() bool {
	return f.desc.IsExtension()
}

func (f *FieldDescriptor) GetDefaultValue() interface{} {
	return f.desc.Default().Interface()
}

func (f *FieldDescriptor) HasPresence() bool {
	return f.desc.HasPresence()
}

func (f *FieldDescriptor) GetMapKeyType() *FieldDescriptor {
	if !f.IsMap() {
		return nil
	}
	msgType := f.desc.Message()
	if msgType == nil {
		return nil
	}
	return WrapField(msgType.Fields().ByNumber(1))
}

func (f *FieldDescriptor) GetMapValueType() *FieldDescriptor {
	if !f.IsMap() {
		return nil
	}
	msgType := f.desc.Message()
	if msgType == nil {
		return nil
	}
	return WrapField(msgType.Fields().ByNumber(2))
}

func (f *FieldDescriptor) IsProto3Optional() bool {
	return f.desc.HasOptionalKeyword()
}

// EnumDescriptor wraps a protoreflect.EnumDescriptor.
type EnumDescriptor struct {
	desc protoreflect.EnumDescriptor
}

// WrapEnum wraps a protoreflect.EnumDescriptor.
func WrapEnum(ed protoreflect.EnumDescriptor) *EnumDescriptor {
	if ed == nil {
		return nil
	}
	return &EnumDescriptor{desc: ed}
}

// Unwrap returns the underlying protoreflect.EnumDescriptor.
func (e *EnumDescriptor) Unwrap() protoreflect.EnumDescriptor {
	return e.desc
}

func (e *EnumDescriptor) GetName() string {
	return string(e.desc.Name())
}

func (e *EnumDescriptor) GetFullyQualifiedName() string {
	return string(e.desc.FullName())
}

func (e *EnumDescriptor) GetFile() *FileDescriptor {
	return WrapFile(e.desc.ParentFile())
}

func (e *EnumDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(e.desc)
}

func (e *EnumDescriptor) GetParent() Descriptor {
	parent := e.desc.Parent()
	switch p := parent.(type) {
	case protoreflect.MessageDescriptor:
		return WrapMessage(p)
	case protoreflect.FileDescriptor:
		return WrapFile(p)
	default:
		return nil
	}
}

func (e *EnumDescriptor) AsProto() proto.Message {
	return protodesc.ToEnumDescriptorProto(e.desc)
}

func (e *EnumDescriptor) GetEnumOptions() *descriptorpb.EnumOptions {
	return protodesc.ToEnumDescriptorProto(e.desc).Options
}

func (e *EnumDescriptor) GetValues() []*EnumValueDescriptor {
	values := e.desc.Values()
	result := make([]*EnumValueDescriptor, values.Len())
	for i := 0; i < values.Len(); i++ {
		result[i] = WrapEnumValue(values.Get(i))
	}
	return result
}

func (e *EnumDescriptor) FindValueByName(name string) *EnumValueDescriptor {
	return WrapEnumValue(e.desc.Values().ByName(protoreflect.Name(name)))
}

func (e *EnumDescriptor) FindValueByNumber(num int) *EnumValueDescriptor {
	return WrapEnumValue(e.desc.Values().ByNumber(protoreflect.EnumNumber(num)))
}

// EnumValueDescriptor wraps a protoreflect.EnumValueDescriptor.
type EnumValueDescriptor struct {
	desc protoreflect.EnumValueDescriptor
}

// WrapEnumValue wraps a protoreflect.EnumValueDescriptor.
func WrapEnumValue(evd protoreflect.EnumValueDescriptor) *EnumValueDescriptor {
	if evd == nil {
		return nil
	}
	return &EnumValueDescriptor{desc: evd}
}

// Unwrap returns the underlying protoreflect.EnumValueDescriptor.
func (e *EnumValueDescriptor) Unwrap() protoreflect.EnumValueDescriptor {
	return e.desc
}

func (e *EnumValueDescriptor) GetName() string {
	return string(e.desc.Name())
}

func (e *EnumValueDescriptor) GetFullyQualifiedName() string {
	return string(e.desc.FullName())
}

func (e *EnumValueDescriptor) GetFile() *FileDescriptor {
	return WrapFile(e.desc.ParentFile())
}

func (e *EnumValueDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(e.desc)
}

func (e *EnumValueDescriptor) GetParent() Descriptor {
	return WrapEnum(e.desc.Parent().(protoreflect.EnumDescriptor))
}

func (e *EnumValueDescriptor) AsProto() proto.Message {
	return protodesc.ToEnumValueDescriptorProto(e.desc)
}

func (e *EnumValueDescriptor) GetEnumValueOptions() *descriptorpb.EnumValueOptions {
	return protodesc.ToEnumValueDescriptorProto(e.desc).Options
}

func (e *EnumValueDescriptor) GetNumber() int32 {
	return int32(e.desc.Number())
}

func (e *EnumValueDescriptor) GetEnum() *EnumDescriptor {
	return WrapEnum(e.desc.Parent().(protoreflect.EnumDescriptor))
}

// ServiceDescriptor wraps a protoreflect.ServiceDescriptor.
type ServiceDescriptor struct {
	desc protoreflect.ServiceDescriptor
}

// WrapService wraps a protoreflect.ServiceDescriptor.
func WrapService(sd protoreflect.ServiceDescriptor) *ServiceDescriptor {
	if sd == nil {
		return nil
	}
	return &ServiceDescriptor{desc: sd}
}

// Unwrap returns the underlying protoreflect.ServiceDescriptor.
func (s *ServiceDescriptor) Unwrap() protoreflect.ServiceDescriptor {
	return s.desc
}

func (s *ServiceDescriptor) GetName() string {
	return string(s.desc.Name())
}

func (s *ServiceDescriptor) GetFullyQualifiedName() string {
	return string(s.desc.FullName())
}

func (s *ServiceDescriptor) GetFile() *FileDescriptor {
	return WrapFile(s.desc.ParentFile())
}

func (s *ServiceDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(s.desc)
}

func (s *ServiceDescriptor) GetParent() Descriptor {
	return WrapFile(s.desc.ParentFile())
}

func (s *ServiceDescriptor) AsProto() proto.Message {
	return protodesc.ToServiceDescriptorProto(s.desc)
}

func (s *ServiceDescriptor) GetServiceOptions() *descriptorpb.ServiceOptions {
	return protodesc.ToServiceDescriptorProto(s.desc).Options
}

func (s *ServiceDescriptor) GetMethods() []*MethodDescriptor {
	methods := s.desc.Methods()
	result := make([]*MethodDescriptor, methods.Len())
	for i := 0; i < methods.Len(); i++ {
		result[i] = WrapMethod(methods.Get(i))
	}
	return result
}

func (s *ServiceDescriptor) FindMethodByName(name string) *MethodDescriptor {
	return WrapMethod(s.desc.Methods().ByName(protoreflect.Name(name)))
}

// MethodDescriptor wraps a protoreflect.MethodDescriptor.
type MethodDescriptor struct {
	desc protoreflect.MethodDescriptor
}

// WrapMethod wraps a protoreflect.MethodDescriptor.
func WrapMethod(md protoreflect.MethodDescriptor) *MethodDescriptor {
	if md == nil {
		return nil
	}
	return &MethodDescriptor{desc: md}
}

// Unwrap returns the underlying protoreflect.MethodDescriptor.
func (m *MethodDescriptor) Unwrap() protoreflect.MethodDescriptor {
	return m.desc
}

func (m *MethodDescriptor) GetName() string {
	return string(m.desc.Name())
}

func (m *MethodDescriptor) GetFullyQualifiedName() string {
	return string(m.desc.FullName())
}

func (m *MethodDescriptor) GetFile() *FileDescriptor {
	return WrapFile(m.desc.ParentFile())
}

func (m *MethodDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(m.desc)
}

func (m *MethodDescriptor) GetParent() Descriptor {
	return WrapService(m.desc.Parent().(protoreflect.ServiceDescriptor))
}

func (m *MethodDescriptor) AsProto() proto.Message {
	return protodesc.ToMethodDescriptorProto(m.desc)
}

func (m *MethodDescriptor) GetMethodOptions() *descriptorpb.MethodOptions {
	return protodesc.ToMethodDescriptorProto(m.desc).Options
}

func (m *MethodDescriptor) GetService() *ServiceDescriptor {
	return WrapService(m.desc.Parent().(protoreflect.ServiceDescriptor))
}

func (m *MethodDescriptor) GetInputType() *MessageDescriptor {
	return WrapMessage(m.desc.Input())
}

func (m *MethodDescriptor) GetOutputType() *MessageDescriptor {
	return WrapMessage(m.desc.Output())
}

func (m *MethodDescriptor) IsStreamingClient() bool {
	return m.desc.IsStreamingClient()
}

func (m *MethodDescriptor) IsClientStreaming() bool {
	return m.desc.IsStreamingClient()
}

func (m *MethodDescriptor) IsStreamingServer() bool {
	return m.desc.IsStreamingServer()
}

func (m *MethodDescriptor) IsServerStreaming() bool {
	return m.desc.IsStreamingServer()
}

// OneofDescriptor wraps a protoreflect.OneofDescriptor.
type OneofDescriptor struct {
	desc protoreflect.OneofDescriptor
}

// WrapOneof wraps a protoreflect.OneofDescriptor.
func WrapOneof(od protoreflect.OneofDescriptor) *OneofDescriptor {
	if od == nil {
		return nil
	}
	return &OneofDescriptor{desc: od}
}

// Unwrap returns the underlying protoreflect.OneofDescriptor.
func (o *OneofDescriptor) Unwrap() protoreflect.OneofDescriptor {
	return o.desc
}

func (o *OneofDescriptor) GetName() string {
	return string(o.desc.Name())
}

func (o *OneofDescriptor) GetFullyQualifiedName() string {
	return string(o.desc.FullName())
}

func (o *OneofDescriptor) GetFile() *FileDescriptor {
	return WrapFile(o.desc.ParentFile())
}

func (o *OneofDescriptor) GetSourceInfo() *descriptorpb.SourceCodeInfo_Location {
	return getSourceLocation(o.desc)
}

func (o *OneofDescriptor) GetParent() Descriptor {
	return WrapMessage(o.desc.Parent().(protoreflect.MessageDescriptor))
}

func (o *OneofDescriptor) AsProto() proto.Message {
	return protodesc.ToOneofDescriptorProto(o.desc)
}

func (o *OneofDescriptor) GetOneofOptions() *descriptorpb.OneofOptions {
	return protodesc.ToOneofDescriptorProto(o.desc).Options
}

func (o *OneofDescriptor) GetOwner() *MessageDescriptor {
	return WrapMessage(o.desc.Parent().(protoreflect.MessageDescriptor))
}

func (o *OneofDescriptor) GetChoices() []*FieldDescriptor {
	fields := o.desc.Fields()
	result := make([]*FieldDescriptor, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		result[i] = WrapField(fields.Get(i))
	}
	return result
}

func (o *OneofDescriptor) IsSynthetic() bool {
	return o.desc.IsSynthetic()
}

// CreateFileDescriptors creates wrapped FileDescriptors from FileDescriptorProtos.
func CreateFileDescriptors(fds []*descriptorpb.FileDescriptorProto) (map[string]*FileDescriptor, error) {
	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{File: fds})
	if err != nil {
		return nil, err
	}

	result := make(map[string]*FileDescriptor)
	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		wrapped := WrapFile(fd)
		result[wrapped.GetName()] = wrapped
		return true
	})

	return result, nil
}

// WrapFiles wraps protoreflect.FileDescriptors into our wrapper type.
func WrapFiles(fds []protoreflect.FileDescriptor) ([]*FileDescriptor, error) {
	result := make([]*FileDescriptor, len(fds))
	for i, fd := range fds {
		result[i] = WrapFile(fd)
	}
	return result, nil
}

// LoadFileDescriptor loads a file descriptor from the global registry.
func LoadFileDescriptor(filename string) (*FileDescriptor, error) {
	fd, err := protoregistry.GlobalFiles.FindFileByPath(filename)
	if err != nil {
		return nil, fmt.Errorf("file %q not found: %w", filename, err)
	}
	return WrapFile(fd), nil
}

// getSourceLocation retrieves the source location for a descriptor.
func getSourceLocation(d protoreflect.Descriptor) *descriptorpb.SourceCodeInfo_Location {
	// Get the source locations from the file descriptor
	file := d.ParentFile()
	if file == nil {
		return nil
	}

	fdp := protodesc.ToFileDescriptorProto(file)
	if fdp.SourceCodeInfo == nil {
		return nil
	}

	// Build the path to this descriptor
	path := buildPath(d)
	if path == nil {
		return nil
	}

	// Find the matching location
	for _, loc := range fdp.SourceCodeInfo.Location {
		if pathEquals(loc.Path, path) {
			return loc
		}
	}

	return nil
}

// buildPath builds the source code info path for a descriptor.
func buildPath(d protoreflect.Descriptor) []int32 {
	if d == nil {
		return nil
	}

	var path []int32
	var buildPathRecursive func(protoreflect.Descriptor) []int32

	buildPathRecursive = func(desc protoreflect.Descriptor) []int32 {
		switch d := desc.(type) {
		case protoreflect.FileDescriptor:
			return nil
		case protoreflect.MessageDescriptor:
			parent := d.Parent()
			var basePath []int32
			if msg, ok := parent.(protoreflect.MessageDescriptor); ok {
				basePath = buildPathRecursive(msg)
				basePath = append(basePath, 3) // nested_type field number
				basePath = append(basePath, int32(d.Index()))
			} else {
				basePath = []int32{4} // message_type field number in FileDescriptorProto
				basePath = append(basePath, int32(d.Index()))
			}
			return basePath
		case protoreflect.FieldDescriptor:
			parent := d.Parent()
			if msg, ok := parent.(protoreflect.MessageDescriptor); ok {
				basePath := buildPathRecursive(msg)
				basePath = append(basePath, 2) // field field number
				basePath = append(basePath, int32(d.Index()))
				return basePath
			}
			// File-level extension
			basePath := []int32{7} // extension field number in FileDescriptorProto
			basePath = append(basePath, int32(d.Index()))
			return basePath
		case protoreflect.OneofDescriptor:
			parent := d.Parent()
			basePath := buildPathRecursive(parent)
			basePath = append(basePath, 8) // oneof_decl field number
			basePath = append(basePath, int32(d.Index()))
			return basePath
		case protoreflect.EnumDescriptor:
			parent := d.Parent()
			var basePath []int32
			if msg, ok := parent.(protoreflect.MessageDescriptor); ok {
				basePath = buildPathRecursive(msg)
				basePath = append(basePath, 4) // enum_type field number in DescriptorProto
				basePath = append(basePath, int32(d.Index()))
			} else {
				basePath = []int32{5} // enum_type field number in FileDescriptorProto
				basePath = append(basePath, int32(d.Index()))
			}
			return basePath
		case protoreflect.EnumValueDescriptor:
			parent := d.Parent()
			basePath := buildPathRecursive(parent)
			basePath = append(basePath, 2) // value field number in EnumDescriptorProto
			basePath = append(basePath, int32(d.Index()))
			return basePath
		case protoreflect.ServiceDescriptor:
			basePath := []int32{6} // service field number in FileDescriptorProto
			basePath = append(basePath, int32(d.Index()))
			return basePath
		case protoreflect.MethodDescriptor:
			parent := d.Parent()
			basePath := buildPathRecursive(parent)
			basePath = append(basePath, 2) // method field number in ServiceDescriptorProto
			basePath = append(basePath, int32(d.Index()))
			return basePath
		}
		return nil
	}

	path = buildPathRecursive(d)
	return path
}

// pathEquals checks if two paths are equal.
func pathEquals(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
