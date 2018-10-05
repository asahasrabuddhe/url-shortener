package datastore

import (
	"errors"
	"github.com/asahasrabuddhe/url-shortener/url/helpers"
	"go.etcd.io/bbolt"
)

type Datastore interface {
	Set(key string, value string) error
	Get(key string) string
	Len() int
	Close()
}

type Database struct {
	db *bbolt.DB
}

var DB *Database

func openDatabase(file string) *bbolt.DB {
	db, err := bbolt.Open(file, 0600, nil)
	helpers.HandleError(err)

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return err
		}
		return nil
	})
	helpers.HandleError(err)

	return db
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) Set(key string, value string) error {
	return d.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		if b == nil {
			return errors.New("bucket not found")
		}

		k := []byte(key)
		v := []byte(value)

		return b.Put(k, v)
	})
}

func (d *Database) Get(key string) (value string) {
	d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		if b == nil {
			return errors.New("bucket not found")
		}

		k := []byte(key)

		v := b.Get(k)
		if v == nil {
			return errors.New("invalid key")
		}

		value = string(v)
		return nil
	})
	return
}

func (d *Database) Len() (num int) {
	d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		if b == nil {
			return errors.New("bucket not found")
		}

		b.ForEach(func(k, v []byte) error {
			num++
			return nil
		})

		return nil
	})
	return
}

func NewDatabase(file string) *Database {
	return &Database{
		db: openDatabase(file),
	}
}
