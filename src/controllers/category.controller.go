package controllers

import (
	"encoding/json"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Get godoc
// @Summary Get categories
// @Description Get all categories
// @Tags category
// @Accept  json
// @Produce  json
// @Router /category [get]
func GetCategories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var category []models.Category

	if err := database.DB.Find(&category).Error; err != nil {
		http.Error(writer, "Category not found", http.StatusNotFound)
		return
	}

	responseMessage := map[string]interface{}{
		"status":  200,
		"message": "fetched successfully",
		"data":    category,
		"count":   len(category),
	}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

// Create godoc
// @Summary Create category
// @Description Create all category
// @Tags category
// @Accept  json
// @Produce  json
// @Router /category [post]
func CreateCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var category models.Category

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&category); err != nil {
		http.Error(writer, "failed to decode JSON", http.StatusBadRequest)
		return
	}

	database.DB.Create(&category)

	responseMessage := map[string]string{"message": "category created successully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

// Delete godoc
// @Summary Delete category
// @Description Delete category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category_id path int true "Category ID"
// @Router /category/{category_id} [delete]
func DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	categoryID := vars["category_id"]

	categoryIDInt, errCategoryID := strconv.Atoi(categoryID)

	if errCategoryID != nil {
		http.Error(writer, "Invalid category_id", http.StatusBadRequest)
		return
	}

	var category models.Category

	result := database.DB.Where("id = ?", categoryIDInt).First(&category)

	if result.Error != nil {
		http.Error(writer, "Category not found", http.StatusNotFound)
		return
	} else {
		deleteResult := database.DB.Delete(&category)

		if deleteResult.Error != nil {
			http.Error(writer, "Error deleting record", http.StatusInternalServerError)
			return
		}
	}

	responseMessage := map[string]string{"message": "Category deleted successfully"}
	jsonResponse, _ := json.Marshal(responseMessage)

	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}
