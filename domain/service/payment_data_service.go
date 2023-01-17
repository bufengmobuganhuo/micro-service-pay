package service

import (
	"github.com/bufengmobuganhuo/micro-service-payment/domain/model"
	"github.com/bufengmobuganhuo/micro-service-payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(*model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}

// NewPaymentDataService 创建
func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService {
	return &PaymentDataService{paymentRepository}
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}

// AddPayment 插入
func (u *PaymentDataService) AddPayment(payment *model.Payment) (int64, error) {
	return u.PaymentRepository.CreatePayment(payment)
}

// DeletePayment 删除
func (u *PaymentDataService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}

// UpdatePayment 更新
func (u *PaymentDataService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdatePayment(payment)
}

// FindPaymentByID 查找
func (u *PaymentDataService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindPaymentByID(paymentID)
}

// FindAllPayment 查找
func (u *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}
