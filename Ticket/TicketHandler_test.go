package TicketHandler

import (
	CustomersHandler "TicketManager/Customer"
	"TicketManager/Utilitys"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTicket_Check(t1 *testing.T) {
	type fields struct {
		id           uuid.UUID
		EventId      uuid.UUID
		CustomerId   uuid.UUID
		ticketSerial string
		sellDate     time.Time
		sellTime     time.Time
		exceptions   *[]Utilitys.Exceptions
		Status       *Utilitys.Exceptions
	}
	type args struct {
		ticketNo string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TicketTest1",
			fields: fields{
				id:           uuid.UUID{},
				EventId:      uuid.UUID{},
				CustomerId:   uuid.UUID{},
				ticketSerial: "",
				sellDate: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				sellTime: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				exceptions: nil,
				Status:     nil,
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ticket{
				id:           tt.fields.id,
				EventId:      tt.fields.EventId,
				CustomerId:   tt.fields.CustomerId,
				ticketSerial: tt.fields.ticketSerial,
				sellDate:     tt.fields.sellDate,
				sellTime:     tt.fields.sellTime,
				exceptions:   tt.fields.exceptions,
				Status:       tt.fields.Status,
			}
			t.Check(tt.args.ticketNo)
		})
	}
}

func TestTicket_Sell(t1 *testing.T) {
	type fields struct {
		id           uuid.UUID
		EventId      uuid.UUID
		CustomerId   uuid.UUID
		ticketSerial string
		sellDate     time.Time
		sellTime     time.Time
		exceptions   *[]Utilitys.Exceptions
		Status       *Utilitys.Exceptions
	}

	tests := []struct {
		name     string
		fields   fields
		Customer CustomersHandler.Customer
		want     *Utilitys.Exceptions
	}{
		{
			name: "",
			fields: fields{
				id:           uuid.UUID{},
				EventId:      uuid.UUID{},
				CustomerId:   uuid.UUID{},
				ticketSerial: "",
				sellDate: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				sellTime: func() time.Time {
					a, _ := time.Parse("2017.09.07", time.Now().String())
					return a
				}(),
				exceptions: nil,
				Status:     nil,
			},
			Customer: CustomersHandler.Customer{
				ID:        nil,
				FirstName: "Mohsen",
				LastName:  "taheri",
				UserName:  "mhthrh",
				Password:  "Qaz@123321",
				CellNo:    "07759448882",
				Email:     "mhthrh@gmail.com",
				Status:    Utilitys.SelectException(0, Utilitys.RaiseError()),
			},
			want: Utilitys.SelectException(0, Utilitys.RaiseError()),
		},
	}
	for _, tt := range tests {
		t := New(Utilitys.RaiseError())
		t.id = tt.fields.id
		t.EventId = tt.fields.EventId
		t.CustomerId = tt.fields.CustomerId
		t.ticketSerial = tt.fields.ticketSerial
		t.sellDate = tt.fields.sellDate
		t.sellTime = tt.fields.sellTime
		t.exceptions = tt.fields.exceptions
		t.Status = tt.fields.Status
		t.Sell(&tt.Customer)
	}
}
