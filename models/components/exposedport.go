// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ExposedPort - Connection details for an active process.
type ExposedPort struct {
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	Port          int           `json:"port"`
	Host          string        `json:"host"`
	Name          string        `json:"name"`
}

func (o *ExposedPort) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *ExposedPort) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ExposedPort) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ExposedPort) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
