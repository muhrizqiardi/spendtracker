package integration

import "testing"

func TestExpenseRepository_Insert(t *testing.T) {
	t.Run("should insert expense", func(t *testing.T) {})
}

func TestExpenseRepository_GetOneByID(t *testing.T) {
	t.Run("should return error if the expense does not exist", func(t *testing.T) {})
	t.Run("should return one expense by ID", func(t *testing.T) {})
}

func TestExpenseRepository_GetMany(t *testing.T) {
	t.Run("should return many expenses", func(t *testing.T) {})
}

func TestExpenseRepository_UpdateOneByID(t *testing.T) {
	t.Run("should return error if the expense does not exist", func(t *testing.T) {})
	t.Run("should update expense and return expense", func(t *testing.T) {})
}

func TestExpenseRepository_DeleteOneByID(t *testing.T) {
	t.Run("should return error if the expense does not exist", func(t *testing.T) {})
	t.Run("should delete expense and return the deleted expense", func(t *testing.T) {})
}
