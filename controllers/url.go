package controllers

import (
    "math/rand"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "url-shortener/models"
    "url-shortener/database"
)

const shortCodeLength = 8
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
    seed := rand.NewSource(time.Now().UnixNano())
    random := rand.New(seed)
    shortCode := make([]byte, shortCodeLength)
    for i := range shortCode {
        shortCode[i] = charset[random.Intn(len(charset))]
    }
    return string(shortCode)
}

func ShortenURL(c *gin.Context) {
    var request struct {
        OriginalURL string `json:"original_url"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    shortCode := generateShortCode()
    url := models.URL{
        ShortCode:   shortCode,
        OriginalURL: request.OriginalURL,
        UserID:      1, // Replace with actual user ID
    }

    if err := database.DB.Create(&url).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"short_code": shortCode})
}

func GetOriginalURL(c *gin.Context) {
    shortCode := c.Param("short_code")
    var url models.URL
    if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    c.Redirect(http.StatusFound, url.OriginalURL)
}