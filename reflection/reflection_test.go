package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Bruno"},
			ExpectedCalls: []string{"Bruno"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Bruno", "Cantagalo"},
			ExpectedCalls: []string{"Bruno", "Cantagalo"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Bruno", 27},
			ExpectedCalls: []string{"Bruno"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name: "Bruno",
				Profile: Profile{
					Age:  27,
					City: "Cantagalo",
				},
			},
			ExpectedCalls: []string{"Bruno", "Cantagalo"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name: "Bruno",
				Profile: Profile{
					Age:  27,
					City: "Cantagalo",
				},
			},
			ExpectedCalls: []string{"Bruno", "Cantagalo"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{27, "Cantagalo"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"Cantagalo", "Reykjavík"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{27, "Cantagalo"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"Cantagalo", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
