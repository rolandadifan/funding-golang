package campaign

import (
	"strings"
	"time"
)

type CampaignFormater struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Perks            string    `json:"perks"`
	BackerCount      int       `json:"backer_count"`
	GoalAmount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	Slug             string    `json:"slug"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ImageURL         string    `json:"image_url"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormatter := CampaignFormater{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.Perks = campaign.Perks
	campaignFormatter.BackerCount = campaign.BackerCount
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter

}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {
	if len(campaigns) == 0 {
		return []CampaignFormater{}
	}
	var campaignsFormatter []CampaignFormater

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	GoalAmount       int                     `json:"goal_amount"`
	CurrentAmount    int                     `json:"current_amount"`
	ImageURL         string                  `json:"image_url"`
	UserID           int                     `json:"user_id"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             CampaignUserFormatter   `json:"user"`
	Images           []CampaignImageFormater `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormater struct {
	ImageURL  string `jspn:"image_url"`
	IsPrimary bool   `jspn:"is_primary"`
}

func FormatCampignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	user := campaign.User
	CampaignUserFormatter := CampaignUserFormatter{}
	CampaignUserFormatter.Name = user.Name
	CampaignUserFormatter.ImageURL = user.Avatar
	campaignDetailFormatter.User = CampaignUserFormatter

	images := []CampaignImageFormater{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormaters := CampaignImageFormater{}
		campaignImageFormaters.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormaters.IsPrimary = isPrimary
		images = append(images, campaignImageFormaters)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
