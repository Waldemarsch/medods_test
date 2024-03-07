package infrastructure

type Repository interface {
	GetToken()
	CreateToken()
}

type Infrastructure struct {
	Repository
}
