package api

import (
	"fmt"
	"net/http"
	"server/db"
	"server/httperrors"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-cid"
	uuid "github.com/satori/go.uuid"
)

func CreateImage(c *gin.Context) {
	var image db.Image
	// check form
	err = c.ShouldBindJSON(&image)
	if err != nil {
		c.AbortWithStatusJSON(httperrors.ImageInvalidForm.Status, httperrors.ImageInvalidForm.Body)
		return
	}
	// check if cid is valid
	_, err = cid.Decode(image.Cid)
	if err != nil {
		c.AbortWithStatusJSON(httperrors.ImageInvalidCid.Status, httperrors.ImageInvalidCid.Body)
		return
	}
	// create image
	err = image.Create()
	if err != nil {
		c.AbortWithStatusJSON(httperrors.ImageAlreadyExist.Status, httperrors.ImageAlreadyExist.Body)
		return
	}
	// if no error return image
	c.JSON(http.StatusOK, image)
}

func GetImageByID(c *gin.Context) {
	// get id param and check if is a valid uuid
	id, _ := c.Params.Get("id")
	_, err = uuid.FromString(id)
	if err != nil {
		c.AbortWithStatusJSON(httperrors.ImageInvalidUUID.Status, httperrors.ImageInvalidUUID.Body)
		return
	}
	// get image by id
	var image = db.Image{
		Id: id,
	}
	err = image.GetById()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(httperrors.ImageNotFound.Status, httperrors.ImageNotFound.Body)
		return
	}
	// if found return image
	c.JSON(http.StatusOK, image)
}
