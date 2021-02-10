package handler

import (
	"fmt"
	"funding/campaign"
	"funding/helper"
	"funding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.FindCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("campaign not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} else if len(campaigns) == 0 {
		response := helper.APIResponse("campaign not found", 404, "error", campaign.FormatCampaigns(campaigns))
		c.JSON(404, response)
		return
	}

	response := helper.APIResponse("list of campaign", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) GetChampaignsById(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("campaign not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaigns, err := h.campaignService.GetCampaignById(input)
	response := helper.APIResponse("list of campaign", http.StatusOK, "success", campaign.FormatCampignDetail(campaigns))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed create campaign", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {

		response := helper.APIResponse("failed create campaign", 422, "error", nil)
		c.JSON(422, response)
		return
	}

	response := helper.APIResponse("list of campaign", http.StatusOK, "success", campaign.FormatCampignDetail(newCampaign))
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("campaign not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed update campaign", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	newUpdatedCampaign, err := h.campaignService.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("failed update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of campaign", http.StatusOK, "success", campaign.FormatCampignDetail(newUpdatedCampaign))
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed update campaign", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.campaignService.SavaeCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success upload campaign image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return
}
