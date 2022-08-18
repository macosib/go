package database

import (
	"context"
	"database/sql"
	"time"
)

//InitDatabase - Функции сощдает таблицы и связи, если их нет.
func InitDatabase(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Users(id SERIAL PRIMARY KEY, user_name varchar(40) NOT NULL, age integer NOT NULL)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS FriendList(
	userId integer REFERENCES Users(id) ON DELETE CASCADE,
	friendId integer REFERENCES Users(id) ON DELETE CASCADE,
	PRIMARY KEY (userId, friendId))`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

}
