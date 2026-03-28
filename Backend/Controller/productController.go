package controllers

import (
	"context"
	database "ecommerce-backend/Database"
	models "ecommerce-backend/Models"
	utils "ecommerce-backend/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection = database.OpenCollection(database.Client, "users")
var productCollection = database.OpenCollection(database.Client, "products")
var validate = validator.New()

// --- AUTH LOGIC ---

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = utils.HashPassword(user.Password)
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	var user models.User
	c.BindJSON(&req)

	err := userCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil || !utils.VerifyPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.Email, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}

// --- CRUD LOGIC ---

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := productCollection.InsertOne(context.TODO(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	cursor, _ := productCollection.Find(context.TODO(), bson.M{})
	var products []models.Product
	cursor.All(context.TODO(), &products)
	c.JSON(http.StatusOK, products)
}
