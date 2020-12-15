/* -*- coding: utf-8 -*-
* ------------------------------------------------------------------------------
*
*   Copyright 2018-2019 Fetch.AI Limited
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*   Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*
* ------------------------------------------------------------------------------
 */

// Package dhtnode (in progress) contains the common interface between dhtpeer and dhtclient
package dhtnode

import (
	"errors"
	utils "libp2p_node/utils"
)

func IsValidProofOfRepresentation(record *AgentRecord, representativePeerPubKey string) (*Status, error) {
	// check public key matches
	if record.PeerPublicKey != representativePeerPubKey {
		err := errors.New("Wrong peer public key, expected " + representativePeerPubKey)
		response := &Status{Code: Status_ERROR_INVALID_PUBLIC_KEY, Msgs: []string{err.Error()}}
		return response, err
	}

	// check that agent address and public key match
	addrFromPubKey, err := utils.FetchAIAddressFromPublicKey(record.PublicKey)
	if err != nil || addrFromPubKey != record.Address {
		if err == nil {
			err = errors.New("Agent address and public key don't match")
		}
		response := &Status{Code: Status_ERROR_INVALID_AGENT_ADDRESS}
		return response, err
	}

	// check that signature is valid
	ok, err := utils.VerifyFetchAISignatureBTC([]byte(record.PeerPublicKey), record.Signature, record.PublicKey)
	if !ok || err != nil {
		if err == nil {
			err = errors.New("Signature is not valid")
		}
		response := &Status{Code: Status_ERROR_INVALID_PROOF}
		return response, err

	}

	// PoR is valid
	response := &Status{Code: Status_SUCCESS}
	return response, nil

}
