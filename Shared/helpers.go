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
