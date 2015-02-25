package data

type Login struct {
	LoginId int
	UserName string `form:"UserName"`
	Password string `form:"Password"`
}

