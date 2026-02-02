package structures_linkedlist_test

import (
	"testing"

	structures_linkedlist "github.com/ferreira-gn/estrutura-de-dados/estruturas/linked_list"
)

func TestFromCreateNewLinkedList( t *testing.T){
	newLinkedList := structures_linkedlist.New[string]()
	
	if newLinkedList == nil {
		t.Errorf("Error creating a new linked list, the constructor should returns a new pointer to the linked list ")
	}
}