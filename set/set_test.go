package set

import "testing"

func TestSet(t *testing.T) {

	set := make(Set)
	set.Add("a")
	set.Add("b")
	set.Add("c")

	hasa := set.Has("a")
	hasb := set.Has("b")
	hasc := set.Has("c")

	t.Log(hasa, hasb, hasc)

	set.Delete("a")
	hasa = set.Has("a")
	t.Log(hasa)

	da := set.Has("a")
	db := set.Has("b")
	dc := set.Has("c")

	t.Log(da, db, dc)

}
