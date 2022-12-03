package expert

import "Consure-App/domain"

type ExpertRepository interface {
	FindByTag(string, *[]*domain.Expert) error
}
