package mongo

import (
	"log"

	"github.com/globalsign/mgo"
)

//ImageRepository to manage the db access
type ImageRepository interface {
}

//ImageDao to operate with the image objects
type ImageDao struct {
	logger       *log.Logger
	mongoSession *mgo.Session
}

//AddNewImage POST
func AddNewImage(url string, logger *log.Logger) (*ImageDao, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	dao := ImageDao{mongoSession: session, logger: logger}
	return &dao, nil
}
