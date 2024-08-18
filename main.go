package main

import (
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	supertokens.Init(supertokens.TypeInput{
		RecipeList: []supertokens.Recipe{
			// TODO: Initialise other recipes
			dashboard.Init(nil),
		},
	})
}
