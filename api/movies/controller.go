package movies

import (
	"assessment/api/common"
	"assessment/db"
	"assessment/db/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(e *gin.Engine) {
	r := e.Group("/movies")

	r.GET("", getMoviesHandler)
	r.GET("/:id", getMovieHandler)
	r.POST("", createMovieHandler)
	r.PATCH("/:id", updateMovieHandler)
	r.DELETE("/:id", deleteMovieHandler)
}

func getMoviesHandler(c *gin.Context) {
	movies := []model.Movie{}
	err := db.Conn.Find(&movies).Error
	if err != nil {
		common.WriteResp(c, nil, "failed to get movies", http.StatusInternalServerError)
		return
	}

	respBody := make([]movie, len(movies))
	for i := range movies {
		respBody[i].parseFrom(&movies[i])
	}

	common.WriteResp(c, respBody, "success get movies", http.StatusOK)
}

func getMovieHandler(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.WriteResp(c, nil, "path param 'id' should be a number", http.StatusBadRequest)
		return
	}

	if movieID < 1 {
		common.WriteResp(c, nil, "invalid number for path param 'id'", http.StatusBadRequest)
		return
	}

	var mov model.Movie
	err = db.Conn.Find(&mov, movieID).Error
	if err != nil || mov.ID == 0 {
		msg := fmt.Sprintf("movie with id: %d does not exist", movieID)
		common.WriteResp(c, nil, msg, http.StatusNotFound)
		return
	}

	var respBody movie
	respBody.parseFrom(&mov)
	common.WriteResp(c, &respBody, "success get movie", http.StatusOK)
}

func createMovieHandler(c *gin.Context) {
	var req movieRequest
	err := c.BindJSON(&req)
	if err != nil {
		common.WriteResp(c, nil, "invalid request body", http.StatusBadRequest)
		return
	}

	err = common.Validate.Struct(&req)
	if err != nil {
		common.WriteResp(c, nil, "invalid request body "+err.Error(), http.StatusBadRequest)
		return
	}

	var mov model.Movie
	err = db.Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(req.parseTo()).Error; err != nil {
			return err
		}

		if err := tx.First(&mov, "title = ?", req.Title).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		common.WriteResp(c, nil, "failed creating movie", http.StatusInternalServerError)
		return
	}

	var respBody movie
	respBody.parseFrom(&mov)
	common.WriteResp(c, &respBody, "success creating movie", http.StatusOK)
}

func updateMovieHandler(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.WriteResp(c, nil, "path param 'id' should be a number", http.StatusBadRequest)
		return
	}

	if movieID < 1 {
		common.WriteResp(c, nil, "invalid number for path param 'id'", http.StatusBadRequest)
		return
	}

	var req movieRequest
	err = c.BindJSON(&req)
	if err != nil {
		common.WriteResp(c, nil, "invalid request body", http.StatusBadRequest)
		return
	}

	err = common.Validate.Struct(&req)
	if err != nil {
		common.WriteResp(c, nil, "invalid request body "+err.Error(), http.StatusBadRequest)
		return
	}

	var mov model.Movie
	err = db.Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&mov, movieID).Error; err != nil {
			return err
		}

		updMov := req.parseTo()
		updMov.Model = mov.Model

		if err := tx.Save(updMov).Error; err != nil {
			return err
		}

		if err := tx.First(&mov).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		common.WriteResp(c, nil, "failed updating movie", http.StatusInternalServerError)
		return
	}

	var respBody movie
	respBody.parseFrom(&mov)
	common.WriteResp(c, &respBody, "success updating movie", http.StatusOK)
}

func deleteMovieHandler(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.WriteResp(c, nil, "path param 'id' should be a number", http.StatusBadRequest)
		return
	}

	if movieID < 1 {
		common.WriteResp(c, nil, "invalid number for path param 'id'", http.StatusBadRequest)
		return
	}

	var rowsAffected int64
	err = db.Conn.Transaction(func(tx *gorm.DB) error {
		tx = tx.Unscoped().Delete(&model.Movie{}, movieID)
		if tx.Error != nil {
			return tx.Error
		}
		rowsAffected = tx.RowsAffected
		return nil
	})
	if err != nil {
		common.WriteResp(c, nil, "failed deleting movie", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		common.WriteResp(c, nil, "movie doesn't exist", http.StatusNotFound)
		return
	}

	common.WriteResp(c, nil, "success deleting movie", http.StatusOK)
}
