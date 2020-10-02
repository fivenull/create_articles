package controllers

import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"mwxy_create_articles/pkg/serializer"
)

func ParamErrorMsg(filed string, tag string) string {
	// 未通过验证的表单域与中文对应
	fieldMap := map[string]string{
		"Name":  "名称",
		"Info":  "简介",
		"Order": "排序",
		"Body":  "内容",
		"CID":   "所属分类",
	}
	// 未通过的规则与中文对应
	tagMap := map[string]string{
		"required": "不能为空",
		"min":      "太短",
		"max":      "太长",
		"email":    "格式不正确",
	}
	fieldVal, findField := fieldMap[filed]
	tagVal, findTag := tagMap[tag]
	if findField && findTag {
		// 返回拼接出来的错误信息
		return fieldVal + tagVal
	}
	return ""
}

func ErrorResponse(err error) serializer.Response {
	fmt.Println("解析内容发生错误", err)
	// 处理 Validator 产生的错误
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.ParamErr(
				ParamErrorMsg(e.Field(), e.Tag()),
				err,
			)
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
