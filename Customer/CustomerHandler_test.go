package CustomersHandler

import (
	"TicketManager/Utilitys"
	"reflect"
	"testing"
)

func TestCustomer_SignIn(t *testing.T) {
	type fields struct {
		UserName string
		Password string
	}
	tests := []struct {
		name   string
		fields *Customer
		want   *Utilitys.Exceptions
	}{
		{
			name: "SignInTest1",
			fields: &Customer{
				UserName: "mhthrh",
				Password: "Qaz@321123",
			},
			want: Utilitys.SelectException(0, Utilitys.RaiseError()),
		},
		{
			name: "SignInTest1",
			fields: &Customer{
				UserName: "mhthrh",
				Password: "Qaz@321123",
			},
			want: Utilitys.SelectException(10000, Utilitys.RaiseError()),
		},
		{
			name: "SignInTest1",
			fields: &Customer{
				UserName: "mhthrh",
				Password: "Qaz@321123",
			},
			want: Utilitys.SelectException(10000, Utilitys.RaiseError()),
		},
	}
	for _, tt := range tests {
		c := New(Utilitys.RaiseError())
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password
		c.SignIn()
	}
}

func TestCustomer_SignUp(t *testing.T) {

	tests := []struct {
		name   string
		fields *Customer
		want   *Utilitys.Exceptions
	}{
		{
			name: "Test1",
			fields: &Customer{
				FirstName: "Mohsen",
				LastName:  "taheri",
				UserName:  "mhthrh",
				Password:  "Qwe@321123",
				CellNo:    "07759448882",
				Email:     "test@test.org",
			},
			want: Utilitys.SelectException(0, Utilitys.RaiseError()),
		},
		{
			name: "Test1",
			fields: &Customer{
				FirstName: "Mohsen",
				LastName:  "taheri",
				UserName:  "mhthrh",
				Password:  "123456",
				CellNo:    "0775---9448882",
				Email:     "testtest.org",
			},
			want: Utilitys.SelectException(0, Utilitys.RaiseError()),
		},
		{
			name: "Test1",
			fields: &Customer{
				FirstName: "Mohsen",
				LastName:  "taheri",
				UserName:  "mhthrh",
				Password:  "123456",
				CellNo:    "07759448882",
				Email:     "test@test.org",
			},
			want: Utilitys.SelectException(0, Utilitys.RaiseError()),
		},
	}
	for _, tt := range tests {
		c := New(Utilitys.RaiseError())
		c.FirstName = tt.fields.FirstName
		c.LastName = tt.fields.LastName
		c.UserName = tt.fields.UserName
		c.Password = tt.fields.Password
		c.CellNo = tt.fields.CellNo
		c.Email = tt.fields.Email
		c.SignUp()

		if got := c.Status.Key; !reflect.DeepEqual(got, tt.want.Key) {
			t.Errorf("Events() = %v, want %v", tt.want.Key, got)
		}
	}
}
