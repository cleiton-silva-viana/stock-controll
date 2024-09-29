package unitofwork

import (
	"database/sql"
	"fmt"
)

type IUnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
}

type UnitOfWork struct {
	db *sql.DB
	tx *sql.Tx
	rollbackAttempt bool
}

func NewUnitOfWork(DB *sql.DB) UnitOfWork {
	return UnitOfWork{db: DB}
}

func (uow *UnitOfWork) Begin() error {
	if uow.tx != nil {
		return transactionAlreadyInProgress
	}
	tx, err := uow.db.Begin()
	if err != nil {
		return &UnitOfworkError{
			Operation: "begin",
			Message:   err.Error(),
		}
	}
	uow.tx = tx
	return nil
}

func (uow *UnitOfWork) Commit() error {
	if uow.rollbackAttempt {
		return CommitNotAllowed
	}

	if uow.tx == nil {
		return noTransactionInProgress("commit")
	}
	err := uow.tx.Commit()
	uow.tx = nil
	return err
}

func (uow *UnitOfWork) Rollback() error {
	if uow.tx == nil {
		return noTransactionInProgress("rollback")
	}
	err := uow.tx.Rollback()
	uow.tx = nil
	return err
}

type UnitOfworkError struct {
	Operation string
	Message   string
	Tip       string
}

func (u *UnitOfworkError) Error() string {
	return fmt.Sprintf("Error for operation %s\nMessage:%s\nTip:%s", u.Operation, u.Message, u.Tip)
}

var transactionAlreadyInProgress = &UnitOfworkError{
	Operation: "Begin",
	Message:   "transaction already in progess",
	Tip:       "finalizer of the ongoing transaction before starting a new transaction",
}

var CommitNotAllowed = &UnitOfworkError{
	Operation: "Commit",
	Message:   "The commit cannot be performed, the rollback command has already been used in this transaction",
	Tip:       "finalizer of the ongoing transaction before starting a new transaction",
}

var noTransactionInProgress = func(transaction string) *UnitOfworkError {
	return &UnitOfworkError{
		Operation: transaction,
		Message:   fmt.Sprintf("no transactions to %s", transaction),
		Tip:       "check if transactions are recorded in the sql.Tx attribute",
	}
}
