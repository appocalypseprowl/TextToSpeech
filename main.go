package main

import (
	"cloud.google.com/go/texttospeech/apiv1"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"log"
	"net/http"
	"strings"
)

const (
	ContentAPIEndpoint = "https://api.ffx.io/api/content/v0/assets/"
)

func processHTML(s string) string {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return ""
	}

	var traversor func(*html.Node)
	output := ""

	traversor = func(n *html.Node) {
		if n.Type == html.TextNode && n.Parent.Data != "x-placeholder" {
			output += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traversor(c)
		}
	}

	traversor(doc)

	return output
}

func fetchAssetData(endpoint string, id string) (*Asset, error) {
	url := fmt.Sprintf("%s/%s", endpoint, id)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var a Asset
	json.NewDecoder(res.Body).Decode(&a)

	return &a, nil
}

func voiceSelectionFor(ctx context.Context, client *texttospeech.Client, locale string, gender texttospeechpb.SsmlVoiceGender) *texttospeechpb.VoiceSelectionParams {

	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	var localVoices []*texttospeechpb.Voice

	// Get all the localized voices
	for _, voice := range resp.Voices {
		for _, langCode := range voice.LanguageCodes {
			if langCode == locale {
				localVoices = append(localVoices, voice)
			}
		}
	}

	var chosenVoice *texttospeechpb.Voice

	// Find the first matching gender
	for _, voice := range localVoices {
		if voice.SsmlGender == gender {
			chosenVoice = voice
			break
		}
	}

	return &texttospeechpb.VoiceSelectionParams{
		Name:         chosenVoice.Name,
		LanguageCode: locale,
		SsmlGender:   gender,
	}
}

func synthesizeToAudio(text string, locale string, gender texttospeechpb.SsmlVoiceGender) (*texttospeechpb.SynthesizeSpeechResponse, error) {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	selectionParams := voiceSelectionFor(ctx, client, locale, gender)

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: selectionParams,
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	return client.SynthesizeSpeech(ctx, &req)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	if len(paths) < 3 {
		http.NotFound(w, r)
		return
	}

	if paths[1] != "synthesize" {
		http.NotFound(w, r)
		return
	}

	articleId := paths[2]

	asset, err := fetchAssetData(ContentAPIEndpoint, articleId)
	if err != nil {
		log.Println(err)
	}

	articleText := processHTML(asset.Data.Body)

	content, err := synthesizeToAudio(articleText, "en-AU", texttospeechpb.SsmlVoiceGender_MALE)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(content.AudioContent)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Print("Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
