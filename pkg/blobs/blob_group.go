package blobs

import (
	"errors"

	"github.com/lylex/drm/pkg/utils"
	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
)

var (
	// ErrGroupUnmatch is returned when add a Blob into a blobgroup which it is unmatch.
	ErrGroupUnmatch = errors.New("group is unmatch")
)

// blobGroup represents a group of blobs with the same file name.
// It is the unit which stores in db.
type blobGroup struct {
	// fileName represents the filename of the file name and it is the key.
	FileName string

	// Count represents how many blob in total have the FileName.
	Count int

	// Blobs represents the blobs data.
	Blobs []Blob
}

func getBlobGroup(filename string) (*blobGroup, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var raw string
	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(filename)
		if err != nil {
			if err == buntdb.ErrNotFound {
				return nil
			}
			return err
		}
		raw = val
		return nil
	})

	if raw == "" {
		return nil, nil
	}

	var bg blobGroup
	if err = utils.Unmarshal(raw, &bg); err != nil {
		return nil, err
	}
	return &bg, nil
}

func getAllBlobGroup() ([]*blobGroup, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rawValues := make([]string, 0)
	err = db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			rawValues = append(rawValues, value)
			return true
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	bgs := make([]*blobGroup, 0)
	for _, value := range rawValues {
		var obj blobGroup
		if err := utils.Unmarshal(value, &obj); err != nil {
			return nil, err
		}
		bgs = append(bgs, &obj)
	}
	return bgs, nil
}

// save stores the data into db, do not care the count.
func (bg *blobGroup) save() error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var json string
	json, err = utils.Marshal(*bg)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(bg.FileName, json, nil)
		return err
	})

	return err
}

func (bg *blobGroup) add(b *Blob) error {
	if bg.FileName != b.FileName {
		return ErrGroupUnmatch
	}

	bg.Blobs = append(bg.Blobs, *b)
	bg.Count++
	return nil
}

func (bg *blobGroup) destroy() error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(bg.FileName)
		return err
	})

	return err
}

// openDB opens a DB, should be explicitly closed later.
func openDB() (*buntdb.DB, error) {
	return buntdb.Open(viper.GetString(CfgDataPathKey) + MetaDataDB)
}
