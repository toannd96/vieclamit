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
func (rp *RepoImpl) FindByUrl(urlJob string, collection string) (int, error) {
	count, err := rp.mg.Db.C(collection).Find(bson.M{"url_job": urlJob}).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
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
