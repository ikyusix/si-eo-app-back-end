package delivery

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"sieo_app/common"
	"sieo_app/event"
	"sieo_app/models"
	"sieo_app/utils"

	"github.com/gorilla/mux"
)

type EventHandler struct {
	eventService event.EventService
}

func CreateEventHandler(r *mux.Router, eventService event.EventService) {

	eventHandler := EventHandler{eventService}

	r.HandleFunc("/event", eventHandler.getAllEvent).Methods(http.MethodGet)
	r.HandleFunc("/event", eventHandler.addEvent).Methods(http.MethodPost)
	r.HandleFunc("/event/{idEvent}", eventHandler.getEventById).Methods(http.MethodGet)
	r.HandleFunc("/listevent", eventHandler.getListEvent).Methods(http.MethodGet)

	r.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		common.Response(writer, common.Message(false, "Url not found"))
		return
	})
}

func (e *EventHandler) getAllEvent(resp http.ResponseWriter, req *http.Request) {
	events, err := e.eventService.ViewAll()
	if err != nil {
		utils.HandleError(resp, "Oppss server somting wrong")
		return
	}
	utils.HandleSuccess(resp, events)
}

func (e *EventHandler) getListEvent(resp http.ResponseWriter, req *http.Request) {
	listEvent, err := e.eventService.ViewListEvent()
	if err != nil {
		utils.HandleError(resp, "Oppss server somting wrong 2")
		return
	}
	utils.HandleSuccess(resp, listEvent)
}

func (e *EventHandler) getEventById(resp http.ResponseWriter, req *http.Request) {

	muxVar := mux.Vars(req)
	strID := muxVar["idEvent"]

	id, err := strconv.Atoi(strID)
	if err != nil {
		utils.HandleError(resp, "ID Harus angka")
		return
	}

	event, err := e.eventService.ViewById(id)
	if err != nil {
		utils.HandleError(resp, "Oppss server someting wrong")
		return
	}

	utils.HandleSuccess(resp, event)
}

func (e *EventHandler) addEvent(resp http.ResponseWriter, req *http.Request) {

	textEoId := req.FormValue("EoID")
	eoId, err := strconv.Atoi(textEoId)
	if err != nil {
		utils.HandleError(resp, "Number only")
		return
	}

	textName := req.FormValue("Name")
	name, err := utils.ValidationName(textName)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}

	textDate := req.FormValue("Date")
	date, err := utils.ValidationNull(textDate)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}

	textLocation := req.FormValue("Location")
	location, err := utils.ValidationNull(textLocation)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}

	textPrince := req.FormValue("Prince")
	prince, err := utils.ValidationNull(textPrince)
	if err != nil {
		utils.HandleError(resp, "Prince Is Required")
		return
	}

	textCapacity := req.FormValue("Capacity")
	capacity, err := utils.ValidationNull(textCapacity)
	if err != nil {
		utils.HandleError(resp, "Capacity Is Required")
		return
	}

	textDescription := req.FormValue("Description")
	description, err := utils.ValidationDescription(textDescription)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}

	fmt.Println(textEoId)
	fmt.Println(textName)
	fmt.Println(textDate)
	fmt.Println(textLocation)
	fmt.Println(textPrince)
	fmt.Println(textCapacity)
	fmt.Println(textDescription)
	fmt.Println("Ini adalah contoh tipe string")

	pathName, paths, err := e.AddBanner(resp, req)
	if err != nil {
		utils.ValidationRollbackImage(paths)
		resp.WriteHeader(http.StatusBadRequest)
		utils.HandleError(resp, err.Error())
		return
	}

	reqEvent := models.Event{
		EoID:        eoId,
		Name:        name,
		Location:    location,
		Date:        date,
		Prince:      prince,
		Capacity:    capacity,
		Description: description,
		Banner:      pathName,
	}

	newEvent, err := e.eventService.InsertEvents(&reqEvent, paths)
	if err != nil {
		utils.HandleError(resp, err.Error())
		return
	}

	utils.HandleSuccess(resp, newEvent)
}

func (e *EventHandler) AddBanner(resp http.ResponseWriter, req *http.Request) (string, []string, error) {

	var paths []string
	// var images []string
	// var size []int64
	err := req.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(resp, err)
		return "", nil, err
	}

	formdata := req.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders,
	files := formdata.File["image"] // grab the filenames
	var pathName string

	for i, _ := range files { // loop through the files one by one

		file, err := files[i].Open()
		path := "C:/xampp/htdocs/sieo_app/banner/"

		if err != nil {
			return "", nil, fmt.Errorf("oops server something wrong")
		}
		defer file.Close()

		_, err = utils.ValidationFileSize(files[i].Size)
		if err != nil {
			return "", nil, err
		}

		_, err = utils.ValidationFormatFile(files[i].Filename)
		if err != nil {
			return "", nil, err
		}

		out, err := os.Create(path + files[i].Filename)
		defer out.Close()

		if err != nil {
			return "", nil, fmt.Errorf("Unable to create the file for writing. Check your write access privilege")
		}

		pathHost := "http://10.0.2.2:80/sieo_app/banner/"
		pathName = pathHost + files[i].Filename

		// images = append(images, files[i].Filename)
		paths = append(paths, pathHost+files[i].Filename)
		// size = append(size, files[i].Size)

		_, err = io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Fprintln(resp, err)
			return "", nil, fmt.Errorf("oops server something wrong")
		}

	}

	// _, err = e.bannerService.AddPath(paths)
	// if err != nil {
	// 	handleError(resp, "Oppss server somting wrong")
	// 	return
	// }

	return pathName, paths, nil
}
