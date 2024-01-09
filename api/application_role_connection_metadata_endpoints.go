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

package api

import (
	"encoding/json"
	"fmt"

	log "github.com/veteran-software/nowlive-logging"
)

// GetApplicationRoleConnectionMetadataRecords - Returns a list of ApplicationRoleConnectionMetadata objects for the given Application.
//
//goland:noinspection GoUnusedExportedFunction
func GetApplicationRoleConnectionMetadataRecords(appID string) ([]*ApplicationRoleConnectionMetadata, error) {
	u := parseRoute(fmt.Sprintf(getApplicationRoleConnectionMetadataRecords, api, appID))

	var m []*ApplicationRoleConnectionMetadata
	responseBytes, err := fireGetRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &m)

	return m, err
}

// UpdateApplicationRoleConnectionMetadataRecords - Updates and returns a list of ApplicationRoleConnectionMetadata objects for the given Application.
//
//goland:noinspection GoUnusedExportedFunction
func UpdateApplicationRoleConnectionMetadataRecords(appID string) ([]*ApplicationRoleConnectionMetadata, error) {
	u := parseRoute(fmt.Sprintf(updateApplicationRoleConnectionMetadataRecords, api, appID))

	var m []*ApplicationRoleConnectionMetadata
	responseBytes, err := firePutRequest(u, nil, nil)
	if err != nil {
		log.Errorln(log.Discord, log.FuncName(), err)
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &m)

	return m, err
}
