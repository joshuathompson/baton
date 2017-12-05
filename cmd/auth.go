package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

const (
	scheme  = "https"
	baseURL = "accounts.spotify.com/"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func getClientCredentials() (id, secret string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Client Id: ")
	scanner.Scan()
	id = scanner.Text()

	fmt.Print("Enter Client Secret: ")
	scanner.Scan()
	secret = scanner.Text()

	return id, secret
}

func printSpotifyAuthorizationURL(id string) {
	v := url.Values{}
	v.Set("client_id", id)
	v.Set("response_type", "code")
	v.Set("redirect_uri", "http://localhost:8080/callback")

	u := &url.URL{
		Scheme:   scheme,
		Path:     baseURL + "authorize",
		RawQuery: v.Encode(),
	}

	fmt.Println(u.String())
}

func getCode() (c string) {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		c = r.URL.Query().Get("code")

		ctx := context.Background()

		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	})

	srv.ListenAndServe()

	return c
}

func getTokens(id, secret, code string) (tokens, error) {
	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("redirect_uri", "http://localhost:8080/callback")

	u := &url.URL{
		Scheme:   scheme,
		Path:     baseURL + "api/token",
		RawQuery: v.Encode(),
	}

	r, err := http.NewRequest("POST", u.String(), nil)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.SetBasicAuth(id, secret)

	fmt.Println(r.URL.String())

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	t := tokens{}

	err = json.NewDecoder(res.Body).Decode(&t)

	return t, err
}

func run(cmd *cobra.Command, args []string) {
	id, secret := getClientCredentials()
	printSpotifyAuthorizationURL(id)
	code := getCode()
	data, err := getTokens(id, secret, code)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "auth",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run:   run,
}
