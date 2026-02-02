package structures_stack_test

import (
	"testing"

	stack "github.com/ferreira-gn/estrutura-de-dados/estruturas/stack"
)

func TestCreateNewStack(t *testing.T) {
	testStack := stack.New[int]()

	if testStack == nil {
		t.Fatalf("Ao criar uma stack com New deve ser retornado um ponteiro válido, recebido %v", testStack)
	}

	if testStack.Len() != 0 {
		t.Errorf("Uma stack recém-criada deve ter tamanho 0, recebido %d", testStack.Len())
	}
}

func TestPushAddsValueOnTop(t *testing.T) {
	testStack := stack.New[int]()

	testStack.Push(2)
	testStack.Push(3)
	testStack.Push(4)

	top := testStack.Peek()

	if top != 4 {
		t.Errorf(
			"O último valor inserido deve estar no topo da stack.\nEsperado: 4\nRecebido: %d",
			top,
		)
	}
}

func TestPopRemovesValueFromTop(t *testing.T) {
	testStack := stack.New[int]()

	testStack.Push(2)
	testStack.Push(3)

	removed := testStack.Pop()
	if removed != 3 {
		t.Errorf(
			"Pop deve retornar o valor removido do topo.\nEsperado: 3\nRecebido: %d",
			removed,
		)
	}

	top := testStack.Peek()
	if top != 2 {
		t.Errorf(
			"Após o Pop, o próximo elemento deve ser o novo topo.\nEsperado: 2\nRecebido: %d",
			top,
		)
	}
}

func TestPopOnEmptyStackReturnsZeroValue(t *testing.T) {
	testStack := stack.New[int]()

	value := testStack.Pop()
	if value != 0 {
		t.Errorf(
			"Pop em stack vazia deve retornar o zero value do tipo.\nEsperado: 0\nRecebido: %d",
			value,
		)
	}
}

func TestPeekDoesNotRemoveValue(t *testing.T) {
	testStack := stack.New[int]()

	testStack.Push(10)

	top := testStack.Peek()
	if top != 10 {
		t.Errorf(
			"Peek deve retornar o topo sem removê-lo.\nEsperado: 10\nRecebido: %d",
			top,
		)
	}

	if testStack.Len() != 1 {
		t.Errorf(
			"Peek não deve alterar o tamanho da stack.\nEsperado: 1\nRecebido: %d",
			testStack.Len(),
		)
	}
}

func TestStackLength(t *testing.T) {
	testStack := stack.New[int]()

	if testStack.Len() != 0 {
		t.Errorf(
			"Stack recém-criada deve ter tamanho 0.\nRecebido: %d",
			testStack.Len(),
		)
	}

	testStack.Push(1)
	testStack.Push(2)
	testStack.Push(3)

	if testStack.Len() != 3 {
		t.Errorf(
			"O tamanho da stack deve crescer linearmente.\nEsperado: 3\nRecebido: %d",
			testStack.Len(),
		)
	}

	testStack.Pop()

	if testStack.Len() != 2 {
		t.Errorf(
			"O tamanho da stack deve diminuir após Pop.\nEsperado: 2\nRecebido: %d",
			testStack.Len(),
		)
	}
}
