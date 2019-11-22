package main

import (
	afs_sdk "./pkg/afs_sdk"
	cmd "./pkg/cmd"
	helper "./pkg/helper"
	"strconv"
)

//sample: result := afs_sdk.Upload("http://39.108.80.53:8074", afs_sdk.AFS, 7, afs_sdk.SEED, "./" , "48_FL1_Prog.pdf" , 1)
func main(){
	cmd.Main()
	if cmd.Configs.Action == "upload" {
		result := afs_sdk.Upload(cmd.Configs.Address, cmd.Configs.Field, cmd.Configs.ExpDays, strconv.Itoa(cmd.Configs.Method), cmd.Configs.FilePath, cmd.Configs.FileName , cmd.Configs.BlockSize)
		print(result + "\n")
	}
	helper.Terminate()
}