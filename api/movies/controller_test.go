package movies

import (
	"assessment/api/common"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func before() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(common.UnseenPanicHandler)
	StartRouter(r)
	return r
}

func getJsonBytes(obj interface{}) []byte {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	return jsonBytes
}

// Purpose for these tests is to check if our endpoints sanitize the request properly
func TestGetMoviesHandler(t *testing.T) {
	router := before()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)
	router.ServeHTTP(w, req)
	// Panic because db connection is not initialized
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"data":null,"message":"unexpected error occured","status_code":500}`, w.Body.String())
}

func TestGetMovieHandler(t *testing.T) {
	router := before()

	// Using string as path parameter for movieID will result in bad request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/hello", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"path param 'id' should be a number","status_code":400}`, w.Body.String())

	// Using number lower than 1 will result in bad request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/movies/-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"invalid number for path param 'id'","status_code":400}`, w.Body.String())

	// Using number bigger or equal than 1 will pass the request validation
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/movies/10", nil)
	router.ServeHTTP(w, req)
	// Panic because db connection is not initialized
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"data":null,"message":"unexpected error occured","status_code":500}`, w.Body.String())
}

func TestCreateMovieHandler(t *testing.T) {
	router := before()

	// Empty title will result in bad request
	reqBody := movieRequest{
		Title:       "",
		Description: "Kisah mengenai...",
		Rating:      6.7,
		Image:       "https://image.com",
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Empty description will result in bad request
	reqBody.Title = "Sebuah judul"
	reqBody.Description = ""
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/movies", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Rating equal to 0 will result in bad request
	reqBody.Description = "Kisah mengenai..."
	reqBody.Rating = 0
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/movies", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Empty image will result in bad request
	reqBody.Rating = 6.7
	reqBody.Image = ""
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/movies", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Proper request body will pass the validation
	reqBody.Image = "https://image.com"
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/movies", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	// Panic because db connection is not initialized
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"data":null,"message":"unexpected error occured","status_code":500}`, w.Body.String())
}

func TestUpdateMovieHandler(t *testing.T) {
	router := before()

	// Using string as path parameter for movieID will result in bad request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/movies/hello", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"path param 'id' should be a number","status_code":400}`, w.Body.String())

	// Using number lower than 1 will result in bad request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"invalid number for path param 'id'","status_code":400}`, w.Body.String())

	// Empty title will result in bad request
	reqBody := movieRequest{
		Title:       "",
		Description: "Kisah mengenai...",
		Rating:      6.7,
		Image:       "https://image.com",
	}
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/10", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Empty description will result in bad request
	reqBody.Title = "Sebuah judul"
	reqBody.Description = ""
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/10", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Rating equal to 0 will result in bad request
	reqBody.Description = "Kisah mengenai..."
	reqBody.Rating = 0
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/10", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Empty image will result in bad request
	reqBody.Rating = 6.7
	reqBody.Image = ""
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/10", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Proper request body will pass the validation
	reqBody.Image = "https://image.com"
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/movies/10", bytes.NewBuffer(getJsonBytes(&reqBody)))
	router.ServeHTTP(w, req)
	// Panic because db connection is not initialized
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"data":null,"message":"unexpected error occured","status_code":500}`, w.Body.String())
}

func TestDeleteMovieHandler(t *testing.T) {
	router := before()

	// Using string as path parameter for movieID will result in bad request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/movies/hello", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"path param 'id' should be a number","status_code":400}`, w.Body.String())

	// Using number lower than 1 will result in bad request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/movies/-1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"data":null,"message":"invalid number for path param 'id'","status_code":400}`, w.Body.String())

	// Using number bigger or equal than 1 will pass the request validation
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/movies/10", nil)
	router.ServeHTTP(w, req)
	// Panic because db connection is not initialized
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"data":null,"message":"unexpected error occured","status_code":500}`, w.Body.String())
}
