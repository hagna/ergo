// Copyright (c) 2017 Daniel Oaks <daniel@danieloaks.net>
// released under the MIT license

// Package sno holds Server Notice masks for easy reference.
package sno

// Mask is a type of server notice mask.
type Mask rune

// Notice mask types
const (
	LocalAccouncements Mask = 'a'
	LocalConnects      Mask = 'c'
	LocalChannels      Mask = 'j'
	LocalKills         Mask = 'k'
	LocalNicks         Mask = 'n'
	LocalOpers         Mask = 'o'
	LocalQuits         Mask = 'q'
	Stats              Mask = 't'
	LocalAccounts      Mask = 'u'
	LocalXline         Mask = 'x'
)

var (
	// NoticeMaskNames has readable names for our snomask types.
	NoticeMaskNames = map[Mask]string{
		LocalAccouncements: "ANNOUNCEMENT",
		LocalConnects:      "CONNECT",
		LocalChannels:      "CHANNEL",
		LocalKills:         "KILL",
		LocalNicks:         "NICK",
		LocalOpers:         "OPER",
		LocalQuits:         "QUIT",
		Stats:              "STATS",
		LocalAccounts:      "ACCOUNT",
		LocalXline:         "XLINE",
	}

	// ValidMasks contains the snomasks that we support.
	ValidMasks = map[Mask]bool{
		LocalAccouncements: true,
		LocalConnects:      true,
		LocalChannels:      true,
		LocalKills:         true,
		LocalNicks:         true,
		LocalOpers:         true,
		LocalQuits:         true,
		Stats:              true,
		LocalAccounts:      true,
		LocalXline:         true,
	}
)
