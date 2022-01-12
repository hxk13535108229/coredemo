package env

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/gohade/hade/framework/contract"
)

type HadeEnv struct {
	//代表.env所在的目录
	folder string

	//保存所有的环境变量
	maps map[string]string
}

func NewHadeEnv(params ...interface{}) (interface{} ,error){
	if len(params)!=1 {
		return nil,errors.New("NewHadeEnv param error")
	}

	//读取folder文件
	folder := params[0].(string)

	//实例化
	hadeEnv := &HadeEnv{
		folder: folder,
		maps: map[string]string{"APP_ENV":contract.EnvDevelopment},
	}

	//解析folder/.env文件
	file := path.Join(folder,".env")
	//读取.env

	//打开文件
	fi,err := os.Open(file)
	if err==nil {
		defer fi.Close()

		//读取文件
		br:= bufio.NewReader(fi)
		for{
			//按照行进行读取
			line,_,c := br.ReadLine()
			if c== io.EOF{
				break
			}
			//按照等号进行解析
			s:= bytes.SplitN(line,[]byte{'='},2)
			//如果不符合规范，则返回
			if len(s)<2 {
				continue
			}
			//保存map
			key := string(s[0])
			val := string(s[1])
			hadeEnv.maps[key]=val
		}
	}

	//获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _,e := range os.Environ() {
		pair := strings.SplitN(e,"=",2)
		if len(pair) <2 {
			continue
		}
		hadeEnv.maps[pair[0]]=pair[1]
	}

	//返回实例
	return hadeEnv,nil
}

func (en *HadeEnv) AppEnv() string {
	return en.Get("APP_ENV")
}

func (en *HadeEnv) IsExist(key string) bool{
	_,ok := en.maps[key]
	return ok
}

func (en *HadeEnv) Get(key string) string {
	if val,ok := en.maps[key] ; ok {
		return val
	}
	return ""
}

func (en *HadeEnv) All() map[string]string {
	return en.maps
}