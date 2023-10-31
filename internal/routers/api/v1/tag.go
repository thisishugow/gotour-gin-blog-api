package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary Get tags
// @Produce json
// @Param name query string false "tag_name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Success 200 {object} model.Tag "Success"
// @Failure 400 {object} errcode.Error "Request Fail"
// @Failure 500 {object} errcode.Error "Interal Error"
// @Router /api/v1/tags [get]
func (t Tag) Get(c *gin.Context)    {}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
