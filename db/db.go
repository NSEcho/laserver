package db

import (
	"encoding/binary"
	"encoding/json"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type bucket []byte

var (
	DataBucket bucket = []byte("data")
)

var buckets = []bucket{
	DataBucket,
}

type DB struct {
	db *bolt.DB
}

func NewDB(dbPath string) *DB {
	db, err := bolt.Open(dbPath, os.ModePerm, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists(bucket)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}

func (d *DB) Close() {
	d.db.Close()
}

type Data struct {
	ID   int
	UUID string
	Time time.Time
}

func (d *DB) Save(data *Data) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(DataBucket)

		id, _ := b.NextSequence()
		data.ID = int(id)

		marshalled, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return b.Put(itob(data.ID), marshalled)
	})
}

func (d *DB) Exists(uuid string) (bool, error) {
	found := false
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(DataBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var dt Data
			if err := json.Unmarshal(v, &dt); err != nil {
				return err
			}
			if dt.UUID == uuid {
				found = true
				return nil
			}
		}
		return nil
	})
	return found, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (d *DB) List() ([]Data, error) {
	var data []Data
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(DataBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var dt Data
			if err := json.Unmarshal(v, &dt); err != nil {
				return err
			}
			data = append(data, dt)
		}
		return nil
	})
	return data, err
}
