/*
 * Vi - A Discord Bot written in Go
 * Copyright (C) 2019  Brett Bender
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package Shared

import "github.com/bwmarrin/discordgo"

func CheckPermissions(s *discordgo.Session, guildid, memberid string, required Permission) bool {
	// Don't even bother checking if no permissions required
	if required == 0 {
		return true
	}

	member, err := s.State.Member(guildid, memberid)
	if err != nil {
		return false
	}

	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildid, roleID)
		if err != nil {
			return false
		}

		// If they have admin they have all permissions
		if role.Permissions&discordgo.PermissionAdministrator != 0 {
			return true
		}

		// If permissions and required isn't 0; return true
		if role.Permissions&int(required) != 0 {
			return true
		}
	}

	// Got nothing so false
	return false
}
