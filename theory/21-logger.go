package theory

import (
	"io"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
)

/* Вариант 1: Логгер в логируемой структуре */
type MetricsRepository struct {
	l *zap.Logger
	// ещё поля
}

type metricModel struct {
	id string
}

func (mr *MetricsRepository) Save(m metricModel) error {
	mr.l.Info("Savind metric to log", zap.String("id", m.id))

	if m.id == "" {
		panic("No id")
	}
	return nil
}

func MainLoggers() {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()
	rep := MetricsRepository{l: logger}
	_ = rep.Save(metricModel{id: "31231"})

	// 1.2, приятный API
	sugar := logger.Sugar()
	sugar.Infoln("New line")
}

// Вариант 2: через синглтон
// Применим, если ко всему в пакете нужен одинаковый подход

var (
	logger         *zap.Logger
	once           sync.Once
	loggerFilePath = "log.json"
)

func GetLogger() zap.Logger {
	once.Do(func() {
		file, err := os.Create(loggerFilePath)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		defer func() {
			if err = file.Close(); err != nil {
				log.Fatalf("%s", err.Error())
			}
		}()

		configureLogger(file, logger)
	})
	return *logger
}

func configureLogger(file io.Writer, l *zap.Logger) {
	// логика конфигурации логгера
}
