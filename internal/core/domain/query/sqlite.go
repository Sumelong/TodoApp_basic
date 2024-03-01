package query

const (
	Add     = "INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) VALUES (?,?,?,?,?,?);"
	Update  = "UPDATE tasks SET updatedAt =?, item=?, done =?, doneAt =? WHERE id =?;"
	FindAll = "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks;"
	FindBy  = "SELECT id,createdAt, updatedAt,item,done,doneAt FROM tasks WHERE id =?;"
	Remove  = "DELETE FROM tasks WHERE id =?;"
)
