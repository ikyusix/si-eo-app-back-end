package common

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"sieo_app/models"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
func RespondWithError(w http.ResponseWriter, status int, message string) {
	var error models.Error
	error.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
func ComparePassword(hashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		return false
	}
	return true
}
func HandleSuccess(resp http.ResponseWriter, data interface{}, user *models.User) {
	returnData:=models.Respons{
		Success:true,
		Message: "Success",
		Data:data,
	}

	jsonData, err := json.Marshal(returnData)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Something when wrong"))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

