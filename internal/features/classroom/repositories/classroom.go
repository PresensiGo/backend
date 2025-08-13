package repositories

import (
	"api/internal/features/classroom/domains"
	"api/internal/features/classroom/models"
	"gorm.io/gorm"
)

type Classroom struct {
	db *gorm.DB
}

func NewClassroom(db *gorm.DB) *Classroom {
	return &Classroom{db}
}

func (r *Classroom) GetOrCreateInTx(tx *gorm.DB, data domains.Classroom) (
	*domains.Classroom, error,
) {
	var classroom models.Classroom
	if err := tx.FirstOrCreate(&classroom, data.ToModel()).Error; err != nil {
		return nil, err
	} else {
		return domains.FromClassroomModel(&classroom), nil
	}
}

func (r *Classroom) Create(data domains.Classroom) (*domains.Classroom, error) {
	classroom := data.ToModel()
	if err := r.db.Create(&classroom).Error; err != nil {
		return nil, err
	}

	return domains.FromClassroomModel(classroom), nil
}

func (r *Classroom) CreateInTx(tx *gorm.DB, data domains.Classroom) (*uint, error) {
	classroom := models.Classroom{
		Name:    data.Name,
		MajorId: data.MajorId,
	}
	if err := tx.Create(&classroom).Error; err != nil {
		return nil, err
	}

	return &classroom.ID, nil
}

func (r *Classroom) CreateInTx2(tx *gorm.DB, data domains.Classroom) (*domains.Classroom, error) {
	classroom := data.ToModel()
	if err := tx.Create(&classroom).Error; err != nil {
		return nil, err
	} else {
		return domains.FromClassroomModel(classroom), nil
	}
}

func (r *Classroom) GetAllByMajorId(majorId uint) (*[]domains.Classroom, error) {
	var classrooms []models.Classroom
	if err := r.db.Where(
		"major_id = ?", majorId,
	).Order("lower(name) asc").Find(&classrooms).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Classroom, len(classrooms))
		for i, classroom := range classrooms {
			result[i] = *domains.FromClassroomModel(&classroom)
		}

		return &result, nil
	}
}

func (r *Classroom) GetManyByMajorId(majorIds []uint) ([]domains.Classroom, error) {
	var classes []models.Classroom
	if err := r.db.Where("major_id in ?", majorIds).
		Order("name asc").
		Find(&classes).
		Error; err != nil {
		return nil, err
	}

	var mappedClasses []domains.Classroom
	for _, class := range classes {
		mappedClasses = append(
			mappedClasses, domains.Classroom{
				Id:      class.ID,
				Name:    class.Name,
				MajorId: class.MajorId,
			},
		)
	}

	return mappedClasses, nil
}

func (r *Classroom) GetManyByIds(classroomIds []uint) (*[]domains.Classroom, error) {
	var classrooms []models.Classroom
	if err := r.db.Model(&models.Classroom{}).
		Where("id in (?)", classroomIds).
		Find(&classrooms).
		Error; err != nil {
		return nil, err
	}

	mappedClassrooms := make([]domains.Classroom, len(classrooms))
	for i, item := range classrooms {
		mappedClassrooms[i] = *domains.FromClassroomModel(&item)
	}

	return &mappedClassrooms, nil
}

func (r *Classroom) Update(classroomId uint, data domains.Classroom) (*domains.Classroom, error) {
	classroom := data.ToModel()
	if err := r.db.Where("id = ?", classroomId).Updates(classroom).Error; err != nil {
		return nil, err
	}

	return domains.FromClassroomModel(classroom), nil
}
