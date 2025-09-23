package storage

import (
	"gott/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func AddToEmptyTest(t *testing.T) {
	jr := JsonRepo{
		fileName: "tests/atet.json",
		MaxId:    0,
		Tasks:    map[int]model.Task{},
	}
	d := "Тест на добавление в пустой список"
	jr.AddTask(d)
	task, ok := jr.Tasks[jr.MaxId]
	assert.EqualValues(t, ok, true)
	assert.EqualValues(t, task.Id, jr.MaxId, 1)
	assert.EqualValues(t, d, task.Description)
	r2, err := NewJsonRepo("tests/atet.json")
	assert.NoError(t, err)
	assert.Equal(t, jr, r2)
}
func AddTest(t *testing.T) {
	jr := JsonRepo{
		fileName: "tests/at.json",
		MaxId:    5,
		Tasks: map[int]model.Task{
			1: {
				Id:          1,
				Description: "До добавления",
				Status:      "in-progress",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}
	d := "Добавление"
	jr.AddTask(d)
	task, ok := jr.Tasks[jr.MaxId]
	assert.EqualValues(t, ok, true)
	assert.EqualValues(t, task.Id, jr.MaxId, 6)
	assert.EqualValues(t, d, task.Description)
	r2, err := NewJsonRepo("tests/at.json")
	assert.NoError(t, err)
	assert.Equal(t, jr, r2)
}
func DeleteTest(t *testing.T) {
	jr := JsonRepo{
		fileName: "tests/dt.json",
		MaxId:    5,
		Tasks: map[int]model.Task{
			1: {
				Id:          1,
				Description: "Сохраню",
				Status:      "in-progress",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, 2: {
				Id:          2,
				Description: "Удалю",
				Status:      "in-progress",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, 3: {
				Id:          3,
				Description: "Тоже сохраню",
				Status:      "in-progress",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}
	jr.DeleteTask(2)
	_, ok := jr.Tasks[2]
	assert.EqualValues(t, ok, false, t)
	r2, err := NewJsonRepo("tests/dt.json")
	assert.NoError(t, err)
	assert.Equal(t, jr, r2)
}
