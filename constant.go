package main

import (
	"errors"
)

const (
	ScrapCommand                   = "instagram-scraper"
	ScrapCommandArgProfileMetadata = "--profile-metadata"
	ScrapCommandArgMediaMetadata   = "--media-metadata"
	ScrapCommandArgAccountUsername = "-u"
	ScrapCommandArgAccountPassword = "-p"
	ScrapCommandArgDataDir         = "-d"
	ScrapCommandArgUsernameSubdir  = "-n"
	ScrapCommandArgCookie          = "--cookiejar"
	ScrapCommandArgLimit           = "-m"
)

// errors
var ErrReqQueryInvalid error = errors.New("request query invalid")
