package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type data struct {
	Key       string `json:"key"`
	Long_url  string `json:"long_url"`
	Short_url string `json:"short_url"`
}

type input_data struct {
	Url string `json:"url"`
}

func Add_data(w http.ResponseWriter, r *http.Request) {
	var full_data data
	// accesed_url := r.URL.RequestURI()
	// fmt.Println("this is the accesed link: ", r.Host, r.URL.Path)

	if r.Method == "POST" {
		var url_data input_data

		if err := json.NewDecoder(r.Body).Decode(&url_data); err != nil {
			panic(err)
		}
		// handle if the user sends the wrong key
		if url_data.Url == "" {
			http.Error(w, "invalid or missing key 'url'", http.StatusBadRequest)
			return
		}
		full_data.Long_url = url_data.Url
		shorted_link := link_generator(6)
		full_data.Short_url = fmt.Sprintf("http://localhost/%s", shorted_link)
		full_data.Key = shorted_link

		// inserting data
		AddLink(full_data.Key, full_data.Long_url)

		json_full_data, err := json.Marshal(full_data)
		if err != nil {
			panic(err)
		}
		w.Header().Set("content/type", "application/json")
		fmt.Printf("Key : %s, long url: %s, short_url: %s\n", full_data.Key, full_data.Long_url, full_data.Short_url)
		w.Write(json_full_data)
	}

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] //get the key by parsing the path
	link_info, err := GetLink(key)
	// handler for get method
	if r.Method == http.MethodGet {
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link_info.Long_url, http.StatusFound)
		// handler for delete the link
	} else if r.Method == http.MethodDelete {
		row_affect := DeleteLink(key)
		if row_affect == 0 {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

	}

}
