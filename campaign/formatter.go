package campaign

import "strings"

type CampaignFormater struct {
	Id                int    `json:"id"`
	User_id           int    `json:"user_id"`
	Name              string `json:"Name"`
	Short_description string `json:"short_description"`
	ImageURL          string `json:"image_url"`
	GoalAmount        int    `json:"goal_amount"`
	CurrentAmount     int    `json:"current_amount"`
	Slug              string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.Id = campaign.Id
	campaignFormater.User_id = campaign.UserId
	campaignFormater.Name = campaign.Name
	campaignFormater.Short_description = campaign.ShortDescription
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	campaignFormater.Slug = campaign.Slug
	campaignFormater.ImageURL = ""

	if len(campaign.CampaignImage) > 0 {
		campaignFormater.ImageURL = campaign.CampaignImage[0].FileName
	}
	return campaignFormater

}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {

	campaignsFormater := []CampaignFormater{}
	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}
	return campaignsFormater
}

type CampaignDetailFormater struct {
	Id                int                  		`json:"id"`
	User_id           int                  		`json:"user_id"`
	Name              string               		`json:"Name"`
	Short_description string               		`json:"short_description"`
	ImageURL          string               		`json:"image_url"`
	GoalAmount        int                  		`json:"goal_amount"`
	CurrentAmount     int                  		`json:"current_amount"`
	Slug              string               		`json:"slug"`
	Perks             []string             		`json:"perks"`
	User              CampaignUserFormater 		`json:"user"`
	Images 			  []CampaignImageFormater 	`json:"images"`
}

type CampaignUserFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormater struct{
	ImageURL string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormater {
	campaignDetailFormater := CampaignDetailFormater{}
	campaignDetailFormater.Id = campaign.Id
	campaignDetailFormater.User_id = campaign.UserId
	campaignDetailFormater.Name = campaign.Name
	campaignDetailFormater.Short_description = campaign.ShortDescription
	campaignDetailFormater.GoalAmount = campaign.GoalAmount
	campaignDetailFormater.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormater.Slug = campaign.Slug
	campaignDetailFormater.ImageURL = ""

	if len(campaign.CampaignImage) > 0 {
		campaignDetailFormater.ImageURL = campaign.CampaignImage[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormater.Perks = perks
	
	user := campaign.User
	campaignUserFormater := CampaignUserFormater{}
	campaignUserFormater.Name = user.Name
	campaignUserFormater.ImageURL = user.AvatarFileName
	campaignDetailFormater.User = campaignUserFormater

	images := []CampaignImageFormater{}

	for _, image := range campaign.CampaignImage{
		campaignImageFormater := CampaignImageFormater{}
		campaignImageFormater.ImageURL = image.FileName
		
		isPrimary := false

		if image.IsPrimary == 1{
			isPrimary = true
		}
		campaignImageFormater.IsPrimary = isPrimary

		images = append(images, campaignImageFormater)
	}
	campaignDetailFormater.Images = images

	return campaignDetailFormater
}