package moviecontroller

import (
	"log"

	mv "github.com/silastgoes/Golang-apirest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesControler struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesControler) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesControler) GetAll() ([]mv.Movie, error) {
	var movies []mv.Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesControler) GetById(ID string) (mv.Movie, error) {
	var movie mv.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(ID)).One(&movie)
	return movie, err
}

func (m *MoviesControler) Create(movie mv.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesControler) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *MoviesControler) Update(id string, movie mv.Movie) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)
	return err
}
