package review

import "Consure-App/domain"

type ReviewRepository interface {
	FindByIdExpert(int, *[]*domain.Review) error
	FindByIdUser(int, *[]*domain.Review) error
	CountAverage(id int) float64
}
