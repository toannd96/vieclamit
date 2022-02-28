package repoimpl

import (
	"fmt"
	"time"

	"vieclamit/common"
	"vieclamit/database"
	"vieclamit/models"
	"vieclamit/repository"

	"gopkg.in/mgo.v2/bson"
)

// RepoImpl struct
type RepoImpl struct {
	mg *database.Mongo
}

// NewRepo new repository mongo
func NewRepo(mg *database.Mongo) repository.Repository {
	return &RepoImpl{
		mg: mg,
	}
}

// Insert insert data recruitment in to mongo
func (rp *RepoImpl) Insert(recruitment models.Recruitment, collection string) error {
	return rp.mg.Db.C(collection).Insert(recruitment)
}

// FindByUrl find url job to check exists
func (rp *RepoImpl) FindByUrl(urlJob, collection string) (int, error) {
	count, err := rp.mg.Db.C(collection).Find(bson.M{"url_job": urlJob}).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByLocation find location
func (rp *RepoImpl) FindByLocation(location, collection string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	err := rp.mg.Db.C(collection).Find(bson.M{"location": bson.M{"$regex": location, "$options": "i"}}).All(&recruitments)
	if err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByTitle find skill
func (rp *RepoImpl) FindBySkill(skill, collection string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	err := rp.mg.Db.C(collection).Find(bson.M{"title": bson.M{"$regex": skill, "$options": "i"}}).All(&recruitments)
	if err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByTitle find company
func (rp *RepoImpl) FindByCompany(company, collection string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	err := rp.mg.Db.C(collection).Find(bson.M{"company": bson.M{"$regex": company, "$options": "i"}}).All(&recruitments)
	if err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindBySkillAndLocation combine find skill and location
func (rp *RepoImpl) FindBySkillAndLocation(skill, location, collection string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	conditions := bson.M{"$and": []bson.M{
		{"title": bson.M{"$regex": skill, "$options": "i"}},
		{"location": bson.M{"$regex": location, "$options": "i"}},
	}}
	err := rp.mg.Db.C(collection).Find(conditions).All(&recruitments)
	if err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByCompanyAndLocation combine find company and location
func (rp *RepoImpl) FindByCompanyAndLocation(company, location, collection string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	conditions := bson.M{"$and": []bson.M{
		{"company": bson.M{"$regex": company, "$options": "i"}},
		{"location": bson.M{"$regex": location, "$options": "i"}},
	}}
	err := rp.mg.Db.C(collection).Find(conditions).All(&recruitments)
	if err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// Delete delete document if expired job deadline
func (rp *RepoImpl) Delete(collection string) (int, error) {
	timeToday, err := common.ParseTime(time.Now().Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	info, errRm := rp.mg.Db.C(collection).RemoveAll(bson.M{"job_deadline": bson.M{"$lt": timeToday}})
	if errRm != nil {
		return 0, errRm
	}
	return info.Removed, nil
}
