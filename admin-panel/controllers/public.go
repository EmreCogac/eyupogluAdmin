package controllers

import (
	"admin-panel/admin-panel/auth"
	"admin-panel/admin-panel/database"
	"admin-panel/admin-panel/models"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

type ParamID struct {
	Id int `json:"id" binding:"required"`
}

// func Update(c *gin.Context)  {
// 	var ilanlars models.Ilanlar
// 	db := database.GlobalDB

// }

func Delete(c *gin.Context) {
	var ID ParamID
	var ilanlars models.Ilanlar
	db := database.GlobalDB
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id ",
		})

		c.Abort()
		return
	}

	err = ilanlars.DeletePost(ID.Id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			" err ": " error deleting post ",
		})
		c.Abort()
		return
	}
	// hatalı
	result := db.Where("id =? ", ID.Id).Find(&ilanlars)

	if result == nil { // burada yanliş yazmışambir daha deneyelim
		log.Println("cant resolve id")
		c.JSON(500, gin.H{
			" err ": " cant find id ",
		})
		c.Abort()
		return

	}

	c.JSON(200, gin.H{
		"sucsess": " deleted ",
	})

}

func CreatePost(c *gin.Context) {
	var post models.Ilanlar
	err := c.ShouldBindJSON(&post)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"err": "ivalid post require",
		})

		c.Abort()
		return
	}
	err = post.ModelCreatePost()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err": "Error Creating Post	",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully created cpost",
	})

}

// func Signup(c *gin.Context) {
// 	var user models.User
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(400, gin.H{
// 			"Error": "Invalid Inputs ",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	err = user.HashPassword(user.Password)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.JSON(500, gin.H{
// 			"Error": "Error Hashing Password",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	err = user.CreateUserRecord()
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(500, gin.H{
// 			"Error": "Error Creating User",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"Message": "Sucessfully Register",
// 	})
// }

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedtoken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedtoken,
	}
	c.JSON(200, tokenResponse)
}

func GetAll(c *gin.Context) {
	db := database.GlobalDB
	ilanlars := []models.Ilanlar{}
	db.Find(&ilanlars)

	c.JSON(http.StatusOK, gin.H{
		"Posts": ilanlars,
	})
}

// func DeletePost(c *gin.Context) {
// 	db := database.GlobalDB
// 	var id = c.Params.ByName("id")
// 	var person models.Ilanlar
// 	err := db.Where("ID = ?", id).Delete(&person, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(200, gin.H{"id #" + id: "deleted"})

// }
// func DeletePayment(db *gorm.DB, id string) (int64, error) {
// 	var deletedPayment models.Ilanlar
// 	result := db.Where("id = ?", id).Delete(&deletedPayment)
// 	if result.RowsAffected == 0 {
// 		return 0, errors.New("payment data not update")
// 	}
// 	return result.RowsAffected, nil
// }
