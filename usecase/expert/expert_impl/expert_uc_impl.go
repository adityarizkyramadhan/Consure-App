package expertimpl

import (
	"Consure-App/domain"
	"Consure-App/repository/expert"
	"Consure-App/repository/general"
	"Consure-App/sdk/upload"
	expertUc "Consure-App/usecase/expert"
)

type ExpertUseCase struct {
	RepoGeneral general.GeneralRepository
	RepoExpert  expert.ExpertRepository
}

func NewExpertUsecase(repoGeneral general.GeneralRepository, repoExpert expert.ExpertRepository) expertUc.ExpertUsecase {
	return &ExpertUseCase{
		RepoGeneral: repoGeneral,
		RepoExpert:  repoExpert,
	}
}

func (ec *ExpertUseCase) SignUp(input *expertUc.InputExpert) error {
	link, err := upload.UploadImage(input.Avatar)
	if err != nil {
		return err
	}
	data := &domain.Expert{
		Email:      input.Email,
		Nama:       input.Nama,
		Price:      input.Price,
		Experience: input.Experience,
		Education:  input.Education,
		Tag:        input.Tag,
		LinkImage:  link,
	}
	err = ec.RepoGeneral.Create(data)
	return err
}

func (ec *ExpertUseCase) FindAll(data *[]*domain.Expert) error {
	return ec.RepoGeneral.FindAll(&data)
}

func (ec *ExpertUseCase) FindById(id int, data interface{}) error {
	return ec.RepoGeneral.FindById(id, data)
}
