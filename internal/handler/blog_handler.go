package handler

import (
	"api-blog-go/internal/models"
	"api-blog-go/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type BlogHandler struct {
	repo *repository.BlogRepository
}

func NewBlogHandler(repo *repository.BlogRepository) *BlogHandler {
	return &BlogHandler{repo: repo}
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tạo slug từ name
	blog.Slug = slug.Make(blog.Name)

	if err := h.repo.Create(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo blog"})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {
	blogs, err := h.repo.FindAll([]string{"id", "name", "slug", "avatar", "created_at"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách blog"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	blog, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy blog"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	blog, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy blog"})
		return
	}

	if err := c.ShouldBindJSON(blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.Slug = slug.Make(blog.Name)

	if err := h.repo.Update(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật blog"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa blog thành công"})
}
