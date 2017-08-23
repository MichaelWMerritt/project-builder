package model

type BuildStatus string

const (

	CREATED BuildStatus = "CREATED"
	IN_PROGRESS BuildStatus = "IN_PROGRESS"
	FAILED BuildStatus = "FAILED"
	COMPLETE BuildStatus = "COMPLETE"

)