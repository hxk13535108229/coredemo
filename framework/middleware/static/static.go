package static

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gohade/hade/framework/gin"
)

const INDEX= "index.html"

type ServeFileSystem interface {
	http.FileSystem

	Exists(prefix string,path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root string
	indexes bool
}

func LocalFile(root string,indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root,indexes),
		root: root,
		indexes: indexes,
	}
}

func (l *localFileSystem) Exists(prefix string,filepath string) bool {
	if p:=strings.TrimPrefix(filepath,prefix);len(p)<len(filepath) {
		name:=path.Join(l.root,p)
		stats,err:=os.Stat(name)
		if err!=nil {
			return false
		}
		if stats.IsDir() {
			if !l.indexes {
				index:=path.Join(name,INDEX)
				_,err:=os.Stat(index)
				if err!=nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func ServeRoot(urlPrefix,root string) gin.HandlerFunc {
	return Serve(urlPrefix,LocalFile(root,false))
}

func Serve(urlPrefix string,fs ServeFileSystem) gin.HandlerFunc {
	fileServe:=http.FileServer(fs) 
	if urlPrefix!="" {
		fileServe=http.StripPrefix(urlPrefix,fileServe)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix,c.Request.URL.Path){
			fileServe.ServeHTTP(c.Writer,c.Request)
			c.Abort()
		}
	}
}