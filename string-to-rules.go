package abolish

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Find key or !key
 * @type {RegExp}
 */
var onlyKey = regexp.MustCompile(`([!a-zA-Z_*0-9]+)`)

/**
 * Find key:name
 * @type {RegExp}
 */
var keyColumnVal = regexp.MustCompile(`([a-zA-Z_*0-9]+:[a-zA-Z_0-9]+)`)

/**
 * Find key:"string"
 * @type {RegExp}
 */
var keyColumnValStringDoubleQuotes = regexp.MustCompile(`([a-zA-Z_*0-9]+:"[^"]+")`)

/**
 * Find key:'string'
 * @type {RegExp}
 */
var keyColumnValStringSingleQuotes = regexp.MustCompile(`([a-zA-Z_*0-9]+:'[^']+')`)

/**
 * Find key:`string`
 * @type {RegExp}
 */
var keyColumnValStringGraveAccent = regexp.MustCompile(`([a-zA-Z_*0-9]+:` + "`" + `[^` + "`" + `]+` + "`" + `)`)

type Rules map[string]any

func StringToRules(str string) Rules {
	rules := make(Rules)

	// split string by |
	s := strings.Split(str, "|")

	for _, pair := range s {
		if keyColumnValStringSingleQuotes.MatchString(pair) ||
			keyColumnValStringDoubleQuotes.MatchString(pair) ||
			keyColumnValStringGraveAccent.MatchString(pair) {

			key := strings.Split(pair, ":")[0]
			value := strings.Split(pair, ":")[1:]

			rules[key] = value
		} else if keyColumnVal.MatchString(pair) {
			key := strings.Split(pair, ":")[0]
			value := strings.Split(pair, ":")[1]

			// check if value is a number
			if _, err := strconv.Atoi(value); err == nil {
				// convert to number
				rules[key], _ = strconv.Atoi(value)
			} else {
				rules[key] = value
			}
		} else if onlyKey.MatchString(pair) {
			/*
			 * If key is like "key|" or "!key|"
			 * ==> {key: true} or {key: false}
			 */

			key := pair
			value := true

			// if !key set value to false
			if key[0:1] == "!" {
				key = key[1:]
				value = false
			}

			rules[key] = value
		}
	}

	return rules
}

func StringToRulesCompiled(str string) (Rules, error) {
	Map := StringToRules(str)

	// get all keys from map
	for k := range StringToRules(str) {
		// check if key is a validator
		if !HasValidator(k) {
			return Map, &ValidationError{
				Code:    "validation",
				Message: fmt.Sprintf("validator [%v] does not exist.", k),
			}
		}
	}

	return Map, nil
}
