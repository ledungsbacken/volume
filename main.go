package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/itchyny/volume-go"
)

// func readVolume(vol *int) {
// 	v, err := volume.GetVolume()
// 	if err != nil {
// 		panic(err)
// 	}
// 	*vol = v
// 	// fmt.Println(*vol)
// 	time.Sleep(1 * time.Second / 5)
// 	go readVolume(vol)
// }

func main() {
	vol := 0
	// v, err := volume.GetVolume()
	// if err != nil {
	// 	panic(err)
	// }
	// vol = v

	go func() {
		for {
			v, err := volume.GetVolume()
			if err != nil {
				return
			}

			vol = v
			time.Sleep(1 * time.Second / 5)
		}
	}()

	// readVolume(&vol)
	// fmt.Println(vol)

	templateFile := "./index.html"
	templ, _ := template.ParseFiles(templateFile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// queryParams := r.URL.Query()

		// diff := 6
		// if _, ok := queryParams["up"]; ok {
		// 	volume.IncreaseVolume(diff)
		// }
		//
		// if _, ok := queryParams["down"]; ok {
		// 	volume.IncreaseVolume(-diff)
		// }

		templ.Execute(w, vol)
	})

	http.HandleFunc("/volume", func(w http.ResponseWriter, r *http.Request) {
		// v, err := volume.GetVolume()
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("Error"))
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%d", vol)))
	})

	http.HandleFunc("/down", func(w http.ResponseWriter, r *http.Request) {
		if err := volume.IncreaseVolume(-6); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}

		v, _ := volume.GetVolume()
		vol = v

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(v)))
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		if err := volume.IncreaseVolume(6); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}

		v, _ := volume.GetVolume()
		vol = v

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(v)))
	})

	http.ListenAndServe(":8080", nil)
}
