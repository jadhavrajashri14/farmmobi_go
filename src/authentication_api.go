package authentication_api

import (
  "os"
  "config"
  "models"
  "entities"
  "encoding/json"
  "net/http"
  "fmt"
  "log"
  "time"
  "github.com/dgrijalva/jwt-go"
  "github.com/davecgh/go-spew/spew"
  "golang.org/x/crypto/bcrypt"
  "net/smtp"
  "html/template"
  "bytes"
  "strconv"
  "github.com/gorilla/mux"
  "encoding/base64"
)


func GetUsers(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetUsers")

  db, err := config.GetDB()

  if err != nil {
  	respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

  	userModel := models.UserModel{
  		Db: db,
  	}
  	users, err2 := userModel.GetUsers()

  	if err2 != nil {
  		respondWithError(response, http.StatusBadRequest, err2.Error())
  	} else {
  		respondWithJson(response, http.StatusOK, users)
  	}
  }
}


func GetUserByID(response http.ResponseWriter, request *http.Request) {

  fmt.Println("API: GetUserByID")
  //var id uint64
  var myerr entities.Error

  vars := mux.Vars(request)
  keyword := vars["keyword"]

  db, err := config.GetDB()

  if err != nil {
    respondWithError(response, http.StatusBadRequest, err.Error())
  } else {

    defer db.Close()

    userModel := models.UserModel{
      Db: db,
    }

    id, err := strconv.ParseUint(fmt.Sprintf("%.f", keyword), 10, 64)

    if err != nil {

     	myerr.Message = "Error while getting user id"
	 	respondWithError(response, http.StatusBadRequest, myerr.Message)
    } else {

    	users, err2 := userModel.GetUserByID(id)

	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, err2.Error())
	    } else {
	      respondWithJson(response, http.StatusOK, users)
	    }
    }
  }
}


func Signup(response http.ResponseWriter, request *http.Request) {

	var user entities.User
	var error entities.Error
	var confirmEmailUrl, absurl string
	var cur_user_id		int64
	var keyword *string


  	err := json.NewDecoder(request.Body).Decode(&user)

  	//spew.Dump(user)

  	db, err := config.GetDB()

    if err != nil {
    	respondWithError(response, http.StatusBadRequest, err.Error())
  	} else {

	    defer db.Close()

	    userModel := models.UserModel{
	      Db: db,
	    }

	    fmt.Println("user.Email : ", *user.Email)
	    keyword = user.Email

	    // check if user with this email is already existing in database,
	    user_id, err2 := userModel.SearchUserEmailExists(*keyword)
	    
	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, err2.Error())
	    } else {

	    	// if user email is already in database,
	    	if user_id != nil  {

	    		error.Message = "Email alreay in use"
	    		respondWithError(response, http.StatusIMUsed, error.Message)

	    	} else {


	    		user_id, err33 := userModel.GetMaxUserId()
	    		if err33!= nil {
			      respondWithError(response, http.StatusBadRequest, err2.Error())
			    } else {

			    	// Creating the user for Signup
			    	fmt.Println("max user_id : ", *user_id)
			    	cur_user_id = (*user_id)+1
		    		user.User_id = &cur_user_id

		    		hash, err3 := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

			    	if err3 != nil {
			    		respondWithError(response, http.StatusBadRequest, err3.Error())
			    	}

			    	//Hashing user password to store in the DB
			    	fmt.Println(hash)

			    	user.Password = string(hash)
			    	user.Is_staff = true
			    	fmt.Println("user.Password after hash : ", user.Password)
		    		
				    err4 := userModel.Create(&user)

				    if err4 != nil {

				      respondWithError(response, http.StatusInternalServerError, err4.Error())
				      return

				    } else {

				    	token, err5 := GenerateToken(user)

						if err5 != nil {
							log.Fatal(err5)
							respondWithError(response, http.StatusBadRequest, err5.Error())
						}

						fmt.Println("register token : ",token)
						
						// sending email to confirm email address
						// comes in the request header as frontendurl to send the email to the user
						confirmEmailUrl = request.Header.Get("confirmEmailUrl")
						fmt.Println("confirmEmailUrl", confirmEmailUrl)

						absurl = confirmEmailUrl+"/?token="+string(token)
						fmt.Println("absurl => " + absurl)

						emailTemplate(absurl, user.Email, "templates/cnfrmemailtemplate.html")

						fmt.Println("returning token = ", token)

				      	respondWithJson(response, http.StatusCreated, token)
				    }
			    }
	    	}
	    }
  	}
}


