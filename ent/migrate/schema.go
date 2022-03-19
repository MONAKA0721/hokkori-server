// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LettersColumns holds the columns for the "letters" table.
	LettersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
	}
	// LettersTable holds the schema information for the "letters" table.
	LettersTable = &schema.Table{
		Name:       "letters",
		Columns:    LettersColumns,
		PrimaryKey: []*schema.Column{LettersColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"in_progress", "completed"}, Default: "in_progress"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "todo_parent", Type: field.TypeInt, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_todos_parent",
				Columns:    []*schema.Column{TodosColumns[5]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LettersTable,
		TodosTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = TodosTable
}
