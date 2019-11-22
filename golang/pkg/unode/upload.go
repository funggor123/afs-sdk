package afs_sdk

import (
	resty "gopkg.in/resty.v1"
	"errors"
	"time"
)

// Constant // 


// Response // 

type UploadResponse struct {
	Afid    string `json:"afid" binding:"required"`
	Status  int    `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// Node Define // 

type UploadNode struct 
{ 
	Address string
}

// EndPoints // 

func (uploadNode UploadNode) getUploadEndPoint() string {
	return uploadNode.Address + "/v1/un/upload/file"
}

// Methods // 

func (uploadNode UploadNode) Upload(field string, expireDays string, uploadMethod string, fileFullPath string) (UploadResponse, error) {
	var uploadResponse UploadResponse
	resty.SetRetryCount(3).
	SetRetryWaitTime(5 * time.Second).
    SetRetryMaxWaitTime(20 * time.Second)
	_, err := resty.R().
		SetFile("file", fileFullPath).
		SetResult(&uploadResponse).
		SetFormData(map[string]string{
			"expire_days": expireDays,
			"upload_type": uploadMethod,
			"field" : field,
		}).
		Post(uploadNode.getUploadEndPoint())
	if err != nil {
		return UploadResponse{}, err
	}
	if uploadNode.isUploadFailed(uploadResponse) {
		return UploadResponse{}, errors.New(uploadResponse.Message)
	}
	return uploadResponse, nil
}

// Helpers // 

func (uploadNode UploadNode) isUploadFailed(uploadResponse UploadResponse) bool {
	return uploadResponse.Status != 1
}

