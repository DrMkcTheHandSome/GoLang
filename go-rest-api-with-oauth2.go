package main;

import(
"encoding/json"
"fmt"
"log"
"net/http"
"gorm.io/gorm"
"io/ioutil"
"github.com/gorilla/mux"
"gorm.io/driver/sqlserver"
"golang.org/x/crypto/bcrypt"
"time"
"os"
"golang.org/x/oauth2"
"golang.org/x/oauth2/google"
)

// TO DO: Refactor

// Global Variables
type Product struct {
 gorm.Model
 Code string `gorm:"column:code"`
 Price uint  `gorm:"column:price"`
}

type User struct{
 gorm.Model
 Email string    `json:"email" gorm:"unique"` 
 Password string `json:"password"`
}

var (
    googleOauthConfig *oauth2.Config
    // TODO: randomize it
	oauthStateString = "pseudo-random"
)




func main() { 
    initializeOauth2Configuration()
	handleRequests()
}

func initializeOauth2Configuration(){
     // Setup Google's example test keys
     os.Setenv("CLIENT_ID", "876220489172-i1msr7n6o01anrcanjg3gqj00h08hain.apps.googleusercontent.com")
     os.Setenv("SECRET_KEY", "H6sWMHe-OiBqC1Nd70prnWvB")
    googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:9000/googlecallback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func handleRequests() {
   initializeRoutes()
   fmt.Println("Hello Go!") 
}

func initializeRoutes(){
   initRoutesByGorillaMux()
}

func initRoutesByGorillaMux(){
   myRouter := mux.NewRouter().StrictSlash(true)
   myRouter.HandleFunc("/", homePage)
   myRouter.HandleFunc("/migration", createDatabaseSchema).Methods("POST")
   myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
   myRouter.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
   myRouter.HandleFunc("/products", returnAllProducts).Methods("GET")
   myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
   myRouter.HandleFunc("/product/{id}",returnSingleProduct).Methods("GET")
   myRouter.HandleFunc("/user", createNewUser).Methods("POST")
   myRouter.HandleFunc("/user/loginViaGoogle", loginUserViaGoogle).Methods("GET")
   myRouter.HandleFunc("/user/loginUser", loginUserWithPassword).Methods("POST")
   myRouter.HandleFunc("/users", returnAllUsers).Methods("GET")
   myRouter.HandleFunc("/googlecallback", handleGoogleCallback).Methods("GET")
   log.Fatal(http.ListenAndServe(":9000", myRouter))
}

// LOGIC

func createDatabaseSchema(w http.ResponseWriter, r *http.Request){
    connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
     db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
        if err != nil {
            fmt.Println("failed to connect database") 
            panic("failed to connect database")
        }
     
        // Migrate the schema
        db.Migrator().CreateTable(&Product{})
        db.Migrator().CreateTable(&User{})	
    }

func homePage(w http.ResponseWriter, r *http.Request){
    var htmlIndex = `<html>
<body>
   <h1>Welcome to the homepage!</h1>
	<a href="/user/loginViaGoogle">Google Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
    fmt.Println("Endpoint Hit: homePage")
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewProduct")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
    json.Unmarshal(reqBody, &product)
	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), product.Code,product.Price)
    json.NewEncoder(w).Encode(product)	 
}

func updateProduct(w http.ResponseWriter, r *http.Request){
 fmt.Println("Endpoint Hit: updateProduct")
 
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)
    var product Product 
   //Update multiple columns
    json.Unmarshal(reqBody, &product)
	db.Exec("UPDATE products SET code=?, price = ? WHERE id = ?", product.Code, product.Price, key)
    json.NewEncoder(w).Encode(product)

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
     fmt.Println("Endpoint Hit: returnAllProducts")
	
   connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
    db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		 fmt.Println("failed to connect database") 
        panic("failed to connect database")
     }
	
    // Get all records
	var products []Product
    db.Exec("select * from products").Scan(&products)
	
    json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: deleteProduct")
  
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }

   vars := mux.Vars(r)
    key := vars["id"]
    
   db.Exec("DELETE FROM products WHERE id = ?", key)
   returnAllProducts(w,r)
} 

func returnSingleProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleProduct")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    
	
    vars := mux.Vars(r)
    key := vars["id"]

	var product Product
    db.Exec("select * from products where id = ?",key).Scan(&product)
		
    //Multiple Query Example
    //db.Raw("select code,price from products; drop table product;").First(&product)

    json.NewEncoder(w).Encode(product)  
}

func loginUserViaGoogle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUserViaGoogle")
 
    url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func loginUserWithPassword(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: loginUserWithPassword")
    
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User 
	var userPayload User
	json.Unmarshal(reqBody, &user)
	userPayload = user
	db.Exec("select * from users where email = ?",user.Email).Scan(&user)
	 err = checkPassword(user.Password,userPayload.Password)
	 if err != nil {
	    fmt.Println("Login Failed")
	 } else {
	    fmt.Println("Login Success")
     }
}

func getUserInfo(state string, code string) ([]byte, error) {
    if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.FormValue("state"))
    content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func createNewUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: createNewUser")
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User 
	var hash_password string = ""
    json.Unmarshal(reqBody, &user)
    hash_password = hashPassword(user.Password)
    db.Exec("INSERT INTO users (created_at,email,password) VALUES (?,?,?)",time.Now(), user.Email,hash_password)
    db.Create(&User{Email: user.Email, Password: hash_password})
	user.Password = hash_password
    json.NewEncoder(w).Encode(user)
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllUsers")
	
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
   db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
    if err != nil {
		fmt.Println("failed to connect database") 
        panic("failed to connect database")
    }
	  var users []User
    db.Exec("select * from users").Scan(&users)
      json.NewEncoder(w).Encode(users)
}

func hashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "failed generate bcrypt password"
	}
    
	var hash_password string = ""
	hash_password = string(bytes)

	return hash_password
}


  func checkPassword(userPasswordfromDB,providedPassword string) error {
	  err := bcrypt.CompareHashAndPassword([]byte(userPasswordfromDB), []byte(providedPassword))
	  if err != nil {
		  return err
	  }

	  return nil
 }





