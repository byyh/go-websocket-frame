// repo init
package repo

import (
	"go-websocket-frame/common/atom"
	"go-websocket-frame/common/global/plugin/log"
	

	"go-websocket-frame/common/repo/bsarea"
	"go-websocket-frame/common/repo/bsshopemployee"
)

var (

	bsareaRepo bsarea.Repository
	bsshopemployeeRepo bsshopemployee.Repository
)

func init() {
	var err error

	if bsareaRepo, err = bsarea.GetRepository(); nil != err {
		log.Error("初始化 bsareaRepo 失败", err)
		panic(atom.NewMyErrorByCode(atom.ErrCodeDb))
	}

	if bsshopemployeeRepo, err = bsshopemployee.GetRepository(); nil != err {
		log.Error("初始化 bsshopemployeeRepo 失败", err)
		panic(atom.NewMyErrorByCode(atom.ErrCodeDb))
	}

}

func GetBsareaRepo() bsarea.Repository {
	return bsareaRepo
}


func GetBsshopemployeeRepo() bsshopemployee.Repository {
	return bsshopemployeeRepo
}

