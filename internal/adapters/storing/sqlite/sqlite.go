package sqlite

import (
	"TodoApp_basic/internal/core/application/ports/query"
	"TodoApp_basic/internal/core/domain/entity"
)

const (
	Add     = "INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) VALUES (?,?,?,?,?,?);"
	Update  = "UPDATE tasks SET updatedAt =?, item=?, done =?, doneAt =? WHERE id =?;"
	FindAll = "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks;"
	FindBy  = "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks WHERE id =?;"
	Remove  = "DELETE FROM tasks WHERE id =?;"
)

func Query() *query.Query {
	return &query.Query{
		Add:     "INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) VALUES (?,?,?,?,?,?);",
		Update:  "UPDATE tasks SET updatedAt =?, item=?, done =?, doneAt =? WHERE id =?;",
		Delete:  "DELETE FROM tasks WHERE id =?;",
		FindOne: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks WHERE id =?;",
		FindAll: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks;",
	}
}

func ForGenericsQuery(e entity.Task) *query.Query {
	return &query.Query{
		Add: `INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) 
				  VALUES (e.id,e.createdAt,e.updatedAt,e.item,e.done,e.doneAt);`,
		Update: `UPDATE tasks 
				  SET updatedAt =e.updatedAt, item=e.item, done =e.done, doneAt =e.doneAt 
				  WHERE id =e.id;`,
		Delete:  "DELETE FROM tasks WHERE id =e.id;",
		FindOne: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks WHERE id =e.id;",
		FindAll: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks;",
	}
}
