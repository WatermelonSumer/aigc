package assistant

import "github.com/gin-gonic/gin"

type AsstServices interface {
	List(c *gin.Context)
	Detail(c *gin.Context)
	Retrieve(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type AsstServiceImpl struct{}

func (s *AsstServiceImpl) List(c *gin.Context) {
	c.JSON(200, gin.H{"message": "List all items"})
}

func (s *AsstServiceImpl) Detail(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Detail of item"})
}

func (s *AsstServiceImpl) Retrieve(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Retrieve item"})
}

func (s *AsstServiceImpl) Delete(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete item"})
}

func (s *AsstServiceImpl) Create(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create new item"})
}

func (s *AsstServiceImpl) Update(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update item"})
}
