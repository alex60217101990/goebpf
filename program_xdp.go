// Copyright (c) 2019 Dropbox, Inc.
// Full license can be found in the LICENSE file.

package goebpf

//#include "bpf_helpers.h"
import "C"

// XdpResult is eBPF program return code enum
type XdpResult int

// XDP program return codes
const (
	XdpAborted  XdpResult = C.XDP_ABORTED
	XdpDrop     XdpResult = C.XDP_DROP
	XdpPass     XdpResult = C.XDP_PASS
	XdpTx       XdpResult = C.XDP_TX
	XdpRedirect XdpResult = C.XDP_REDIRECT
)



// XdpAttachParams used to pass parameters to Attach() call.
type XdpAttachParams struct {
	// Interface is string name of interface to attach program to
	Interface string
	// Mode is one of XdpAttachMode.
	Mode XdpAttachMode
}

func (t XdpResult) String() string {
	switch t {
	case XdpAborted:
		return "XDP_ABORTED"
	case XdpDrop:
		return "XDP_DROP"
	case XdpPass:
		return "XDP_PASS"
	case XdpTx:
		return "XDP_TX"
	case XdpRedirect:
		return "XDP_REDIRECT"
	}

	return "UNKNOWN"
}




