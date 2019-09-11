/*
 * Copyright 2019 Pengcheg Cai
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package clashcaller

import(
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	log "github.com/sirupsen/logrus"
)

var switchNotify = make(chan int)

func start() {
	log.Info("Clash starting.")

	if err := config.Init(constant.Path.HomeDir()); err != nil {
		log.Fatalf("Initial configuration directory error: %s", err.Error())
	}

	if err := hub.Parse(); err != nil {
		log.Fatalf("Parse config error: %s", err.Error())
	}

	log.Info("Clash started.")

	<-switchNotify
	log.Info("Clash stopped.")
}

func stop()  {
	switchNotify <- 0
}

func restart()  {
	stop()
	switchNotify <- 0
	<-switchNotify
	start()
}
