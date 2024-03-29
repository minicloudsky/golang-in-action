// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/minicloudsky/golang-in-action/ent-demo/db_instance/ent/dbinstance"
)

// DbInstanceCreate is the builder for creating a DbInstance entity.
type DbInstanceCreate struct {
	config
	mutation *DbInstanceMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dic *DbInstanceCreate) SetCreateTime(t time.Time) *DbInstanceCreate {
	dic.mutation.SetCreateTime(t)
	return dic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableCreateTime(t *time.Time) *DbInstanceCreate {
	if t != nil {
		dic.SetCreateTime(*t)
	}
	return dic
}

// SetUpdateTime sets the "update_time" field.
func (dic *DbInstanceCreate) SetUpdateTime(t time.Time) *DbInstanceCreate {
	dic.mutation.SetUpdateTime(t)
	return dic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableUpdateTime(t *time.Time) *DbInstanceCreate {
	if t != nil {
		dic.SetUpdateTime(*t)
	}
	return dic
}

// SetInstanceID sets the "instance_id" field.
func (dic *DbInstanceCreate) SetInstanceID(s string) *DbInstanceCreate {
	dic.mutation.SetInstanceID(s)
	return dic
}

// SetInstanceName sets the "instance_name" field.
func (dic *DbInstanceCreate) SetInstanceName(s string) *DbInstanceCreate {
	dic.mutation.SetInstanceName(s)
	return dic
}

// SetNillableInstanceName sets the "instance_name" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableInstanceName(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetInstanceName(*s)
	}
	return dic
}

// SetHost sets the "host" field.
func (dic *DbInstanceCreate) SetHost(s string) *DbInstanceCreate {
	dic.mutation.SetHost(s)
	return dic
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableHost(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetHost(*s)
	}
	return dic
}

// SetEnv sets the "env" field.
func (dic *DbInstanceCreate) SetEnv(s string) *DbInstanceCreate {
	dic.mutation.SetEnv(s)
	return dic
}

// SetInstanceType sets the "instance_type" field.
func (dic *DbInstanceCreate) SetInstanceType(s string) *DbInstanceCreate {
	dic.mutation.SetInstanceType(s)
	return dic
}

// SetNillableInstanceType sets the "instance_type" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableInstanceType(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetInstanceType(*s)
	}
	return dic
}

// SetEngine sets the "engine" field.
func (dic *DbInstanceCreate) SetEngine(s string) *DbInstanceCreate {
	dic.mutation.SetEngine(s)
	return dic
}

// SetNillableEngine sets the "engine" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableEngine(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetEngine(*s)
	}
	return dic
}

// SetEngineVersion sets the "engine_version" field.
func (dic *DbInstanceCreate) SetEngineVersion(s string) *DbInstanceCreate {
	dic.mutation.SetEngineVersion(s)
	return dic
}

// SetNillableEngineVersion sets the "engine_version" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableEngineVersion(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetEngineVersion(*s)
	}
	return dic
}

// SetSpecification sets the "specification" field.
func (dic *DbInstanceCreate) SetSpecification(s string) *DbInstanceCreate {
	dic.mutation.SetSpecification(s)
	return dic
}

// SetNillableSpecification sets the "specification" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableSpecification(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetSpecification(*s)
	}
	return dic
}

// SetInstanceStatus sets the "instance_status" field.
func (dic *DbInstanceCreate) SetInstanceStatus(s string) *DbInstanceCreate {
	dic.mutation.SetInstanceStatus(s)
	return dic
}

// SetNillableInstanceStatus sets the "instance_status" field if the given value is not nil.
func (dic *DbInstanceCreate) SetNillableInstanceStatus(s *string) *DbInstanceCreate {
	if s != nil {
		dic.SetInstanceStatus(*s)
	}
	return dic
}

// SetInstanceCreateTime sets the "instance_create_time" field.
func (dic *DbInstanceCreate) SetInstanceCreateTime(s string) *DbInstanceCreate {
	dic.mutation.SetInstanceCreateTime(s)
	return dic
}

// SetAnnotations sets the "annotations" field.
func (dic *DbInstanceCreate) SetAnnotations(m map[string]string) *DbInstanceCreate {
	dic.mutation.SetAnnotations(m)
	return dic
}

// SetLabels sets the "labels" field.
func (dic *DbInstanceCreate) SetLabels(m map[string]string) *DbInstanceCreate {
	dic.mutation.SetLabels(m)
	return dic
}

// Mutation returns the DbInstanceMutation object of the builder.
func (dic *DbInstanceCreate) Mutation() *DbInstanceMutation {
	return dic.mutation
}

