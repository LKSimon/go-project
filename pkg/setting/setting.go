package setting

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
	"gopkg.in/yaml.v2"
)

var (
	DOCKER_REPO_NAME string
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath               string
	LogSaveName               string
	LogFileExt                string
	TimeFormat                string
	OutputConsole             bool
	IntelligenceOpinionSearch string
}

var AppSetting = &App{}

type LogConfig struct {
	Level           string
	JsonFormat      bool
	StacktraceLevel string
	Stdout          bool
	RuntimeRootPath string
	LogOutDir       string
	LogOutFileName  string
	LogFileExt      string
	LogOutMaxSize   int
	LogOutMaxBackup int
	LogOutMaxAge    int
	LogOutLocalTime bool
	LogOutCompress  bool
}

var LogConfigSetting = &LogConfig{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ENV          string
}

var ServerSetting = &Server{}

type Database struct {
	Url       string
	OpenConns int
	IdleConns int
}

var DatabaseSetting = &Database{}

type CompanyDB struct {
	Url       string
	OpenConns int
	IdleConns int
}

var IntelligenceDBSetting = &CompanyDB{}

type DashboardDb struct {
	Url       string
	OpenConns int
	IdleConns int
}

var DashboardDbSetting = &DashboardDb{}

type PostgresqlDb struct {
	Url      string
	Schema   string // postgrespl schema, 默认public
	Username string
}

var PostgresqlDbSetting = &PostgresqlDb{}

type ClickhouseDb struct {
	Url       string
	OpenConns int
	IdleConns int
}

var ClickhouseDbSetting = &ClickhouseDb{}

type ClickHouseSync struct {
	Cycle int
}

var ClickHouseSyncSetting = &ClickHouseSync{}

type LabelSync struct {
	Cycle int
}

var LabelSyncSetting = &LabelSync{}

type SMTPConfig struct {
	Host       string
	Port       int32
	NickName   string
	Username   string
	Password   string
	MaxSendJob int32
}

var SMTPSetting = &SMTPConfig{}

type RedisCacheConfig struct {
	Host        string
	Port        string
	Password    string
	Db          int
	KeyPrefix   string
	CacheEnable bool
}

var ResourceSetting = &ResourceConfig{}

type ResourceConfig struct {
	Url       string
	SecretID  string
	SecretKey string
}

var TranslateSetting = &TranslateConfig{}

type TranslateConfig struct {
	Url       string
	SecretId  string
	SecretKey string
}

var DataCacheSetting = &RedisCacheConfig{}
var InterfaceCacheSetting = &RedisCacheConfig{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup(env string) {
	//var err error
	//DOCKER_REPO_NAME = os.Getenv("DOCKER_REPO_NAME")
	//if DOCKER_REPO_NAME == "svr" {
	//	log.Fatalf("setting.Setup, fail to get DOCKER_REPO_NAME: %s", DOCKER_REPO_NAME)
	//	panic(err)
	//}
	var err error
	iniPath := fmt.Sprintf("conf/app_%s.ini", env)
	cfg, err = ini.Load(iniPath)
	if err != nil {
		//log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		panic(err)
		return
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("intelligenceDB", IntelligenceDBSetting)
	mapTo("dashboardDb", DashboardDbSetting)
	mapTo("postgresqlDb", PostgresqlDbSetting)
	mapTo("clickhouseDb", ClickhouseDbSetting)
	mapTo("clickHouseSync", ClickHouseSyncSetting)
	mapTo("labelSync", LabelSyncSetting)

	mapTo("smtpSrv", SMTPSetting)
	mapTo("dataCache", DataCacheSetting)
	mapTo("interfaceCache", InterfaceCacheSetting)
	mapTo("resourceManager", ResourceSetting)
	mapTo("translate", TranslateSetting)
	mapTo("log", LogConfigSetting)

	if PostgresqlDbSetting.Schema == "" {
		PostgresqlDbSetting.Schema = "public"
	}
	log.Printf("PostgresqlDbSetting.Schema: %s", PostgresqlDbSetting.Schema)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

// LoadBizConf 读取业务配置
func LoadBizConf(dataStruct interface{}) {
	var err error
	DOCKER_REPO_NAME = os.Getenv("DOCKER_REPO_NAME")
	if DOCKER_REPO_NAME == "svr" {
		log.Fatalf("setting.LoadBizConf, fail to get DOCKER_REPO_NAME: %s", DOCKER_REPO_NAME)
		panic(err)
	}
	iniPath := fmt.Sprintf("conf/biz_conf_%s.yaml", DOCKER_REPO_NAME)
	content, err := ioutil.ReadFile(iniPath)
	if err != nil {
		log.Fatalf("setting.LoadBizConf, fail to read '%v': %v", iniPath, err)
	}

	err = yaml.Unmarshal(content, dataStruct)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("load biz conf: %v", dataStruct)
}
