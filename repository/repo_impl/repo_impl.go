package repoimpl

import (
	"context"
	"fmt"
	"os"
	"time"

	"vieclamit/common"
	"vieclamit/database"
	"vieclamit/models"
	"vieclamit/repository"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = os.Getenv("COLLECTION")

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
func (rp *RepoImpl) Insert(recruitment models.Recruitment) error {
	_, err := rp.mg.Db.Collection(collection).InsertOne(context.TODO(), recruitment)
	if err != nil {
		return err
	}
	return nil
}

// FindByUrl find url job to check exists
func (rp *RepoImpl) FindByUrl(urlJob string) (int64, error) {
	count, err := rp.mg.Db.Collection(collection).CountDocuments(context.TODO(), bson.M{"url_job": urlJob})
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByLocation find location
func (rp *RepoImpl) FindByLocation(location string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	cursor, err := rp.mg.Db.Collection(collection).Find(context.TODO(), bson.M{"location": bson.M{"$regex": location, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &recruitments); err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByTitle find skill
func (rp *RepoImpl) FindBySkill(skill string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	cursor, err := rp.mg.Db.Collection(collection).Find(context.TODO(), bson.M{"title": bson.M{"$regex": skill, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &recruitments); err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByTitle find company
func (rp *RepoImpl) FindByCompany(company string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	cursor, err := rp.mg.Db.Collection(collection).Find(context.TODO(), bson.M{"company": bson.M{"$regex": company, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &recruitments); err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindBySkillAndLocation combine find skill and location
func (rp *RepoImpl) FindBySkillAndLocation(skill, location string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	conditions := bson.M{"$and": []bson.M{
		{"title": bson.M{"$regex": skill, "$options": "i"}},
		{"location": bson.M{"$regex": location, "$options": "i"}},
	}}
	cursor, err := rp.mg.Db.Collection(collection).Find(context.TODO(), conditions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &recruitments); err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// FindByCompanyAndLocation combine find company and location
func (rp *RepoImpl) FindByCompanyAndLocation(company, location string) (*models.Recruitments, error) {
	var recruitments models.Recruitments
	conditions := bson.M{"$and": []bson.M{
		{"company": bson.M{"$regex": company, "$options": "i"}},
		{"location": bson.M{"$regex": location, "$options": "i"}},
	}}
	cursor, err := rp.mg.Db.Collection(collection).Find(context.TODO(), conditions)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &recruitments); err != nil {
		return nil, err
	}
	return &recruitments, nil
}

// Delete delete document if expired job deadline
func (rp *RepoImpl) Delete() (int64, error) {
	timeToday, err := common.ParseTime(time.Now().Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	info, errRm := rp.mg.Db.Collection(collection).DeleteMany(context.TODO(), bson.M{"job_deadline": bson.M{"$lt": timeToday}})
	if errRm != nil {
		return 0, errRm
	}
	return info.DeletedCount, nil
}
