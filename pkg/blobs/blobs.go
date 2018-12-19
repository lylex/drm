package blobs

import (
	"errors"
	"fmt"
	"time"

	"github.com/lylex/drm/pkg/files"
	"github.com/lylex/drm/pkg/utils"
)

const (
	// MaxRandomNumber represents the random number range [0, MaxRandomNumber).
	MaxRandomNumber int = 10000

	// MetaDataDB represents the name of the file used for storing metadata.
	MetaDataDB string = "data.db"

	// CfgDataPathKey represents the dataPath key in config.
	CfgDataPathKey = "dataPath"
)

var (
	// ErrOperateDB is returned when unknow error hanppend from database.
	ErrOperateDB = errors.New("error happended when operating database")

	// ErrNotFound represents no record found.
	ErrNotFound = errors.New("not found")
)

// Blob represents the blob instance.
type Blob struct {
	// FileName represents the file name of the deleted file.
	FileName string

	// Dir represent where the blob stored before deleted.
	Dir string

	// CreatedAt represents when the blob is created, i.e. when it is deleted.
	CreatedAt time.Time

	// ID used to indentify blob when two or more files has the same filename.
	ID string
}

// Create makes a blob metadate.
func Create(fullPath string) *Blob {
	return &Blob{
		FileName:  files.Name(fullPath),
		Dir:       files.Dir(fullPath),
		CreatedAt: time.Now(),
		ID:        utils.GenerateRandString(6),
	}
}

// Name gets the name of the blob.
func (b *Blob) Name() string {
	return fmt.Sprintf("%d_%s_%s",
		b.CreatedAt.UnixNano(),
		b.ID,
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

// Destroy deletes a blob from database.
// Only Filename and ID are compared.
func (b *Blob) Destroy() error {
	bg, err := getBlobGroup(b.FileName)
	if err != nil {
		return ErrOperateDB
	}

	if bg == nil {
		return ErrNotFound
	}

	found := false
	for i, c := range bg.Blobs {
		if c.ID == b.ID {
			found = true
			bg.Blobs[i] = bg.Blobs[len(bg.Blobs)-1]
			bg.Blobs = bg.Blobs[:len(bg.Blobs)-1]
			bg.Count--
			break
		}
	}
	if !found {
		return ErrNotFound
	}

	if bg.Count == 0 {
		err = bg.destroy()
	} else {
		err = bg.save()
	}

	return err
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
