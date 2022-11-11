package postgresmanager

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresManager struct {
	db *gorm.DB
}

var postgresmanager = &PostgresManager{}

func Open(host, dbname, port, username, password string) error {
	var err error
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable", host, dbname, port, username, password)
	postgresmanager.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	return err
}

func AutoCreateStruct(data interface{}) error {
	err := postgresmanager.db.AutoMigrate(data)
	return err
}

func Save(data interface{}) error {
	res := postgresmanager.db.Create(data)
	return res.Error
}

func Query(data, store interface{}) error {
	res := postgresmanager.db.Where(data).First(store)
	return res.Error
}

func GroupQuery(model, store interface{}) error {
	res := postgresmanager.db.Where(model).Find(store)
	return res.Error
}

func QueryAll(store interface{}) error {
	res := postgresmanager.db.Find(store)
	return res.Error
}

func Update(model, data interface{}) error {
	res := postgresmanager.db.Model(model).Updates(data)
	return res.Error
}

func Delete(data interface{}) error {
	res := postgresmanager.db.Delete(data)
	return res.Error
}

func CreateAssociation(model interface{}, key string, value interface{}) error {
	err := postgresmanager.db.Model(model).Association(key).Append(value)
	return err
}

func ReadAssociation(model interface{}, key string, store interface{}) error {
	err := postgresmanager.db.Model(model).Association(key).Find(store)
	return err
}

func UpdateAssociation(model interface{}, key string, value interface{}) error {
	err := postgresmanager.db.Model(model).Association(key).Replace(value)
	return err
}

func DeleteAssociation(model interface{}, key string, value interface{}) error {
	err := postgresmanager.db.Model(model).Association(key).Delete(value)
	return err
}

func ClearAssociations(model interface{}, key string) error {
	err := postgresmanager.db.Model(model).Association(key).Clear()
	return err
}
