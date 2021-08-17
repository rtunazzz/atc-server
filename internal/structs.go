package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
)

// Represents a locale
type Locale struct {
	DisplayName         string `json:"display_name"`          // Display name of the locale
	AtcUrl              string `json:"atc_url"`               // Localized ATC URL
	AtcRedirectUrl      string `json:"atc_redirect_url"`      // Localized ATC redirect URL
	WishlistUrl         string `json:"wishlist_url"`          // Localized Wishlist URL
	WishlistRedirectUrl string `json:"wishlist_redirect_url"` // Localized Wishlist redirect URL
}

// Represents a configuration for one site
type SiteConfig struct {
	AtcSupported      bool              `json:"atc_supported"`      // Whether or not atc request type is supported
	WishlistSupported bool              `json:"wishlist_supported"` // Whether or not wishlist request type is supported
	WaitForSubmit     bool              `json:"wait_for_submit"`    // Whether or not JS will wait with redirecting until request was successfully submitted
	Locales           map[string]Locale `json:"locales"`            // Map of locales, there always needs to be at least one `default` one
	RunScript         bool              `json:"run_script"`         // Runs a script instead of the default form handler
	ScriptName        string            `json:"script_name"`        // name of the script to run, the script then need to be localed under /web/scripts/<scriptname>.js
}

// Loads the default config
func loadConfig(filename string) (config map[string]SiteConfig, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	return
}

// Represents formdata
type FormData struct {
	Message       string     // Message which will be displayed to the user before rediredting
	Website       string     // Website the user is being redirected to
	SubmitUrl     string     // Endpoint to which the the form will submit
	WaitForSubmit bool       // Whether or not will JS wait for subission before redirecting
	Redirect      bool       // Whether or not to redirect user after form submit
	RedirectUrl   string     // Where to redirect the user
	Fields        url.Values // Form fields
	RunScript     bool       // Runs a script instead of the default form handler
	ScriptName    string     // Name of the script to run
}
