package theory

import "fmt"

type Config struct {
	Timeout float64
	StaticPath string
}

type HttpService struct {
	Name string
	Url string
	Config
}

func MainStructEmbedding() {
	MainStructEmbeddingExternal()
	MainStructEmbeddingConflict()
}

func MainStructEmbeddingExternal() {
	// Пример встраивания. Все поля встроенного объекта (Config) можно получить/установить из внешнего (HttpService)
	service := HttpService{
		Name: "service_name",
		Url: "http://random/host/url",
		Config: Config{
			Timeout: 10.0,
			StaticPath: "/opt/app/static",
		},
	}
	fmt.Printf("%#v\n", service)

	fmt.Println(service.Config.StaticPath) // /opt/app/static
	fmt.Println(service.StaticPath) // /opt/app/static
}

type KafkaService struct {
	Host string
	Port int32
	StaticPath string // StaticPath уже есть в Config. Проверка, что будет при конфликте.
	Config
}

func MainStructEmbeddingConflict() {
	// Пример встраивания, когда одно и то же поле/метод есть в обоих объектах.
	// Расширения не происходит. Каждое поле выполняет свою функциональность
	service := KafkaService{
		Host: "192.168.0.1",
		Port: 9092,
		StaticPath: "/kafka/service/static/path",
		Config: Config{
			Timeout: 10.0,
			StaticPath: "/config/static/path",
		},
	}
	fmt.Printf("%#v\n", service)
	fmt.Println(service.Config.StaticPath) // /config/static/path
	fmt.Println(service.Config.ImagePath("picture")) // /config/static/path/picture.png

	fmt.Println(service.StaticPath) // /kafka/service/static/path
	fmt.Println(service.ImagePath("picture")) // /kafka/service/static/path/picture.jpg
}

func (c *Config) ImagePath (filename string) string {
	return c.StaticPath + "/" + filename + ".png"
}

func (k *KafkaService) ImagePath (filename string) string {
	return k.StaticPath + "/" + filename + ".jpg"
}
