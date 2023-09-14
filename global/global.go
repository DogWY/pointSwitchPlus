package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"project/lora"
	"time"
)

var (
	Logger *zap.SugaredLogger
	Port   *lora.Port
	DB     *gorm.DB
	Addrs  []int
	States map[int]int
	MDis   map[int][]int
	MTimes map[int][]string
)

const (
	Length = 5
)

func init() {
	InitLog()
	InitConfig()
	InitDB()
	InitPort()
	Addrs = viper.GetIntSlice("addrs")
	States = make(map[int]int)
	MDis = make(map[int][]int)
	MTimes = make(map[int][]string)
	for _, addr := range Addrs {
		States[addr] = 0
		MDis[addr] = make([]int, Length)
		MTimes[addr] = make([]string, Length)
	}
}

func InitLog() {
	logName := time.Now().Format("2006-01-02-15-04-05")
	file, err := os.Create("./" + logName + ".log")
	if err != nil {
		log.Fatalf("create log fail")
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)

	logger := zap.New(core)
	Logger = logger.Sugar()
}

func InitConfig() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Errorw("read config fail")
	} else {
		Logger.Infow("read config success")
	}
}

func InitDB() {
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")))
	if err != nil {
		Logger.Errorw("open mysql fail")
	} else {
		Logger.Infow("open mysql success")
	}
	DB = db
}

func InitPort() {
	cfg := lora.Config{
		Name:        viper.GetString("port.Name"),
		Baud:        viper.GetInt("port.Baud"),
		ReadTimeout: viper.GetDuration("port.ReadTimeout"),
		Addr:        viper.GetInt("port.Addr"),
	}
	port, err := lora.OpenPort(cfg)
	if err != nil {
		Logger.Errorw("open port fail")
	} else {
		Logger.Infow("open port success")
	}
	Port = port
}
