// Copyright (C) 2019 Luiz de Milon (kori)

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package listenbrainz

// source: https://listenbrainz.readthedocs.io/en/latest/dev/api.html#constants
const (
	// Maximum overall listen size in bytes, to prevent egregious spamming.
	MaxListenSize = 10240
	// The maximum number of listens returned in a single GET request.
	MaxItemsPerGet = 100
	// The default number of listens returned in a single GET request.
	DefaultItemsPerGet = 25
	// The maximum number of tags per listen.
	MaxTagsPerListen = 50
	// The maximum length of a tag
	MaxTagSize = 64
)
