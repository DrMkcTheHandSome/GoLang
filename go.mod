module main

go 1.15

require (
	docs v1.2.3
	entities v1.2.3
	controllers v1.2.3
	repositories v1.2.3
	services v1.2.3
	connections v1.2.3
	constants v1.2.3
	github.com/gorilla/mux v1.8.0
	github.com/swaggo/http-swagger v1.0.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3
	gorm.io/driver/sqlserver v1.0.5
	gorm.io/gorm v1.20.11
)

replace docs v1.2.3 => ./docs
replace entities v1.2.3 => ./data/entities
replace controllers v1.2.3 => ./presentation/controllers
replace services v1.2.3 => ./business/services
replace repositories v1.2.3 => ./data/repositories
replace connections v1.2.3 => ./data/connections
replace constants v1.2.3 => ./business/constants


