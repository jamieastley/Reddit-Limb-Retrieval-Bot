package main

import "fmt"

type Shrug string

var LiteralCodeShrugPattern = fmt.Sprintf(`<code>.*%s.*<\/code>`, MissingLeftArmPattern)

const literalShrug = `¯\_(ツ)_/¯`

const (
	// NoShrug defines that no matching shrug was found
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

func (s Shrug) matchType() string {
	switch s {
	case MissingLeftArmPattern:
		return "MissingLeftArm"
	case MissingShouldersPattern:
		return "MissingShoulders"
	}
	return ""
}
