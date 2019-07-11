package service

import (
	"PhoneBook/interface1"
	"errors"
	"strings"
)

type phoneBookService struct {
	repo interface1.ContactRepository
}

func NewPhoneBookService(r interface1.ContactRepository) *phoneBookService {
	return &phoneBookService{repo:r}
}
func (s phoneBookService) AddContactService(c interface1.Contact) error{
	c.Name = strings.Trim(c.Name, " ")
	if c.Name == "" {
		return errors.New("Name empty")
	}
	if s.repo.IsExist(c.Name){
		return errors.New("Exist")
	}
	s.repo.AddContact(c)
	return nil
}

func (s phoneBookService) GetContactService(id string) interface1.Contact {
	return s.repo.GetContact(id)
}

func (s phoneBookService) GetContactListService() []interface1.Contact {
	return s.repo.GetContactList()
}

func (s phoneBookService) UpdateContactService(id string, c interface1.Contact) {
	s.repo.UpdateContact(id, c)
}

func (s *phoneBookService) DeleteContactService(id string) {
	s.repo.DeleteContact(id)
}

