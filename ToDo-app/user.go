package todo

type User struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	Userrname string `json:"username"`
	Password  string `json:"password"`
}
