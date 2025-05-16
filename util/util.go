/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package util

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func GetHomeDir() string {
	return getUserHomeDir()
}

func getUserHomeDir() string {
	dir := "symfony5"
	dotDir := "." + dir

	if InCloud() {
		u, err := user.Current()
		if err != nil {
			return filepath.Join(os.TempDir(), dir)
		}
		return filepath.Join(os.TempDir(), u.Username, dir)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	// use the old path if it exists already
	fallback := filepath.Join(home, dotDir)
	if _, err := os.Stat(fallback); !os.IsNotExist(err) {
		return fallback
	}

	// macos only: if $HOME/.config exist, prefer that over 'Library/Application Support'
	if runtime.GOOS == "darwin" {
		dotconf := filepath.Join(home, ".config")
		if _, err := os.Stat(dotconf); !os.IsNotExist(err) {
			return filepath.Join(dotconf, dir)
		}
	}

	userCfg, err := os.UserConfigDir()
	if err != nil {
		return fallback
	}

	return filepath.Join(userCfg, dir)
}
