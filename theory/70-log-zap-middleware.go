package theory

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var sugar zap.SugaredLogger

func initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	sugar = *logger.Sugar()
}

func MainZapMiddleware() {
	initLogger()
	defer sugar.Sync()

	http.Handle("/ping", withLogging2(pingHandler()))

	addr := "127.0.0.1:8080"
	sugar.Infow("Starting server", "addr", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		sugar.Fatalw(err.Error(), "event", "start server failed")
	}
}

func pingHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, "pong\n")
	}
	return http.HandlerFunc(fn)
}

//func withLogging(next http.Handler) http.Handler {
//	logFn := func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//		next.ServeHTTP(w, r)
//		duration := time.Since(start)
//
//		sugar.Infoln("uri", r.RequestURI, "method", r.Method, "duration", duration)
//	}
//	return http.HandlerFunc(logFn)
//}

func withLogging2(next http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseData := &responseData{status: 0, size: 0}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}
		next.ServeHTTP(&lw, r)
		duration := time.Since(start)

		sugar.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"duration", duration,
			"status", responseData.status,
			"size", responseData.size,
		)
	}
	return http.HandlerFunc(logFn)
}

// встраиваение в оригинальный http.ResponseWriter логирование статуса ответа и размера
type (
	// берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}
)

// Write переопределение
func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := w.ResponseWriter.Write(b)
	w.responseData.size += size
	return size, err
}

// WriteHeader переопределение
func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.responseData.status = statusCode
}
