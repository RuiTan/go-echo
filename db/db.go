package db

import "github.com/globalsign/mgo"

type Database struct {
	session *mgo.Session
	database string
}

var Global *Database;

const (
	CUser	= "user"
)

type GetCollection func() (collection *mgo.Collection, closeConn func())

func InitializeGlobalDB(url string, database string) error {
	d, err := NewDatabase(url, database)
	if err != nil {
		return err
	}
	err = d.EnsureIndex()
	if err != nil {
		return err
	}
	Global = d
	return nil
}

func NewDatabase(url string, database string) (*Database, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	d := &Database{
		session:  session,
		database: database,
	}
	return d, nil
}

func (d *Database) DB() (*mgo.Database, func())  {
	conn := d.session.Copy()
	return conn.DB(d.database), func() {
		conn.Close()
	}
}

func (d *Database) EnsureIndex() (err error)  {
	database, closeConn := d.DB()
	defer closeConn()
	err = database.C(CUser).EnsureIndex(mgo.Index{
		Key:              []string{"username"},
		Unique:           true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DropDatabase() (err error) {
	database, closeConn := d.DB()
	defer closeConn()
	return database.DropDatabase()
}

func (d *Database) User() (collection *mgo.Collection, closeConn func()) {
	database, closeConn := d.DB()
	collection = database.C(CUser)
	return collection, closeConn
}

func (d *Database) Collection(collectionName string) GetCollection {
	return func() (collection *mgo.Collection, closeConn func()) {
		database, closeConn := d.DB()
		collection = database.C(collectionName)
		return collection, closeConn
	}
}