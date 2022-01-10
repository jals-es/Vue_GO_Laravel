package users

import (
	"net/http"
	"appbar/common"
	"errors"
	"fmt"
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado correctamente"})
}

func Login(c *gin.Context){
	userModelValidator := NewLoginUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userModel, err := FindOne(&UserModel{Email: userModelValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(userModelValidator.User.Passwd) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	// fmt.Println(userModel.ID)
	// fmt.Println(userModel.Email)
	// fmt.Println(userModel.Passwd)

	erro := CheckSuperAdmin(userModel.ID)

	if erro == nil {
		fmt.Println("is superadmin")

		postBody, _ := json.Marshal(map[string]string{
			"user":  userModel.Email,
			"passwd": userModel.Passwd,
		})
		responseBody := bytes.NewBuffer(postBody)
	  	//Leverage Go's HTTP Post function to make request
		resp, err := http.Post("http://127.0.0.1:8000/api/auth", "application/json", responseBody)
	  	//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string([]byte(body))
		
		fmt.Println(sb)

		var raw map[string]string
		json.Unmarshal([]byte(sb), &raw)

		UpdateContextUserModel(c, userModel.ID)
		serializer := UserSerializer{c}
		c.JSON(http.StatusCreated, gin.H{ "user": serializer.SuperResponse(raw["data"])})
	}else{
		UpdateContextUserModel(c, userModel.ID)
		serializer := UserSerializer{c}
		c.JSON(http.StatusCreated, gin.H{ "user": serializer.Response()})
	}
}

func CheckToken(c *gin.Context){
	// myUserModel := c.MustGet("my_user_model").(UserModel)

	c.JSON(http.StatusCreated, gin.H{ "data": "Usuario verificado" })
}