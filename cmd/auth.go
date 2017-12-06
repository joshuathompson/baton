package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func getClientCredentials() (id, secret string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nFollow these instructions to authenticate the Baton CLI to change your songs, volume, etc:\n" +
		"1. Go to https://beta.developer.spotify.com/dashboard\n" +
		"2. Log in with your Spotify username/password\n" +
		"3. Create a new app\n" +
		"4. Click the newly created app\n" +
		"5. Click 'Edit Settings'\n" +
		"6. Add 'http://localhost:15298/callback' as a redirect URI, don't forget to save\n" +
		"7. Copy the Client Id and Client Secret\n" +
		"8. Input the items as the CLI asks for them\n")

	fmt.Print("Enter Client Id: ")
	scanner.Scan()
	id = scanner.Text()

	fmt.Print("Enter Client Secret: ")
	scanner.Scan()
	secret = scanner.Text()

	return id, secret
}

func serverManager(srv *http.Server, keepAlive chan bool) {
	for {
		select {
		case <-keepAlive:
			ctx := context.Background()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal(err)
			}

			return
		default:
		}
	}
}

func getCode(id string) (c string) {
	m := api.GetAuthorizationURL(id)
	fmt.Printf("\nNavigate to the following URL to Authorize Baton:\n%s\n", m)
	keepAlive := make(chan bool)

	srv := &http.Server{Addr: ":15298"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		keepAlive <- true
		fmt.Fprintf(w, "<h1>Almost done!</h1><p>Baton has been approved, just copy the following code back to the CLI: <span style=\"color: #FF0000\">%s</span></p>", code)
	})

	go serverManager(srv, keepAlive)
	srv.ListenAndServe()

	fmt.Print("\nEnter Code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c = scanner.Text()

	return c
}

func authenticate(cmd *cobra.Command, args []string) {
	id, secret := getClientCredentials()
	code := getCode(id)
	api.AuthorizeWithCode(id, secret, code)
	fmt.Println("\nAuthentication successful!")
}

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize Baton to access API on your behalf",
	Long:  `Authorize Baton to access the Spotify API on your behalf by obtaining a long-lasting refresh token using your client_id, client_secret, and approval`,
	Run:   authenticate,
}
