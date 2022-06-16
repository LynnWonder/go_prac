package person

import (
	"github.com/LynnWonder/gin_prac/biz/common"
	"github.com/LynnWonder/gin_prac/biz/dal"
	personDao "github.com/LynnWonder/gin_prac/biz/dal/person"
	"github.com/LynnWonder/gin_prac/biz/model"
	"gorm.io/gorm"
)


var log = common.SugaredLogger

// ListPersons
// @param condition available key list: [name, type, status]
// @return [persons, row count, error]
func ListPersons(page, pageSize int, condition map[string]interface{}) ([]*model.Person, int64, error) {
	log.Debugf("Query person list param: page=%d, pageSize=%d, condition=%v", page, pageSize, condition)
	rowCount, err := personDao.CountPersons(condition)
	if err != nil {
		log.Errorf("Query person count error: %v", err)
		return nil, 0, err
	}
	if rowCount == 0 || int64((page-1)*pageSize) > rowCount {
		log.Infof("No more person record, row count=%d", rowCount)
		return []*model.Person{}, 0, nil
	}

	persons, err := personDao.ListPersons(page, pageSize, condition)
	if err != nil {
		log.Errorf("Query person list error: %v", err)
		return nil, 0, err
	}
	return persons, rowCount, nil
}


func CreatePerson(personName string, personDesc string) (*model.Person, error) {
	log.Debugf("Create personDao [%s] type [%s] cronExpress [%s]", personName, personDesc)

	var person = &model.Person{
		Name:              personName,
		Description:       &personDesc,
	}

	err := dal.DBConnection.Transaction(func(tx *gorm.DB) error {

		if result := personDao.Create(tx, person); result.Error != nil {
			log.Errorf("Insert new person failed: %v", result.Error)
			return result.Error
		}

		personId := person.Id
		log.Infof("Insert new person id: %v", personId)

		return nil
	})

	if err != nil {
		log.Errorf("Commint create person transcation err: %v", err)
		return nil, err
	}
	return person, nil
}

func UpdatePerson(personId int, personName string, personDesc string) (*model.Person, error) {
	log.Debugf("Update personDao [%s] Description [%s]", personName, personDesc)

	person, err := personDao.FindById(personId)
	if err != nil {
		log.Errorf("find person [%d] failed: %v", personId, err)
		return nil, err
	}
	if person.Id == 0 {
		log.Warnf("Not find target person [%d]", personId)
		return nil, ErrPersonNotFound
	}
	updatedPerson := &model.Person{
		Id: personId,
		Name: personName,
		Description: &personDesc,
	}
	result :=personDao.UpdatePerson(updatedPerson)
	if result.Error !=nil {
		log.Errorf("update person failed: %v", result.Error)
		return nil, result.Error
	}
	return updatedPerson, nil
}

func DeletePerson(personId int) error {
	log.Infof("Delete person [%v]", personId)
	person, err := personDao.FindById(personId)
	if err != nil {
		log.Errorf("find person [%d] failed: %v", personId, err)
		return err
	}
	if person.Id == 0 {
		log.Warnf("Not find target person [%d]", personId)
		return ErrPersonNotFound
	}

	// 这里只是示例用事务来做
	err = dal.DBConnection.Transaction(func(tx *gorm.DB) error {
		if err = personDao.DeleteById(tx, personId); err != nil {
			log.Errorf("Delete person by id [%d] failed: %v", personId, err)
			return err
		}
		return nil
	})

	if err != nil {
		log.Errorf("Commint delete person transcation err: %v", err)
		return err
	}
	return nil
}
