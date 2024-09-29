package unitofwork

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type UnitOfWorkTest struct {
	description   string
	query         string
	wantError     bool
	expectedError error
}

func Test_UnitOfWork_Begin(t *testing.T) {
	tests := []UnitOfWorkTest{
		{
			description: "Valid operation - init begin",
			query:       "",
			wantError:   false,
		},
		{
			description:   "Invalid operation - have operation in progress",
			query:         "SELECT (fist_name, last_name, gender) FROM users WHERE ID = ?",
			wantError:     true,
			expectedError: transactionAlreadyInProgress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Arrange
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			mock.ExpectBegin()
			uow := NewUnitOfWork(db)

			// Act
			if tt.query != "" {
				uow.tx = &sql.Tx{}
			}
			err = uow.Begin()

			// Assert
			if tt.wantError {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_UnitOfWork_Commit(t *testing.T) {
	tests := []UnitOfWorkTest{
		{
			description: "Valid commit operation - realized successfully ",
			query:       "SELECT * FROM users",
			wantError:   false,
		},
		{
			description:   "Invalid commit operation - no transaction for commit",
			wantError:     true,
			expectedError: noTransactionInProgress("commit"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Arrange
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			mock.ExpectBegin()
			mock.ExpectCommit()
			uow := NewUnitOfWork(db)
			uow.Begin()

			if tt.wantError {
				uow.tx = nil
			} else {
				uow.tx.Exec(tt.query)
			}

			// Act
			err = uow.Commit()

			// Assert
			if tt.wantError {
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_UnitOfWork_Rollback(t *testing.T) {
	// Arrange
	tests := []UnitOfWorkTest{
		{
			description: "Valid rollback operation",
			query: "SELECT * FROM address LIMIT 10",
			wantError: false,
		},
		{
			description: "Invalid rollback operation - no transactions for rollback",
			wantError: true,
			expectedError: noTransactionInProgress("rollback"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			mock.ExpectBegin()
			mock.ExpectRollback()
			uow := NewUnitOfWork(db)
			uow.Begin()
			
			// Act
			if tt.wantError {
				uow.tx = nil
			}
			err = uow.Rollback()

			// Assert
			if tt.wantError {
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	} 
}
