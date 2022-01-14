package EventHandler

import (
	"TicketManager/Utilitys"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestEvent_AddEvent(t *testing.T) {

	type fields struct {
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

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test1",
			fields: fields{
				ID:       uuid.New(),
				Name:     "World Cup",
				Duration: 120,
				Count:    1000,
				Amount:   12.3,
				Date: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				Time: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				StartSellDate: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				EndSellDate: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				creationDate: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				exceptions: Utilitys.RaiseError(),
				Status:     nil,
			},
		},
	}
	for _, tt := range tests {
		e := &Event{
			ID:            tt.fields.ID,
			Name:          tt.fields.Name,
			Duration:      tt.fields.Duration,
			Count:         tt.fields.Count,
			Amount:        tt.fields.Amount,
			Date:          tt.fields.Date,
			Time:          tt.fields.Time,
			StartSellDate: tt.fields.StartSellDate,
			EndSellDate:   tt.fields.EndSellDate,
			creationDate:  tt.fields.creationDate,
			exceptions:    tt.fields.exceptions,
			Status:        tt.fields.Status,
		}
		e.AddEvent()
	}
}

func TestEvent_Events(t *testing.T) {

	tests := []struct {
		name string
		want *[]Event
	}{
		// TODO: Add test cases.
	}
	e := New(Utilitys.RaiseError())
	for _, tt := range tests {

		if got := e.Events(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Events() = %v, want %v", got, tt.want)
		}
	}
}
