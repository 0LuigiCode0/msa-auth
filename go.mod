module x-msa-auth

go 1.16

replace x-msa-core v0.0.0 => ./modules/x-msa-core

require (
	github.com/0LuigiCode0/go-utill v1.0.9
	github.com/0LuigiCode0/logger v1.1.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.7.0
	x-msa-core v0.0.0
)
