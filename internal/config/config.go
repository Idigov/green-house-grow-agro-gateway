package config

import (
	"os"
	"strconv"
	"time"
)

// Config: главная структура конфигурации Gateway.
// Передаётся в main и далее в server, middleware, clients.
type Config struct {
	Server    Server    // параметры HTTP-сервера
	JWT       JWT       // проверка JWT (секрет или публичный ключ)
	CORS      CORS      // разрешённые origins для браузера
	RateLimit RateLimit // лимиты запросов по IP и user
	Services  Services  // адреса внутренних сервисов
}

// Server: настройки HTTP-сервера.
type Server struct {
	Port            string        // порт
	ReadTimeout     time.Duration // макс. время чтения запроса; защита от медленных клиентов
	WriteTimeout    time.Duration // макс. время записи ответа
	ShutdownTimeout time.Duration // время на graceful shutdown при остановке
}

// JWT: проверка подписи токена.
// Используется либо Secret (HS256), либо PublicKey (RS256), в зависимости от Auth.
type JWT struct {
	Secret    string // секрет для HS256; Auth и Gateway должны использовать одно значение
	PublicKey string // публичный ключ для RS256; путь к файлу или PEM-строка
}

// CORS: разрешённые источники cross-origin запросов.
type CORS struct {
	Origins string // список через запятую, например "http://localhost:3000,https://app.example.com"
}

// RateLimit: ограничение частоты запросов.
type RateLimit struct {
	PerIP   int // макс. запросов в минуту с одного IP (до авторизации и защита от флуда)
	PerUser int // макс. запросов в минуту с одного user_id (после JWT)
}

// Services: адреса внутренних сервисов (gRPC или HTTP).
type Services struct {
	Auth       string // аутентификация и пользователи
	Directory  string // справочники (теплицы, блоки, циклы, культуры)
	Production string // урожай, планы, культурооборот
	Nutrition  string // питание растений, заправки, лабораторные анализы
	Protection string // СЗР, обработки, мониторинг болезней
	Warehouse  string // склад (остатки, документы, инвентаризация)
	Payroll    string // учёт работ, табель, расчёт оплаты
	Sales      string // реализация, отгрузки, контрагенты
	Lighting   string // досветка (режимы, потребление)
	Reporting  string // KPI, отчёты, своды для дашборда
	MinIO      string // S3-совместимое хранилище для pre-signed ссылок
}

// getEnv возвращает значение переменной окружения или defaultVal, если переменная пуста.
func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// getDuration парсит длительность из env; при ошибке или пустом значении возвращает defaultVal.
func getDuration(key string, defaultVal time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return defaultVal
	}
	return d
}

// getInt парсит int из env; при ошибке или пустом значении возвращает defaultVal.
func getInt(key string, defaultVal int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return n
}

// LoadConfig читает переменные окружения (префикс GATEWAY_) и возвращает Config.
// Перед вызовом можно вызвать godotenv.Load(), чтобы подтянуть .env.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: Server{
			Port:            getEnv("GATEWAY_PORT", "8080"),
			ReadTimeout:     getDuration("GATEWAY_READ_TIMEOUT", 10*time.Second),
			WriteTimeout:    getDuration("GATEWAY_WRITE_TIMEOUT", 10*time.Second),
			ShutdownTimeout: getDuration("GATEWAY_SHUTDOWN_TIMEOUT", 15*time.Second),
		},
		JWT: JWT{
			Secret:    getEnv("GATEWAY_JWT_SECRET", ""),
			PublicKey: getEnv("GATEWAY_JWT_PUBLIC_KEY", ""),
		},
		CORS: CORS{
			Origins: getEnv("GATEWAY_CORS_ORIGINS", "http://localhost:3000"),
		},
		RateLimit: RateLimit{
			PerIP:   getInt("GATEWAY_RATE_LIMIT_PER_IP", 100),
			PerUser: getInt("GATEWAY_RATE_LIMIT_PER_USER", 200),
		},
		Services: Services{
			Auth:       getEnv("GATEWAY_SERVICE_AUTH", ""),
			Directory:  getEnv("GATEWAY_SERVICE_DIRECTORY", ""),
			Production: getEnv("GATEWAY_SERVICE_PRODUCTION", ""),
			Nutrition:  getEnv("GATEWAY_SERVICE_NUTRITION", ""),
			Protection: getEnv("GATEWAY_SERVICE_PROTECTION", ""),
			Warehouse:  getEnv("GATEWAY_SERVICE_WAREHOUSE", ""),
			Payroll:    getEnv("GATEWAY_SERVICE_PAYROLL", ""),
			Sales:      getEnv("GATEWAY_SERVICE_SALES", ""),
			Lighting:   getEnv("GATEWAY_SERVICE_LIGHTING", ""),
			Reporting:  getEnv("GATEWAY_SERVICE_REPORTING", ""),
			MinIO:      getEnv("GATEWAY_SERVICE_MINIO", ""),
		},
	}
	return cfg, nil
}
