package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	Env               string      `mapstructure:"env"`
	Port              int         `mapstructure:"port"`
	MySQL             MySQLConfig `mapstructure:"mysql"`
	SchedulerTokenKey string      `mapstructure:"schedulerTokenKey"`
	CallbackTokenKey  string      `mapstructure:"callbackTokenKey"`
	APITokenKey       string      `mapstructure:"apiTokenKey"`
	AppCorsDomain     string      `mapstructure:"appCorsDomain"`
	SecretKey         string      `mapstructure:"secretKey"`
	JwtAuth           JwtAuth     `mapstructure:"jwtAuth"`
}

type JwtAuth struct {
	JwtSecretKey         string `mapstructure:"jwtSecretKey"`
	JwtRefreshSecretKey  string `mapstructure:"jwtRefreshSecretKey"`
	JwtExpireTime        int    `mapstructure:"jwtExpireTime"`
	JwtRefreshExpireTime int    `mapstructure:"jwtRefreshExpireTime"`
}

type MySQLConfig struct {
	ConnMaxLifetime    int   `mapstructure:"connMaxLifetime"`
	MaxOpenConnections int   `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int   `mapstructure:"maxIdleConnections"`
	MaxIdleLifetime    int   `mapstructure:"maxIdleLifetime"`
	ConnectTimeout     int   `mapstructure:"connectTimeout"`
	Master             MySQL `mapstructure:"master"`
	Slave              MySQL `mapstructure:"slave"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

var (
	AppName        = "GO_MANAGEMENT_AUTH_SCHOOL"
	consulEndpoint = "127.0.0.1:8500"
	consulPath     = "GO_MANAGEMENT_AUTH_SCHOOL"
	Env            string
	MasterDB       *sqlx.DB
	SlaveDB        *sqlx.DB
)

func New() (conf Config) {
	var once sync.Once
	once.Do(func() {
		v := viper.New()
		retried := 0
		err := InitialiseRemote(v, retried)
		if err != nil {
			log.Printf("No remote server configured will load configuration from file and environment variables: %+v", err)
			if err := InitialiseFileAndEnv(v, "config.local"); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					configFileName := fmt.Sprintf("%s.yaml", "config.local")
					log.Printf("No '" + configFileName + "' file found on search paths. Will either use environment variables or defaults")
				} else {
					log.Fatalf("Error occured during loading config: %s", err.Error())
				}
			}
		}
		err = v.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
	})
	return conf
}

func InitialiseRemote(v *viper.Viper, retried int) error {
	if consulEnv := os.Getenv("CONSUL_URL"); consulEnv != "" {
		consulEndpoint = consulEnv
	}
	log.Printf("Initialising remote config, consul endpoint: %s, consul path: %s, retried: %d", consulEndpoint, consulPath, retried)
	v.AddRemoteProvider("consul", consulEndpoint, consulPath)
	v.SetConfigType("yaml")
	err := v.ReadRemoteConfig()
	if err != nil && retried < 1 {
		time.Sleep(500 * time.Millisecond)
		return InitialiseRemote(v, retried+1)
	}
	return err
}

func SetupMasterDB(conf Config) (db *sqlx.DB) {
	// db, err := sqlx.Connect("mysql", "username:password@(localhost:3306)/dbname")
	log.Printf("Connecting to postgresql (master) %s", fmt.Sprintf("host=%s user=%s dbname=%s port=%d",
		conf.MySQL.Master.Host, conf.MySQL.Master.Username, conf.MySQL.Master.Database, conf.MySQL.Master.Port))

	var (
		masterDB *sqlx.DB
		err      error
	)

	masterDB, err = sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s",
			conf.MySQL.Master.Username, conf.MySQL.Master.Password, conf.MySQL.Master.Host, conf.MySQL.Master.Port, conf.MySQL.Master.Database))
	if err != nil {
		log.Fatalf("open db connection failed (master) %v", err)
	}

	// masterDB.SetMaxOpenConns(conf.Postgres.MaxOpenConnections)
	// masterDB.SetMaxIdleConns(conf.Postgres.MaxIdleConnections)
	// masterDB.SetConnMaxIdleTime(time.Duration(conf.Postgres.MaxIdleLifetime) * time.Millisecond)
	// if err := masterDB.Ping(); err != nil && !unitTest {
	// 	log.Fatalf("ping db connection failed (master) %v", err)
	// }

	return masterDB
}

func SetupSlaveDB(conf Config) (db *sqlx.DB) {
	// db, err := sqlx.Connect("mysql", "username:password@(localhost:3306)/dbname")
	log.Printf("Connecting to postgresql (master) %s", fmt.Sprintf("host=%s user=%s dbname=%s port=%d",
		conf.MySQL.Slave.Host, conf.MySQL.Slave.Username, conf.MySQL.Slave.Database, conf.MySQL.Slave.Port))

	var (
		slaveDB *sqlx.DB
		err     error
	)

	slaveDB, err = sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s",
			conf.MySQL.Slave.Username, conf.MySQL.Slave.Password, conf.MySQL.Slave.Host, conf.MySQL.Slave.Port, conf.MySQL.Slave.Database))
	if err != nil {
		log.Fatalf("open db connection failed (master) %v", err)
	}

	// masterDB.SetMaxOpenConns(conf.Postgres.MaxOpenConnections)
	// masterDB.SetMaxIdleConns(conf.Postgres.MaxIdleConnections)
	// masterDB.SetConnMaxIdleTime(time.Duration(conf.Postgres.MaxIdleLifetime) * time.Millisecond)
	// if err := masterDB.Ping(); err != nil && !unitTest {
	// 	log.Fatalf("ping db connection failed (master) %v", err)
	// }
	return slaveDB
}

func InitialiseFileAndEnv(v *viper.Viper, configName string) error {
	var searchPath = []string{
		"/etc/go-management-auth-school",
		"$HOME/.go-management-auth-school",
		".",
	}
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	v.SetEnvPrefix("go-management-auth-school")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return v.ReadInConfig()
}
