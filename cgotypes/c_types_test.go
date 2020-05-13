package cgotypes

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestYamlPortKey(t *testing.T) {
	port := GetPortKey(DestinationPort, UDPPort, 8552)
	bts, err := yaml.Marshal(&port)

	if err != nil {
		t.Error(err)
		return
	}
	var tmp PortKey
	err = yaml.Unmarshal(bts, &tmp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tmp, bts, port)
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
