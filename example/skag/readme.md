# How to create massive SKAGs


## Prerequesites:
- have go installed - https://golang.org/doc/install

## Set up:
1) in downloads create folder 'skag'
2) copy into  that folder:  main.go, keywords.csv from example file and create a /cert folder for your certificates
3) install search ads package, in your terminal type:
```
go get -u github.com/fgbytes/go-apple-search-ads/
```
4) edit main.go file for your orgID (organisation id- from search ads dashboard)
5) put certificates in /cert folder
6) edit keywords.csv - if your just testing - create a paused campaign and use its id for skag creation
7) run your skag builder. In terminal navigate to your folder:
```
cd ~/Downloads/skag
```
and then
```
go build && ./gog
```

<!-- go get -u github.com/fgbytes/go-apple-search-ads/ -->
 
## Structure
 > /cert - need to put .pem and .key files
 > keywords.csv - file with all your keyword,bid,campaignID
 > main.go - copy from example

## Changes to main.go:
``` 
orgID := int64(1111111)
csvFile, err := os.Open(fmt.Sprintf("%s", "keywords.csv"))
```
Organisation id and path to document with keywords. Organisation ID can be found in apple search ads dashboard.


## Running:
in terminal:
``` $ go build  && ./{{folder_name}}```