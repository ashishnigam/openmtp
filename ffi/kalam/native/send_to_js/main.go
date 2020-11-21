package send_to_js

/*
	#include "stdio.h"
	#include "stdlib.h"
	#include "stdbool.h"

 	typedef void (* cb_result_t)(char*);
	void send_result(int64_t ptr, char* json) {
		cb_result_t cb = (cb_result_t) ptr;

		cb(json);
	}

 	typedef void (* cb_upload_files_t)(char*, double);
	void send_upload_files_result(int64_t ptr, char* name, double speed) {
		cb_upload_files_t cb = (cb_upload_files_t) ptr;

		cb(name, speed);
	}
*/
import "C"
import (
	"github.com/ganeshrvel/go-mtpfs/mtp"
)

func SendError(ptr int64, err error) {
	pErr, pErrorMsg := processError(err)

	o := ErrorResult{
		Error:    pErr,
		ErrorMsg: pErrorMsg,
		Data:     nil,
	}

	json := toJson(o)

	C.send_result(C.int64_t(ptr), C.CString(json))
}

func SendInitialize(ptr int64, deviceInfo *mtp.DeviceInfo) {
	o := InitializeResult{
		Error:    "",
		ErrorMsg: "",
		Data:     *deviceInfo,
	}

	json := toJson(o)

	C.send_result(C.int64_t(ptr), C.CString(json))
}

func SendDeviceInfo(ptr int64, deviceInfo *mtp.DeviceInfo) {
	o := DeviceInfoResult{
		Error:    "",
		ErrorMsg: "",
		Data:     *deviceInfo,
	}

	json := toJson(o)

	C.send_result(C.int64_t(ptr), C.CString(json))
}

//func UploadFilesCb(ptr int64, fi *mtpx.ProgressInfo) {
//	C.send_upload_files_result(C.int64_t(ptr), C.CString(fi.FileInfo.Name), C.double(fi.Speed))
//}

//// List the file list
//func SendWalkResult(ptr int64, device *mtp.Device, storageId uint32, fullPath string) {
//	var fList []FileInfo
//
//	_, _, _ = mtpx.Walk(device, storageId, fullPath, false, false, func(objectId uint32, fi *mtpx.FileInfo, err error) error {
//		f := FileInfo{
//			Size:       fi.Size,
//			IsDir:      fi.IsDir,
//			ModTime:    fi.ModTime.Format(DateTimeFormat),
//			Name:       fi.Name,
//			FullPath:   fi.FullPath,
//			ParentPath: fi.ParentPath,
//			Extension:  fi.Extension,
//			ParentId:   fi.ParentId,
//			ObjectId:   fi.ObjectId,
//		}
//
//		fList = append(fList, f)
//
//		return nil
//	})
//
//	start := time.Now()
//	var json = jsoniter.ConfigFastest
//	w, err := json.Marshal(&fList)
//	if err != nil {
//		return err
//	}
//
//	pretty.Println("Native elapsed: ", time.Since(start).Microseconds())
//	pretty.Println("fList length: ", len(fList))
//
//	C.send_file_info_result(C.int64_t(ptr), C.CString(string(w)))
//}