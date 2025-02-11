package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func IssueCategory(db *gorm.DB) error {
	issueCategories := []entities.IssueCategory{
		{Name: "Billing & Account Issue", IsAdvanced: true},
		{Name: "Usage Limits", IsAdvanced: true},
		{Name: "Connectivity", IsAdvanced: true},
		{Name: "Geo-Blocking", IsAdvanced: true},
		{Name: "Protocol Compatibility", IsAdvanced: true},
		{Name: "IP Rotation Issue", IsAdvanced: true},
		{Name: "Performance Variability", IsAdvanced: true},
		{Name: "New Feature Request", IsAdvanced: false},
		{Name: "Error Message", IsAdvanced: false},
	}
	for _, issueCategory := range issueCategories {
		if err := db.FirstOrCreate(&issueCategory, issueCategory).Error; err != nil {
			return err
		}
	}
	return nil
}
