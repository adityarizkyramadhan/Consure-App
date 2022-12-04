package transaction

import "Consure-App/domain"

type TransactionRepository interface {
	History(id int, data *[]*domain.Transaction) error
}
