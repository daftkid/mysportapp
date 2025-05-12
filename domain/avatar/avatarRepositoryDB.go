package avatar

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

const DATABASE_URL = "postgres://root:root@localhost:5432/mysportapp"

type AvatarRepositoryDB struct {
	client *pgx.Conn
}

func (d AvatarRepositoryDB) FindAll() ([]Avatar, error) {
	findAllSql := "SELECT * FROM avatars"

	rows, err := d.client.Query(context.Background(), findAllSql)

	if err != nil {
		log.Println("Error while querying users table " + err.Error())
		return nil, err
	}

	avatars := make([]Avatar, 0)

	for rows.Next() {
		var a Avatar
		err := rows.Scan(&a.Id, &a.Path, &a.Alt)
		if err != nil {
			log.Println("Error while scanning avatars " + err.Error())
			return nil, err
		}

		avatars = append(avatars, a)
	}

	return avatars, nil
}

func (d AvatarRepositoryDB) GetById(id string) (*Avatar, error) {
	avatarSql := "SELECT * FROM avatars WHERE id = $1"

	row := d.client.QueryRow(context.Background(), avatarSql, id)

	var a Avatar

	err := row.Scan(&a.Id, &a.Path, &a.Alt)

	if err != nil {
		log.Println("Error scanning avatar " + err.Error())
		return nil, err
	}

	return &a, nil
}

func NewAvatarRepositoryDB() AvatarRepositoryDB {
	client, err := pgx.Connect(context.Background(), DATABASE_URL)

	if err != nil {
		log.Println("Unable to connect to database: %v\n", err.Error())
		panic(err)
	}

	//defer client.Close(context.Background())

	return AvatarRepositoryDB{client}
}
