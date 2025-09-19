package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := NewService(make([]Task, 0))
	d := "qwerty"
	s.add(d)
	task := s.getTask("1")
	assert.Equal(t, task.Id, 1)
	assert.Equal(t, task.Description, d)
	assert.Equal(t, task.Status, "todo")
	s.delete("1")
	assert.Empty(t, s.tasks)

}
