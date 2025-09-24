package DAO

import (
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// 用户基本操作
func (dao *UserDAO) CreateUser(user *User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAO) GetUserByEmail(email string) (*User, error) {
	var user User
	err := dao.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (dao *UserDAO) GetUserByID(id string) (*User, error) {
	var user User
	err := dao.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (dao *UserDAO) UpdateUser(user *User) error {
	return dao.db.Save(user).Error
}

// 认证相关操作
func (dao *UserDAO) CreateCertification(cert *Certification) error {
	return dao.db.Create(cert).Error
}

func (dao *UserDAO) GetCertificationByUserID(userID string) (*Certification, error) {
	var cert Certification
	err := dao.db.Where("user_id = ?", userID).First(&cert).Error
	return &cert, err
}

func (dao *UserDAO) UpdateCertification(cert *Certification) error {
	return dao.db.Save(cert).Error
}

// AI陪伴相关操作
func (dao *UserDAO) CreateAICompanion(companion *AICompanion) error {
	return dao.db.Create(companion).Error
}

func (dao *UserDAO) GetAICompanionsByUserID(userID string) ([]AICompanion, error) {
	var companions []AICompanion
	err := dao.db.Where("user_id = ?", userID).Find(&companions).Error
	return companions, err
}

// 虚拟疗愈导师相关操作
func (dao *UserDAO) CreateVirtualTherapist(therapist *VirtualTherapist) error {
	return dao.db.Create(therapist).Error
}

func (dao *UserDAO) GetVirtualTherapistsByUserID(userID string) ([]VirtualTherapist, error) {
	var therapists []VirtualTherapist
	err := dao.db.Where("user_id = ?", userID).Find(&therapists).Error
	return therapists, err
}

// 儿童档案相关操作
func (dao *UserDAO) CreateChildArchive(archive *ChildArchive) error {
	return dao.db.Create(archive).Error
}

func (dao *UserDAO) GetChildArchivesByUserID(userID string) ([]ChildArchive, error) {
	var archives []ChildArchive
	err := dao.db.Where("user_id = ?", userID).Find(&archives).Error
	return archives, err
}

func (dao *UserDAO) UpdateChildArchive(archive *ChildArchive) error {
	return dao.db.Save(archive).Error
}

func (dao *UserDAO) DeleteChildArchive(archiveID string) error {
	return dao.db.Delete(&ChildArchive{}, "id = ?", archiveID).Error
}

// 收藏相关操作
func (dao *UserDAO) AddFavorite(favorite *UserFavorite) error {
	return dao.db.Create(favorite).Error
}

func (dao *UserDAO) GetFavoritesByUserID(userID string) ([]UserFavorite, error) {
	var favorites []UserFavorite
	err := dao.db.Where("user_id = ?", userID).Find(&favorites).Error
	return favorites, err
}

func (dao *UserDAO) RemoveFavorite(userID, resourceType, resourceID string) error {
	return dao.db.Where("user_id = ? AND resource_type = ? AND resource_id = ?", 
		userID, resourceType, resourceID).Delete(&UserFavorite{}).Error
}

// 课程相关操作
func (dao *UserDAO) GetCourses() ([]Course, error) {
	var courses []Course
	err := dao.db.Where("status = ?", "published").Find(&courses).Error
	return courses, err
}

func (dao *UserDAO) GetCourseByID(courseID string) (*Course, error) {
	var course Course
	err := dao.db.Where("id = ?", courseID).First(&course).Error
	return &course, err
}

// 游戏相关操作
func (dao *UserDAO) GetGames() ([]Game, error) {
	var games []Game
	err := dao.db.Where("status = ?", "published").Find(&games).Error
	return games, err
}

func (dao *UserDAO) GetGameByID(gameID string) (*Game, error) {
	var game Game
	err := dao.db.Where("id = ?", gameID).First(&game).Error
	return &game, err
}