package repository

import (
	"PhoneBook/interface1"
)

//type ContactRepository interface {
//	GetContact(id string) interface1.Contact
//	GetContactList() []interface1.Contact
//	AddContact(contact interface1.Contact)
//	UpdateContact(id string, contact interface1.Contact) interface1.Contact
//	DeleteContact(id string)
//	FindById(id string) bool
//}

//type Contact struct {
//	Name  string
//	Phone string
//}
type phoneBookRepository struct {
	list map[string]interface1.Contact
}

var ai = 0

func NewPhoneBookRepository() *phoneBookRepository {
	return &phoneBookRepository{list: map[string]interface1.Contact{}}
}

func (pb *phoneBookRepository) FindById(key string) bool {
	_, ok := pb.list[key]
	return ok
}

func (pb *phoneBookRepository) GetContact(key string) interface1.Contact {
	return pb.list[key]
}

func (pb *phoneBookRepository) GetContactList() []interface1.Contact {
	var list []interface1.Contact
	for _, item := range pb.list {
		list = append(list,item)
	}
	return list
}

func (pb *phoneBookRepository) AddContact(c interface1.Contact) {
	pb.list[c.Name] = c
	ai++
}

func (pb *phoneBookRepository) DeleteContact(key string) {
	delete(pb.list, key)
}

func (pb *phoneBookRepository) UpdateContact(key string, c interface1.Contact) {
	pb.list[key] = c
}
func (pb *phoneBookRepository) IsExist(name string) bool  {
	_, ok := pb.list[name]
	return  ok
}


