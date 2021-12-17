package bootstrap

import (
	"StudentAssignment/config"
	"StudentAssignment/models"
	"StudentAssignment/pkg/Db"
	"StudentAssignment/pkg/util"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

func InitDatabase() bool {
	database := config.Conf.Database
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username, database.Password, database.Host, database.Port, database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: database.TablePrefix,
		},
	})
	if err != nil {
		log.Fatalf("连接数据库出现错误:%s", err.Error())
		return false
	}
	Db.Db = db
	log.Info("迁移数据库...")

	tables := []interface{}{
		&models.Jobs{},
		&models.Tasks{},
		&models.Users{},
		&models.Classes{},
		&models.Admins{},
	}

	for _, table := range tables {
		err := Db.Db.AutoMigrate(table)
		if err != nil {
			log.Fatalf("数据库迁移失败:%s", err.Error())
			return false
		}
	}
	initData()
	if config.Conf.Debug {
		Db.Db = Db.Db.Debug()
	}
	return true
}
func initData() {
	if util.FileExist("install.lock") {
		return
	}
	defer os.Create("install.lock")
	class := models.Classes{
		Model: gorm.Model{
			ID:        23301909,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Name:  fmt.Sprintf("%v级移动本9", time.Now().Year()),
		Users: nil,
	}
	username, _ := base64.StdEncoding.DecodeString("6ZmI54Oo")
	password, _ := base64.StdEncoding.DecodeString("Y2hlbnll")
	admin := models.Admins{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: "admin",
		Password: util.HashAndSalt(string(password)),
	}
	user := models.Users{
		Model: gorm.Model{
			ID:        2330190902,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ClassId:  23301909,
		Username: string(username),
		Password: util.HashAndSalt(string(password)),
	}
	Db.Db.Create(&class)
	Db.Db.Create(&admin)
	Db.Db.Create(&user)
}
