package delivery

import (
	"net/http"
	"sieo_app/banner"
	"sieo_app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type BannerHandler struct {
	bannerService banner.BannerService
}

func CreateBannerHandler(r *mux.Router, bannerService banner.BannerService) {

	bannerHandler := BannerHandler{bannerService}

	r.HandleFunc("/banner", bannerHandler.getAllBanner).Methods(http.MethodGet)
	r.HandleFunc("/banner/{idEvent}", bannerHandler.getBannerByIdEvent).Methods(http.MethodGet)
	// r.HandleFunc("/event/{idEvent}", eventHandler.getEventById).Methods(http.MethodGet)
}

func (e *BannerHandler) getAllBanner(resp http.ResponseWriter, req *http.Request) {
	banner, err := e.bannerService.ViewAll()
	if err != nil {
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	utils.HandleSuccess(resp, banner)
}

func (e *BannerHandler) getBannerByIdEvent(resp http.ResponseWriter, req *http.Request) {
	muxVar := mux.Vars(req)
	strID := muxVar["idEvent"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		utils.HandleError(resp, "ID Harus angka")
		return
	}
	banner, err := e.bannerService.ViewByIdEvent(id)
	if err != nil {
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	utils.HandleSuccess(resp, banner)
}

// func (e *EventHandler) getEventById(resp http.ResponseWriter, req *http.Request) {

// 	muxVar := mux.Vars(req)

// 	strID := muxVar["idEvent"]

// 	id, err := strconv.Atoi(strID)
// 	if err != nil {
// 		handleError(resp, "ID Harus angka")
// 		return
// 	}

// 	event, err := e.eventService.ViewById(id)
// 	if err != nil {
// 		handleError(resp, "Oppss server someting wrong")
// 		return
// 	}

// 	handleSuccess(resp, event)
// }
