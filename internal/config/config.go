package config

type ServerConfig struct {
	Port string
	DSN  string
}

const JwtSecret = "AboutVaccine"

const SuccessStatus = 1
const FailureStatus = -1

const UserClaimCookie = "user_token"
