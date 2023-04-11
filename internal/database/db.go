package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type weatherdb struct {
	conn *pgx.Conn
}

var dbinstance *weatherdb

func GetInstance() *weatherdb {
	if dbinstance != nil {
		return dbinstance
	}
	dbinstance = &weatherdb{}
	return dbinstance
}

func (d *weatherdb) CreateConnection(parentCtx context.Context) error {
	if d.conn != nil {
		return nil
	}
	conn, err := pgx.Connect(parentCtx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	dbinstance.conn = conn
	return nil
}

func (d *weatherdb) Close(parentCtx context.Context) error {
	return d.conn.Close(parentCtx)
}

func (d *weatherdb) CreateNecessaryTables(parentCtx context.Context) error {
	err := d.conn.QueryRow(parentCtx,
		"CREATE TABLE IF NOT EXISTS weather ("+
			"id integer,"+
			"lon decimal,"+
			"lat decimal,"+
			"temp decimal,"+
			"PRIMARY KEY (id)"+
			")",
	).Scan()
	if err != nil && err != pgx.ErrNoRows {
		return err
	}
	return nil
}
