package activity

import (
	"awesomeProject2/payload"
	"context"
	"log"
)

type User payload.User

type PostItem payload.PostItem

// Login activity
func Login(ctx context.Context, user User) error {
	log.Printf("Performing login activity for user: %s\n", user.Username)
	// Perform login logic here
	return nil
}

// PostImage activity
func PostImage(ctx context.Context, post PostItem) error {
	log.Printf("Performing post image activity")
	// Perform post image logic here
	return nil
}

// SharePost activity
func SharePost(ctx context.Context, post PostItem) error {
	log.Printf("Performing share post activity")
	// Perform share post logic here
	return nil
}
