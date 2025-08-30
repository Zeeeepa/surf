package main

import (
	"github.com/enetx/g"
	"github.com/enetx/surf"
)

func main() {
	const URL = "https://httpbingo.org/anything"

	// with file path
	surf.NewClient().
		FileUpload(URL, "filefield", "/path/to/file.txt").
		// FileUpload(URL, "filefield", "/Users/user/Desktop/1.txt").
		Do()

	// without physical file
	surf.NewClient().
		FileUpload(URL, "filefield", "info.txt", "Hello from surf!").
		Do().Unwrap().Body.String().Print()

	// with multipart data
	multipartData := g.NewMapOrd[string, string]()

	multipartData.Set("_wpcf7", "36484")
	multipartData.Set("_wpcf7_version", "5.4")
	multipartData.Set("_wpcf7_locale", "ru_RU")
	multipartData.Set("_wpcf7_unit_tag", "wpcf7-f36484-o1")
	multipartData.Set("_wpcf7_container_post", "0")
	multipartData.Set("_wpcf7_posted_data_hash", "")
	multipartData.Set("your-name", "name")
	multipartData.Set("retreat", "P48")
	multipartData.Set("your-message", "message")

	// with file path
	surf.NewClient().
		FileUpload(URL, "filefield", "/path/to/file.txt", multipartData).
		Do()

	// without physical file
	r := surf.NewClient().
		FileUpload(URL, "filefield", "info.txt", "Hello from surf again!", multipartData).
		Do().
		Unwrap()

	r.Debug().Request(true).Response(true).Print()
}
