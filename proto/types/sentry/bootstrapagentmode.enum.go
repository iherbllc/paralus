// Code generated by go generate; DO NOT EDIT.
package sentry

import (
	bytes "bytes"
	driver "database/sql/driver"
	"fmt"
)

// Scan converts database string to BootstrapAgentMode
func (e *BootstrapAgentMode) Scan(value interface{}) error {
	s := value.([]byte)
	*e = BootstrapAgentMode(BootstrapAgentMode_value[string(s)])
	return nil
}

// Value converts BootstrapAgentMode into database string
func (e BootstrapAgentMode) Value() (driver.Value, error) {
	return BootstrapAgentMode_name[int32(e)], nil
}

// MarshalJSON converts BootstrapAgentMode to JSON
func (e BootstrapAgentMode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	buffer.WriteString(e.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

// UnmarshalJSON converts BootstrapAgentMode from JSON
func (e *BootstrapAgentMode) UnmarshalJSON(b []byte) error {
	if b != nil {
		*e = BootstrapAgentMode(BootstrapAgentMode_value[string(b[1:len(b)-1])])
	}
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface
func (e BootstrapAgentMode) MarshalYAML() (interface{}, error) {
	return BootstrapAgentMode_name[int32(e)], nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (e *BootstrapAgentMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	if err := unmarshal(&name); err != nil {
		return err
	}

	value, ok := BootstrapAgentMode_value[name]
	if !ok {
		return fmt.Errorf("invalid BootstrapAgentMode: %s", name)
	}

	*e = BootstrapAgentMode(value)
	return nil
}

// implement proto enum interface
func (e BootstrapAgentMode) IsEnum() {
}
