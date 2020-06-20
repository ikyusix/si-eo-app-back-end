package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sieo_app/models"
)

func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	// returnData := models.Respons{
	// 	Success: true,
	// 	Message: "SUCCESS",
	// 	Data:    data,
	// }

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[HandleSucess.utils] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

func HandleError(resp http.ResponseWriter, message string) {
	data := models.Respons{
		Success: false,
		Message: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong"))
		fmt.Printf("[HandleError.utils] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}
