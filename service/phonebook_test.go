package service

import (
	"PhoneBook/interface1"
	"PhoneBook/repository"
	"testing"
)

func TestPhoneBookService_AddContactService(t *testing.T) {
	phonebook := repository.NewPhoneBookRepository()
	srv := NewPhoneBookService(phonebook)
	err := srv.AddContactService(interface1.Contact{})
	if err != nil{
		t.Error()
	}
}
