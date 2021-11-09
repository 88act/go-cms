package system

import (
	"database/sql"
	"fmt"
	"path/filepath"

	uuid "github.com/satori/go.uuid"

	"github.com/88act/go-cms/server/config"
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/system"
	"github.com/88act/go-cms/server/model/system/request"
	"github.com/88act/go-cms/server/utils"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//@author: [88act-5](https://github.com/88act)
//@function: writeConfig
//@description: 回写配置
//@param: viper *viper.Viper, mysql config.Mysql
//@return: error

type InitDBService struct {
}

func (initDBService *InitDBService) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	viper.Set("jwt.signing-key", uuid.NewV4())
	return viper.WriteConfig()
}

//@author: [88act-5](https://github.com/88act)
//@function: createTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error

func (initDBService *InitDBService) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (initDBService *InitDBService) initDB(InitDBFunctions ...system.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//@author: [88act-5](https://github.com/88act)
//@function: InitDB
//@description: 创建数据库并初始化
//@param: conf request.InitDB
//@return: error

func (initDBService *InitDBService) InitDB(conf request.InitDB) error {

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if MysqlConfig.Dbname == "" {
		return nil
	}

	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.DB = db
	}
	var err error
	//自动迁移为给定模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据
	//修改 by  ljd ，不再需要initDB 20211105
	fmt.Println(" 不再需要initDB ")
	/*
			err := global.DB.AutoMigrate(
				system.SysUser{},
				system.SysAuthority{},
				system.SysApi{},
				system.SysBaseMenu{},
				system.SysBaseMenuParameter{},
				system.JwtBlacklist{},
				system.SysDictionary{},
				system.SysDictionaryDetail{},
				example.ExaFileUploadAndDownload{},
				example.ExaFile{},
				example.ExaFileChunk{},
				example.ExaCustomer{},
				system.SysOperationRecord{},
				system.SysSuperBuilderHistory{},
			)
			if err != nil {
				global.DB = nil
				return err
			}

		// 修改 by  ljd ，不再需要initDB 20211105
		err = initDBService.initDB(
			source.Admin,
			source.Api,
			source.AuthorityMenu,
			source.Authority,
			source.AuthoritiesMenus,
			source.Casbin,
			source.DataAuthorities,
			source.Dictionary,
			source.DictionaryDetail,
			source.File,
			source.BaseMenu,
			source.UserAuthority,
		)
		if err != nil {
			global.DB = nil
			return err
		}
	*/

	if err = initDBService.writeConfig(global.VP, MysqlConfig); err != nil {
		return err
	}
	global.CONFIG.SuperBuilder.Root, _ = filepath.Abs("..")
	return nil
}
