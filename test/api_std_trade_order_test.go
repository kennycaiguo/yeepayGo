package test

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/sdk"
	"testing"
)

func TestApiStdTradeOrder(t *testing.T) {
	config, err := GetYeepayConfig()
	if err != nil {
		t.Error(err)
		return
	}
	logger := yeepayGo.NewYeepayLogger()
	apiStdTradeOrder := sdk.NewApiStdTradeOrder(config, logger)
	req := apiStdTradeOrder.GetRequest("TEST002", "0.01", "商品名称", "商品描述", "")
	//goodsName := "abcdefghijklmnopqrstuvwxyz0123456789"
	//goodsName = "~!@#$^&*()_`-={}|:\"<>?[]\\;,./"
	//goodsName = "" //%+'
	//req := apiStdTradeOrder.GetRequest("DS191016_20280839", "0.01", goodsName, "")
	resp, err := apiStdTradeOrder.GetResponse(req)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.IsSuccess() {
		t.Error("resp.Result.Code", resp.Result.Code)
		t.Error("resp.Result.Message", resp.Result.Message)
		return
	}
	t.Log("resp.Result.Code", resp.Result.Code)
	t.Log("resp.Result.Message", resp.Result.Message)
	t.Log("resp.Result.ParentMerchantNo", resp.Result.ParentMerchantNo)
	t.Log("resp.Result.MerchantNo", resp.Result.MerchantNo)
	t.Log("resp.Result.OrderId", resp.Result.OrderId)
	t.Log("resp.Result.UniqueOrderNo", resp.Result.UniqueOrderNo)
	t.Log("resp.Result.GoodsParamExt", resp.Result.GoodsParamExt)
	t.Log("resp.Result.Token", resp.Result.Token)
	t.Log("resp.Result.FundProcessType", resp.Result.FundProcessType)
	t.Log("resp.Result.ParentMerchantName", resp.Result.ParentMerchantName)
	t.Log("resp.Result.MerchantName", resp.Result.MerchantName)
}