func emailTemplate(Url string, ToUser *string, templateLocation string) () {

	    // Sender data.
	  from := os.Getenv("FromEmailAddr")
	  password := os.Getenv("SMTPPwd")

	  // Receiver email address.
	  to_user := []string{
	    *ToUser,
	  }

	 // smtp server configuration.
	  smtpHost := "smtp.gmail.com"
	  smtpPort := "587"

	  // Authentication.
	  auth := smtp.PlainAuth("", from, password, smtpHost)

	  templateData := struct {
			Name string
			URL  string
		}{
			Name: *ToUser,
			URL:  Url,
		}

	  t, err := template.ParseFiles(templateLocation)
	  if err != nil {
	  	fmt.Println(err)
	  	return
	  }

	  var body bytes.Buffer

	  mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	  body.Write([]byte(fmt.Sprintf("Subject: Verify your email \n%s\n\n", mimeHeaders)))

	  t.Execute(&body, templateData)

	  fmt.Println("html/template => ", string(body.Bytes()))

	  // Sending email.
	  err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to_user, body.Bytes())

	  if err != nil {
	    fmt.Println(err)
	    return
	  }

	  fmt.Println("Email Sent Successfully!")
	  return
}


func GenerateToken(user entities.User) (string, error) {
	var err error
	secret := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.User_id,
		"username" : user.Username,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
	})

	spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}


func VerifyEmail(response http.ResponseWriter, request *http.Request) {

	var myerr entities.Error

    fmt.Println("Endpoint Hit: VerifyEmail")

    requestToken := request.Header.Get("Authorization")

    //requestToken := request.URL.Query().Get("token")
    fmt.Println("requestToken : ", requestToken)

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		 //Make sure that the token method conform to "SigningMethodHMAC"
		 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		 	myerr.Message = "unexpected signing method with token"
		 	respondWithError(response, http.StatusBadRequest, myerr.Message)
		 }
		 return []byte(os.Getenv("SECRET_KEY")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		 myerr.Message = "Activation Link expired"
		 respondWithError(response, http.StatusBadRequest, myerr.Message)
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		myerr.Message = "Invalid Token"
	 	respondWithError(response, http.StatusBadRequest, myerr.Message)
	}

	//Since token is valid, get the user_id:
  	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

  	if ok && token.Valid {

  		userName := claims["username"].(string)

	     userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userid"]), 10, 64)

	     if err != nil {

	     	myerr.Message = "Error while extracting token details"
		 	respondWithError(response, http.StatusBadRequest, myerr.Message)
	     }

		    fmt.Println("userid, username from token : ", userId, userName)


		    db, err := config.GetDB()

			  if err != nil {
			  	respondWithError(response, http.StatusBadRequest, err.Error())
			  } else {

			defer db.Close()

		  	userModel := models.UserModel{
		  		Db: db,
		  	}

		    // check if user with this email is already existing in database,
			user, err2 := userModel.GetUserByID(userId)

			if err2 != nil {
				respondWithError(response, http.StatusBadRequest, err2.Error())
			} else {

				if user != nil {

					// User present in the database so making it as is_verified = true & is_active=true

					_, err3 := userModel.UpdateFlags(uint64(userId))

				    if err3 != nil {
				      respondWithError(response, http.StatusBadRequest, err3.Error())
				    } else {
				      respondWithJson(response, http.StatusOK, "User Successfully Activated")
				    }
				}
			}
		}
	}
}


func GenerateTokenPair(user entities.User) (map[string]string, error) {
	var err error
	secret := os.Getenv("SECRET_KEY")

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = user.User_id
	claims["username"] = user.Username
	claims["is_admin"] = user.Is_admin
	claims["is_superuser"] = user.Is_superuser
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["userid"] = user.User_id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	fmt.Println("----------------------------------")
	fmt.Println("GenerateTokenPair => access_token : ", t, "refresh_token : ", rt)
	fmt.Println("----------------------------------")


	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil

}


