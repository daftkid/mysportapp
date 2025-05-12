package avatar

type Avatar struct {
	Id   int
	Path string
	Alt  string
}

type AvatarRepository interface {
	FindAll() ([]Avatar, error)
	GetById(string) (*Avatar, error)
	//CreateAvatar(avatar Avatar) (int, error)
	//ReadAvatar(int) (Avatar, error)
	//UpdateAvatar(avatar Avatar) (int, error)
	//DeleteAvatar(int) error
}
