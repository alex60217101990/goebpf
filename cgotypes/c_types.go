package cgotypes

/*
#include <linux/types.h>
#include <string.h>
#include <stdlib.h>

struct lpm_v4_key
{
    __u32 prefixlen;
    __u8 address[4];
};

struct lpm_v6_key
{
    __u32 prefixlen;
    __u8 address[16];
};

struct mac_key
{
	unsigned char address [6];
};

enum port_type
{
    source_port,
    destination_port,
};

enum port_protocol
{
    tcp_port,
    udp_port,
};

struct port_key
{
    enum port_type type_p;
    enum port_protocol proto;
    __u32 port;
};

struct counters
{
    __u64 packets;
    __u64 bytes;
};

static __always_inline struct port_key gen_port_key(enum port_type t, enum port_protocol p, __u32 port) {
    return ((struct port_key) {
        .type_p = t,
        .proto = p,
        .port = port,
    });
};

*/
import "C"
import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"unsafe"

	"gopkg.in/yaml.v3"
)

type PortKey C.struct_port_key

func PrintEnums() {
	pt := new(C.enum_port_type)
	fmt.Println(*pt)

	fmt.Println(C.destination_port)
}

type LpmV4Key C.struct_lpm_v4_key

func ParseFromSrtV4(ipStr string) (key LpmV4Key, err error) {
	var byteToInt int = 0
	var ipnet *net.IPNet
	// Check if given address is CIDR
	if strings.Contains(ipStr, "/") {
		_, ipnet, err = net.ParseCIDR(ipStr)
	} else {
		// Without mask
		ipnet.IP = net.ParseIP(ipStr)
	}
	if err != nil {
		return key, err
	}
	if ipnet.Mask != nil {
		byteToInt, _ = ipnet.Mask.Size()
	}
	var addr [4]uint8
	copy(addr[:], ipnet.IP.To4())
	key = GetLpmV4Key(uint8(byteToInt), addr)
	return key, err
}

func GetLpmV4Key(prefix uint8, address [4]uint8) LpmV4Key {
	//oUnsafePointer := C.CBytes(address[:])
	resp := C.struct_lpm_v4_key{
		prefixlen: C.__u32(uint32(prefix)),
		//address:   (*C.__u8)(goUnsafePointer),
	}
	C.memcpy(unsafe.Pointer(&resp.address[0]), unsafe.Pointer(&address[0]), C.size_t(4))
	return LpmV4Key(resp)
}

type MacKey [6]C.uchar

// // DestructorMacKey
// func DestructorMacKey(m MacKey) {
// 	C.free(unsafe.Pointer(m))
// }

func GetMacKey(address [6]uint8) MacKey {
	var mac MacKey
	C.memcpy(unsafe.Pointer(&mac[0]), unsafe.Pointer(&address[0]), C.size_t(6))
	return mac
}

func (m MacKey) GetCharByIndex(index uint8) C.uchar {
	return m[index]
}

func ParseFromSrtMac(macStr string) (key MacKey, err error) {
	var ha net.HardwareAddr
	ha, err = net.ParseMAC(macStr)
	if err != nil {
		return key, err
	}
	var addr [6]uint8
	copy(addr[:], ha)
	key = GetMacKey(addr)
	fmt.Println(key)
	return key, err
}

type LpmV6Key C.struct_lpm_v6_key

func GetLpmV6Key(prefix uint8, address [16]uint8) LpmV6Key {
	resp := C.struct_lpm_v6_key{
		prefixlen: C.__u32(uint32(prefix)),
	}
	C.memcpy(unsafe.Pointer(&resp.address[0]), unsafe.Pointer(&address[0]), C.size_t(16))
	return LpmV6Key(resp)
}

func ParseFromSrtV6(ipStr string) (key LpmV6Key, err error) {
	var byteToInt int = 0
	var ipnet *net.IPNet
	// Check if given address is CIDR
	if strings.Contains(ipStr, "/") {
		_, ipnet, err = net.ParseCIDR(ipStr)
	} else {
		// Without mask
		ipnet.IP = net.ParseIP(ipStr)
	}
	if err != nil {
		return key, err
	}
	if ipnet.Mask != nil {
		byteToInt, _ = ipnet.Mask.Size()
	}
	var addr [16]uint8
	copy(addr[:], ipnet.IP.To16())
	key = GetLpmV6Key(uint8(byteToInt), addr)
	return key, err
}

type PortType C.enum_port_type

const (
	DestinationPort PortType = C.destination_port
	SourcePort      PortType = C.source_port
)

var (
	_PortTypeNameToValue = map[string]PortType{
		"source_port":      SourcePort,
		"destination_port": DestinationPort,
	}

	_PortTypeValueToName = map[PortType]string{
		SourcePort:      "source_port",
		DestinationPort: "destination_port",
	}
)

func (p PortType) MarshalYAML() (interface{}, error) {
	s, ok := _PortTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortType: %d", p)
	}
	return s, nil
}

func (p *PortType) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _PortTypeNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid PortType %q", value.Value)
	}
	*p = v
	return nil
}

