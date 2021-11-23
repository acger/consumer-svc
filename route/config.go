package route

type Config struct {
	ChatDataSource string
	Queue struct{
		Group string
		Host []string
		Topic []string
	}
	Log struct{
		Path string
	}
}
