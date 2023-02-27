package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Data struct {
	Project string
	Region  string

	GaeApplication string
	GaeVersion     string

	CrService  string
	CrRevision string
}

func main() {
	// Get project ID from metadata server
	project := ""
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/project/project-id", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			responseBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			project = string(responseBody)
		}
	}

	// Get region from metadata server
	region := ""
	req, _ = http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/region", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err = client.Do(req)
	if err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			responseBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			region = regexp.MustCompile(`projects/[^/]*/regions/`).ReplaceAllString(string(responseBody), "")
		}
	}

	data := Data{
		Project: project,
		Region:  region,

		GaeApplication: os.Getenv("GAE_APPLICATION"),
		GaeVersion:     os.Getenv("GAE_VERSION"),

		CrService:  os.Getenv("K_SERVICE"),
		CrRevision: os.Getenv("K_REVISION"),
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Printf("The server started successfully and is listening for HTTP requests on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
