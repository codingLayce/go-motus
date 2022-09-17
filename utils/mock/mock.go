package mock

import "testing"

func Do[T any](t *testing.T, f1 *T, f2 T) {
	origin := *f1

	*f1 = f2

	t.Cleanup(func() {
		*f1 = origin
	})
}