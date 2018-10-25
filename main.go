package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"cloud.google.com/go/texttospeech/apiv1"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func main() {
	router := httprouter.New()
	router.GET("/synthesize/:articleId", handle)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	fileName := fmt.Sprintf("%s%s", ps.ByName("articleId"), ".mp3")
	writeToStorage(fileName, resp.AudioContent)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func writeToStorage(fileName string, audioContent []byte) error {
	log.Print("YAY START!")
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

	log.Print("Bucket created.\n", bucketName)

	wc := bucket.Object(fileName).NewWriter(ctx)
	wc.ContentType = "audio/mpeg"
	wc.Write(audioContent)

	if err := wc.Close(); err != nil {
		return err
	}

	acl := bucket.Object(fileName).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil
}
