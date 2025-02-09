// Copyright © 2017 Alessandro Sanino <saninoale@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package examples

import (
	"github.com/llazzaro/golang-crypto-trading-bot/environment"
	"github.com/llazzaro/golang-crypto-trading-bot/exchanges"
	"github.com/llazzaro/golang-crypto-trading-bot/strategies"
	"github.com/sirupsen/logrus"
)

// Websocket strategy
var Websocket = strategies.WebsocketStrategy{
	Model: strategies.StrategyModel{
		Name: "Websocket",
		Setup: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			for _, wrapper := range wrappers {
				err := wrapper.FeedConnect(markets)
				if err == exchanges.ErrWebsocketNotSupported || err == nil {
					continue
				}
				return err
			}
			return nil
		},
		OnUpdate: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			// do something
			return nil
		},
		TearDown: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			return nil
		},
		OnError: func(err error) {
			logrus.Error(err)
		},
	},
}
