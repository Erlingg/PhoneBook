package interface1

type Contact struct {
	Name  string
	Phone string
}

type ContactRepository interface {
	GetContact(id string) Contact
	GetContactList() []Contact
	AddContact(contact Contact)
	UpdateContact(id string, contact Contact)
	DeleteContact(id string)
	FindById(id string) bool
	IsExist(name string) bool
}

type ContactService interface {
	GetContactService(id string) Contact
	GetContactListService() []Contact
	AddContactService(contact Contact) error
	UpdateContactService(id string, contact Contact)
	DeleteContactService(id string)
}