// This is the api to refresh tokens
func TokenRefresh(response http.ResponseWriter, request *http.Request) {

    vars := mux.Vars(request)
    keyword := vars["refresh_token"]

	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	tokenReq.RefreshToken = keyword

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("secret"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Get the user record from database or
			// run through your business logic to verify if the user can log in

			db, err := config.GetDB()

			if err != nil {
				respondWithError(response, http.StatusBadRequest, err.Error())
				return
			} else {

				defer db.Close()

				userModel := models.UserModel{
					Db: db,
				}

				userId := uint64(claims["userid"].(float64))

				// check if user with this email is already existing in database,
				user, err2 := userModel.GetUserByID(userId)

				if err2 != nil {
				respondWithError(response, http.StatusBadRequest, err2.Error())
				} else {

				if user != nil {

					if int(claims["sub"].(float64)) == 1 {

						newTokenPair, err := GenerateTokenPair(user[0])
						if err != nil {
							return
						}

						respondWithJson(response, http.StatusOK, newTokenPair)
						return
					}
					
				} else {
					respondWithError(response, http.StatusBadRequest, "Unauthorized to refresh tokens")
					return
				}
			}
			respondWithError(response, http.StatusBadRequest, "Unable to refresh tokens")
			return
		}
	} else {
		if err != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		}
	}
}


func Authenticate(reqUser *entities.User, user *entities.User) bool {

	//fmt.Println("reqUser.Password : ", reqUser.Password)
	fmt.Println("user.Password : ", user.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		fmt.Println("password hashes are not same")
		return false
	}
	return true

}


func Login(response http.ResponseWriter, request *http.Request) {

	fmt.Println("API: Login")
	var reqUser entities.User

  	err := json.NewDecoder(request.Body).Decode(&reqUser)

  	//spew.Dump(reqUser)

  	db, err := config.GetDB()

    if err != nil {
    	respondWithError(response, http.StatusBadRequest, err.Error())
  	} else {

	    defer db.Close()

	    userModel := models.UserModel{
	      Db: db,
	    }

	    keyword := reqUser.Username

	    // check if user with this email is already existing in database,
	    user, err2 := userModel.GetUserByName(keyword)
	    
	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, "User does not exist")
	    } else {

	    	if user != nil {

	    		// User in the DB so better authenticate the user
		    	valid := Authenticate(&reqUser, &user[0])

			  	if !valid {
			  		respondWithError(response, http.StatusBadRequest, "Invalid Password, try again")
			  		return
			  	}

			  	active := user[0].Is_active

			  	if !active {
			  		respondWithError(response, http.StatusBadRequest, "Account disabled, contact admin")
			  		return
			  	}

			  	verified := user[0].Is_verified

			  	if !verified {
			  		respondWithError(response, http.StatusBadRequest, "Email is not verified")
			  		return
			  	}

				tokenmap, err := GenerateTokenPair(user[0])

				if err != nil {
					log.Fatal(err)
				}

				//fmt.Println("tokenmap : ", tokenmap)

				accessToken := tokenmap["access_token"]
				refreshToken := tokenmap["refresh_token"]

				tokens := "access_token:" + accessToken + " refresh_token:" + refreshToken

				fmt.Println("tokens : ", tokens)
				respondWithJson(response, http.StatusOK, tokens)
				return
	    	} else {
	    		// User does not exist in DB
	    		respondWithError(response, http.StatusBadRequest, "User does not exist")
	    		return
	    	}
	    }
	}
}


