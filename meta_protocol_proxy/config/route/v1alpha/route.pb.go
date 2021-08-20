// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: api/meta_protocol_proxy/config/route/v1alpha/route.proto

package aeraki_meta_protocol_proxy_config_route_v1alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Match *RouteMatch  `protobuf:"bytes,2,opt,name=match,proto3" json:"match,omitempty"`
	Route *RouteAction `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Route) GetMatch() *RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Route) GetRoute() *RouteAction {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*v3.HeaderMatcher `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescGZIP(), []int{2}
}

func (x *RouteMatch) GetMetadata() []*v3.HeaderMatcher {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RouteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
}

func (x *RouteAction) Reset() {
	*x = RouteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAction) ProtoMessage() {}

func (x *RouteAction) ProtoReflect() protoreflect.Message {
	mi := &file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAction.ProtoReflect.Descriptor instead.
func (*RouteAction) Descriptor() ([]byte, []int) {
	return file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescGZIP(), []int{3}
}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (x *RouteAction) GetCluster() string {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *RouteAction) GetWeightedClusters() *v3.WeightedCluster {
	if x, ok := x.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *v3.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

var File_api_meta_protocol_proxy_config_route_v1alpha_route_proto protoreflect.FileDescriptor

var file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDesc = []byte{
	0x0a, 0x38, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x61, 0x65, 0x72, 0x61,
	0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x2c, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x78, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x61, 0x65,
	0x72, 0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x05,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x65, 0x72, 0x61, 0x6b,
	0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x5c, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61, 0x65,
	0x72, 0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x4e, 0x0a, 0x0a, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x11, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x18, 0x0a,
	0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x5b, 0x0a, 0x32, 0x69, 0x6f, 0x2e, 0x61, 0x65,
	0x72, 0x61, 0x6b, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1b, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescOnce sync.Once
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescData = file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDesc
)

func file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescGZIP() []byte {
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescOnce.Do(func() {
		file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescData)
	})
	return file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDescData
}

var file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil), // 0: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteConfiguration
	(*Route)(nil),              // 1: aeraki.meta_protocol_proxy.config.route.v1alpha.Route
	(*RouteMatch)(nil),         // 2: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteMatch
	(*RouteAction)(nil),        // 3: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteAction
	(*v3.HeaderMatcher)(nil),   // 4: envoy.config.route.v3.HeaderMatcher
	(*v3.WeightedCluster)(nil), // 5: envoy.config.route.v3.WeightedCluster
}
var file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_depIdxs = []int32{
	1, // 0: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteConfiguration.routes:type_name -> aeraki.meta_protocol_proxy.config.route.v1alpha.Route
	2, // 1: aeraki.meta_protocol_proxy.config.route.v1alpha.Route.match:type_name -> aeraki.meta_protocol_proxy.config.route.v1alpha.RouteMatch
	3, // 2: aeraki.meta_protocol_proxy.config.route.v1alpha.Route.route:type_name -> aeraki.meta_protocol_proxy.config.route.v1alpha.RouteAction
	4, // 3: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteMatch.metadata:type_name -> envoy.config.route.v3.HeaderMatcher
	5, // 4: aeraki.meta_protocol_proxy.config.route.v1alpha.RouteAction.weighted_clusters:type_name -> envoy.config.route.v3.WeightedCluster
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_init() }
func file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_init() {
	if File_api_meta_protocol_proxy_config_route_v1alpha_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
		file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
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
		file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAction); i {
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
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_goTypes,
		DependencyIndexes: file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_depIdxs,
		MessageInfos:      file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_msgTypes,
	}.Build()
	File_api_meta_protocol_proxy_config_route_v1alpha_route_proto = out.File
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_rawDesc = nil
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_goTypes = nil
	file_api_meta_protocol_proxy_config_route_v1alpha_route_proto_depIdxs = nil
}
