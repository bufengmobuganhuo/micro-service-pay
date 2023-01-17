package handler

import (
	"context"
	common "github.com/bufengmobuganhuo/micro-service-common"
	"github.com/bufengmobuganhuo/micro-service-payment/domain/model"
	"github.com/bufengmobuganhuo/micro-service-payment/domain/service"
	payment "github.com/bufengmobuganhuo/micro-service-payment/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (p Payment) AddPayment(ctx context.Context, req *payment.PaymentInfo, resp *payment.PaymentID) error {
	payment := &model.Payment{}
	err := common.SwapTo(req, payment)
	if err != nil {
		return err
	}
	paymentId, err := p.PaymentDataService.AddPayment(payment)
	if err != nil {
		return err
	}
	resp.PaymentId = paymentId
	return nil
}

func (p Payment) UpdatePayment(ctx context.Context, req *payment.PaymentInfo, resp *payment.Response) error {
	payment := &model.Payment{}
	err := common.SwapTo(req, payment)
	if err != nil {
		return err
	}
	if err = p.PaymentDataService.UpdatePayment(payment); err != nil {
		return err
	}
	resp.Msg = "更新成功"
	return nil
}

func (p Payment) DeletePaymentByID(ctx context.Context, id *payment.PaymentID, resp *payment.Response) error {
	err := p.PaymentDataService.DeletePayment(id.PaymentId)
	if err != nil {
		return err
	}
	resp.Msg = "删除成功"
	return nil
}

func (p Payment) FindPaymentByID(ctx context.Context, id *payment.PaymentID, resp *payment.PaymentInfo) error {
	model, err := p.PaymentDataService.FindPaymentByID(id.PaymentId)
	if err != nil {
		return err
	}
	info := &payment.PaymentInfo{}
	return common.SwapTo(model, info)
}

func (p Payment) FindAllPayment(ctx context.Context, req *payment.All, resp *payment.PaymentAll) error {
	allPayment, err := p.PaymentDataService.FindAllPayment()
	if err != nil {
		return err
	}
	for _, v := range allPayment {
		paymentInfo := &payment.PaymentInfo{}
		if err := common.SwapTo(v, paymentInfo); err != nil {
			continue
		}
		resp.PaymentInfo = append(resp.PaymentInfo, paymentInfo)
	}
	return nil
}
