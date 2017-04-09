package main

import (
	"fmt"

	"github.com/iopred/discordgo"
)

// GetUserInfo takes an OAuth2 generated token and returns the user's info
func GetUserInfo(token string) (*discordgo.User, error) {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil, err
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
		return nil, err
	}

	return u, nil
}
