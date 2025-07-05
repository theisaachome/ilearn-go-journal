package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Contact represents a single entry in our phone book
type Contact struct{
	Name string
	Phone string
	Email string
}

// PhoneBook manages our collection of contacts
type PhoneBook struct{
	contacts map[string]Contact // Using map for quick lookup by name
	order []string // To maintain the order of contacts
}



// NewPhoneBook creates and returns a new empty PhoneBook
func NewPhoneBook() *PhoneBook{
	return  &PhoneBook{
		contacts: make(map[string]Contact),
		order: make([]string, 0),
	}
}

func (pb *PhoneBook)AddContact(name,phone,email string)(bool,error){
	// Check if the contact already exists
	if _ , exists :=pb.contacts[name]; exists{
		return false, fmt.Errorf("contact with name %s already exists", name)
	}
	newContact := Contact{
		Name:  name,
		Phone: phone,
		Email: email,	
	}
	pb.contacts[name]= newContact
	pb.order = append(pb.order, name)	
	return true,nil
}

func (pb *PhoneBook) ListAllContacts() []Contact{
	var allContacts []Contact
	for _, name := range pb.order {
		if contact, exists := pb.contacts[name]; exists {
			allContacts = append(allContacts, contact)
		}
	}
	return allContacts
} 
func (pb *PhoneBook) GetContactByName(name string) (Contact, bool) {
	contact, found := pb.contacts[name]
	if !found {
		return Contact{}, false // Return empty contact if not found
	}
	return contact, true // Return the found contact
}
func (pb *PhoneBook) DeleteContact(name string) bool {
	if _,exist := pb.contacts[name]; !exist {
		return false
	}
	delete(pb.contacts, name)
	newOrder :=make([]string, 0,len(pb.order)-1)
	for _,n :=range pb.order {
		if n != name {
			newOrder = append(newOrder, n)
		}
	}
	pb.order = newOrder
	return true
}
func main() {
	fmt.Println("Hello, Phone book App!")
	pb := NewPhoneBook()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to your Go Phone Book!")
	fmt.Println("------------------------------")

	operatePhoneBook(reader, pb)
}

func operatePhoneBook(reader *bufio.Reader,phoneBook *PhoneBook) {
	for{
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List All Contacts")
		fmt.Println("3. Search Contact by Name")
		fmt.Println("4. Delete Contact by Name")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice,err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 5.")
			continue
		}
		switch choice {
		case 1:
			fmt.Print("Enter contact name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter phone number: ")
			phone, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phone)

			fmt.Print("Enter email (optional, press Enter to skip): ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			added, err := phoneBook.AddContact(name, phone, email) // Multiple return values
			if added {
				fmt.Printf("Contact '%s' added successfully.\n", name)
			} else {
				fmt.Printf("Error adding contact: %v\n", err)
			}
		case 2:
			contacts := phoneBook.ListAllContacts() // Returns a slice
			if len(contacts) == 0 {
				fmt.Println("No contacts in the phone book.")
				continue
			}
			fmt.Println("\n--- All Contacts ---")
			for i, contact := range contacts {
				fmt.Printf("%d. Name: %s, Phone: %s", i+1, contact.Name, contact.Phone)
				if contact.Email != "" {
					fmt.Printf(", Email: %s", contact.Email)
				}
				fmt.Println()
			}
			fmt.Println("--------------------")
		case 3:
			fmt.Print("Enter name to search: ")
			nameToSearch, _ := reader.ReadString('\n')
			nameToSearch = strings.TrimSpace(nameToSearch)

			contact, found := phoneBook.GetContactByName(nameToSearch) // Multiple return values
			if found {
				fmt.Printf("Found: Name: %s, Phone: %s", contact.Name, contact.Phone)
				if contact.Email != "" {
					fmt.Printf(", Email: %s", contact.Email)
				}
				fmt.Println()
			} else {
				fmt.Printf("Contact '%s' not found.\n", nameToSearch)
			}
		case 4:
			fmt.Print("Enter name to delete: ")
			nameToDelete, _ := reader.ReadString('\n')
			nameToDelete = strings.TrimSpace(nameToDelete)

			deleted := phoneBook.DeleteContact(nameToDelete) // Single boolean return
			if deleted {
				fmt.Printf("Contact '%s' deleted successfully.\n", nameToDelete)
			} else {
				fmt.Printf("Contact '%s' not found or could not be deleted.\n", nameToDelete)
			}
		case 5:
			fmt.Println("Exiting Phone Book. Goodbye!")
			return // Exit the program

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}			
		
	}
}