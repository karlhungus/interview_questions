package translate

/*Stripe phone screen
4 part question, building on the result of the previous

given input of string of requested languages and the output of supported languages,
 return a list of the supported languages requested.
example: function("en-US, en-GB, fr-FR", ["en-US", "fr-FR"])   should return ["en-US", "fr-FR"]
                  ^ requested             ^ supported                          ^ requested that matched supported
also have to use unittest library in python to do unit tests. that tripped me up as i never used that library before
but first part took 10 mins for me, question was easy, just unittest tripped me up
2) support the case where a more general language is input as string ex: function("en-US, en", ["en-GR", "en-US", "en-CA", "fr-FR") should return ["en-US", "en-GR", "en-CA"] since "en" catches leftover matches of en
3) support wildcard * in input to catch all languages not already in the result
4) support priority, so each input language has a priority attached, want to return the highest priority value first, then the lowest.
function("en-US;q=1, en-CA;q=0, *;q=0.5", ["en-US", "en-CA", "fr-FR", "gr-GR"]) should return
["en-US", "fr-Fr", "gr-GR", "en-CA"]. note order is important, since we want to match en-US first, then any values that match *, then en-CA as last resort, so en-CA has to come last
4 tripped me up and I couldn't think straight, he was nice and gave me ample 20 minutes to think about it. I said to sort it, then add it in like in 1-3 but if already in result remove it and append to back to get the right order. He said I was along the right path but I didn't finish and I'm sure it would have issues.
Anyways good interview, guy was calm and nice. 45 minute coding, 10 minutes for questions that you can ask
*/

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExactMatchListOfLangugaes(t *testing.T) {
	got := languages([]string{"en-US", "en-GB", "fr-FR"}, []string{"en-US", "fr-FR"})
	want := []string{"en-US", "fr-FR"}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestPrefixListOfLanguages(t *testing.T) {
	//function("en-US, en", ["en-GR", "en-US", "en-CA", "fr-FR") should return ["en-US", "en-GR", "en-CA"] since "en" catches leftover matches of en
	got := languages([]string{"en-US", "en"}, []string{"en-GR", "en-US", "en-CA", "fr-FR"})
	want := []string{"en-US", "en-GR", "en-CA"}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestWildCardListOfLanguages(t *testing.T) {
	//function("en-US, en", ["en-GR", "en-US", "en-CA", "fr-FR") should return ["en-US", "en-GR", "en-CA"] since "en" catches leftover matches of en
	got := languages([]string{"en-US", "*"}, []string{"en-GR", "en-US", "en-CA", "fr-FR"})
	want := []string{"en-US", "en-GR", "en-CA", "fr-FR"}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestWithPriority(t *testing.T) {
	//function("en-US;q=1, en-CA;q=0, *;q=0.5", ["en-US", "en-CA", "fr-FR", "gr-GR"]) should return ["en-US", "fr-Fr", "gr-GR", "en-CA"
	got := languagesPriority(PriorityPairs{
		PriorityPair{
			language: "en-US",
			priority: 1,
		},
		PriorityPair{
			language: "en-CA",
			priority: 0.5,
		},
		PriorityPair{
			language: "*",
			priority: 0,
		},
	}, []string{"en-US", "en-CA", "fr-FR", "gr-GR"})
	want := []string{"en-US", "en-CA", "fr-FR", "gr-GR"}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
