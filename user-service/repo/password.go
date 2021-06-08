package repo

import (
	"github.com/jinzhu/gorm"
	pb "github.com/zero-dora/laracom/user-service/proto/user"
)

type PasswordResetInterface interface {
	CreateReset(reset *pb.PasswordReset) error
	GetByToken(token string) (*pb.PasswordReset, error)
	DeleteReset(reset *pb.PasswordReset) error
	GetByEmail(email string) (*pb.PasswordReset, error)
}

type PasswordResetRepository struct {
	Db *gorm.DB
}

func (repo *PasswordResetRepository) CreateReset(reset *pb.PasswordReset) error {
	if err := repo.Db.Create(reset).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PasswordResetRepository) GetByToken(token string) (*pb.PasswordReset, error) {
	reset := &pb.PasswordReset{}
	if err := repo.Db.Where("token = ?", token).First(&reset).Error; err != nil {
		return nil, err
	}
	return reset, nil
}

func (repo *PasswordResetRepository) GetByEmail(email string) (*pb.PasswordReset, error) {
	reset := &pb.PasswordReset{}
	if err := repo.Db.Where("email = ?", email).First(&reset).Error; err != nil {
		return nil, err
	}
	return reset, nil
}

func (repo *PasswordResetRepository) DeleteReset(reset *pb.PasswordReset) error {
	err := repo.Db.Delete(reset).Error
	return err
}
