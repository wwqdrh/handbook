package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SimpleSugarLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "https://wwqdrh.github.io",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "https://wwqdrh.github.io")
}

// 比 SugaredLogger 更快并且分配的更少，但它只支持结构化日志记录
// Info中使用zap.String、zap.Int等等
func SimpleBasicLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger.Info("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		zap.String("url", "https://wwqdrh.github.io"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// 配置zapcore
func ConfigureLogger() {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	topicDebugging := zapcore.AddSync(ioutil.Discard) // ioutil.DIscard不会真实写入到某个位置，直接返回其长度
	topicErrors := zapcore.AddSync(ioutil.Discard)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}

// 使用json字符串配置logger
func ConfigureJsonLogger() {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
}

func MultiLogger() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	logger := zap.NewExample() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func LoggerWithColor() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		//EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, //这里可以指定颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.InfoLevel)
	config := zap.Config{
		Level:       atom, // 日志级别
		Development: true, // 开发模式，堆栈跟踪
		//Encoding:         "json",                                              // 输出格式 console 或 json
		Encoding:         "console",                                            // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                        // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "wisdom_park"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"},                                   // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder //这里可以指定颜色

	// infolevel 构建日志
	ZapLog_V1, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}

	ZapLog_V1.Info("log 初始化成功")

	// 动态调整
	config.Level.SetLevel(zap.WarnLevel)
	warnLogger, _ := config.Build()
	warnLogger.Info("log 初始化成功")
}
