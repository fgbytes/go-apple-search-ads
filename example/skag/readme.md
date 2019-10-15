# How to create massive SKAGs


## Prerequesites:
- have go installed - https://golang.org/doc/install
- set up your project with the necesary proejct structure
- create a csv file whihc will then create and populate your SKAG
 
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