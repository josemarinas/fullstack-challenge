package db

func (i *Image) Create() error {
	var err error
	if err != nil {
		return err
	}
	sqlString := `INSERT INTO images (cid) VALUES ($1) returning id`
	prep, err := db.Prepare(sqlString)
	if err != nil {
		return err
	}
	defer prep.Close()
	err = prep.QueryRow(i.Cid).Scan(&i.Id)
	if err != nil {
		return err
	}
	return nil
}
func (i *Image) GetById() error {
	var err error
	sqlString := `SELECT cid FROM images WHERE id=$1 LIMIT 1`
	prep, err := db.Prepare(sqlString)
	if err != nil {
		return err
	}
	defer prep.Close()
	row, err := prep.Query(i.Id)
	if !row.Next() {
		return err
	}

	err = row.Scan(&i.Cid)
	if err != nil {
		return err
	}
	return nil
}
