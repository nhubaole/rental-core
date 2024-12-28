package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"golang.org/x/oauth2/google"
)

type NotificationServiceImpl struct {
	repo               *dataaccess.Queries
	projectID          string
	serviceAccountPath string
}

func NewNotificationServiceImpl() NotificationService {
	return &NotificationServiceImpl{
		repo:               dataaccess.New(global.Db),
		projectID:          global.Config.FCM.ProjectID,
		serviceAccountPath: global.Config.FCM.ServiceAccountPath,
	}
}

func (n *NotificationServiceImpl) GetByUserID(userID int) *responses.ResponseData {
	notifications, err := n.repo.GetNotificationsByUserID(context.Background(), int32(userID))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	if len(notifications) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       notifications,
	}
}

// Create implements NotificationService.
func (n *NotificationServiceImpl) Create(req dataaccess.CreateNotificationParams) (string, error) {
	err := n.repo.CreateNotification(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to create notification: %w", err)
	}

	return "Notification created successfully", nil
}

// SendNotification implements NotificationService.
func (n *NotificationServiceImpl) SendNotification(userId int, message string, referenceId *int, referenceType string) error {
	//create noti record
	createNotificationParams := dataaccess.CreateNotificationParams{
		UserID:        int32(userId),
		ReferenceID:   common.ConvertIntToPointerInt32(referenceId),
		ReferenceType: &referenceType,
		Title:         message,
	}
	_, err := n.Create(createNotificationParams)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	// send noti
	url := fmt.Sprintf("https://fcm.googleapis.com/v1/projects/%s/messages:send", n.projectID)
	deviceToken, err := n.repo.GetDeviceTokenByUserID(context.Background(), int32(userId))
	if err != nil || deviceToken == nil {
		fmt.Printf("Failed to get device token: %v", err)
		return err
	}

	payload := requests.NotificationRequest{}
	payload.Message.Token = *deviceToken
	payload.Message.Notification.Title = "Bạn có 1 thông báo mới"
	payload.Message.Notification.Body = message

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	bearerToken, err := getBearerToken(n)
	if err != nil {
		fmt.Printf("Failed to get bearer token: %v", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to send notification: %s", string(body))
	}

	fmt.Println("Notification sent successfully!")

	return nil
}

func getBearerToken(n *NotificationServiceImpl) (string, error) {
	ctx := context.Background()

	saData, err := ioutil.ReadFile(n.serviceAccountPath)
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(saData, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return "", err
	}

	token, err := conf.TokenSource(ctx).Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
