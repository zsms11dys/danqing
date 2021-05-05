package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"project/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestResponse struct {
	Message string             `json:"message"`
	Count   int                `json:"count"`
	Results []model.TestResult `json:"test_results"`
}

type ProductResponse struct {
	Message  string          `json:"message"`
	Count    int             `json:"count"`
	Products []model.Product `json:"products"`
}

type TicketResponse struct {
	Message string         `json:"message"`
	Count   int            `json:"count"`
	Tickets []model.Ticket `json:"tickets"`
}

type Standard struct {
	Name  string  `json:"name"`
	Lower float64 `json:"lower"`
	Upper float64 `json:"upper"`
	Unit  string  `json:"unit"`
}

type Process struct {
	Name      string     `json:"name"`
	Key       int64      `json:"key"`
	Standards []Standard `json:"standards"`
}

type Test struct {
	TicketId int64     `json:"ticket_id"`
	NodeKey  int64     `json:"node_key"`
	Tester   int64     `json:"tester"`
	Values   []float64 `json:"values"`
}

var r *gin.Engine
var DB *gorm.DB
var err error

func addProduct(c *gin.Context) {
	db := DB
	var product model.Product
	c.BindJSON(&product)
	product.Version = 1
	err = db.Create(&product).Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func updateProduct(c *gin.Context) {
	db := DB
	var product model.Product
	c.BindJSON(&product)
	var old_product model.Product
	err = db.Table("product").Where("product_key = ?", product.ProductKey).Order("version_id desc").First(&old_product).Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": "product not found",
		})
		return
	}
	product.Version = old_product.Version + 1
	err = db.Create(&product).Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func getProduct(c *gin.Context) {
	key, err := strconv.ParseInt(c.Query("product_key"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	db := DB
	response := ProductResponse{Count: 0}
	var products []model.Product
	var count int
	err = db.Table("product").Where("product_key = ?", key).Find(&products).Count(&count).Error
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = "success"
		response.Count = count
		response.Products = products
	}
	c.JSON(200, response)
}

func getProducts(c *gin.Context) {
	db := DB
	response := ProductResponse{Count: 0}
	var products []model.Product
	var count int
	err := db.Find(&products).Count(&count).Error
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = "success"
		response.Count = count
		response.Products = products
	}
	c.JSON(200, response)
}

func getTickets(c *gin.Context) {
	db := DB
	response := TicketResponse{Count: 0}
	var tickets []model.Ticket
	var count int
	err := db.Find(&tickets).Count(&count).Error
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = "success"
		response.Count = count
		response.Tickets = tickets
	}
	c.JSON(200, response)
}

func getTests(c *gin.Context) {
	db := DB
	response := TestResponse{Count: 0}
	var results []model.TestResult
	var count int
	err := db.Find(&results).Count(&count).Error
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = "success"
		response.Count = count
		response.Results = results
	}
	c.JSON(200, response)
}

func addTicket(c *gin.Context) {
	db := DB
	var ticket model.Ticket
	c.BindJSON(&ticket)
	var product model.Product
	err = db.Table("product").Where("product_key = ?", ticket.ProductKey).Where("version_id = ?", ticket.Version).First(&product).Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": "product not found",
		})
		return
	} else {
		ticket.ProductId = product.Id
		err = db.Create(&ticket).Error
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		var processes []Process
		err = json.Unmarshal([]byte(product.Processes), &processes)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
		db1 := db.Begin()
		defer func() {
			if err != nil {
				db1.Rollback()
			}
		}()
		for _, process := range processes {
			for standard_index, standard := range process.Standards {
				result := model.TestResult{
					TicketId: ticket.Id,
					NodeKey:  process.Key,
					NodeName: process.Name,
					NodeNum:  int64(standard_index),
					Upper:    standard.Upper,
					Lower:    standard.Lower,
					Unit:     standard.Unit,
				}
				err = db1.Create(&result).Error
				if err != nil {
					c.JSON(200, gin.H{
						"message": err.Error(),
					})
					return
				}
			}
		}
		err = db1.Commit().Error
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
			})
		}
	}
}

func addTest(c *gin.Context) {
	db := DB
	var test Test
	c.BindJSON(&test)
	db1 := db.Begin()
	defer func() {
		if err != nil {
			db1.Rollback()
		}
	}()
	for index, value := range test.Values {
		err = db1.Table("test_result").Where("ticket_id = ?", test.TicketId).Where("node_key = ?", test.NodeKey).Where("node_num = ?", int64(index)).Updates(model.TestResult{Value: value, Tester: test.Tester}).Error
		if err != nil {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	err = db1.Commit().Error
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func main() {
	r = gin.Default()

	DB, err = gorm.Open("mysql", "root:danqing@(82.156.191.198)/project?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println(err)
	}
	DB = DB.Debug()
	defer DB.Close()

	//localhost:8080
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	r.GET("/product", getProduct)
	r.GET("/products", getProducts)
	r.GET("/tickets", getTickets)
	r.GET("/results", getTests)
	r.POST("/addproduct", addProduct)
	r.POST("/addticket", addTicket)
	r.POST("/addtest", addTest)
	r.POST("/updateproduct", updateProduct)
	r.Run()
}
