package orm

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sync"
	"time"
)

type HadeGorm struct {
	container framework.Container // 服务容器
	dbs       map[string]*gorm.DB // key为dsn, value为gorm.DB（连接池）

	lock *sync.RWMutex
}

func NewHadeGorm(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	dbs := make(map[string]*gorm.DB)
	lock := &sync.RWMutex{}
	return &HadeGorm{
		container: container,
		dbs:       dbs,
		lock:      lock,
	}, nil
}

func (app *HadeGorm) GetDB(option ...contract.DBOption) (*gorm.DB, error) {
	logger := app.container.MustMake(contract.LogKey).(contract.Log)

	config := GetBaseConfig(app.container)
	fmt.Printf("conf %v    \n &config %v \n", *config, &config)
	fmt.Println("read  " + config.ReadTimeout)
	fmt.Println("wri  " + config.WriteTimeout)

	logService := app.container.MustMake(contract.LogKey).(contract.Log)

	ormLogger := NewOrmLogger(logService)
	config.Config = &gorm.Config{
		Logger: ormLogger,
	}

	for _, opt := range option {
		if err := opt(app.container, config); err != nil {
			return nil, err
		}
	}

	if config.Dsn == "" {
		dsn, err := config.FormatDsn()
		if err != nil {
			return nil, err
		}
		config.Dsn = dsn
	}

	fmt.Printf("config.Dsn %s \n", config.Dsn)

	app.lock.RLock()
	if db, ok := app.dbs[config.Dsn]; ok {
		app.lock.RUnlock()
		return db, nil
	}
	app.lock.RUnlock()

	app.lock.Lock()
	defer app.lock.Unlock()

	// 实例化gorm.DB
	var db *gorm.DB
	var err error
	switch config.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.Dsn), config)
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.Dsn), config)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Dsn), config)
	case "sqlserver":
		db, err = gorm.Open(sqlserver.Open(config.Dsn), config)
	case "clickhouse":
		db, err = gorm.Open(clickhouse.Open(config.Dsn), config)
	}

	fmt.Printf("config.Dsn %s \n", config.Dsn)
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	if config.ConnMaxIdle > 0 {
		sqlDB.SetMaxIdleConns(config.ConnMaxIdle)
	}
	if config.ConnMaxOpen > 0 {
		sqlDB.SetMaxOpenConns(config.ConnMaxOpen)
	}
	if config.ConnMaxLifetime != "" {
		lifeTime, err := time.ParseDuration(config.ConnMaxLifetime)
		if err != nil {
			logger.Error(context.Background(), "conn max life time error", map[string]interface{}{
				"err": err,
			})
		} else {
			sqlDB.SetConnMaxLifetime(lifeTime)
		}
	}

	if config.ConnMaxIdletime != "" {
		idleTime, err := time.ParseDuration(config.ConnMaxIdletime)
		if err != nil {
			logger.Error(context.Background(), "conn max idle time error", map[string]interface{}{
				"err": err,
			})
		} else {
			sqlDB.SetConnMaxIdleTime(idleTime)
		}
	}

	if err != nil {
		app.dbs[config.Dsn] = db
	}

	return db, err
}
