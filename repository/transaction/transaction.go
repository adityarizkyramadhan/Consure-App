package transaction

import "Consure-App/domain"

type TransactionRepository interface {
	History(id int, status string, data *[]*domain.Transaction) error
}
