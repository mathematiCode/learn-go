package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but didn't get one.")
	}

	if err.Error() != want {
		t.Errorf("got error: %s but wanted error: %s", err.Error(), want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one. Error: %s", err.Error())
	}
}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known-word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)

		assertNoError(t, err)
	})

	t.Run("unknown-word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound.Error()

		assertError(t, err, want)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		addErr := d.Add("test", "this is just a test")

		assertNoError(t, addErr)
		got, err := d.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{"existing": "existing definition"}
		err := d.Add("existing", "this is just a test")

		assertError(t, err, ErrWordExists.Error())

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update definition of existing word", func(t *testing.T) {
		dictionary := Dictionary{"existing": "I'm a definition"}

		err := dictionary.Update("existing", "new definition")

		assertNoError(t, err)
	})

	t.Run("update definition of a word not in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"existing": "I'm a definition"}

		err := dictionary.Update("new", "new definition")

		assertError(t, err, ErrWordDoesNotExist.Error())
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "test definition"}
		err := dictionary.Delete("test")

		assertNoError(t, err)
	})

	t.Run("delete a nonexistent word", func(t *testing.T) {
		dictionary := Dictionary{"test": "test definition"}
		err := dictionary.Delete("unknown")

		assertError(t, err, ErrWordDoesNotExist.Error())
	})

}
