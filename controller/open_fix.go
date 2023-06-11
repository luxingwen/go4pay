package controller

import (
	"go4pay/model"
	log "go4pay/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PixTransactionReceived struct {
	Event     string        `json:"event"`
	Charge    interface{}   `json:"charge"`
	PixQRCode PixQRCode     `json:"pixQrCode"`
	Pix       Pix           `json:"pix"`
	Company   Company       `json:"company"`
	Account   Account       `json:"account"`
	Refunds   []interface{} `json:"refunds"`
}

type PixQRCode struct {
	Name           string      `json:"name"`
	Value          interface{} `json:"value"`
	Comment        string      `json:"comment"`
	Identifier     string      `json:"identifier"`
	CorrelationID  string      `json:"correlationID"`
	PaymentLinkID  string      `json:"paymentLinkID"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
	BRCode         string      `json:"brCode"`
	PaymentLinkURL string      `json:"paymentLinkUrl"`
	QRCodeImage    string      `json:"qrCodeImage"`
}

type Payer struct {
	Name          string `json:"name"`
	TaxID         TaxID  `json:"taxID"`
	CorrelationID string `json:"correlationID"`
}

type TaxID struct {
	ID   string `json:"taxID"`
	Type string `json:"type"`
}

type Pix struct {
	Payer         Payer     `json:"payer"`
	Value         float64   `json:"value"`
	Time          string    `json:"time"`
	EndToEndID    string    `json:"endToEndId"`
	TransactionID string    `json:"transactionID"`
	Type          string    `json:"type"`
	PixQRCode     PixQRCode `json:"pixQrCode"`
	CreatedAt     string    `json:"createdAt"`
	PixKey        string    `json:"pixKey"`
	GlobalID      string    `json:"globalID"`
}

type Company struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	TaxID string `json:"taxID"`
}

type Account struct {
	ClientID string `json:"clientId"`
}

func OpenFixCallBack(c *gin.Context) {

	event := c.Param("event")

	log.Infof("event:%s", event)
	var pixTransaction PixTransactionReceived
	if err := c.ShouldBindJSON(&pixTransaction); err != nil {
		log.Errorf("ShouldBindJSON err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pixTransaction.Pix.Value = pixTransaction.Pix.Value / 100
	fee, err := model.CalculationFee(pixTransaction.Pix.Value)

	if err != nil {
		log.Errorf("CalculationFee err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payOrder := &model.PayOrder{
		OrderId:    pixTransaction.Pix.Payer.CorrelationID,
		Channel:    "open_fix",
		Name:       "",
		Comment:    pixTransaction.PixQRCode.Comment,
		ValueIn:    pixTransaction.Pix.Value,
		ValueType:  pixTransaction.Pix.Payer.TaxID.Type,
		Identifier: pixTransaction.PixQRCode.Identifier,
		Data:       c.GetString("body"),
		Status:     pixTransaction.Event,
		CreateTime: pixTransaction.PixQRCode.CreatedAt,
		UpdateTime: pixTransaction.PixQRCode.UpdatedAt,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	payOrder.Fee = fee
	payOrder.Value = payOrder.ValueIn - fee

	model.CreateOrUpdatePayOrder(payOrder)
	// Return a response
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
