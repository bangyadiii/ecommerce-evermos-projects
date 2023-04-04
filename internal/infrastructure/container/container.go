package container

import (
	"ecommerce-evermos-projects/internal/helper"
	"ecommerce-evermos-projects/internal/infrastructure/mysql"
	"ecommerce-evermos-projects/internal/infrastructure/storage"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

const currentfilepath = "internal/infrastructure/container/container.go"

type (
	Container struct {
		Mysqldb *gorm.DB
		Apps    *Apps
		Storage storage.Storage
	}

	Apps struct {
		Name                string `mapstructure:"name"`
		Host                string `mapstructure:"host"`
		Version             string `mapstructure:"version"`
		Address             string `mapstructure:"address"`
		HttpPort            int    `mapstructure:"httpport"`
		SecretJwt           string `mapstructure:"secretJwt"`
		FileSystem          string `mapstructure:"fileSystem"`
		GCSBucketName       string `mapstructure:"gcs_bucketName"`
		GCSPublicUrl        string `mapstructure:"gcs_publicUrl"`
		BasePathFileStorage string `mapstructure:"basePathFileStorage"`
	}
)

func loadEnv() {
	projectDirName := "ecommerce-evermos-projects"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()))
	}

	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read configuration file")
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)

	log.Println("apps", apps)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()))
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when unmarshal configuration file")
	return
}

func InitContainer() *Container {
	apps := AppsInit(v)
	mysqldb := mysql.DatabaseInit(v)
	var strg storage.Storage

	if apps.FileSystem == "gcss" {
		strg = storage.NewCloudStorage(apps.GCSBucketName, apps.BasePathFileStorage, apps.GCSPublicUrl)
	} else {
		strg = storage.NewLocalStorage(apps.BasePathFileStorage)
	}

	return &Container{
		Apps:    &apps,
		Mysqldb: mysqldb,
		Storage: strg,
	}

}
