// Copyright 2019 0xrgb. All rights reserved.
// Use of this source code is governed by a MIT license that can be found
// in the LICENSE file.

package main

// Blocked Domain Lists
var blockedDomainList = [...]string{
	// Twitch
	"www.twitch.tv",
	"api.twitch.tv",
	"gql.twitch.tv",

	// Afreeca
	"www.afreecatv.com",
	"res.afreecatv.com",
	"live.afreecatv.com",
	"play.afreecatv.com",
	"vod.afreecatv.com",
	"st.afreecatv.com",
	"live-stream-manager.afreecatv.com",

	// League of legends (KR)
	"auth.riotgames.com",
	"status.leagueoflegends.com",
	"kr.patchdata.lolstatic.com",
	"chat.kr.lol.riotgames.com",
	"prod.kr.lol.riotgames.com",

	// Battle.net (KR)
	"kr.version.battle.net",
	"kr.patch.battle.net",
	"kr.actual.battle.net",

	// Community
	"www.dcinside.com",
	"gall.dcinside.com",
}

const (
	hosts       string = `C:\Windows\System32\drivers\etc\hosts`
	hostsBackup string = `C:\Windows\System32\drivers\etc\hosts-nomore`
)
