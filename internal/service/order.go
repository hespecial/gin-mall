package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/constant"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/model"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"github.com/hespecial/gin-mall/internal/repository/kafka"
	"github.com/hespecial/gin-mall/pkg/random"
	"go.uber.org/zap"
	"time"
)

type orderService struct{}

var OrderService = new(orderService)

func (*orderService) GetOrderList(c *gin.Context, _ *request.GetOrderListReq) (*response.GetOrderListResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	list, err := dao.GetOrderList(userID)
	if err != nil {
		global.Log.Error("获取用户订单列表失败", zap.Error(err))
		return nil, e.ErrorGetOrderList, e.IsLogicError
	}

	var orders []*response.Order
	for _, order := range list {
		var items []*response.OrderItem
		for _, item := range order.Items {
			items = append(items, &response.OrderItem{
				OrderItemID: item.ID,
				ProductID:   item.ProductID,
				Quantity:    item.Quantity,
				ProductName: item.Product.Title,
				UnitPrice:   item.Product.Price,
				TotalPrice:  item.Product.Price * float64(item.Quantity),
			})
		}
		orders = append(orders, &response.Order{
			OrderID:       order.ID,
			OrderNumber:   order.OrderNumber,
			PaymentStatus: order.PaymentStatus,
			TotalAmount:   order.TotalAmount,
			Address:       order.Address.Address,
			Items:         items,
			CreatedAt:     order.CreatedAt,
		})
	}

	resp := &response.GetOrderListResp{
		Orders: orders,
	}

	return resp, e.Success, e.NotLogicError
}

func (*orderService) GetOrderInfo(c *gin.Context, req *request.GetOrderInfoReq) (*response.GetOrderInfoResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	order, err := dao.GetOrderInfo(userID, req.OrderID)
	if err != nil {
		global.Log.Error("获取订单详情失败", zap.Error(err))
		return nil, e.ErrorGetOrderInfo, e.IsLogicError
	}

	var items []*response.OrderItem
	for _, item := range order.Items {
		items = append(items, &response.OrderItem{
			OrderItemID: item.ID,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			ProductName: item.Product.Title,
			UnitPrice:   item.Product.Price,
			TotalPrice:  item.Product.Price * float64(item.Quantity),
		})
	}

	resp := &response.GetOrderInfoResp{
		Order: &response.Order{
			OrderID:       order.ID,
			OrderNumber:   order.OrderNumber,
			PaymentStatus: order.PaymentStatus,
			TotalAmount:   order.TotalAmount,
			Address:       order.Address.Address,
			Items:         items,
			CreatedAt:     order.CreatedAt,
		},
	}

	return resp, e.Success, e.NotLogicError
}

func (*orderService) CreateOrder(c *gin.Context, req *request.CreateOrderReq) (*response.CreateOrderResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	totalAmount, err := calculateTotalAmount(req.Items)
	if err != nil {
		global.Log.Error("计算总金额失败", zap.Error(err))
		return nil, e.ErrorCalculateTotalAmount, e.IsLogicError
	}

	var items []model.OrderItem
	for _, item := range req.Items {
		items = append(items, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	order := &model.Order{
		OrderNumber:     random.GenerateOrderNumber(),
		UserID:          userID,
		TotalAmount:     totalAmount,
		AddressID:       req.AddressID,
		Items:           items,
		PaymentStatus:   model.PaymentStatusPending,
		PaymentDeadline: time.Now().Add(constant.OrderDueTime),
	}
	err = dao.CreateOrder(order)
	if err != nil {
		global.Log.Error("创建订单失败", zap.Error(err))
		return nil, e.ErrorCreateOrder, e.IsLogicError
	}

	producer, err := kafka.NewKafkaProducer([]string{constant.KafkaBrokers})
	if err != nil {
		global.Log.Error("新建kafka生产者失败", zap.Error(err))
		return nil, e.ErrorCreateOrder, e.IsLogicError
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		global.Log.Error("订单序列化失败", zap.Error(err))
		return nil, e.ErrorCreateOrder, e.IsLogicError
	}

	err = producer.SendMessage(constant.KafkaOrderTopic, order.OrderNumber, string(orderJson))
	if err != nil {
		global.Log.Error("订单发送至消息队列失败", zap.Error(err))
		return nil, e.ErrorCreateOrder, e.IsLogicError
	}

	resp := &response.CreateOrderResp{
		OrderNumber:   order.OrderNumber,
		PaymentStatus: order.PaymentStatus,
		TotalAmount:   totalAmount,
	}

	return resp, e.Success, e.NotLogicError
}

func (*orderService) DeleteOrder(_ *gin.Context, req *request.DeleteOrderReq) (*response.DeleteOrderResp, e.Code, bool) {
	err := dao.DeleteOrder(req.OrderID)
	if err != nil {
		global.Log.Error("删除订单失败", zap.Error(err))
		return nil, e.ErrorDeleteOrder, e.IsLogicError
	}

	resp := &response.DeleteOrderResp{}

	return resp, e.Success, e.NotLogicError
}

func calculateTotalAmount(items []*request.OrderItem) (float64, error) {
	// 获取所有商品ID
	var productIDs []uint
	for _, item := range items {
		productIDs = append(productIDs, item.ProductID)
	}

	// 根据商品ID获取商品信息
	products, err := dao.GetProductsByIDs(productIDs)
	if err != nil {
		return 0, err
	}

	// 创建一个 map 来方便商品 ID 与商品记录的映射
	productMap := make(map[uint]*model.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	// 计算总金额
	var totalAmount float64
	for _, item := range items {
		product, exists := productMap[item.ProductID]
		if !exists {
			return 0, fmt.Errorf("product with ID %d not found", item.ProductID)
		}
		totalAmount += float64(item.Quantity) * product.Price
	}

	return totalAmount, nil
}
