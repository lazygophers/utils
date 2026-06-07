package fake_test

import (
	"fmt"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

func ExampleNew() {
	f := fake.New(country.UnitedStates, fake.WithSeed(42), fake.WithGender(fake.GenderFemale))
	fmt.Println(f.Country().Alpha2(), f.DefaultGender())
	// Output: US female
}
