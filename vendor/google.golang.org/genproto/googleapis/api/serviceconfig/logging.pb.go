// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/logging.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Logging configuration of the service.
//
// The following example shows how to configure logs to be sent to the
// producer and consumer projects. In the example,
// the `library.googleapis.com/activity_history` log is
// sent to both the producer and consumer projects, whereas
// the `library.googleapis.com/purchase_history` log is only sent to the
// producer project:
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     logs:
//     - name: library.googleapis.com/activity_history
//       labels:
//       - key: /customer_id
//     - name: library.googleapis.com/purchase_history
//     logging:
//       producer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         logs:
//         - library.googleapis.com/activity_history
//         - library.googleapis.com/purchase_history
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         logs:
//         - library.googleapis.com/activity_history
//
type Logging struct {
	// Logging configurations for sending logs to the producer project.
	// There can be multiple producer destinations, each one must have a
	// different monitored resource type. A log can be used in at most
	// one producer destination.
	ProducerDestinations []*Logging_LoggingDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations" json:"producer_destinations,omitempty"`
	// Logging configurations for sending logs to the consumer project.
	// There can be multiple consumer destinations, each one must have a
	// different monitored resource type. A log can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Logging_LoggingDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations" json:"consumer_destinations,omitempty"`
}

func (m *Logging) Reset()                    { *m = Logging{} }
func (m *Logging) String() string            { return proto.CompactTextString(m) }
func (*Logging) ProtoMessage()               {}
func (*Logging) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *Logging) GetProducerDestinations() []*Logging_LoggingDestination {
	if m != nil {
		return m.ProducerDestinations
	}
	return nil
}

func (m *Logging) GetConsumerDestinations() []*Logging_LoggingDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific logging destination (the producer project
// or the consumer project).
type Logging_LoggingDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,3,opt,name=monitored_resource,json=monitoredResource" json:"monitored_resource,omitempty"`
	// Names of the logs to be sent to this destination. Each name must
	// be defined in the [Service.logs][google.api.Service.logs] section.
	Logs []string `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
}

func (m *Logging_LoggingDestination) Reset()                    { *m = Logging_LoggingDestination{} }
func (m *Logging_LoggingDestination) String() string            { return proto.CompactTextString(m) }
func (*Logging_LoggingDestination) ProtoMessage()               {}
func (*Logging_LoggingDestination) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0, 0} }

func init() {
	proto.RegisterType((*Logging)(nil), "google.api.Logging")
	proto.RegisterType((*Logging_LoggingDestination)(nil), "google.api.Logging.LoggingDestination")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/logging.proto", fileDescriptor10)
}

var fileDescriptor10 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xae, 0x28, 0x1b, 0x45, 0x30, 0x28, 0x2c, 0x7b, 0x2a, 0x1e, 0x64, 0x2f, 0x36,
	0xa0, 0x6f, 0xb0, 0xe8, 0x41, 0xf0, 0xb0, 0xf4, 0xe2, 0xc1, 0xc3, 0x52, 0xd3, 0x71, 0x08, 0xb4,
	0x33, 0x4b, 0x92, 0xfa, 0x34, 0x3e, 0xac, 0xd9, 0x26, 0xb5, 0x45, 0x4f, 0x7a, 0x49, 0xc2, 0xcc,
	0x3f, 0xdf, 0xfc, 0xf9, 0xc5, 0x23, 0x32, 0x63, 0x03, 0x05, 0x72, 0x53, 0x11, 0x16, 0x6c, 0x51,
	0x21, 0xd0, 0xde, 0xb2, 0x67, 0x15, 0x5b, 0xd5, 0xde, 0x38, 0x15, 0x0e, 0xe5, 0xc0, 0x7e, 0x18,
	0x0d, 0x9a, 0xe9, 0xdd, 0xa0, 0x6a, 0x18, 0xd1, 0x84, 0x89, 0x5e, 0x2a, 0x45, 0xc2, 0x04, 0xdd,
	0xea, 0xe9, 0xbf, 0xc8, 0x8a, 0x88, 0x7d, 0xe5, 0x0d, 0x93, 0x8b, 0xd8, 0xeb, 0xcf, 0x99, 0x38,
	0x79, 0x8e, 0x8b, 0xe4, 0xab, 0xb8, 0x0a, 0xc5, 0xba, 0xd3, 0x60, 0x77, 0x35, 0x38, 0x6f, 0x28,
	0x4a, 0x97, 0x59, 0x3e, 0x5f, 0x9f, 0xde, 0xdd, 0x14, 0xa3, 0x85, 0x22, 0xcd, 0x0c, 0xf7, 0xc3,
	0x28, 0x2f, 0x2f, 0x07, 0xc8, 0xa4, 0xe8, 0x0e, 0xf0, 0x60, 0xc2, 0x75, 0xed, 0x4f, 0xf8, 0xec,
	0x6f, 0xf0, 0x01, 0x32, 0x85, 0xaf, 0x5e, 0x84, 0xfc, 0xad, 0x95, 0xb7, 0x42, 0xb6, 0x4c, 0xc6,
	0xb3, 0x85, 0x7a, 0x67, 0xc1, 0x71, 0x67, 0x35, 0x2c, 0xe7, 0x79, 0xb6, 0x5e, 0x94, 0x17, 0xdf,
	0x9d, 0x32, 0x35, 0xa4, 0x14, 0x47, 0x21, 0xf2, 0xf8, 0xdb, 0x45, 0xd9, 0xbf, 0x37, 0xb9, 0x38,
	0xd7, 0xdc, 0x4e, 0xbc, 0x6d, 0xce, 0xd2, 0xa2, 0xed, 0x21, 0xbe, 0x6d, 0xf6, 0x76, 0xdc, 0xe7,
	0x78, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x54, 0xda, 0x29, 0xe7, 0x01, 0x00, 0x00,
}
