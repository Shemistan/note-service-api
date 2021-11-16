package api

import (
	"fmt"
)

// Note...
type Note struct {
	Id          int64
	UserId      int64
	ClassroomId int64
	DocumentId  int64
}

func (n *Note) String() {
	fmt.Printf(
		"Id = %v; UserId = %v; ClasssroomId = %v; DocumentId = %v\n",
		n.Id, n.UserId, n.ClassroomId, n.DocumentId)
}
