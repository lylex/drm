package blobs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lylex/drm/pkg/files"
	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
)

const (
	MaxRandomNumber int = 10000

	MetaDataDB string = "data.db"

	CfgDataPathKey = "dataPath"
)

type Blob struct {
	FileName string
	Dir      string

	// CreatedAt represents when the blob is created, i.e. when it is deleted.
	CreatedAt time.Time
}

// Create make a blob metadate.
func Create(fullPath string) *Blob {
	return &Blob{
		FileName:  files.Name(fullPath),
		Dir:       files.Dir(fullPath),
		CreatedAt: time.Now(),
	}
}

func (b *Blob) Name() string {
	return fmt.Sprintf("%d_%4d_%s",
		b.CreatedAt.UnixNano(),
		rand.Intn(MaxRandomNumber),
		b.FileName)
}

// TODO
func (b *Blob) Save() error {
	db, err := buntdb.Open(viper.GetString(CfgDataPathKey) + MetaDataDB)
	if err != nil {
		return err
	}
	defer db.Close()

	var json string
	json, err = utils.Marshal(*b)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(b.FileName, json, nil)
		return err
	})

	return nil
}
