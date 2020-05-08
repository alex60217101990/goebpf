package goebpf

import (
	"errors"
	"fmt"

	"github.com/vishvananda/netlink"
)

// XDP eBPF program (implements Program interface)
type xdpProgram struct {
	BaseProgram

	// Interface name and attach mode
	ifname string
	mode   XdpAttachMode
}

func newXdpProgram(name, license string, bytecode []byte) Program {
	return &xdpProgram{
		BaseProgram: BaseProgram{
			name:        name,
			license:     license,
			bytecode:    bytecode,
			programType: ProgramTypeXdp,
		},
	}
}

// Attach attaches eBPF(XDP) program to network interface.
// There are 2 possible ways to do that:
//
// 1. Pass interface name as parameter, e.g.
//    xdpProgram.Attach("eth0")
//
// 2. Using XdpAttachParams structure:
//    xdpProgram.Attach(
//			&XdpAttachParams{Mode: XdpAttachModeSkb, Interface: "eth0"
//    })
func (p *xdpProgram) Attach(data interface{}) error {
	var ifaceName string
	var attachMode = XdpAttachModeNone // AutoSelect

	switch x := data.(type) {
	case string:
		ifaceName = x
	case *XdpAttachParams:
		ifaceName = x.Interface
		attachMode = x.Mode
	default:
		return fmt.Errorf("%T is not supported for Attach()", data)
	}

	// Lookup interface by given name, we need to extract iface index
	link, err := netlink.LinkByName(ifaceName)
	if err != nil {
		if err != nil {
			link, err = netlink.LinkByAlias(ifaceName)
			if err != nil {
				link, err = netlink.LinkByIndex(0)
			} else {
				// Most likely no such interface
				return fmt.Errorf("LinkByName() failed: %v", err)
			}
		}
	}

	// Attach program
	if err := netlink.LinkSetXdpFdWithFlags(link, p.fd, int(attachMode)); err != nil {
		return fmt.Errorf("LinkSetXdpFd() failed: %v", err)
	}

	p.ifname = ifaceName
	p.mode = attachMode

	return nil
}

// Detach detaches program from network interface
// Must be previously attached by Attach() call.
func (p *xdpProgram) Detach() error {
	if p.ifname == "" {
		return errors.New("Program isn't attached")
	}
	// Lookup interface by given name, we need to extract iface index
	link, err := netlink.LinkByName(p.ifname)
	if err != nil {
		// Most likely no such interface
		return fmt.Errorf("LinkByName() failed: %v", err)
	}

	// Setting eBPF program with FD -1 actually removes it from interface
	if err := netlink.LinkSetXdpFdWithFlags(link, -1, int(p.mode)); err != nil {
		return fmt.Errorf("LinkSetXdpFd() failed: %v", err)
	}
	p.ifname = ""

	return nil
}
