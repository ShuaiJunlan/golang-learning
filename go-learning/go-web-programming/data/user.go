package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id       int
	Uuid     string
	Email    string
	UserId   int
	CreatedAt time.Time
}

func (user *User) CreateSession(session Session, err error) {
	statement := ""
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Uuid, &session.Email, &session.Id, &session.CreatedAt)
}
//Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("").
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("").Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete "
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(session.Uuid)
	return
}
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("").Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func SessionDeleteAll() (err error) {
	statement := ""
	_, err = Db.Exec(statement)
	return
}
func (user *User) Create() (err error) {
	statement := ""
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}
