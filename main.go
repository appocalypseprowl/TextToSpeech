package main

import (
	"cloud.google.com/go/storage"
	"cloud.google.com/go/texttospeech/apiv1"
	"fmt"
	"golang.org/x/net/context"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Print("Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("not found")
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")

	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: "Hello, World!"},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(resp.AudioContent)

	writeToStorage(resp.AudioContent)
}

func writeToStorage(audioContent []byte) error {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "cloud-tts-220423.appspot.com"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	fmt.Printf("Bucket %v created.\n", bucketName)

	wc := bucket.Object("article1").NewWriter(ctx)
	wc.ContentType = "audio/mpeg"
	wc.Write(audioContent)

	if err := wc.Close(); err != nil {
		return err
	}

  acl := bucket.Object("article1").ACL()
  if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
    return err
  }

	return nil
}
