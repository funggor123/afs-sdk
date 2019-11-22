package afs_sdk

import (
	"io/ioutil"
	unode "../unode"
	helper "../helper"
	"strconv"
)
/*

Upload file to Upload Node
params: address  string   	  upload node ip + ":" + upload node port
		field        	string   AFS / ARFS / BOTH
		expireDays   	int	  	 The expire days
		uploadMethod 	string   RAW / SEED 
		uploadFilePath 	string   File Path of upload file
		uploadFileName 	string   File Name of upload file

resp:   afid 		 string   upload file afid
		status       bool     upload status
		message      string   if status = false, it will show the error message
*/
func Upload(address string, field string, expireDays int, uploadMethod string, uploadFilePath string, uploadFileName string, blockSize int, maxUploadThread int) string {

	expDays := strconv.Itoa(expireDays)

	tempDirPath, err := ioutil.TempDir(uploadFilePath, "temp_")
	if err !=nil {
		return "_r=false;_message=cannot create temporary directory" + ";"
	}
	defer helper.DeleteDirectory(tempDirPath)
	uploadNode := unode.UploadNode{ Address: address}

	switch uploadMethod {
	case RAW:
		uploadResponse, err := uploadNode.Upload(field, expDays, uploadMethod, uploadFilePath + uploadFileName)
		if err !=nil {
			return "_r=false;_message=upload file fail:" + err.Error() + ";"
		}
		return "_r=true;_afid=" + uploadResponse.Afid + ";"
	case SEED:
		blocks, err := uploadBlocks(uploadNode, blockSize, uploadFilePath, uploadFileName, tempDirPath, expDays, field, maxUploadThread)
		if err !=nil {
			return "_r=false;_message=blocks create/upload fail:" +  err.Error() + ";"
		}
		seedFileFullPath, err := createSeedFile(blocks, tempDirPath)
		if err !=nil {
			return "_r=false;_message=create seed file fail:" +  err.Error() + ";"
		}
		uploadResponse, err := uploadNode.Upload(field, expDays, RAW, seedFileFullPath)
		if err !=nil {
			return "_r=false;_message=upload seed file fail:" + err.Error() + ";"
		}
		return "_r=true;_afid=" + uploadResponse.Afid + ";"
	}
	return "_r=false;_message=invalid upload method " + uploadMethod + ";"
}

