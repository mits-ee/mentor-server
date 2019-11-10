package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
	"strconv"
)

func GetUserSelf(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)

	data := models.GetUser(userId, true)

	resp := utils.Message(true, "Success")
	resp["data"] = data

	utils.Respond(w, resp)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid user id"))
		return
	}

	user := models.GetUser(uint(userId), true)
	if user == nil {
		utils.Respond(w, utils.Message(false, fmt.Sprintf("Cannot find user with id %d", userId)))
		return
	}

	resp := utils.Message(true, "Success")

	resp["data"] = user.GetPublicInfo()

	utils.Respond(w, resp)
	return
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	//

	type profile struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		ImageUrl  string `json:"imageUrl"`
		Bio       string `json:"bio"`
	}

	request := profile{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
	}

	user := models.GetUser(userId, false)

	if len(request.FirstName) > 0 {
		user.FirstName = request.FirstName
	}
	if len(request.LastName) > 0 {
		user.LastName = request.LastName
	}
	if len(request.Bio) > 0 {
		user.Bio = request.Bio
	}

	models.GetDB().Save(user)

	resp := utils.Message(true, "Profile successfully edited")

	utils.Respond(w, resp)
}

func EditGroupProfile(w http.ResponseWriter, r *http.Request) {
	groupId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid group id"))
		return
	}

	group := models.GetGroup(uint(groupId))

	if group == nil {
		utils.Respond(w, utils.Message(false, "Group not found"))
		return
	}

	type payload struct {
		Title       string `json:"title"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
	}
	profile := payload{}
	err = json.NewDecoder(r.Body).Decode(&profile)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	if len(profile.Title) > 0 {
		group.Title = profile.Title
	}
	if len(profile.Tagline) > 0 {
		group.Title = profile.Tagline
	}
	if len(profile.Description) > 0 {
		group.Title = profile.Description
	}

	models.GetDB().Save(group)

	utils.Respond(w, utils.Message(true, "Profile successfully edited"))
}
