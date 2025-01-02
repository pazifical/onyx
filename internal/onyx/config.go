package onyx

type Config struct {
	Port              int
	MarkdownDirectory string
	MatrixConfig      MatrixConfig
}

type MatrixConfig struct {
	Username string
	Password string
	RoomID   string
}
