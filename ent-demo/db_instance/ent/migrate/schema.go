// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TDbInstanceColumns holds the columns for the "t_db_instance" table.
	TDbInstanceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "update_time", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "instance_id", Type: field.TypeString, Unique: true},
		{Name: "instance_name", Type: field.TypeString, Default: ""},
		{Name: "host", Type: field.TypeString, Default: ""},
		{Name: "env", Type: field.TypeString},
		{Name: "instance_type", Type: field.TypeString, Default: ""},
		{Name: "engine", Type: field.TypeString, Default: ""},
		{Name: "engine_version", Type: field.TypeString, Default: ""},
		{Name: "specification", Type: field.TypeString, Default: ""},
		{Name: "instance_status", Type: field.TypeString, Default: ""},
		{Name: "instance_create_time", Type: field.TypeString, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "annotations", Type: field.TypeJSON, Nullable: true},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
	}
	// TDbInstanceTable holds the schema information for the "t_db_instance" table.
	TDbInstanceTable = &schema.Table{
		Name:       "t_db_instance",
		Columns:    TDbInstanceColumns,
		PrimaryKey: []*schema.Column{TDbInstanceColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TDbInstanceTable,
	}
)

func init() {
	TDbInstanceTable.Annotation = &entsql.Annotation{
		Table: "t_db_instance",
	}
}
