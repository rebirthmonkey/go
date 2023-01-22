/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugina

import (
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/scheme"
	"github.com/rebirthmonkey/go/96_dp/88_scheme-builder/pkg/scheme/registry"
)

func init() {
	registry.Register(func(s *scheme.Scheme) error {
		s.Plugins["pluginB"] = "yyy"
		return nil
	})
}
