package httperrors

import "net/http"

var ImageInvalidForm = Error{
	Status: http.StatusNotAcceptable,
	Body: Body{
		Code:    "000001",
		Message: "invalid form",
	},
}
var ImageInvalidCid = Error{

	Status: http.StatusNotAcceptable,
	Body: Body{
		Code:    "000002",
		Message: "invalid cid",
	},
}
var ImageAlreadyExist = Error{
	Status: http.StatusConflict,
	Body: Body{
		Code:    "000003",
		Message: "already exist",
	},
}
var ImageInvalidUUID = Error{
	Status: http.StatusNotAcceptable,
	Body: Body{
		Code:    "000004",
		Message: "invalid uuid",
	},
}
var ImageNotFound = Error{
	Status: http.StatusNotFound,
	Body: Body{
		Code:    "000005",
		Message: "image not found",
	},
}
