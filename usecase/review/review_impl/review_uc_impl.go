package reviewimpl

import (
	"Consure-App/domain"
	"Consure-App/dto"
	"Consure-App/repository/general"
	repoRev "Consure-App/repository/review"
	ucRev "Consure-App/usecase/review"
)

type ReviewUseCase struct {
	RepoGeneral general.GeneralRepository
	RepoReview  repoRev.ReviewRepository
}

func NewReviewUsecase(repoGeneral general.GeneralRepository, repoReview repoRev.ReviewRepository) ucRev.ReviewUsecase {
	return &ReviewUseCase{
		RepoGeneral: repoGeneral,
		RepoReview:  repoReview,
	}
}

func (ec *ReviewUseCase) FindAll(data *[]*domain.Review) error {
	return ec.RepoGeneral.FindAll(&data)
}

func (ec *ReviewUseCase) FindById(id int, data interface{}) error {
	return ec.RepoGeneral.FindById(id, data)
}

func (ec *ReviewUseCase) FindByIdExpert(id int, data *dto.DataExpertWithReview) error {
	var review []*domain.Review

	if err := ec.RepoReview.FindByIdExpert(id, &review); err != nil {
		return err
	}

	expert := new(domain.Expert)

	if err := ec.RepoGeneral.FindById(id, expert); err != nil {
		return err
	}

	data.Reviews = review
	data.Expert = expert
	avg := ec.RepoReview.CountAverage(id)

	data.AverageStar = avg

	return nil
}
func (ec *ReviewUseCase) FindByIdUser(id int, data *[]*domain.Review) error {
	return ec.RepoReview.FindByIdUser(id, data)
}

func (ec *ReviewUseCase) Create(idUser int, input *ucRev.InputReview) error {
	data := &domain.Review{
		Star:     input.Star,
		Komentar: input.Komentar,
		IdExpert: input.IdExpert,
		IdUser:   idUser,
	}
	return ec.RepoGeneral.Create(data)
}
