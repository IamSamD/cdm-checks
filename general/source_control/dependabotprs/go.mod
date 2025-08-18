module dependabotprs

go 1.23.0

toolchain go1.23.12

require (
	github.com/google/go-github/v74 v74.0.0
	github.com/iamsamd/cdm_framework v0.0.1
)

require (
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)

// replace github.com/iamsamd/cdm_framework => ../../../../cdm_framework // for local development - to be removed on release v1 of cdm_framework
