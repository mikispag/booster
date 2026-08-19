package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseProperties(t *testing.T) {
	got := parseProperties("PROP1=VAL1\nPROP2= \"VAL2\"\nPROP3=VAL3\nFONT=cp866-8x14\n", false)

	expect := map[string]string{
		"PROP1": "VAL1",
		"PROP2": "\"VAL2\"",
		"PROP3": "VAL3",
		"FONT":  "cp866-8x14",
	}
	require.Equal(t, expect, got)
}

func TestParseAndStripProperties(t *testing.T) {
	got := parseProperties("PROP1=VAL1\nPROP2= VAL2\nPROP3=VAL3\nFONT=  \"cp866-8x14\" \nKEYMAP='us'\n", true)

	expect := map[string]string{
		"PROP1":  "VAL1",
		"PROP2":  "VAL2",
		"PROP3":  "VAL3",
		"FONT":   "cp866-8x14",
		"KEYMAP": "us",
	}
	require.Equal(t, expect, got)
}

func TestStripQuotes(t *testing.T) {
	require.Equal(t, "", stripQuotes(""))
	require.Equal(t, "\"", stripQuotes("\""))
	require.Equal(t, "'", stripQuotes("'"))
	require.Equal(t, "", stripQuotes("\"\""))
	require.Equal(t, "", stripQuotes("''"))
	require.Equal(t, "hello", stripQuotes("\"hello\""))
	require.Equal(t, "hello", stripQuotes("'hello'"))
	require.Equal(t, "\"hello'", stripQuotes("\"hello'"))
	require.Equal(t, "hello", stripQuotes("hello"))
}
