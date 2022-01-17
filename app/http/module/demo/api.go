package demo

import (
	demoService "github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/app/provider/user"
	"github.com/gohade/hade/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	r.GET("/demo/user",api.DemoUser)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{
		service: service,
	}
}

func (api *DemoApi) Demo(c *gin.Context) {
	c.JSON(200,"sleep!")
}

func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

func (api *DemoApi) DemoUser(c *gin.Context) {
	userProvider:=c.MustMake(user.UserKey).(user.UService)
	s:=userProvider.Foo()
	c.JSON(200,s)
}