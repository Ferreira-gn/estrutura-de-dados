package stack_test

import (
	"testing"

	"github.com/ferreira-gn/estrutura-de-dados/estruturas/stack"
)

func TestFromCreateNewStack(t *testing.T) {
	testStack := stack.StackConstructor()

	if testStack == nil {
		t.Errorf("Ao criar uma stack usando o construtor deve ser retorenado um ponteiro para a Stack.\nFoi retornado %v ", testStack)
	}
}

func TestFromAddNewValueInStack(t *testing.T) {
	testStack := stack.StackConstructor()

	testStack.Push(2)
	testStack.Push(3)
	testStack.Push(4)

	top := testStack.ViewTheTop()

	if top != 4 {
		t.Errorf("Ao adicionar um valor na pilha esse valor deve ser armazenado no topo da pilha.\nValor esperado : 4 \nValor recebido : %v ", top)
	}
}

func TestFromRemoveValueInStack(t *testing.T) {
	testStack := stack.StackConstructor()

	testStack.Push(2)
	testStack.Push(3)
	testStack.Pop()
	top := testStack.ViewTheTop()

	if top != 2 {
		t.Errorf("Ao remover um valor do topo da pilha o proximo nó deve ser alocado no topo da pilha\nValor esperado : 2 \nValor recebido : %v ", top)
	}

	testStack.Pop()
	top = testStack.ViewTheTop()

	if top != 0 {
		t.Errorf("Deve ser possível remover todos os valores de uma pilha.\nValor esperado : 0 \nValor recebido : %v ", top)
	}
}

func TestFromViewTheTopStack(t *testing.T) {
	testStack := stack.StackConstructor()
	top := testStack.ViewTheTop()

	if top != 0 {
		t.Errorf("Ao inicializar uma pilha o valor do topo deve ser sempre nil ou zero.\nValor esperado : 0 \nValor recebido : %v ", top)
	}

	testStack.Push(3)
	top = testStack.ViewTheTop()

	if top != 3 {
		t.Errorf("Ao visualizar o topo de uma pilha deve ser retornado o ultimo valor adicionado na pilha.\nValor esperado : 3 \nValor recebido : %v ", top)
	}
}

func TestFromCaptureTheLengthOfStack (t *testing.T){
	testStack := stack.StackConstructor()
	length := testStack.Len()

	if length != 0 {
		t.Errorf("Ao inicializar uma pilha o seu tamanho deve ser sempre zero.\nValor esperado : 0 \nValor recebido : %v ", length)
	}

	testStack.Push(3)
	testStack.Push(4)
	testStack.Push(5)
	testStack.Push(6)
	length = testStack.Len()

	if length != 4 {
		t.Errorf("Ao adicionar um novo valor em uma pilha o seu tamanho deve continuar aumentando linearmente.\nValor esperado : 4 \nValor recebido : %v ", length)
	}
}