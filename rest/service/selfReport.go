package service

import (
	"fmt"
	"time"

	"github.com/RedHatInsights/chrome-service-backend/rest/database"
	"github.com/RedHatInsights/chrome-service-backend/rest/models"
)

func GetSelfReport(accountID uint)(models.SelfReport, error) {
  var selfReport models.SelfReport
  fmt.Println("This is our accountID on GetSelfReport: ", accountID)
  
  err := database.DB.Where("user_identity_id = ?", accountID).Find(&selfReport).Error
  return selfReport, err
}

func HandlePatchSelfReport(accountID uint, newSelfReport models.SelfReport)(models.SelfReport, error) {
  var selfReport models.SelfReport

  err := database.DB.Where("user_identity_id = ?", accountID).Find(&selfReport).Error
  

  return selfReport, err
}

func HandleNewSelfReport(accountID uint, newSelfReport models.SelfReport)(error) {
  var selfReport models.SelfReport 

  err := database.DB.Statement.DB.Where("user_identity_id = ?", accountID).Create(&selfReport).Error
  selfReport.UserIdentityID = accountID

  fmt.Println("checking our accountID at HandleNewSelfReport: ", accountID)
  
  if err != nil {
    return err
  }

  // Query logic following last-visited 
  return database.DB.Model(&selfReport).
    Update("job_role", newSelfReport.JobRole).
    Update("products_of_interest", newSelfReport.ProductsOfInterest).
    Update("updated_at", time.Now()).
    Error
}
