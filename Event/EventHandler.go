package EventHandler

import (
	"TicketManager/Utilitys"
	DbUtils "TicketManager/Utilitys/DbUtilitys"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	db *DbUtils.GreSQLResult
)

type EventInterface interface {
	AddEvent()
	Events() *[]Event
}

type Event struct {
	ID            uuid.UUID
	Name          string
	Duration      int
	Count         int
	Amount        float32
	Date          time.Time
	Time          time.Time
	StartSellDate time.Time
	EndSellDate   time.Time
	creationDate  time.Time
	exceptions    *[]Utilitys.Exceptions
	Status        *Utilitys.Exceptions
}

func init() {
	db = DbUtils.NewConnection(nil)
}
func New(e *[]Utilitys.Exceptions) *Event {
	result := new(Event)
	result.ID, _ = uuid.NewRandom()
	result.creationDate = time.Now()
	result.exceptions = e
	return result
}
func (e *Event) AddEvent() {
	if db.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}
	db.Command = fmt.Sprintf("INSERT INTO public.\"Events\"(\"ID\", name, \"Duration\", \"Count\", \"Amount\", \"Date\", \"Time\", \"StartSellDate\", \"EndSellDate\", \"creationDate\", \"IsOpen\")VALUES ('%s', '%d', '%d', '%2f', '%s', '%s', '%s', '%s')", e.ID, e.Duration, e.Count, e.Amount, e.Date, e.Time, e.StartSellDate, e.creationDate)
	db.PgExecuteNonQuery()

	if db.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return
	}
	Utilitys.SelectException(0, e.exceptions)
}

func (e *Event) Events() *[]Event {
	var result []Event
	if db.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return nil
	}
	db.Command = fmt.Sprintf("SELECT \"ID\", name, \"Duration\", \"Count\", \"Amount\", \"Date\", \"Time\", \"StartSellDate\", \"EndSellDate\" FROM public.\"Events\" where \"IsOpen\" ='%s'", "0")
	db.PgExecuteNonQuery()

	if db.Status.Key != 0 {
		e.Status = Utilitys.SelectException(10000, e.exceptions)
		return nil
	}
	for db.ResultSet.(*sql.Rows).Next() {
		if err := db.ResultSet.(*sql.Rows).Scan(&e.ID, &e.Name, &e.Duration, &e.Count, &e.Amount, &e.Date, &e.Time, &e.StartSellDate, &e.EndSellDate); err != nil {
			e.Status = Utilitys.SelectException(10000, e.exceptions)
			return nil
		}
		result = append(result, *e)
	}
	e.Status = Utilitys.SelectException(0, e.exceptions)
	return &result
}
