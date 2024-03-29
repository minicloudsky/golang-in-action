// Code generated by ent, DO NOT EDIT.

package dbinstance

import (
	"time"
)

const (
	// Label holds the string label denoting the dbinstance type in the database.
	Label = "db_instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldInstanceID holds the string denoting the instance_id field in the database.
	FieldInstanceID = "instance_id"
	// FieldInstanceName holds the string denoting the instance_name field in the database.
	FieldInstanceName = "instance_name"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldEnv holds the string denoting the env field in the database.
	FieldEnv = "env"
	// FieldInstanceType holds the string denoting the instance_type field in the database.
	FieldInstanceType = "instance_type"
	// FieldEngine holds the string denoting the engine field in the database.
	FieldEngine = "engine"
	// FieldEngineVersion holds the string denoting the engine_version field in the database.
	FieldEngineVersion = "engine_version"
	// FieldSpecification holds the string denoting the specification field in the database.
	FieldSpecification = "specification"
	// FieldInstanceStatus holds the string denoting the instance_status field in the database.
	FieldInstanceStatus = "instance_status"
	// FieldInstanceCreateTime holds the string denoting the instance_create_time field in the database.
	FieldInstanceCreateTime = "instance_create_time"
	// FieldAnnotations holds the string denoting the annotations field in the database.
	FieldAnnotations = "annotations"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// Table holds the table name of the dbinstance in the database.
	Table = "t_db_instance"
)

// Columns holds all SQL columns for dbinstance fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldInstanceID,
	FieldInstanceName,
	FieldHost,
	FieldEnv,
	FieldInstanceType,
	FieldEngine,
	FieldEngineVersion,
	FieldSpecification,
	FieldInstanceStatus,
	FieldInstanceCreateTime,
	FieldAnnotations,
	FieldLabels,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// InstanceIDValidator is a validator for the "instance_id" field. It is called by the builders before save.
	InstanceIDValidator func(string) error
	// DefaultInstanceName holds the default value on creation for the "instance_name" field.
	DefaultInstanceName string
	// InstanceNameValidator is a validator for the "instance_name" field. It is called by the builders before save.
	InstanceNameValidator func(string) error
	// DefaultHost holds the default value on creation for the "host" field.
	DefaultHost string
	// HostValidator is a validator for the "host" field. It is called by the builders before save.
	HostValidator func(string) error
	// EnvValidator is a validator for the "env" field. It is called by the builders before save.
	EnvValidator func(string) error
	// DefaultInstanceType holds the default value on creation for the "instance_type" field.
	DefaultInstanceType string
	// InstanceTypeValidator is a validator for the "instance_type" field. It is called by the builders before save.
	InstanceTypeValidator func(string) error
	// DefaultEngine holds the default value on creation for the "engine" field.
	DefaultEngine string
	// EngineValidator is a validator for the "engine" field. It is called by the builders before save.
	EngineValidator func(string) error
	// DefaultEngineVersion holds the default value on creation for the "engine_version" field.
	DefaultEngineVersion string
	// EngineVersionValidator is a validator for the "engine_version" field. It is called by the builders before save.
	EngineVersionValidator func(string) error
	// DefaultSpecification holds the default value on creation for the "specification" field.
	DefaultSpecification string
	// SpecificationValidator is a validator for the "specification" field. It is called by the builders before save.
	SpecificationValidator func(string) error
	// DefaultInstanceStatus holds the default value on creation for the "instance_status" field.
	DefaultInstanceStatus string
	// InstanceStatusValidator is a validator for the "instance_status" field. It is called by the builders before save.
	InstanceStatusValidator func(string) error
	// InstanceCreateTimeValidator is a validator for the "instance_create_time" field. It is called by the builders before save.
	InstanceCreateTimeValidator func(string) error
)
