package cgotypes

func PortKeyVal(v PortKey) *PortKey {
	return &v
}

func PortKeyPtr(v *PortKey) PortKey {
	if v != nil {
		return *v
	}
	return PortKey{}
}
