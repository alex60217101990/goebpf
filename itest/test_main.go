package itest

// package main

// import (
// 	"log"

// 	"github.com/alex60217101990/goebpf"
// 	"github.com/alex60217101990/goebpf/cgotypes"
// )

// // func main() {
// // 	var err error
// // 	// var ipnet *net.IPNet
// // 	// var byteToInt int = 16

// // 	bpf := goebpf.NewDefaultEbpfSystem()
// // 	err = bpf.LoadElf("/home/alex/work/src/github.com/alex60217101990/packets-dump/tmp/fw.elf")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	m := bpf.GetMapByName("v4_blacklist")
// // 	if m == nil {
// // 		log.Fatal("nil map")
// // 	}
// // 	defer m.Close()

// // 	// Create map
// // 	// m := &goebpf.EbpfMap{
// // 	// 	Type:       goebpf.MapTypeLPMTrie,
// // 	// 	KeySize:    8, // prefix size + ipv4
// // 	// 	ValueSize:  8,
// // 	// 	MaxEntries: 10,
// // 	// }
// // 	// err = m.Create()
// // 	// t.Log(err)

// // 	ipStr := "79.122.225.166/4"
// // 	key, err := cgotypes.ParseFromSrtV4(ipStr)

// // 	// if strings.Contains(ipStr, "/") {
// // 	// 	_, ipnet, err = net.ParseCIDR(ipStr)
// // 	// 	fmt.Println(ipnet.String())
// // 	// 	// byteToInt, err = strconv.Atoi(ipnet.Mask.String())
// // 	// 	if err != nil {
// // 	// 		log.Fatal(err)
// // 	// 	}
// // 	// } else {
// // 	// 	// IPv4
// // 	// 	ipnet.IP = net.ParseIP(ipStr)
// // 	// }
// // 	// var addr [4]uint8
// // 	// copy(addr[:], ipnet.IP.To4())
// // 	// key := cgotypes.GetLpmV4Key(uint8(byteToInt), addr)
// // 	// port := cgotypes.GetPortKey(cgotypes.DestinationPort, cgotypes.TCPPort, 8552)
// // 	// fmt.Println(key, port)

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	err = m.Upsert(key, 11)
// // 	// Perform lookup (it is usually done from XDP program - here is only for integration tests)
// // 	val1, err := m.LookupInt(key)
// // 	log.Println(val1, err)
// // 	//assert.Equal(t, 11, val1, "The two words should be the same.")
// // }

// func main() {
// 	var err error

// 	bpf := goebpf.NewDefaultEbpfSystem()
// 	err = bpf.LoadElf("/home/alex/work/src/github.com/alex60217101990/packets-dump/tmp/fw.elf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	m := bpf.GetMapByName("port_blacklist")
// 	if m == nil {
// 		log.Fatal("nil map")
// 	}
// 	defer m.Close()

// 	key := cgotypes.GetPortKey(cgotypes.DestinationPort, cgotypes.UDPPort, 8552)

// 	// val1, err := m.LookupInt(key)
// 	// log.Println(val1, err)

// 	// err = m.Upsert(key, 11)
// 	// Perform lookup (it is usually done from XDP program - here is only for integration tests)
// 	val1, err := m.LookupInt(key)
// 	log.Println(val1, err)
// }
