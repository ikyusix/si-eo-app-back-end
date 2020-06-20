package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"sieo_app/common"
	"sieo_app/models"
	"sieo_app/user"
)

type UserHandler struct {
	UserService user.UseCaseInterface
	DB *gorm.DB
}

func CreateUserHandler(r *mux.Router, userService user.UseCaseInterface) {
	userHandler := &UserHandler{UserService: userService}
	r.HandleFunc("/user", userHandler.FindAllUser).Methods(http.MethodGet)
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/{id}", userHandler.FindUserById).Methods(http.MethodGet)
	s.HandleFunc("/signup", userHandler.SignUp).Methods(http.MethodPost)
	s.HandleFunc("/signin", userHandler.Login).Methods(http.MethodPost)
	s.HandleFunc("/upload/{id}", userHandler.UploadImage).Methods(http.MethodPost)
	s.HandleFunc("/{id}", userHandler.UpdateUser).Methods(http.MethodPut)
	s.HandleFunc("/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)
	s.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		common.Response(writer, common.Message(false, "Url not found"))
		return
	})
}

func (h UserHandler) FindAllUser(writer http.ResponseWriter, request *http.Request) {
	response, err := h.UserService.FindAll()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, response)
		return
	}
	common.Response(writer, response)
	return
}
func (h UserHandler) FindUserById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, common.Message(false, "Please provide valid id"))
		return
	}

	response, err := h.UserService.FindById(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, response)
		return
	}

	common.Response(writer, response)
	return
}
func (h UserHandler) SignUp(writer http.ResponseWriter, request *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, common.Message(false, "Invalid Request "+err.Error()))
		return
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logrus.Error(err)
		fmt.Errorf("[UserHandler.SignUp] Hash password: %w", err)
	}

	user.Password = string(hashPass)
	response, err := h.UserService.SignUp(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, response)
		return
	}
	common.Response(writer, response)
	return
}
func (h UserHandler) Login(writer http.ResponseWriter, request *http.Request) {
	dataLogin, err := ioutil.ReadAll(request.Body)
	if err != nil {
		common.RespondWithError(writer, http.StatusBadRequest, "Ooops something wrong")
		logrus.Error(err)
	}
	user := models.User{}

	err = json.Unmarshal(dataLogin, &user)
	if user.Email == ""{
		common.RespondWithError(writer, http.StatusBadRequest, "Email is missing")
		return
	}
	if user.Password == ""{
		common.RespondWithError(writer, http.StatusBadRequest, "Password is missing")
		return
	}
	password := user.Password
	dataUser, err := h.UserService.GetUserByEmail(user.Email)
	if err != nil {
		logrus.Error(err)
		return
	}

	hashedPassword := dataUser.Password
	checkPassword := common.ComparePassword(hashedPassword, []byte(password))
	if checkPassword {
		common.HandleSuccess(writer, dataUser, nil)
		return
	}
	common.RespondWithError(writer, http.StatusBadRequest, "Invalid password")
}
func (h UserHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	user := new(models.User)
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, common.Message(false, "Please provide valid id"))
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, common.Message(false, "Invalid request "+err.Error()))
		return
	}
	response, err := h.UserService.Update(id, user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, response)
		return
	}
	common.Response(writer, response)
	return
}
func (h UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, common.Message(false, "Please provide valid id"))
		return
	}

	response, err := h.UserService.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		common.Response(writer, response)
		return
	}
	common.Response(writer, response)
	return
}
func (h UserHandler) UploadImage(writer http.ResponseWriter, request *http.Request) {
	pathvar := mux.Vars(request)
	id, err := strconv.Atoi(pathvar["id"])
	if err != nil {
		logrus.Error(err)
		fmt.Errorf("[UserHandler.UploadImage] Upload profile: %w", err)
		common.RespondWithError(writer, http.StatusBadRequest, "Please provide valid id")
		return
	}

	uploadImage, handler, err := request.FormFile("UrlImage")
	if err != nil {
		logrus.Error(err)
		fmt.Errorf("[UserHandler.UploadImage] Upload profile: %w", err)
	}
	defer uploadImage.Close()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fileName := handler.Filename
	tempFile := filepath.Join("assets", "user", fileName)
	targetFile, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logrus.Error(err)
		fmt.Errorf("[UserHandler.UploadImage]: %w", err)
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadImage); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user, message, err := h.UserService.UploadImage(id, tempFile)
	if err != nil {
		common.RespondWithError(writer, http.StatusBadRequest, "Ooops something wrong")
		logrus.Error(err)
		fmt.Errorf("[UserHandler.UploadImage]: %w", err)
		return
	}
	if message != "" {
		common.RespondWithError(writer, http.StatusBadRequest, message)
		return
	}
	common.HandleSuccess(writer, http.StatusOK, user)
}