func RequestPasswordResetEmail(response http.ResponseWriter, request *http.Request) {

	fmt.Println("API: RequestPasswordResetEmail")

	changePasswordUrl := request.Header.Get("changePasswordUrl")
    fmt.Println("changePasswordUrl : ", changePasswordUrl)

    var emailAddress map[string]string

  	json.NewDecoder(request.Body).Decode(&emailAddress)
  	spew.Dump(emailAddress)
  	requestEmail := emailAddress["email"]
  	fmt.Println("reset password email = ", requestEmail)

  	db, err := config.GetDB()

    if err != nil {
    	respondWithError(response, http.StatusBadRequest, err.Error())
  	} else {

	    defer db.Close()

	    userModel := models.UserModel{
	      Db: db,
	    }

	    // check if user with this email is already existing in database,
	    user_id, err2 := userModel.SearchUserEmailExists(requestEmail)
	    
	    if err2 != nil {
	      respondWithError(response, http.StatusBadRequest, err2.Error())
	    } else {

	    	if user_id != nil {

				secret := os.Getenv("SECRET_KEY")

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"userid": *user_id,
					"email" : requestEmail,
					"exp" : time.Now().Add(time.Hour * 72).Unix(),
				})

				spew.Dump(token)

				tokenString, err := token.SignedString([]byte(secret))

				if err != nil {
					respondWithError(response, http.StatusBadRequest, "Error while getting token")
				} else {

					// String to encode
    				intVal := *user_id

    				str := strconv.FormatInt(intVal, 10)

    				// base64.StdEncoding: Standard encoding with padding
				    // It requires a byte slice so we cast the string to []byte
				    encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
				    fmt.Println("Encoded uidb64:", encodedStr)

					absurl := changePasswordUrl+ "?uidb64=" + encodedStr + "&token=" + string(tokenString)
					fmt.Println("absurl => " + absurl)

					emailTemplate(absurl, &requestEmail, "templates/resetpwdtemplate.html")

			      	respondWithJson(response, http.StatusOK, "Check Email to Reset Password")
				}
	    	}
		}
	}
}


func PasswordTokenCheck(response http.ResponseWriter, request *http.Request) {

	fmt.Println("API: PasswordTokenCheck")
	var myerr entities.Error

	vars := mux.Vars(request)
    uidb64 := vars["uidb64"]
    checkToken := vars["token"]

    fmt.Println("URL : ", request.URL.String())
    fmt.Println("uidb64 : ", uidb64)
    fmt.Println("token : ", checkToken)

	// Decoding may return an error, in case the input is not well formed
    decodedStr, err := base64.StdEncoding.DecodeString(uidb64)
    if err != nil {
        respondWithError(response, http.StatusBadRequest, "Error while decoding uidb64")
    }
    fmt.Println("Decoded uidb64:", string(decodedStr))

    byteToInt, err := strconv.Atoi(string(decodedStr))

    if err != nil {
    	respondWithError(response, http.StatusBadRequest, err.Error())
    	return
    }

    userId := uint64(byteToInt)

    fmt.Println("userId:", userId)

    db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {

		defer db.Close()

		userModel := models.UserModel{
			Db: db,
		}

		// check if user with this email is already existing in database,
		user, err2 := userModel.GetUserByID(userId)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {

			if user != nil {

				token, err := jwt.Parse(checkToken, func(token *jwt.Token) (interface{}, error) {
				 //Make sure that the token method conform to "SigningMethodHMAC"
					 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					 	myerr.Message = "unexpected signing method with token"
					 	respondWithError(response, http.StatusBadRequest, myerr.Message)
					 }
					 return []byte(os.Getenv("SECRET_KEY")), nil
				})

				//if there is an error, the token must have expired
				if err != nil {
					 myerr.Message = "Token expired, please request a new one"
					 respondWithError(response, http.StatusBadRequest, myerr.Message)
				}

				//is token valid?
				if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
					myerr.Message = "Invalid Token, please request a new one"
				 	respondWithError(response, http.StatusBadRequest, myerr.Message)
				}

				//Since token is valid, get the user_id:
				claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

				if ok && token.Valid {

					emailAddress := claims["email"].(string)
			     	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userid"]), 10, 64)

				     if err != nil {

				     	myerr.Message = "Error while extracting token details"
					 	respondWithError(response, http.StatusBadRequest, myerr.Message)
				     }

					fmt.Println("userid, emailAddress from token : ", userId, emailAddress)
					fmt.Println("PasswordTokenCheck success : Credentials Valid")
					respondWithJson(response, http.StatusOK, "Credentials valid")
					return
				}
			}
		}
	}
}



