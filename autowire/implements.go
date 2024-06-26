/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package autowire

import (
	"fmt"

	"github.com/cc-cheunggg/ioc-golang/autowire/util"
	"github.com/cc-cheunggg/ioc-golang/logger"
)

const (
	defaultProfileKey = "_default"
)

// implementsMap a map from interface SDID to implement struct SDID
var implementsMap = make(map[string]map[string]map[string]struct{})

func registerImplements(sd *StructDescriptor) {
	// get and register impl interfaces metadata
	allImpledIntefaces := parseCommonImplementsMetadataFromSDMetadata(sd.Metadata)
	activeProfile := parseCommonActiveProfileMetadataFromSDMetadata(sd.Metadata)
	for _, impledInteface := range allImpledIntefaces {
		interfaceSDID := util.GetSDIDByStructPtr(impledInteface)
		if err := registerImplementMapping(interfaceSDID, activeProfile, sd.ID()); err != nil {
			logger.Error("[Autowire Singleton] RegisterImplementMapping failed with error = %s", err.Error())
		}
	}
}

func registerImplementMapping(interfaceSDID, activeProfile, implStructSDID string) error {
	if activeProfile == "" {
		activeProfile = defaultProfileKey
	}

	// 1. assure interfaceSDID map exists
	if _, ok := implementsMap[interfaceSDID]; !ok {
		implementsMap[interfaceSDID] = make(map[string]map[string]struct{})
	}

	// 2. assure activeProfile slice exist
	_, ok := implementsMap[interfaceSDID][activeProfile]
	if !ok {
		implementsMap[interfaceSDID][activeProfile] = make(map[string]struct{}, 0)
	}

	// 3. register implements mapping
	implementsMap[interfaceSDID][activeProfile][implStructSDID] = struct{}{}
	return nil
}

// GetBestImplementMapping get best impl struct sdids slice and profile of given interfaceSDID and activitedOrderedProfiles
// if there isn't any matched implementation, even "_default" impl is not found, an error occurs
// return values are bestMatchesStructImplSDIDs, bestMatchProfile, error
func GetBestImplementMapping(interfaceSDID string, activitedOrderedProfiles []string) ([]string, string, error) {
	interfaceImplsMap, ok := implementsMap[interfaceSDID]
	if !ok {
		err := fmt.Errorf("[Autowire Implement] Any of interface %s implements not found", interfaceSDID)
		logger.Error(err.Error())
		return nil, "", err
	}

	// validate default fallback
	if activitedOrderedProfiles == nil {
		activitedOrderedProfiles = make([]string, 0)
	}
	activitedOrderedProfiles = append([]string{defaultProfileKey}, activitedOrderedProfiles...)
	for k, v := range activitedOrderedProfiles {
		if v == "" {
			activitedOrderedProfiles[k] = defaultProfileKey
		}
	}

	// get with best profile
	bestMatchesStructImplSDIDsMap := make(map[string]struct{}, 0)
	bestMatchProfile := ""
	for i := len(activitedOrderedProfiles) - 1; i >= 0; i-- {
		currentProfile := activitedOrderedProfiles[i]
		bestMatchesStructImplSDIDsMap, ok = interfaceImplsMap[currentProfile]
		if !ok || len(bestMatchesStructImplSDIDsMap) == 0 {
			logger.Info("[Autowire Implement] Interface %s implements with profile %s not found", interfaceSDID, currentProfile)
			continue
		}
		logger.Info("[Autowire Implement] Interface %s implements SDID %s with profile %s bast matches activited profiles, select it(them)", interfaceSDID, bestMatchesStructImplSDIDsMap, currentProfile)
		bestMatchProfile = currentProfile
		break
	}
	if bestMatchProfile == "" {
		allImplementedProfiles := make([]string, 0)
		for k := range interfaceImplsMap {
			allImplementedProfiles = append(allImplementedProfiles, k)
		}
		err := fmt.Errorf("[Autowire Implement] Interface %s has implemented profile %+v, but activited profiles %+v doesn't match any",
			interfaceSDID, allImplementedProfiles, activitedOrderedProfiles)
		logger.Error(err.Error())
		return nil, "", err
	}
	bestMatchesStructImplSDIDs := make([]string, 0)
	for k := range bestMatchesStructImplSDIDsMap {
		bestMatchesStructImplSDIDs = append(bestMatchesStructImplSDIDs, k)
	}
	return bestMatchesStructImplSDIDs, bestMatchProfile, nil
}
