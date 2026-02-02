package structures_queue_test

import (
	"testing"

	structures_queue "github.com/ferreira-gn/estrutura-de-dados/estruturas/queue"
)

func TestFromCreateNewQueue( t *testing.T){
	newQueue := structures_queue.New[string]()
	
	if newQueue == nil {
		t.Errorf("Error creating a new queue, the constructor should returns a new pointer to queue.")
	}
}