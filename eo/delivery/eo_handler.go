package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sieo_app/common"
	"sieo_app/eo"
	"sieo_app/models"
	"sieo_app/utils"
	"strconv"

	"github.com/davecgh/go-spew/spew"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type EoHandler struct {
	EoService eo.EoServiceInterface
}

func CreateEoHandler(r *mux.Router, eoService eo.EoServiceInterface) {
	eoHandler := &EoHandler{EoService: eoService}

	v1 := r.PathPrefix("/eo").Subrouter()

	v1.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(eoHandler.readAllEo))).Methods(http.MethodGet)
	v1.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(eoHandler.insertEo))).Methods(http.MethodPost)
	v1.Handle("/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(eoHandler.updateEo))).Methods(http.MethodPut)
	// v1.Handle("/upgrade/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(eoHandler.upgradeEo))).Methods(http.MethodPut)
	v1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		common.Response(w, common.Message(false, "URL Not Found"))
		return
	})
}

func (e *EoHandler) readAllEo(w http.ResponseWriter, r *http.Request) {
	response, err := e.EoService.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}

func (e *EoHandler) insertEo(w http.ResponseWriter, r *http.Request) {

	textId := r.FormValue("UserID")
	userId, _ := strconv.Atoi(textId)

	txtName := r.FormValue("EoName")
	eoname, err := utils.ValidationNull(txtName)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}

	txtIdentity := r.FormValue("Identity")
	identity, err := utils.ValidationNull(txtIdentity)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}

	uploadIdentity, handler, err := r.FormFile("IdentityImg")
	if err != nil {
		logrus.Error(err)
		common.Message(false, "[EoHandler.uploadIdentity] Upload IdentityImg")
	}
	defer uploadIdentity.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileName := handler.Filename
	identityPath := filepath.Join("assets", "identity", fileName)
	targetFile, err := os.OpenFile(identityPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logrus.Error(err)
		common.Message(false, "[EoHandler.uploadIdentity] Upload IdentityImg")
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadIdentity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	txtLicense := r.FormValue("License")
	license, err := utils.ValidationNull(txtLicense)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}

	uploadLicense, handler, err := r.FormFile("LicenseImg")
	if err != nil {
		logrus.Error(err)
		common.Message(false, "[EoHandler.uploadLicense] Upload uploadLicense")
	}
	defer uploadIdentity.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nameLicense := handler.Filename
	licensePath := filepath.Join("assets", "license", nameLicense)
	targetFileLicense, err := os.OpenFile(licensePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		logrus.Error(err)
		common.Message(false, "[EoHandler.UploadImage] Upload uploadLicense")
	}
	defer targetFileLicense.Close()

	if _, err := io.Copy(targetFile, uploadLicense); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	txtAddress := r.FormValue("Address")
	address, err := utils.ValidationNull(txtAddress)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}
	txtPhone := r.FormValue("Phone")
	phone, err := utils.ValidationNull(txtPhone)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}
	txtWebsite := r.FormValue("Website")
	website, err := utils.ValidationNull(txtWebsite)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}
	txtInstagram := r.FormValue("Instagram")
	instagram, err := utils.ValidationNull(txtInstagram)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}
	txtFacebook := r.FormValue("Facebook")
	facebook, err := utils.ValidationNull(txtFacebook)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}
	txtTwitter := r.FormValue("Twitter")
	twitter, err := utils.ValidationNull(txtTwitter)
	if err != nil {
		utils.HandleError(w, err.Error())
		return
	}

	reqEo := models.Eo{
		UserID:      userId,
		Name:        eoname,
		Identity:    identity,
		ImgIdentity: identityPath,
		License:     license,
		ImgLicense:  licensePath,
		Address:     address,
		Phone:       phone,
		Website:     website,
		Instagram:   instagram,
		Facebook:    facebook,
		Twitter:     twitter,
	}

	spew.Dump(reqEo)

	response, err := e.EoService.Insert(&reqEo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}
	common.Response(w, response)
	return
}

func (e *EoHandler) updateEo(w http.ResponseWriter, r *http.Request) {
	Eo := new(models.Eo)
	User := new(models.User)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		common.Response(w, common.Message(false, "Please Provide Valid ID"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&Eo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		common.Response(w, common.Message(false, "invalid Request"))
		return
	}
	User.Status = Eo.Status
	response, responseUser, err := e.EoService.UpdateEo(id, Eo, User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		common.Response(w, response)
		return
	}
	common.Response(w, response)
	return
}
