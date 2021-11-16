package db

type Image struct {
	Id  string `form:"-" json:"id"`
	Cid string `form:"cid" json:"cid" binding:"required"`
}
