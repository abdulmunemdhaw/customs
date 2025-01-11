package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username     string `yaml:"username"`
		Hostname     string `yaml:"hostname"`
		Port         uint16 `yaml:"port"`
		Databasename string `yaml:"databasename"`
		Password     string `yaml:"password"`
	} `yaml:"database"`
}

var ConfigData *Config
var Db *gorm.DB
const configPath = "./config/";
var configFiles [2]string = [2]string{"database","server"};
func Init() {
	ConfigData = &Config{}
	for _,configFile := range configFiles {
		loadConfig(configFile);
	}
	databaseConnect();
}
func loadConfig(configFileName string) {
	file, err := os.Open(configPath + configFileName +".yaml");
	if err != nil {
		log.Fatalln(err.Error());
	}
	defer file.Close();
	d := yaml.NewDecoder(file);
	if err := d.Decode(ConfigData); err != nil {
		log.Fatalln(err.Error());
	}
}
func databaseConnect(){
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host="+ConfigData.Database.Hostname+" user=" + ConfigData.Database.Username + " password=" + ConfigData.Database.Password + " dbname=" + ConfigData.Database.Databasename + " port=" + fmt.Sprintf("%v",ConfigData.Database.Port) + " sslmode=disable TimeZone=Asia/Shanghai";
	var err error;
	Db,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	sqlDB,err := Db.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = sqlDB.Ping();
	if err != nil {
		log.Fatalln("Feild Connect To database")
	}else {
		log.Println("database connected succfully")
	}
}