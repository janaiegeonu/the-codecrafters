package helpers

func Complier(text string) string {

	holder := Trimmer(text)

	holder = LowToUp(holder)

	ReplaceWord(holder)
	CapToTitle(holder)

	return Trimmer(holder)

}
