package main

import (
	"fmt"
)

type ContaBase struct {
	Nome  string
	Saldo float64
}

type ContaPoupanca struct {
	ContaBase
}

type ContaCorrente struct {
	ContaBase
}

func NewContaCorrente(nome string, saldo float64) *ContaCorrente {
	return &ContaCorrente{ContaBase{Nome: nome, Saldo: saldo}}
}

func NewContaPoupanca(nome string, saldo float64) *ContaPoupanca {
	return &ContaPoupanca{ContaBase{Nome: nome, Saldo: saldo}}
}

func NewTransactionError(message string) *TransactionError {
	return &TransactionError{message: message}
}

type TransactionError struct {
	message string
}

func (t *TransactionError) Error() string {
	return "Erro ao tentar transacionar, " + t.message
}

func (c *ContaBase) sacar(valor float64) error {
	if valor < 0 {
		return NewTransactionError("valor inválido!")
	}
	if c.Saldo < valor {
		return NewTransactionError("Saldo insuficiente!")
	}
	c.Saldo -= valor
	return nil
}

func (c *ContaBase) depositar(valor float64) error {
	if valor < 0 {
		return NewTransactionError("valor inválido!")
	}
	c.Saldo += valor
	return nil
}

func (c *ContaBase) transferir(valor float64, contaDestino IContaBase) error {
	err := c.sacar(valor)
	if err == nil {
		return err
	}
	contaDestino.depositar(valor)
	if err == nil {
		return err
	}
	return nil
}

func (c *ContaCorrente) emprestimo(valor float64) error {
	if valor < 0 {
		return NewTransactionError("valor inválido!")
	}
	c.Saldo += valor
	return nil
}

func (c *ContaPoupanca) poupar(valor float64) error {
	if valor < 0 {
		return NewTransactionError("valor inválido!")
	}
	c.Saldo += valor
	return nil
}

type IContaBase interface {
	sacar(valor float64) error
	depositar(valor float64) error
	transferir(valor float64, contaDestino IContaBase) error
}

type IContaPoupanca interface {
	poupar(valor float64) error
}

type IContaCorrente interface {
	emprestimo(valor float64) error
}

func Depositar(conta IContaBase, valor float64) {
	conta.depositar(valor)
}

func Sacar(conta IContaBase, valor float64) {
	conta.sacar(valor)
}

func Transferir(contaOrigem IContaBase, contaDestino IContaBase, valor float64) {
	contaOrigem.transferir(valor, contaDestino)
}

func Poupar(conta IContaPoupanca, valor float64) {
	conta.poupar(valor)
}

func Emprestimo(conta IContaCorrente, valor float64) {
	conta.emprestimo(valor)
}

func main() {
	conta1 := NewContaCorrente("berilo queiroz", 20.0)
	conta2 := NewContaPoupanca("richelly sousa", 10.0)
	Depositar(conta1, 12.5)
	fmt.Println("Saldo conta 1", conta1.Saldo)
	Sacar(conta1, 12.25)
	fmt.Println("Saldo conta 1", conta1.Saldo)
	fmt.Println("Saldo conta 2", conta2.Saldo)
	Transferir(conta1, conta2, 50.0)
	fmt.Println("Saldo conta 1", conta1.Saldo)
	fmt.Println("Saldo conta 2", conta2.Saldo)
	Poupar(conta2, 10.5)
	Emprestimo(conta1, 11)
	fmt.Println("Saldo conta 1", conta1.Saldo)
	fmt.Println("Saldo conta 2", conta2.Saldo)
}
