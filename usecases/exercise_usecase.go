package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

type ExerciseUseCase struct {
	exerciseRepo repository.ExerciseRepository
}

func NewExerciseUseCase(exerciseRepo repository.ExerciseRepository) *ExerciseUseCase {
	return &ExerciseUseCase{
		exerciseRepo: exerciseRepo,
	}
}

func (uc *ExerciseUseCase) Create(exercise *entity.Exercise) error {
	return uc.exerciseRepo.Create(exercise)
}
func (uc *ExerciseUseCase) GetByID(id uint) (*entity.Exercise, error) {
	return uc.exerciseRepo.GetByID(id)
}
func (uc *ExerciseUseCase) GetAll() ([]*entity.Exercise, error) {
	return uc.exerciseRepo.GetAll()
}
func (uc *ExerciseUseCase) Update(exercise *entity.Exercise) error {
	return uc.exerciseRepo.Update(exercise)
}
func (uc *ExerciseUseCase) Delete(id uint) error {
	return uc.exerciseRepo.Delete(id)
}
func (uc *ExerciseUseCase) GetByGroupID(groupID uint) ([]*entity.Exercise, error) {
	return uc.exerciseRepo.GetByGroupID(groupID)
}
func (uc *ExerciseUseCase) GetByTrackID(trackID uint) ([]*entity.Exercise, error) {
	return uc.exerciseRepo.GetByTrackID(trackID)
}
func (uc *ExerciseUseCase) GetByProblemID(problemID uint) ([]*entity.Exercise, error) {
	return uc.exerciseRepo.GetByProblemID(problemID)
}