func SetNewPassword(response http.ResponseWriter, request *http.Request) {

	fmt.Println("API: SetNewPassword")
	var myerr entities.Error
	var requestData map[string]string

  	json.NewDecoder(request.Body).Decode(&requestData)

  	uidb64 := requestData["uidb64"]
  	checkToken := requestData["token"]
  	password := requestData["password"]


    fmt.Println("URL : ", request.URL.String())
    fmt.Println("uidb64 : ", uidb64)
    fmt.Println("token : ", checkToken)
    fmt.Println("password : ", password)


    // Decoding may return an error, in case the input is not well formed
    decodedStr, err := base64.StdEncoding.DecodeString(uidb64)
    if err != nil {
        respondWithError(response, http.StatusBadRequest, "Error while decoding uidb64")
    }
    fmt.Println("Decoded uidb64:", string(decodedStr))

    byteToInt, err := strconv.Atoi(string(decodedStr))

    if err != nil {
    	respondWithError(response, http.StatusBadRequest, err.Error())
    	return
    }

    userId := uint64(byteToInt)
    fmt.Println("userId:", userId)

    db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {

		defer db.Close()

		userModel := models.UserModel{
			Db: db,
		}

		// check if user with this email is already existing in database,
		user, err2 := userModel.GetUserByID(userId)

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {

			if user != nil {

				token, err := jwt.Parse(checkToken, func(token *jwt.Token) (interface{}, error) {
				 //Make sure that the token method conform to "SigningMethodHMAC"
					 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					 	myerr.Message = "unexpected signing method with token"
					 	respondWithError(response, http.StatusBadRequest, myerr.Message)
					 }
					 return []byte(os.Getenv("SECRET_KEY")), nil
				})

				//if there is an error, the token must have expired
				if err != nil {
					 myerr.Message = "Token expired, please request a new one"
					 respondWithError(response, http.StatusBadRequest, myerr.Message)
				}

				//is token valid?
				if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
					myerr.Message = "Invalid Token, please request a new one"
				 	respondWithError(response, http.StatusBadRequest, myerr.Message)
				}

				//Since token is valid, get the user_id:
				claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

				if ok && token.Valid {

					userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userid"]), 10, 64)

				     if err != nil {

				     	myerr.Message = "Error while extracting token details"
					 	respondWithError(response, http.StatusBadRequest, myerr.Message)
				     }

					fmt.Println("userid from token : ", userId)

					hash, err3 := bcrypt.GenerateFromPassword([]byte(password), 10)

			    	if err3 != nil {
			    		respondWithError(response, http.StatusBadRequest, err3.Error())
			    	}

			    	//Hashing user password to store in the DB
			    	fmt.Println("user.Password before hash : ", password)
			    	user[0].Password = string(hash)
			    	fmt.Println("user.Password after hash : ", user[0].Password)
		    		
				    _, err4 := userModel.UpdatePwd(userId, user[0].Password)

				    if err4 != nil {

				      respondWithError(response, http.StatusInternalServerError, err4.Error())
				      return

				    } else {
				    	fmt.Println("Password reset Successfully")
				    	respondWithJson(response, http.StatusOK, "Password reset Successfully")
				    }
				}
			}
		}
	}
}


func Get_User_Menu(response http.ResponseWriter, request *http.Request) {

	fmt.Println("API: Get_User_Menu")

	vars := mux.Vars(request)
    userId := vars["user_id"]

  	id, _ := strconv.ParseInt(userId, 10, 64)

	db, err := config.GetDB()

	if err != nil {
	respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		defer db.Close()

		fmt.Println("Getting User Profile Menu")

		profileMenuModel := models.ProfileMenuModel{
  			Db: db,
  		}

		menuItems, err2 := profileMenuModel.Get_User_Menu(id)

	  	if err2 != nil {
	  		respondWithError(response, http.StatusBadRequest, err2.Error())
	  	} else {
	  		fmt.Println("menuItems : ", menuItems)
	  		respondWithJson(response, http.StatusOK, menuItems)
	  	}
	}
}


func ProtectedEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("ProtectedEndpoint invoked")
}


func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked")
	return nil
}


func respondWithError(w http.ResponseWriter, code int, msg string) {

	respondWithJson(w, code, map[string]string{"error": msg})
}


func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
}
