package db

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func (i *Image) Create() error {
	u := uuid.NewV4()
	if err != nil {
		return err
	}
	sqlString := fmt.Sprintf(`INSERT INTO images (id, cid) VALUES ('%s','%s')`, u.String(), i.Cid)
	prep, err := db.Prepare(sqlString)
	if err != nil {
		return err
	}
	_, err = prep.Exec()
	if err != nil {
		return err
	}
	i.Id = u.String()
	return nil
}
func (i *Image) GetById() error {
	sqlString := fmt.Sprintf(`SELECT cid FROM images WHERE id='%s' LIMIT 1`, i.Id)
	prep, err := db.Prepare(sqlString)
	if err != nil {
		return err
	}
	row, err := prep.Query()
	if err != nil {
		return err
	}
	row.Next()
	err = row.Scan(&i.Cid)
	if err != nil {
		return err
	}
	return nil
}
