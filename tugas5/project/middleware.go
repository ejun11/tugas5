package project

import (
	"log"
	"strconv"
	"strings"
	"tugas5/project/controller"
	"tugas5/project/model"
	respone "tugas5/project/respone"
	"tugas5/project/service"
	"tugas5/tokenjwt"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userSvc *service.UserServices
}

func NewMiddleware(userSvc *service.UserServices) *Middleware {
	return &Middleware{
		userSvc: userSvc,
	}
}

func (m *Middleware) Auth(c *gin.Context) {
	tokenn := c.GetHeader("Authorization")
	tokenArr := strings.Split(tokenn, "Token ")
	if len(tokenArr) != 2 {
		c.Set("ERROR", "no token")
		controller.WriteErrorJsonResponseGin(c, respone.ErrUnauthorized())
		return
	}
	myTok, err := tokenjwt.VerifyToken(tokenArr[1])
	if err != nil {
		c.Set("ERROR", err.Error())
		controller.WriteErrorJsonResponseGin(c, respone.ErrUnauthorized())
		return
	}
	user := m.userSvc.FindByEmail(myTok.Email)
	userDetail := user.Payload.(*model.User)
	c.Set("USER_ID", strconv.FormatUint(uint64(userDetail.ID), 10))
	c.Set("USER_EMAIL", userDetail.Email)
	c.Set("USER_NAME", userDetail.Fullname)

	c.Next()

}

func (m *Middleware) Check(next gin.HandlerFunc, roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("USER_EMAIL")
		user := m.userSvc.FindByEmail(email)
		userDetail := user.Payload.(*model.User)

		isExist := false

		for _, role := range roles {
			if role == userDetail.Role {
				isExist = true
				break
			}
		}

		if !isExist {
			controller.WriteErrorJsonResponseGin(ctx, respone.ErrUnauthorized())
			return
		}

		next(ctx)
	}
}

func (m *Middleware) Trace(c *gin.Context) {
	log.Printf("Get request :%v url :%v\n", c.Request.Method, c.Request.URL)
	c.Next()
	isError := c.GetString("ERROR")
	if isError != "" {
		log.Printf("error get all :%v\n", isError)
	}
	log.Printf("Finised request :%v url :%v\n", c.Request.Method, c.Request.URL)
}