func (p PortType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PortTypeValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortType: %d", p)
	}
	return json.Marshal(s)
}

func (p *PortType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PortType should be a string, got %s", data)
	}
	v, ok := _PortTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PortType %q", s)
	}
	*p = v
	return nil
}

func (p PortType) Val() uint8 {
	return uint8(p)
}

// it's for using with flag package
func (p *PortType) Set(val string) error {
	if p == nil {
		var defaultType PortType
		p = &defaultType
	}
	if at, ok := _PortTypeNameToValue[val]; ok {
		*p = at
		return nil
	}
	return fmt.Errorf("invalid PortType value: %s", val)
}

func (p PortType) String() string {
	return _PortTypeValueToName[p]
}

type PortProtocol C.enum_port_protocol

const (
	TCPPort PortProtocol = C.tcp_port
	UDPPort PortProtocol = C.udp_port
)

var (
	_PortProtocolNameToValue = map[string]PortProtocol{
		"tcp_port": TCPPort,
		"udp_port": UDPPort,
	}

	_PortProtocolValueToName = map[PortProtocol]string{
		TCPPort: "tcp_port",
		UDPPort: "udp_port",
	}
)

func (p PortProtocol) MarshalYAML() (interface{}, error) {
	s, ok := _PortProtocolValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortProtocol: %d", p)
	}
	return s, nil
}

func (p *PortProtocol) UnmarshalYAML(value *yaml.Node) error {
	v, ok := _PortProtocolNameToValue[value.Value]
	if !ok {
		return fmt.Errorf("invalid PortProtocol %q", value.Value)
	}
	*p = v
	return nil
}

func (p PortProtocol) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PortProtocolValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid PortProtocol: %d", p)
	}
	return json.Marshal(s)
}

func (p *PortProtocol) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PortProtocol should be a string, got %s", data)
	}
	v, ok := _PortProtocolNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PortProtocol %q", s)
	}
	*p = v
	return nil
}

func (p PortProtocol) Val() uint8 {
	return uint8(p)
}

// it's for using with flag package
func (p *PortProtocol) Set(val string) error {
	if p == nil {
		var defaultType PortProtocol
		p = &defaultType
	}
	if at, ok := _PortProtocolNameToValue[val]; ok {
		*p = at
		return nil
	}
	return fmt.Errorf("invalid PortProtocol type: %s", val)
}

func (p PortProtocol) String() string {
	return _PortProtocolValueToName[p]
}

func GetPortKey(tp PortType, p PortProtocol, port uint32) PortKey {
	return PortKey(C.struct_port_key{
		type_p: C.enum_port_type(tp),
		proto:  C.enum_port_protocol(p),
		port:   C.__u32(uint32(port)),
	})
}

func (p *PortKey) MarshalJSON() ([]byte, error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint32 `yaml:"port" json:"port"`
	}
	if p == nil {
		p = &PortKey{}
	}
	return json.Marshal(alias{
		Type:  PortType(p.type_p).String(),
		Proto: PortProtocol(p.proto).String(),
		Port:  uint32(p.port),
	})
}

func (p *PortKey) UnmarshalJSON(data []byte) (err error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint32 `yaml:"port" json:"port"`
	}
	var tmp alias
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if p == nil {
		p = &PortKey{}
	}
	tmpPortType := PortType(p.type_p)
	err = tmpPortType.Set(tmp.Type)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortType: %v", tmp.Type, err)
	}
	p.type_p = C.enum_port_type(tmpPortType)
	tmpPortProtocol := PortProtocol(p.proto)
	err = tmpPortProtocol.Set(tmp.Proto)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortProtocol: %v", tmp.Proto, err)
	}
	p.proto = C.enum_port_protocol(tmpPortProtocol)
	p.port = C.__u32(uint32(tmp.Port))

	return nil
}

func (p *PortKey) MarshalYAML() (interface{}, error) {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint32 `yaml:"port" json:"port"`
	}
	if p == nil {
		p = &PortKey{}
	}
	return alias{
		Type:  PortType(p.type_p).String(),
		Proto: PortProtocol(p.proto).String(),
		Port:  uint32(p.port),
	}, nil
}

func (p *PortKey) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type alias struct {
		Type  string `yaml:"type" json:"type"`
		Proto string `yaml:"proto" json:"proto"`
		Port  uint32 `yaml:"port" json:"port"`
	}
	var tmp alias
	if err := unmarshal(&tmp); err != nil {
		return err
	}
	if p == nil {
		p = &PortKey{}
	}
	tmpPortType := PortType(p.type_p)
	err := tmpPortType.Set(tmp.Type)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortType: %v", tmp.Type, err)
	}
	p.type_p = C.enum_port_type(tmpPortType)
	tmpPortProtocol := PortProtocol(p.proto)
	err = tmpPortProtocol.Set(tmp.Proto)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to PortProtocol: %v", tmp.Proto, err)
	}
	p.proto = C.enum_port_protocol(tmpPortProtocol)
	p.port = C.__u32(uint32(tmp.Port))

	return nil
}

type PortKeyGo struct {
	Type uint16
	Port uint16
}
