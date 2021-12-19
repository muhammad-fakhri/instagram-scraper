# instagram-scraper

## how to use

1. install [instagram-scraper](https://github.com/arc298/instagram-scraper)
2. use command below to scrap the data. scraped data will be in `/data` directory, and the subdirectory will be the username of scraped instagram profiles

```
instagram-scraper {instagram_username} --profile-metadata --media-metadata -u {account_username} -p {account_password} -d data -n --cookiejar cookie -m 10
```

3. start the app http server, and access it on localhost:8080  
   `go run main.go`
