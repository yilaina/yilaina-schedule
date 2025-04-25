package dao

import "go_code/ginStudy/gin01_b/bubble/models"

// crud 操作方法
/*
ToDo这个model的增删改查操作都放在这里
*/
// CreateOneTodo 创建todo
func CreateOneTodo(todo *models.Todo) (err error) {
	err = DB.Create(&todo).Error
	return err
}
func GetAllTodo() (todolist []*models.Todo, err error) {
	if err := DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetOneTodo(id string) (todo *models.Todo, err error) {
	todo = new(models.Todo)
	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateOneTodo(todo *models.Todo) (err error) {
	err = DB.Save(todo).Error
	return err
}

func DeleteOneTodo(id string) (err error) {
	err = DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return err
}
