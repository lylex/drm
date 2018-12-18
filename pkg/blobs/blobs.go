package blobs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lylex/drm/pkg/files"
)

const (
	// MaxRandomNumber represents the random number range [0, MaxRandomNumber).
	MaxRandomNumber int = 10000

	// MetaDataDB represents the name of the file used for storing metadata.
	MetaDataDB string = "data.db"

	// CfgDataPathKey represents the dataPath key in config.
	CfgDataPathKey = "dataPath"
)

// Blob represents the blob instance.
type Blob struct {
	// FileName represents the file name of the deleted file.
	FileName string

	// Dir represent where the blob stored before deleted.
	Dir string

	// CreatedAt represents when the blob is created, i.e. when it is deleted.
	CreatedAt time.Time
}

// Create makes a blob metadate.
func Create(fullPath string) *Blob {
	return &Blob{
		FileName:  files.Name(fullPath),
		Dir:       files.Dir(fullPath),
		CreatedAt: time.Now(),
	}
}

// Name gets the name of the blob.
func (b *Blob) Name() string {
	return fmt.Sprintf("%d_%4d_%s",
		b.CreatedAt.UnixNano(),
		rand.Intn(MaxRandomNumber),
		b.FileName)
}

// Save stores the blob metadata to disk.
func (b *Blob) Save() error {
	var bg *blobGroup
	var err error
	if bg, err = getBlobGroup(b.FileName); err != nil {
		return err
	}
	if bg == nil {
		bg = &blobGroup{
			FileName: b.FileName,
		}
	}

	if err = bg.add(b); err != nil {
		return err
	}
	if err = bg.save(); err != nil {
		return err
	}

	return nil
}

// GetAll returns all the blobs in database.
func GetAll() ([]*Blob, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var bgs []*blobGroup
	if bgs, err = getAllBlobGroup(); err != nil {
		return nil, err
	}

	blobs := make([]*Blob, 0)
	for _, bg := range bgs {
		if bg.Count == 0 {
			continue
		}
		for _, b := range bg.Blobs {
			blobs = append(blobs, &b)
		}
	}

	return blobs, nil
}
