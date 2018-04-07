package main

import (
	"fmt"
)

func printUser(username string) error {

	if username == "" {
		return fmt.Errorf("user argument required")
	}

	fmt.Printf("\nGetting user: %s\n\n", username)
	usr, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return err
	}

	fmt.Printf("ID: %d\n", usr.ID)
	fmt.Printf("Name: %s\n", usr.GetName())
	fmt.Printf("Login: %s\n", usr.GetLogin())
	fmt.Printf("Email: %s\n", usr.GetEmail())
	fmt.Printf("Location: %s\n", usr.GetLocation())
	fmt.Printf("Created: %v\n", usr.GetCreatedAt())
	fmt.Printf("Company: %s\n", usr.GetCompany())

	fmt.Println()

	return nil
}
