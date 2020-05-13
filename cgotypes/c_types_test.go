package cgotypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMacKey(t *testing.T) {
	var ha net.HardwareAddr
	ha, _ = net.ParseMAC("02:00:5e:10:00:00:00:01")
	fmt.Printf("%2X:%2X:%2X:%2X:%2X:%2X\n", ha[0], ha[1], ha[2], ha[3], ha[4], ha[5])

	key, err := ParseFromSrtMac("02:00:5e:10:00:00:00:01")
	// defer DestructorMacKey(key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%2X:%2X:%2X:%2X:%2X:%2X\n", key.GetCharByIndex(0), key.GetCharByIndex(1), key.GetCharByIndex(2), key.GetCharByIndex(3), key.GetCharByIndex(4), key.GetCharByIndex(5))
	t.Log(key)
}

func TestYamlPortKey(t *testing.T) {
	ports := []*PortKey{
		PortKeyVal(GetPortKey(DestinationPort, UDPPort, 8552)),
		PortKeyVal(GetPortKey(DestinationPort, TCPPort, 5287)),
	}
	var tmp PortKey
	for _, port := range ports {
		bts, err := yaml.Marshal(port)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(port, bts)
		err = yaml.Unmarshal(bts, &tmp)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(tmp, bts, port)
	}
	bts, err := yaml.Marshal(&ports)
	if err != nil {
		t.Error(err)
		return
	}
	err = ioutil.WriteFile("ttt.yaml", bts, 0664)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestJsonPortKey(t *testing.T) {
	port := GetPortKey(DestinationPort, UDPPort, 8552)
	bts, err := json.Marshal(&port)
	if err != nil {
		t.Error(err)
		return
	}

	var tmp PortKey
	err = json.Unmarshal(bts, &tmp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tmp)
}
