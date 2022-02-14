package conf

import (
	"fmt"
	"strings"

	"code.huawink.com/beiwanglu/user/model"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

//服务配置，初始化
func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("配置文件错误，请检查文件路径，err:", err)
	}
	//加载mysql数据和配置信息
	LoadMysqlData(file)
	//拼接成具体路径字符串，传入Database
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	//与mysql数据库进行连接
	model.Database(path)
}

//加载mysql数据和配置信息
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
