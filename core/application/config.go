// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package application

import (
	"github.com/juju/errors"
	"github.com/juju/schema"
	"github.com/juju/utils/set"
	"gopkg.in/juju/environschema.v1"
)

// TODO(caas) - these are CAAS specific, figure out a better way
const (
	// JujuExternalHostNameKey specifies the hostname of a CAAS application.
	JujuExternalHostNameKey = "juju-external-hostname"

	// JujuApplicationPath specifies the relative http path used to access a CAAS application.
	JujuApplicationPath = "juju-application-path"

	defaultApplicationPath = "/"
)

var configFields = environschema.Fields{
	JujuExternalHostNameKey: {
		Description: "the external hostname of an exposed application",
		Type:        environschema.Tstring,
		Group:       environschema.EnvironGroup,
	},
	JujuApplicationPath: {
		Description: "the relative http path used to access an application",
		Type:        environschema.Tstring,
		Group:       environschema.EnvironGroup,
	},
}

// ConfigAttributes is the config for an application.
type ConfigAttributes map[string]interface{}

// Config encapsulates config for an application.
type Config struct {
	attributes map[string]interface{}
	schema     environschema.Fields
	defaults   schema.Defaults
}

// NewConfig returns a new config instance with the given attributes and
// allowing for the extra provider attributes.
func NewConfig(attrs map[string]interface{}, extra environschema.Fields, extraDefaults schema.Defaults) (*Config, error) {
	cfg := &Config{}
	var err error
	if cfg.schema, err = configSchema(extra); err != nil {
		return nil, errors.Trace(err)
	}
	cfg.defaults = ConfigDefaults(extraDefaults)
	if err := cfg.setAttributes(attrs); err != nil {
		return nil, errors.Trace(err)
	}
	return cfg, nil
}

func (c *Config) setAttributes(attrs map[string]interface{}) error {
	checker, err := c.schemaChecker()
	if err != nil {
		return errors.Trace(err)
	}
	m := make(map[string]interface{})
	for k, v := range attrs {
		m[k] = v
	}
	result, err := checker.Coerce(m, nil)
	if err != nil {
		return errors.Trace(err)
	}
	c.attributes = result.(map[string]interface{})
	return nil
}

type ConfigFields environschema.Fields

// KnownConfigKeys returns the valid application config keys.
func (c ConfigFields) KnownConfigKeys() set.Strings {
	result := set.NewStrings()
	for name := range c {
		result.Add(name)
	}
	return result
}

// Fields casts c to environSchema.Fields.
func (c ConfigFields) Fields() environschema.Fields {
	return environschema.Fields(c)
}

// ConfigSchema returns the valid fields for an application config.
func ConfigSchema(extra environschema.Fields) (ConfigFields, error) {
	fields, err := configSchema(extra)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return ConfigFields(fields), nil
}

// ConfigDefaults returns the default values for an application config.
func ConfigDefaults(extra schema.Defaults) schema.Defaults {
	defaults := schema.Defaults{JujuApplicationPath: defaultApplicationPath}
	for key, value := range extra {
		defaults[key] = value
	}
	return defaults
}

func configSchema(extra environschema.Fields) (environschema.Fields, error) {
	fields := make(environschema.Fields)
	for name, field := range configFields {
		fields[name] = field
	}
	for name, field := range extra {
		if _, ok := configFields[name]; ok {
			return nil, errors.Errorf("config field %q clashes with common config", name)
		}
		fields[name] = field
	}
	return fields, nil
}

func (c *Config) schemaChecker() (schema.Checker, error) {
	schemaFields, schemaDefaults, err := c.schema.ValidationSchema()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for key, value := range c.defaults {
		schemaDefaults[key] = value
	}
	return schema.StrictFieldMap(schemaFields, schemaDefaults), nil
}

// Validate returns an error if the config is not valid.
func (c *Config) Validate() error {
	return nil
}

// Attributes returns all the config attributes.
func (c *Config) Attributes() ConfigAttributes {
	if c == nil {
		return nil
	}
	result := make(ConfigAttributes)
	for k, v := range c.attributes {
		result[k] = v
	}
	return result
}

// Get gets the specified attribute.
func (c ConfigAttributes) Get(attrName string, defaultValue interface{}) interface{} {
	if val, ok := c[attrName]; ok {
		return val
	}
	return defaultValue
}

// GetInt gets the specified bool attribute.
func (c ConfigAttributes) GetBool(attrName string, defaultValue bool) bool {
	if val, ok := c[attrName]; ok {
		return val.(bool)
	}
	return defaultValue
}

// GetInt gets the specified int attribute.
func (c ConfigAttributes) GetInt(attrName string, defaultValue int) int {
	if val, ok := c[attrName]; ok {
		if value, ok := val.(float64); ok {
			return int(value)
		}
		return val.(int)
	}
	return defaultValue
}

// GetString gets the specified string attribute.
func (c ConfigAttributes) GetString(attrName string, defaultValue string) string {
	if val, ok := c[attrName]; ok {
		return val.(string)
	}
	return defaultValue
}
