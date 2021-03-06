// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/knq/xo/examples/pgcatalog/pgtypes"

// GENERATED BY XO. DO NOT EDIT.

// SQLLanguage represents a row from information_schema.sql_languages.
type SQLLanguage struct {
	Tableoid                       pgtypes.Oid           // tableoid
	Cmax                           pgtypes.Cid           // cmax
	Xmax                           pgtypes.Xid           // xmax
	Cmin                           pgtypes.Cid           // cmin
	Xmin                           pgtypes.Xid           // xmin
	Ctid                           pgtypes.Tid           // ctid
	SQLLanguageSource              pgtypes.CharacterData // sql_language_source
	SQLLanguageYear                pgtypes.CharacterData // sql_language_year
	SQLLanguageConformance         pgtypes.CharacterData // sql_language_conformance
	SQLLanguageIntegrity           pgtypes.CharacterData // sql_language_integrity
	SQLLanguageImplementation      pgtypes.CharacterData // sql_language_implementation
	SQLLanguageBindingStyle        pgtypes.CharacterData // sql_language_binding_style
	SQLLanguageProgrammingLanguage pgtypes.CharacterData // sql_language_programming_language
}
