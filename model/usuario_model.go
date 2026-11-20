package model

type User struct {
	Id    string
	First string
	Last  string
	Born  int
}

type UserDTO struct {
	Id    string `firestore:"id"`
	First string `firestore:"first"`
	Last  string `firestore:"last"`
	Born  int    `firestore:"born"`
}

func (u UserDTO) ToEntity() (User, error) {
	return User{
		Id:    u.Id,
		First: u.First,
		Last:  u.Last,
		Born:  u.Born,
	}, nil
}