// Save creates the DbInstance in the database.
func (dic *DbInstanceCreate) Save(ctx context.Context) (*DbInstance, error) {
	var (
		err  error
		node *DbInstance
	)
	dic.defaults()
	if len(dic.hooks) == 0 {
		if err = dic.check(); err != nil {
			return nil, err
		}
		node, err = dic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbInstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dic.check(); err != nil {
				return nil, err
			}
			dic.mutation = mutation
			if node, err = dic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dic.hooks) - 1; i >= 0; i-- {
			if dic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DbInstance)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DbInstanceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dic *DbInstanceCreate) SaveX(ctx context.Context) *DbInstance {
	v, err := dic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dic *DbInstanceCreate) Exec(ctx context.Context) error {
	_, err := dic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dic *DbInstanceCreate) ExecX(ctx context.Context) {
	if err := dic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dic *DbInstanceCreate) defaults() {
	if _, ok := dic.mutation.CreateTime(); !ok {
		v := dbinstance.DefaultCreateTime()
		dic.mutation.SetCreateTime(v)
	}
	if _, ok := dic.mutation.UpdateTime(); !ok {
		v := dbinstance.DefaultUpdateTime()
		dic.mutation.SetUpdateTime(v)
	}
	if _, ok := dic.mutation.InstanceName(); !ok {
		v := dbinstance.DefaultInstanceName
		dic.mutation.SetInstanceName(v)
	}
	if _, ok := dic.mutation.Host(); !ok {
		v := dbinstance.DefaultHost
		dic.mutation.SetHost(v)
	}
	if _, ok := dic.mutation.InstanceType(); !ok {
		v := dbinstance.DefaultInstanceType
		dic.mutation.SetInstanceType(v)
	}
	if _, ok := dic.mutation.Engine(); !ok {
		v := dbinstance.DefaultEngine
		dic.mutation.SetEngine(v)
	}
	if _, ok := dic.mutation.EngineVersion(); !ok {
		v := dbinstance.DefaultEngineVersion
		dic.mutation.SetEngineVersion(v)
	}
	if _, ok := dic.mutation.Specification(); !ok {
		v := dbinstance.DefaultSpecification
		dic.mutation.SetSpecification(v)
	}
	if _, ok := dic.mutation.InstanceStatus(); !ok {
		v := dbinstance.DefaultInstanceStatus
		dic.mutation.SetInstanceStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dic *DbInstanceCreate) check() error {
	if _, ok := dic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "DbInstance.create_time"`)}
	}
	if _, ok := dic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "DbInstance.update_time"`)}
	}
	if _, ok := dic.mutation.InstanceID(); !ok {
		return &ValidationError{Name: "instance_id", err: errors.New(`ent: missing required field "DbInstance.instance_id"`)}
	}
	if v, ok := dic.mutation.InstanceID(); ok {
		if err := dbinstance.InstanceIDValidator(v); err != nil {
			return &ValidationError{Name: "instance_id", err: fmt.Errorf(`ent: validator failed for field "DbInstance.instance_id": %w`, err)}
		}
	}
	if _, ok := dic.mutation.InstanceName(); !ok {
		return &ValidationError{Name: "instance_name", err: errors.New(`ent: missing required field "DbInstance.instance_name"`)}
	}
	if v, ok := dic.mutation.InstanceName(); ok {
		if err := dbinstance.InstanceNameValidator(v); err != nil {
			return &ValidationError{Name: "instance_name", err: fmt.Errorf(`ent: validator failed for field "DbInstance.instance_name": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "DbInstance.host"`)}
	}
	if v, ok := dic.mutation.Host(); ok {
		if err := dbinstance.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "DbInstance.host": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Env(); !ok {
		return &ValidationError{Name: "env", err: errors.New(`ent: missing required field "DbInstance.env"`)}
	}
	if v, ok := dic.mutation.Env(); ok {
		if err := dbinstance.EnvValidator(v); err != nil {
			return &ValidationError{Name: "env", err: fmt.Errorf(`ent: validator failed for field "DbInstance.env": %w`, err)}
		}
	}
	if _, ok := dic.mutation.InstanceType(); !ok {
		return &ValidationError{Name: "instance_type", err: errors.New(`ent: missing required field "DbInstance.instance_type"`)}
	}
	if v, ok := dic.mutation.InstanceType(); ok {
		if err := dbinstance.InstanceTypeValidator(v); err != nil {
			return &ValidationError{Name: "instance_type", err: fmt.Errorf(`ent: validator failed for field "DbInstance.instance_type": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Engine(); !ok {
		return &ValidationError{Name: "engine", err: errors.New(`ent: missing required field "DbInstance.engine"`)}
	}
	if v, ok := dic.mutation.Engine(); ok {
		if err := dbinstance.EngineValidator(v); err != nil {
			return &ValidationError{Name: "engine", err: fmt.Errorf(`ent: validator failed for field "DbInstance.engine": %w`, err)}
		}
	}
	if _, ok := dic.mutation.EngineVersion(); !ok {
		return &ValidationError{Name: "engine_version", err: errors.New(`ent: missing required field "DbInstance.engine_version"`)}
	}
	if v, ok := dic.mutation.EngineVersion(); ok {
		if err := dbinstance.EngineVersionValidator(v); err != nil {
			return &ValidationError{Name: "engine_version", err: fmt.Errorf(`ent: validator failed for field "DbInstance.engine_version": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Specification(); !ok {
		return &ValidationError{Name: "specification", err: errors.New(`ent: missing required field "DbInstance.specification"`)}
	}
	if v, ok := dic.mutation.Specification(); ok {
		if err := dbinstance.SpecificationValidator(v); err != nil {
			return &ValidationError{Name: "specification", err: fmt.Errorf(`ent: validator failed for field "DbInstance.specification": %w`, err)}
		}
	}
	if _, ok := dic.mutation.InstanceStatus(); !ok {
		return &ValidationError{Name: "instance_status", err: errors.New(`ent: missing required field "DbInstance.instance_status"`)}
	}
	if v, ok := dic.mutation.InstanceStatus(); ok {
		if err := dbinstance.InstanceStatusValidator(v); err != nil {
			return &ValidationError{Name: "instance_status", err: fmt.Errorf(`ent: validator failed for field "DbInstance.instance_status": %w`, err)}
		}
	}
	if _, ok := dic.mutation.InstanceCreateTime(); !ok {
		return &ValidationError{Name: "instance_create_time", err: errors.New(`ent: missing required field "DbInstance.instance_create_time"`)}
	}
	if v, ok := dic.mutation.InstanceCreateTime(); ok {
		if err := dbinstance.InstanceCreateTimeValidator(v); err != nil {
			return &ValidationError{Name: "instance_create_time", err: fmt.Errorf(`ent: validator failed for field "DbInstance.instance_create_time": %w`, err)}
		}
	}
	return nil
}

func (dic *DbInstanceCreate) sqlSave(ctx context.Context) (*DbInstance, error) {
	_node, _spec := dic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dic *DbInstanceCreate) createSpec() (*DbInstance, *sqlgraph.CreateSpec) {
	var (
		_node = &DbInstance{config: dic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbinstance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbinstance.FieldID,
			},
		}
	)
	if value, ok := dic.mutation.CreateTime(); ok {
		_spec.SetField(dbinstance.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dic.mutation.UpdateTime(); ok {
		_spec.SetField(dbinstance.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dic.mutation.InstanceID(); ok {
		_spec.SetField(dbinstance.FieldInstanceID, field.TypeString, value)
		_node.InstanceID = value
	}
	if value, ok := dic.mutation.InstanceName(); ok {
		_spec.SetField(dbinstance.FieldInstanceName, field.TypeString, value)
		_node.InstanceName = value
	}
	if value, ok := dic.mutation.Host(); ok {
		_spec.SetField(dbinstance.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := dic.mutation.Env(); ok {
		_spec.SetField(dbinstance.FieldEnv, field.TypeString, value)
		_node.Env = value
	}
	if value, ok := dic.mutation.InstanceType(); ok {
		_spec.SetField(dbinstance.FieldInstanceType, field.TypeString, value)
		_node.InstanceType = value
	}
	if value, ok := dic.mutation.Engine(); ok {
		_spec.SetField(dbinstance.FieldEngine, field.TypeString, value)
		_node.Engine = value
	}
	if value, ok := dic.mutation.EngineVersion(); ok {
		_spec.SetField(dbinstance.FieldEngineVersion, field.TypeString, value)
		_node.EngineVersion = value
	}
	if value, ok := dic.mutation.Specification(); ok {
		_spec.SetField(dbinstance.FieldSpecification, field.TypeString, value)
		_node.Specification = value
	}
	if value, ok := dic.mutation.InstanceStatus(); ok {
		_spec.SetField(dbinstance.FieldInstanceStatus, field.TypeString, value)
		_node.InstanceStatus = value
	}
	if value, ok := dic.mutation.InstanceCreateTime(); ok {
		_spec.SetField(dbinstance.FieldInstanceCreateTime, field.TypeString, value)
		_node.InstanceCreateTime = value
	}
	if value, ok := dic.mutation.Annotations(); ok {
		_spec.SetField(dbinstance.FieldAnnotations, field.TypeJSON, value)
		_node.Annotations = value
	}
	if value, ok := dic.mutation.Labels(); ok {
		_spec.SetField(dbinstance.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	return _node, _spec
}

// DbInstanceCreateBulk is the builder for creating many DbInstance entities in bulk.
type DbInstanceCreateBulk struct {
	config
	builders []*DbInstanceCreate
}

// Save creates the DbInstance entities in the database.
func (dicb *DbInstanceCreateBulk) Save(ctx context.Context) ([]*DbInstance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dicb.builders))
	nodes := make([]*DbInstance, len(dicb.builders))
	mutators := make([]Mutator, len(dicb.builders))
	for i := range dicb.builders {
		func(i int, root context.Context) {
			builder := dicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbInstanceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dicb *DbInstanceCreateBulk) SaveX(ctx context.Context) []*DbInstance {
	v, err := dicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dicb *DbInstanceCreateBulk) Exec(ctx context.Context) error {
	_, err := dicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dicb *DbInstanceCreateBulk) ExecX(ctx context.Context) {
	if err := dicb.Exec(ctx); err != nil {
		panic(err)
	}
}
