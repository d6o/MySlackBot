package answerers

var (
	userPrefs map[string]string
)

func init() {
	userPrefs = make(map[string]string)
}
