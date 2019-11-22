
package afs_sdk

import (
	"io/ioutil"
	"strconv"
	unode "../unode"
	helper "../helper"
	"math"
	"sync"
	"errors"
	"fmt"
)

type Block struct {
	fileName string
	afid     string
}

type UploadBlockChannelStruct struct {
	uploadResponse    unode.UploadResponse
	blockIndex        int
	transmissionError error
}

func uploadBlock(responseChannel chan UploadBlockChannelStruct, uploadNode unode.UploadNode, block Block, exp_days string,
	blockFileIndex int, field string) {

	uploadBlockChannelStruct := UploadBlockChannelStruct{transmissionError: nil}

	uploadResponse, err := uploadNode.Upload(field, exp_days, RAW, block.fileName)

	if err != nil {
		uploadBlockChannelStruct.transmissionError = err
	}

	uploadBlockChannelStruct.uploadResponse = uploadResponse
	uploadBlockChannelStruct.blockIndex = blockFileIndex
	responseChannel <- uploadBlockChannelStruct
}

func createSeedFile(blocks []Block, fileDir string) (string, error) {

	f, err := ioutil.TempFile(fileDir, "temp_")
	defer f.Close()
	if err != nil {
		return "", err
	}

	f.WriteString(strconv.Itoa(1) + "\r\n")
	for _, block := range blocks {
		f.WriteString(block.afid + "\r\n")
	}
	return f.Name(), nil
}

func uploadBlocks(uploadNode unode.UploadNode, blockSizeInt int, uploadFilePath string, uploadFileName string, tempDirPath string, expDays string, field string, maxUploadThread int) ([]Block, error) {

	var blockSize int64
	blockSize = mbToBytes(blockSizeInt)

	file, err := helper.OpenFile(uploadFilePath, uploadFileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	fileSize, err := helper.GetFileSize(file)
	if err != nil {
		return nil, err
	}

	numberOfSegment := int(math.Ceil(float64(fileSize) / float64(blockSize)))

	c := make(chan UploadBlockChannelStruct)
	c2 := make(chan error)

	uploadComplete := 0
	uploaderRemain := maxUploadThread

	var i int64 = 1
	b := make([]byte, blockSize)
	var Error error

	var blocks []Block
	wg := sync.WaitGroup{}
	waitGroupAllComplete := sync.WaitGroup{}
	waitGroupAllComplete.Add(1)

	go func() {
		defer waitGroupAllComplete.Done()
		for uploadComplete != numberOfSegment {
			select {
			case msg := <-c:
				uploaderRemain++
				if msg.transmissionError == nil {
					if msg.uploadResponse.Status == 1 {
						blocks[msg.blockIndex-1].afid = msg.uploadResponse.Afid
						uploadComplete++
						fmt.Println(strconv.Itoa(uploadComplete)+"/"+strconv.Itoa(numberOfSegment))
					} else {
						Error = errors.New(msg.uploadResponse.Message)
					}
				} else {
					Error = msg.transmissionError
				}
				if uploaderRemain == 1 {
					wg.Done()
				}
			case <-c2:

			}
			if Error != nil && uploaderRemain == maxUploadThread {
				break
			}
		}
	}()

	for ; i <= int64(numberOfSegment); i++ {
		wg.Wait()
		if Error != nil {
			break
		}

		file.Seek((i-1)*(blockSize), 0)

		if len(b) > int((fileSize - (i-1)*blockSize)) {
			b = make([]byte, fileSize-(i-1)*blockSize)
		}

		file.Read(b)

		f, err := ioutil.TempFile(tempDirPath, "temp_")
		defer f.Close()
		block := Block{fileName: f.Name()}

		_, err = f.Write(b)
		if err != nil {
			Error = err
			c2 <- Error
			break
		}

		blocks = append(blocks, block)
		uploaderRemain--
		go uploadBlock(c, uploadNode, block, expDays, int(i), field)

		if uploaderRemain == 0 {
			wg.Add(1)
		}
	}
	waitGroupAllComplete.Wait()

	if Error != nil {
		return nil, Error
	}

	return blocks, nil
}


