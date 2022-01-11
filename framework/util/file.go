package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

)
func Exists(path string) bool {
	//os.Stat获取文件信息
	_,err:=os.Stat(path)
	if err!=nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}

func IsHiddenDirectory(path string) bool{
	return len(path)>1&&strings.HasPrefix(filepath.Base(path),".")
}

func SubDir(folder string) ([]string,error) {
	subs,err:=ioutil.ReadDir(folder)
	if err!=nil{
		return nil,err
	}

	ret:=[]string{}
	for _,sub :=range subs{
		if sub.IsDir(){
			ret = append(ret, sub.Name())
		}
	}
	return ret,nil
}

func DownloadFile(filepath string,url string) error{
	resp,err:=http.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()

	out,err:=os.Create(filepath)
	if err!=nil{
		return err
	}
	defer out.Close()

	_,err=io.Copy(out,resp.Body)
	return err
}