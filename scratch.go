package main

import (
	"cloud.google.com/go/texttospeech/apiv1"
	"context"
	"fmt"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"log"
)

func ListVoices() {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Performs the list voices request.
	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, voice := range resp.Voices {
		// Display the voice's name. Example: tpc-vocoded
		fmt.Printf("Name: %v\n", voice.Name)

		// Display the supported language codes for this voice. Example: "en-US"
		for languageCode := range voice.LanguageCodes {
			fmt.Printf("  Supported language: %v\n", languageCode)
		}

		// Display the SSML Voice Gender.
		fmt.Printf("  SSML Voice Gender: %v\n", voice.SsmlGender.String())

		// Display the natural sample rate hertz for this voice. Example: 24000
		fmt.Printf("  Natural Sample Rate Hertz: %v\n", voice.NaturalSampleRateHertz)
	}
}
