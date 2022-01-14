package DbUtils

import (
	"TicketManager/Utilitys"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DatabaseInterfaces interface {
	NewConnection(*GreSQL)
	CloseConnection()
	PgExecuteNonQuery()
	PgLastInsertId()
}

type GreSQLResult struct {
	db        *sql.DB
	Command   string
	Exception *[]Utilitys.Exceptions
	Status    *Utilitys.Exceptions
	ResultSet interface{}
}
type GreSQL struct {
	Host      string
	Port      int32
	User      string
	Pass      string
	Dbname    string
	Driver    string
	Exception *[]Utilitys.Exceptions
}

func NewConnection(g *GreSQL) *GreSQLResult {
	r := new(GreSQLResult)
	if g == nil {
		g = &GreSQL{
			Host:      "localhost",
			Port:      5432,
			User:      "postgres",
			Pass:      "123456",
			Dbname:    "Ticket",
			Driver:    "postgres",
			Exception: Utilitys.RaiseError(),
		}
	}
	r.Status = Utilitys.SelectException(0, g.Exception)
	db, err := sql.Open(g.Driver, fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		g.Host, g.Port, g.User, g.Pass, g.Dbname))
	if err != nil {
		r.Status = Utilitys.SelectException(10005, g.Exception)
	}
	r.Exception = g.Exception
	r.db = db
	r.Command = ""
	r.ResultSet = nil
	return r
}

func (d *GreSQLResult) CloseConnection() {
	d.Status = Utilitys.SelectException(0, d.Exception)
	if err := d.db.Close(); err != nil {
		d.Status = Utilitys.SelectException(10006, d.Exception)
	}
}
