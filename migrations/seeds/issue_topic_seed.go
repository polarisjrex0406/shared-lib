package seeds

import (
	"github.com/omimic12/shared-lib/entities"
	"gorm.io/gorm"
)

func IssueTopic(db *gorm.DB) error {
	issueTopics := []entities.IssueTopic{
		{Name: "Billing & Account Issue"},
		{Name: "Connectivity"},
		{Name: "Geo-Blocking"},
		{Name: "Performance Variability"},
		{Name: "New Feature Request"},
		{Name: "Error Message"},
	}
	for _, issueTopic := range issueTopics {
		if err := db.FirstOrCreate(&issueTopic, issueTopic).Error; err != nil {
			return err
		}
	}
	return nil
}
