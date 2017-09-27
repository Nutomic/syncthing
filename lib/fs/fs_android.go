// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package fs

import (
	"os"
	"time"
)

type AndroidFilesystem struct {
	root string
}

func newAndroidFilesystem(root string) *AndroidFilesystem {
	return &AndroidFilesystem{
		root: root,
	}
}

// TODO: implement functions in java using JNI

func (f *AndroidFilesystem) rooted(rel string) (string, error) {
	return nil, nil
}

func (f *AndroidFilesystem) unrooted(path string) string {
	return nil
}

func (f *AndroidFilesystem) Chmod(name string, mode FileMode) error {
	return nil
}

func (f *AndroidFilesystem) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return nil
}

func (f *AndroidFilesystem) Mkdir(name string, perm FileMode) error {
	return nil
}

func (f *AndroidFilesystem) Lstat(name string) (FileInfo, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Remove(name string) error {
	return nil
}

func (f *AndroidFilesystem) RemoveAll(name string) error {
	return nil
}

func (f *AndroidFilesystem) Rename(oldpath, newpath string) error {
	return nil
}

func (f *AndroidFilesystem) Stat(name string) (FileInfo, error) {
	return nil, nil
}

func (f *AndroidFilesystem) DirNames(name string) ([]string, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Open(name string) (File, error) {
	return nil, nil
}

func (f *AndroidFilesystem) OpenFile(name string, flags int, mode FileMode) (File, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Create(name string) (File, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Walk(root string, walkFn WalkFunc) error {
	return nil
}

func (f *AndroidFilesystem) Glob(pattern string) ([]string, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Usage(name string) (Usage, error) {
	return nil, nil
}

func (f *AndroidFilesystem) Type() FilesystemType {
	return FileSystemTypeAndroid
}

func (f *AndroidFilesystem) URI() string {
	return nil
}

type fsFile struct {
	*os.File
	name string
}

func (f fsFile) Name() string {
	return nil
}

func (f fsFile) Stat() (FileInfo, error) {
	return nil, nil
}

func (f fsFile) Sync() error {
	return nil
}

type fsFileInfo struct {
	os.FileInfo
}

func (e fsFileInfo) Mode() FileMode {
	return false
}

func (e fsFileInfo) IsRegular() bool {
	return false
}

func (e fsFileInfo) IsSymlink() bool {
	return false
}
