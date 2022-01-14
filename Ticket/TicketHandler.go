package TicketHandler

import (
	CustomersHandler "TicketManager/Customer"
	EventHandler "TicketManager/Event"
	"TicketManager/Utilitys"
	DbUtils "TicketManager/Utilitys/DbUtilitys"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TicketInterface interface {
	Sell()
	Check()
}

var (
	db *DbUtils.GreSQLResult
)

type Ticket struct {
	id           uuid.UUID
	EventId      uuid.UUID
	CustomerId   uuid.UUID
	ticketSerial string
	sellDate     time.Time
	sellTime     time.Time
	exceptions   *[]Utilitys.Exceptions
	Status       *Utilitys.Exceptions
}

func init() {
	db = DbUtils.NewConnection(nil)
}
func New(a *[]Utilitys.Exceptions) *Ticket {
	r := new(Ticket)
	r.id, _ = uuid.NewRandom()
	r.ticketSerial = ""
	r.sellDate = func() time.Time {
		a, _ := time.Parse("2017.09.07", time.Now().String())
		return a
	}()
	r.sellTime = func() time.Time {
		a, _ := time.Parse("2017.09.07", time.Now().String())
		return a
	}()
	r.exceptions = a
	return r
}
func (t *Ticket) Sell(c *CustomersHandler.Customer) {
	for _, e := range *EventHandler.New(nil).Events() {
		if e.ID == t.EventId {
			if t.CustomerId.String() == c.ID.String() {
				db.Command = fmt.Sprintf("INSERT INTO public.\"Tickets\"(\"ID\", \"EventId\", \"CustomerId\", \"ticketSerial\", \"sellDate\", \"sellTime\")VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", t.id, t.EventId, t.CustomerId, t.ticketSerial, t.sellDate, t.sellTime)
				db.PgExecuteNonQuery()
				if db.Status.Key != 0 {
					t.Status = Utilitys.SelectException(10000, t.exceptions)
					return
				}
			}
			t.Status = Utilitys.SelectException(10000, t.exceptions)
			return
		}
		t.Status = Utilitys.SelectException(10000, t.exceptions)
		return
	}
	t.Status = Utilitys.SelectException(10000, t.exceptions)
}

func (t *Ticket) Check(ticketNo string) {
	if db.Status.Key != 0 {
		t.Status = Utilitys.SelectException(10000, t.exceptions)
		return
	}
	db.Command = fmt.Sprintf("select count(*) from public.\"Tickets\" where \"ticketSerial\" = '%s'", ticketNo)
	db.PgExecuteNonQuery()

	var counter int = 0
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&counter); err != nil {
			t.Status = Utilitys.SelectException(10000, t.exceptions)
			return
		}
	}

	if db.Status.Key != 1 {
		t.Status = Utilitys.SelectException(10000, t.exceptions)
		return
	}
	t.Status = Utilitys.SelectException(0, t.exceptions)

}
