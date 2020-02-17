// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func apiAllMessagesHandler(r *http.Request) interface{} {
	messages, err := core.SelectMessages()
	if err != nil {
		log.Error(err)
		return nil
	}
	return joinMessages(messages)
}

func joinMessages(messages []*core.Message) []*types.Message {
	var m []*types.Message
	for _, v := range messages {
		m = append(m, v.Message)
	}
	return m
}

func apiMessageCreateHandler(w http.ResponseWriter, r *http.Request) {
	var message *types.Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	msg := core.ReturnMessage(message)
	_, err = msg.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(msg, "create", w, r)
}

func apiMessageGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(message, w, r)
}

func apiMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = message.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(message, "delete", w, r)
}

func apiMessageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(fmt.Errorf("message #%v was not found", vars["id"]), w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&message)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	_, err = message.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(message, "update", w, r)
}
