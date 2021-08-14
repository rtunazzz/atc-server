package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var sites map[string]SiteConfig // Global variable for holding all site configurations

// Load config before starting
func init() {
	log.Printf("Loading config")
	c, err := LoadConfig("./config.json")
	if err != nil {
		panic(err)
	}
	sites = c
}

func defaultHandler(title string, description string, err bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, e := template.ParseFiles("./web/templates/default.html")
		if e != nil {
			log.Println(e)
		}

		m := map[string]interface{}{
			"Title":       title,
			"Description": description,
			"IsError":     err,
		}

		if err {
			w.WriteHeader(http.StatusNotFound)
		}

		exer := t.Execute(w, m)
		if exer != nil {
			log.Println(exer)
		}
	}
}

// Handler for serving a template for autosubmitting a form
func submitFormHandler(formType string) http.HandlerFunc {
	form := strings.ToLower(formType)
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		sitesParam, _ := q["site"]
		currentSite := strings.ToLower(sitesParam[0])

		if form == "addtocart" && !sites[currentSite].AtcSupported {
			errorPage := defaultHandler("Not supported", fmt.Sprintf("Sorry, %s is not supported.", strings.Title(currentSite)), true)
			errorPage(w, r)
			return
		} else if form == "wishlist" && !sites[currentSite].WishlistSupported {
			errorPage := defaultHandler("Not supported", fmt.Sprintf("Sorry, %s is not supported.", strings.Title(currentSite)), true)
			errorPage(w, r)
			return
		}

		// check if cookie exists for the current site
		cookie, cErr := r.Cookie(currentSite + "_locale")

		// if it doesn't, set the cookie to 'default'
		if cErr != nil {
			newCookie := http.Cookie{Name: currentSite + "_locale", Value: "default", Expires: time.Now().Add(365 * 24 * time.Hour)}
			cookie = &newCookie
		}

		localeData := sites[currentSite].Locales[cookie.Value]

		var formUrl, redirectUrl string
		if form == "addtocart" {
			formUrl = localeData.AtcUrl
			redirectUrl = localeData.AtcRedirectUrl
		} else if form == "wishlist" {
			formUrl = localeData.WishlistUrl
			redirectUrl = localeData.WishlistRedirectUrl
		}

		if formUrl == "" {
			log.Printf("ERROR: Form URL for site %s and locale %s not found!", currentSite, localeData.DisplayName)
			errorPage := defaultHandler("Something went wrong.", "Internal server error. Please report this to the administrator.", true)
			errorPage(w, r)
			return
		}

		t, err := template.ParseFiles("./web/templates/form.html")
		if err != nil {
			log.Println(err)
			errorPage := defaultHandler("Something went wrong.", "Internal server error. Please report this to the administrator.", true)
			errorPage(w, r)
			return
		}

		q.Del("site")

		// add some form fields if we're ATCing on snipes/ sbx/ onygo
		if (form == "addtocart") && (currentSite == "snipes" || currentSite == "solebox" || currentSite == "onygo") {
			quantities, okQuant := q["quantity"]
			if !okQuant || len(quantities[0]) < 1 {
				q.Add("quantity", "1")
			}

			options, okOpt := q["options"]
			if !okOpt || len(options[0]) < 1 {
				q.Add("options", "")
			}
		}

		msg := "Adding to cart"
		if form == "wishlist" {
			msg = "Adding to wishlist"
		}

		exer := t.Execute(w, FormData{
			Message:       msg,
			Website:       strings.Title(currentSite),
			SubmitUrl:     formUrl,
			Redirect:      redirectUrl != "",
			RedirectUrl:   redirectUrl,
			WaitForSubmit: sites[currentSite].WaitForSubmit,
			Fields:        q,
		})

		if exer != nil {
			log.Println(exer)
			errorPage := defaultHandler("Something went wrong.", "Internal server error. Please report this to the administrator.", true)
			errorPage(w, r)
		}
	}
}

// Handler for setting a locale
func setLocaleHandler(site string) http.HandlerFunc {
	siteLower := strings.ToLower(site)
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./web/templates/setlocale.html")
		if err != nil {
			log.Println(err)
			errorPage := defaultHandler("Something went wrong.", "Internal server error. Please report this to the administrator.", true)
			errorPage(w, r)
			return
		}

		redirectBase := "https://" + siteLower + ".com/"
		if siteLower == "snipes" {
			redirectBase = "https://" + siteLower + "."
		}

		// Create a map of the found locales
		var localeMap = make(map[string]string)
		for symbol, locale := range sites[siteLower].Locales {
			if symbol != "default" {
				localeMap[locale.DisplayName] = symbol
			}
		}

		m := map[string]interface{}{
			"Site":         strings.Title(siteLower),
			"CookieName":   siteLower + "_locale",
			"RedirectBase": redirectBase,
			"RedirectURI":  "/login",
			"Locales":      localeMap,
		}

		t.Execute(w, m)
	}
}

// Handler for autoredirecting
func redirectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./web/templates/redirect.html")
		if err != nil {
			log.Println(err)
			errorPage := defaultHandler("Something went wrong.", "Internal server error. Please report this to the administrator.", true)
			errorPage(w, r)
			return
		}

		q := r.URL.Query()
		urls, _ := q["url"]

		m := map[string]interface{}{
			"Title": "Redirecting...",
			"URL":   urls[0],
		}
		t.Execute(w, m)
	}
}
