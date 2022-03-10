package bot

type Shrug string

const (
	NoShrug Shrug = ""
	// MissingLeftArmPattern ¯\_(ツ)_/¯
	MissingLeftArmPattern Shrug = `(¯\\_\(ツ\)_\/¯)`
	// MissingShouldersPattern ¯\(ツ)/¯
	MissingShouldersPattern Shrug = `(¯\\\\_\(ツ\)_\/¯)`
)

var invalidShrugBodies = [...]Shrug{
	MissingLeftArmPattern,
	MissingShouldersPattern,
}

func (s Shrug) commentResponse() string {
	switch s {
	case MissingLeftArmPattern:
		return `You dropped this \`
	case MissingShouldersPattern:
		return `You dropped these _ _`
	}
	return ""
}
