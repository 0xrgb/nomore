// Copyright 2019 0xrgb. All rights reserved.
// Use of this source code is governed by a MIT license that can be found
// in the LICENSE file.

// Generate resource file using windres (MinGW).
//go:generate windres -o nomore.syso resource/nomore.rc

package main

import (
	// Walk

	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// DLL
var (
	dnsapiDLL                 = syscall.NewLazyDLL("dnsapi.dll")
	procDnsFlushResolverCache = dnsapiDLL.NewProc("DnsFlushResolverCache")
)

func flushDNS() error {
	r1, _, err := procDnsFlushResolverCache.Call()
	if r1 != 0 {
		return nil
	}

	// Failed because of ...
	val := err.(syscall.Errno)
	return fmt.Errorf("DnsFlushResolver failed: windows error #%d", val)
}

// 실제 함수
var (
	ErrAlreadyBanned     = errors.New("이미 차단되어 있습니다.")
	ErrCannotBackup      = errors.New("백업 실패")
	ErrCannotModifyHosts = errors.New("hosts 파일을 열 수 없습니다.")
	ErrCannotFlushDNS    = errors.New("DNS 초기화 실패. 컴퓨터를 재부팅해 주세요.")
)

func backupFile() error {
	src, err := os.Open(hosts)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(hostsBackup)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}

func addHostWithError() error {
	// 이미 차단되이 있는 경우
	if _, err := os.Stat(hostsBackup); err == nil {
		return ErrAlreadyBanned
	}

	// 복사
	if err := backupFile(); err != nil {
		return ErrCannotBackup
	}

	// 파일에 추가
	f, err := os.OpenFile(hosts, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return ErrCannotModifyHosts
	}

	fb := bufio.NewWriter(f)
	for _, val := range blockedDomainList {
		fb.WriteString("\r\n127.0.0.1 ")
		fb.WriteString(val)
	}
	fb.Flush()

	f.Sync()
	f.Close()

	// DNS flush
	if err := flushDNS(); err != nil {
		return ErrCannotFlushDNS
	}

	return nil
}

// UI에 사용될 함수들
var mw *walk.MainWindow

func addHost() {
	if err := addHostWithError(); err != nil {
		walk.MsgBox(mw, "실패", err.Error(), walk.MsgBoxIconWarning)
	} else {
		walk.MsgBox(mw, "성공", "차단 완료", walk.MsgBoxIconInformation)
	}
}

func seeUsage() {
	walk.MsgBox(mw, "NOMORE 정보", "NOMORE은 게임과 커뮤니티 사이트를 차단하여 집중력을 향상시켜주는 프로그램입니다.", walk.MsgBoxIconInformation)
}

func main() {
	// Load icon
	icon, _ := walk.NewIconFromResourceId(2)
	MainWindow{
		AssignTo: &mw,
		Title:    "NOMORE",
		Icon:     icon,
		MinSize:  Size{250, 150},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text:      "게임 차단하기",
				OnClicked: addHost,
			},
			PushButton{
				Text:      "프로그램 정보",
				OnClicked: seeUsage,
			},
		},
	}.Create()

	mw.Run()
}
