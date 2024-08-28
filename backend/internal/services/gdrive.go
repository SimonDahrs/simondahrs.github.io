package services

import (
	"backend/pkg/models"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const ServiceAccountJsonFile = "personalwebsite-433813-7b35b396c4a8.json" // TODO move to config
const cmsJsonFileID = "1BmOIwnIcKg64znvNPGdat9PhLohBFOiZ"
const cmsJsonFilePath = "cms.json"

func RetrievePosts() []models.Post {
	var posts []models.Post
	if checkPostsIsFresh() {
		posts = retrievePostsFromFile()
	} else {
		posts = retrievePostsFromGdrive()
		go storePosts(posts)
	}
	return posts
}

func checkPostsIsFresh() bool {
	fileInfo, err := os.Stat(cmsJsonFilePath)
	if err != nil {
		return false
	}
	modificationTime := fileInfo.ModTime()
	now := time.Now()
	return now.Sub(modificationTime) <= 60*time.Minute // TODO move to config
}

func retrievePostsFromFile() []models.Post {
	log.Println("Reading cms.json from disk")
	jsonFile, _ := os.Open(cmsJsonFilePath)
	defer jsonFile.Close()
	posts, err := unmarshalPosts(jsonFile)
	if err != nil {
		log.Fatalf("Error unmarshalling posts from file: %v", err)
	}
	return posts
}

func retrievePostsFromGdrive() []models.Post {
	log.Println("Downloading cms.json from Google Drive")
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccountJsonFile), option.WithScopes(drive.DriveScope))
	if err != nil {
		log.Fatalln("Error while creating Drive client")
	}
	resp, err := srv.Files.Get(cmsJsonFileID).Download()

	if resp.StatusCode >= 300 || err != nil {
		log.Fatalf("CMS returned error")
	}
	defer resp.Body.Close()
	posts, err := unmarshalPosts(resp.Body)
	if err != nil {
		log.Fatalf("Error unmarshalling posts from file: %v", err)
	}
	return posts
}

func unmarshalPosts(reader io.Reader) ([]models.Post, error) {
	byteValue, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	err = json.Unmarshal(byteValue, &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func storePosts(posts []models.Post) {
	jsonEnc, _ := json.Marshal(posts)
	os.WriteFile(cmsJsonFilePath, jsonEnc, 0770)
}
