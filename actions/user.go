package actions

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/shreeharsha-factly/chi-swagger-go/models"
)

// single user returned in response
// swagger:response userResponse
type userResponse struct {
	// single user
	// in: body
	Body models.User
}

// swagger:parameters addUser updateUser
type userBody struct {
	// single user
	// in: body
	User struct {
		// the email for this user
		Email string `gorm:"column:email" json:"email"`
		// the name for this user
		Name string `gorm:"column:name" json:"name"`
		// the age for this user
		Age int `gorm:"column:age" json:"age"`
	}
}

// swagger:parameters getUser updateUser deleteUser
type userIdParam struct {
	// in: path
	// Required: true
	UserID string `gorm:"primary_key" json:"userId"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{userId} users getUser
	//
	// Get user by id
	//
	// This will show a user by id.
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//   200: userResponse
	userId := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal(err)
	}

	req := &models.User{
		ID: uint(id),
	}
	json.NewDecoder(r.Body).Decode(&req)

	models.DB.Model(&models.User{}).First(&req)

	json.NewEncoder(w).Encode(req)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /users users addUser
	//
	// Get user by id
	//
	// This will show a user by id.
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//   200: userResponse
	req := &models.User{}
	json.NewDecoder(r.Body).Decode(&req)

	if validErrs := req.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err := models.DB.Model(&models.User{}).Create(&req).Error

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(req)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route PUT /users/{userId} users updateUser
	//
	// Get user by id
	//
	// This will show a user by id.
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//   200: userResponse
	userId := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal(err)
	}

	req := &models.User{
		ID: uint(id),
	}

	json.NewDecoder(r.Body).Decode(&req)
	user := &models.User{
		ID: uint(id),
	}
	models.DB.First(&user)

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Age != 0 {
		user.Age = req.Age
	}

	models.DB.Model(&user).Updates(models.User{Email: req.Email, Name: req.Name, Age: req.Age})

	models.DB.First(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route DELETE /users/{userId} users deleteUser
	//
	// Get user by id
	//
	// This will show a user by id.
	//
	// Produces:
	// - application/json
	//
	// Responses:
	//   200: userResponse
	userId := chi.URLParam(r, "userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal(err)
	}

	user := &models.User{
		ID: uint(id),
	}
	json.NewDecoder(r.Body).Decode(&user)

	models.DB.First(&user)
	models.DB.Delete(&user)

	json.NewEncoder(w).Encode(user)
}
