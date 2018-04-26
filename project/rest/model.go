package main

import (
	"database/sql"
	"fmt"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("select name, age from users where id=%d", u.ID)

	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("update users set name='%s', age=%d where id=%d", u.Name, u.Age, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("delete from users where id='%d'", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("insert into users(name, age) values('%s', '%d')", u.Name, u.Age)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("select last_insert_id()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("select id, name, age from users limit %d offset %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
