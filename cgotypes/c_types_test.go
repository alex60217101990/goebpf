package cgotypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v3"
)

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
	t.Log("=-->", string(bts))

	var tmp PortKey
	err = json.Unmarshal(bts, &tmp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tmp)
}
