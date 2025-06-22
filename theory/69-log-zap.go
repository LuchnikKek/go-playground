// zap - скорость (x5 от logrus)
// использует рефлексию

package theory

import (
	"go.uber.org/zap"
)

func MainLogZap() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	const url = "https://www.google.com"

	// SugaredLogger
	sugar := logger.Sugar()

	sugar.Infow("Failed to fetch URL", "url", url, "attempt", 3)

	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Errorf("Failed to fetch URL: %s", url)

	// переводим в обычный Logger
	plain := sugar.Desugar()

	plain.Info("Hello, World!")
	plain.Error("Failed to fetch URL:", zap.String("url", url))
}
