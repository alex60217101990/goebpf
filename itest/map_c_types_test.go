package itest

import (
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/alex60217101990/goebpf"
	"github.com/alex60217101990/goebpf/cgotypes"
)

// type keyTestSuite struct {
// 	suite.Suite
// }

// // TestMapLPMTrieStruct
// func (ts *keyTestSuite) TestMapLPMTrieStruct() {
// 	var err error
// 	var ipnet *net.IPNet
// 	var byteToInt int = 16

// 	// Create map
// 	m := &goebpf.EbpfMap{
// 		Type:       goebpf.MapTypeLPMTrie,
// 		KeySize:    8, // prefix size + ipv4
// 		ValueSize:  1,
// 		MaxEntries: 10,
// 	}
// 	err = m.Create()
// 	ts.NoError(err)

// 	ipStr := "187.162.11.94/16"
// 	if strings.Contains(ipStr, "/") {
// 		_, ipnet, err = net.ParseCIDR(ipStr)
// 		fmt.Println(ipnet.String())
// 		// byteToInt, err = strconv.Atoi(ipnet.Mask.String())
// 		if err != nil {
// 			ts.Error(err)
// 			return
// 		}
// 	} else {
// 		// IPv4
// 		ipnet.IP = net.ParseIP(ipStr)
// 	}
// 	if err != nil {
// 		ts.Error(err)
// 		return
// 	}
// 	//fmt.Println(byteToInt, ipnet.String())
// 	// PrintEnums()
// 	var addr [8]uint8
// 	copy(addr[:], ipnet.IP.To4())
// 	fmt.Println(GetLpmV4Key(uint8(byteToInt), addr))
// 	// err = m.Insert(, "value16")
// }

// TestMapLPMTrieStruct
func TestMapLPMTrieStruct(t *testing.T) {
	var err error
	var ipnet *net.IPNet
	var byteToInt int = 16

	// Create map
	m := &goebpf.EbpfMap{
		Type:       goebpf.MapTypeLPMTrie,
		KeySize:    8, // prefix size + ipv4
		ValueSize:  8,
		MaxEntries: 10,
	}
	err = m.Create()
	t.Log(err)

	ipStr := "187.163.11.94/16"
	if strings.Contains(ipStr, "/") {
		_, ipnet, err = net.ParseCIDR(ipStr)
		fmt.Println(ipnet.String())
		// byteToInt, err = strconv.Atoi(ipnet.Mask.String())
		if err != nil {
			t.Fatal(err)
		}
	} else {
		// IPv4
		ipnet.IP = net.ParseIP(ipStr)
	}

	//fmt.Println(byteToInt, ipnet.String())
	// PrintEnums()
	var addr [4]uint8
	copy(addr[:], ipnet.IP.To4())
	key := cgotypes.GetLpmV4Key(uint8(byteToInt), addr)
	port := cgotypes.GetPortKey(cgotypes.DestinationPort, cgotypes.TCPPort, 8552)
	fmt.Println(key, port)
}
