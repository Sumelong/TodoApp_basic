package postgres

import "TodoApp_basic/internal/adapters/storing"

func Query() *storing.Query {
	return &storing.Query{
		Add:     "INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) VALUES ($1,$2,$3,$4,$5,$6);",
		Update:  "UPDATE tasks SET updatedAt =$1, item=$2, done =$3, doneAt =$4  WHERE id =$5;",
		Delete:  "DELETE FROM tasks WHERE id =$1;",
		FindOne: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks WHERE id =$1;",
		FindAll: "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks;",
	}
}
