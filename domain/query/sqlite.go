package query

const (
	Add     = "INSERT INTO tasks(id,createdAt,updatedAt,item,done,doneAt) VALUES (?,?,?,?,?,?);"
	Update  = "UPDATE tasks SET id=?,updatedAt =?, item=?, done =?, doneAt =? WHERE id =?;"
	FindAll = "SELECT id,item,done,doneAt FROM tasks;"
	FindBy  = "SELECT id,item,done,doneAt FROM tasks WHERE id =? AND done =? AND doneAt =?;"
	Remove  = "DELETE FROM tasks WHERE id =?;"
)
