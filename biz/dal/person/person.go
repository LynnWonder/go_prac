package person

import (
	"github.com/LynnWonder/gin_prac/biz/common"
	"github.com/LynnWonder/gin_prac/biz/dal"
	"github.com/LynnWonder/gin_prac/biz/model"
	"gorm.io/gorm"
)

var db = dal.DBConnection
var log = common.SugaredLogger

func FindById(id int) (*model.Person, error) {
	var person model.Person
	result := db.Find(&person, id)
	return &person, result.Error
}

func Create(tx *gorm.DB, person *model.Person) *gorm.DB {
	return tx.Create(person)
}

func UpdatePerson(person *model.Person) *gorm.DB {
	updateResult := dal.DBConnection.Model(&model.Person{}).Where("id = ?", person.Id).
		Update("name", person.Name).Update("description", person.Description)
	return updateResult
}

func DeleteById(tx *gorm.DB, personId int) error {
	result := tx.Delete(&model.Person{}, personId)
	return result.Error
}

func ListPersons(page, pageSize int, condition map[string]interface{}) ([]*model.Person, error) {
	var querySession = dal.DBConnection.Model(&model.Person{})
	queryPersonCondition(querySession, condition)

	var persons []*model.Person
	offset := (page - 1) * pageSize
	// 这里的使用方法注意一下，结果值是第一个入参，函数返回会包含是否失败的信息
	result := querySession.Offset(offset).Limit(pageSize).Find(&persons)

	return persons, result.Error
}

func CountPersons(condition map[string]interface{}) (int64, error) {
	var querySession = dal.DBConnection.Model(&model.Person{})
	// 直接给挂上查询条件
	queryPersonCondition(querySession, condition)

	var count int64
	result := querySession.Count(&count)

	return count, result.Error
}

func queryPersonCondition(session *gorm.DB, condition map[string]interface{}) {
	if condition != nil && len(condition) > 0 {
		name, ok := condition["name"]
		if ok {
			session.Or("name LIKE ?", "%"+name.(string)+"%")
		}
	}
}