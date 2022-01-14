package DbUtils

import (
	"TicketManager/Utilitys"
)

func (d *GreSQLResult) PgExecuteNonQuery() {
	var err error
	d.Status = Utilitys.SelectException(0, d.Exception)
	d.ResultSet, err = d.db.Query(d.Command)
	if err != nil {
		d.Status = Utilitys.SelectException(10000, d.Exception)
	}
}

func (d *GreSQLResult) PgLastInsertId() {
	d.Status = Utilitys.SelectException(0, d.Exception)
	err := d.db.QueryRow(d.Command).Scan(&d.ResultSet)
	if err != nil {
		d.Status = Utilitys.SelectException(10007, d.Exception)
	}
}
