package reviewimpl

import (
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
