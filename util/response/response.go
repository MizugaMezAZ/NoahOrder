package response

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Gin struct for bind method to *gin.context
type Gin struct {
	Ctx *gin.Context
}

// Response struct define return json struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Response function for quickly using my define return sturct
func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	g.Ctx.Set("statuscode", code)
	return
}

// ErrResponse ...
func (g *Gin) ErrResponse(err interface{}, args ...interface{}) {
	switch err := err.(type) {
	case validator.ValidationErrors:
		var errlist []string
		// ermap := err.Translate(ginvalidator.DefaultTranslator)
		// for _, v := range ermap {
		// 	errlist = append(errlist, v)
		// }

		g.Response(400, "參數驗證錯誤", gin.H{
			"errors": errlist,
		})
		fmt.Println(err)

	case string:
		fmt.Println(err, args)

		if strings.Contains(err, "NULL") {
			g.Response(400, "查無資料", gin.H{
				// todo: lang.Get()...
				"errors": err,
			})
			return
		}

		if strings.Contains(err, "1054") {
			g.Response(400, "資料庫欄位有誤", gin.H{
				// todo: lang.Get()...
				"errors": err,
			})
			return
		}

		g.Response(400, err, gin.H{
			"errors": args,
		})

	case error:
		errstring := err.Error()
		fmt.Println(errstring)

		if strings.Contains(errstring, "cannot unmarshal") {
			g.Response(400, "請求格式解析錯誤", gin.H{
				// todo: lang.Get()...
				"errors": "json資料格式不正確",
			})
			return
		}

		if strings.Contains(errstring, "invalid character") {
			g.Response(400, "請求格式解析錯誤", gin.H{
				// todo: lang.Get()...
				"errors": "json格式錯誤",
			})
			return
		}

		if strings.Contains(errstring, "EOF") {
			g.Response(400, "請求格式解析錯誤", gin.H{
				// todo: lang.Get()...
				"errors": "json缺少body文本",
			})
			return
		}

		if strings.Contains(errstring, "1062") {
			e := strings.Fields(errstring)
			errs := ""
			for i, v := range e {
				if v == "Duplicate" {
					errs = fmt.Sprintf("%v 已經被使用", strings.ReplaceAll(e[i+2], "'", ""))
					break
				}
			}

			g.Response(400, errs, gin.H{
				"errors": errs,
			})
			return
		}

		if strings.Contains(errstring, "1054") {
			g.Response(400, "資料庫欄位有誤", gin.H{
				"errors": err,
			})
			return
		}

		if strings.Contains(errstring, "1406") {
			g.Response(400, "輸入資料長度過長", gin.H{
				"errors": err,
			})
			return
		}

		if strings.Contains(errstring, "no rows") {
			g.Response(400, "查無資料", gin.H{
				"errors": "查無資料",
			})
			return
		}

		if strings.Contains(errstring, "converting NULL") {
			g.Response(400, "查無資料", gin.H{
				"errors": "查無資料",
			})
			return
		}

		if strings.Contains(errstring, "playing electronic games") {
			g.Response(400, "玩家在遊戲中", gin.H{
				"errors": "玩家在遊戲中",
			})
			return
		}

		if strings.Contains(errstring, "餘額不足") {
			g.Response(400, "帳號餘額不足", gin.H{
				"errors": "帳號餘額不足",
			})
			return
		}

		if strings.Contains(errstring, "rpc error") {

			s := strings.Split(errstring, ":")
			str := s[len(s)-1]
			g.Response(400, str, gin.H{
				"errors": str,
			})
			return
		}

		g.Response(400, "服務器很忙", gin.H{
			"errors": errstring,
		})
	}
}
