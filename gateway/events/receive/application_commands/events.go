/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package application_commands

import (
	"github.com/veteran-software/discord-api-wrapper/v10/api"
)

// ApplicationCommandPermissionsUpdate - APPLICATION_COMMAND_PERMISSIONS_UPDATE event, sent when an application command's permissions are updated.
//
// The inner payload is an application command permissions object.
type (
	ApplicationCommandPermissionsUpdate api.ApplicationCommandPermissions
)
