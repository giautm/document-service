// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"giautm.dev/awesome/ent/schema/pulid"
	"giautm.dev/awesome/ent/todo"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges       TodoEdges `json:"edges"`
	todo_parent *pulid.ID
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Children holds the value of the children edge.
	Children []*Todo `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Todo `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) ChildrenOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) ParentOrErr() (*Todo, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			values[i] = new(pulid.ID)
		case todo.FieldPriority:
			values[i] = new(sql.NullInt64)
		case todo.FieldText, todo.FieldStatus:
			values[i] = new(sql.NullString)
		case todo.FieldCreateTime, todo.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case todo.ForeignKeys[0]: // todo_parent
			values[i] = &sql.NullScanner{S: new(pulid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Todo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case todo.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case todo.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case todo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = todo.Status(value.String)
			}
		case todo.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int(value.Int64)
			}
		case todo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field todo_parent", values[i])
			} else if value.Valid {
				t.todo_parent = new(pulid.ID)
				*t.todo_parent = *value.S.(*pulid.ID)
			}
		}
	}
	return nil
}

// QueryChildren queries the "children" edge of the Todo entity.
func (t *Todo) QueryChildren() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryChildren(t)
}

// QueryParent queries the "parent" edge of the Todo entity.
func (t *Todo) QueryParent() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryParent(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return (&TodoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", text=")
	builder.WriteString(t.Text)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (t Todo) IsEntity() {}

func (t Todos) PluckIDs() []pulid.ID {
	ids := make([]pulid.ID, len(t))
	for _i := range t {
		ids[_i] = t[_i].ID
	}
	return ids
}

// Todos is a parsable slice of Todo.
type Todos []*Todo

func (t Todos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
