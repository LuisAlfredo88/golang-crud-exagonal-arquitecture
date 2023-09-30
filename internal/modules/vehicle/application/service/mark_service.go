package service

import (
	"errors"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
	markRepo "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/repository"
	markService "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/service"
)

type mark struct {
	markRepository markRepo.MarkRepository
}

func NewMarkService(markR markRepo.MarkRepository) markService.MarkService {
	return &mark{
		markRepository: markR,
	}
}

func (p *mark) DeleteMark(markId int32) error {
	if markId <= 0 {
		return errors.New("must specify vehcile id")
	}

	logError := p.markRepository.DeleteMark(markId)
	return logError
}

func (l *mark) GetMark(markId int32) (entity.Mark, error) {
	return l.markRepository.GetMark(markId)
}

func (l *mark) GetAllMarks() ([]entity.Mark, error) {
	return l.markRepository.GetAllMarks()
}

func (p *mark) SaveMark(markEntity entity.Mark) error {
	err := error(nil)

	if err != nil {
		return err
	}

	logError := p.markRepository.SaveMark(markEntity)
	return logError
}

func (p *mark) DeleteModel(modelId int32) error {
	if modelId <= 0 {
		return errors.New("must specify vehcile id")
	}

	logError := p.markRepository.DeleteModel(modelId)
	return logError
}

func (l *mark) GetModel(modelId int32) (entity.VehicleModel, error) {
	return l.markRepository.GetModel(modelId)
}

func (l *mark) GetAllModels() ([]entity.VehicleModel, error) {
	return l.markRepository.GetAllModels()
}

func (p *mark) SaveModel(modelEntity entity.VehicleModel) error {
	err := error(nil)

	if err != nil {
		return err
	}

	logError := p.markRepository.SaveModel(modelEntity)
	return logError
}
