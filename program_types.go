package goebpf

// ProgramType is eBPF program types enum
type ProgramType int

// Must be in sync with enum bpf_prog_type from <linux/bpf.h>
const (
	ProgramTypeUnspec ProgramType = iota
	ProgramTypeSocketFilter
	ProgramTypeKprobe
	ProgramTypeSchedCls
	ProgramTypeSchedAct
	ProgramTypeTracepoint
	ProgramTypeXdp
	ProgramTypePerfEvent
	ProgramTypeCgroupSkb
	ProgramTypeCgroupSock
	ProgramTypeLwtIn
	ProgramTypeLwtOut
	ProgramTypeLwtXmit
	ProgramTypeSockOps
)

func (t ProgramType) String() string {
	switch t {
	case ProgramTypeSocketFilter:
		return "SocketFilter"
	case ProgramTypeKprobe:
		return "Kprobe"
	case ProgramTypeSchedCls:
		return "SchedCLS"
	case ProgramTypeSchedAct:
		return "SchedACT"
	case ProgramTypeTracepoint:
		return "Tracepoint"
	case ProgramTypeXdp:
		return "XDP"
	case ProgramTypePerfEvent:
		return "PerfEvent"
	case ProgramTypeCgroupSkb:
		return "CgroupSkb"
	case ProgramTypeCgroupSock:
		return "CgroupSock"
	case ProgramTypeLwtIn:
		return "LWTin"
	case ProgramTypeLwtOut:
		return "LWTout"
	case ProgramTypeLwtXmit:
		return "LWTxmit"
	case ProgramTypeSockOps:
		return "SockOps"
	}

	return "Unknown"
}

// Map defines interface to interact with eBPF maps
type Map interface {
	Create() error
	GetFd() int
	GetName() string
	GetType() MapType
	Close() error
	// Makes a copy of map definition. This will NOT create map, just copies definition, "template".
	// Useful for array/map of maps use case
	CloneTemplate() Map
	// Generic lookup. Accepts any type which will be
	// converted to []byte eventually, returns bytes
	Lookup(interface{}) ([]byte, error)
	// The same, but does casting of return value to int / uint64
	LookupInt(interface{}) (int, error)
	LookupUint64(interface{}) (uint64, error)
	// The same, but does casting of return value to string
	LookupString(interface{}) (string, error)
	Insert(interface{}, interface{}) error
	Update(interface{}, interface{}) error
	Upsert(interface{}, interface{}) error
	Delete(interface{}) error
}

// BaseProgram is common shared fields of eBPF programs
type BaseProgram struct {
	fd            int // File Descriptor
	name          string
	programType   ProgramType
	license       string // License
	bytecode      []byte // eBPF instructions (each instruction - 8 bytes)
	kernelVersion int    // Kernel requires version to match running for "kprobe" programs
}

// XdpAttachMode selects a way how XDP program will be attached to interface
type XdpAttachMode int

const (
	// XdpAttachModeNone stands for "best effort" - kernel automatically
	// selects best mode (would try Drv first, then fallback to Generic).
	// NOTE: Kernel will not fallback to Generic XDP if NIC driver failed
	//       to install XDP program.
	XdpAttachModeNone XdpAttachMode = 0
	// XdpAttachModeSkb is "generic", kernel mode, less performant comparing to native,
	// but does not requires driver support.
	XdpAttachModeSkb XdpAttachMode = (1 << 1)
	// XdpAttachModeDrv is native, driver mode (support from driver side required)
	XdpAttachModeDrv XdpAttachMode = (1 << 2)
	// XdpAttachModeHw suitable for NICs with hardware XDP support
	XdpAttachModeHw XdpAttachMode = (1 << 3)
)