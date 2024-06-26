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

package protocol_impl

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/cc-cheunggg/ioc-golang/logger"
)

// ParseArgs @data should be []interface{}'s marshal result, @argsType should be each object's reflect type
func ParseArgs(argsType []reflect.Type, data []byte) (finalArgument []interface{}, err error) {
	finalArgument = make([]interface{}, 0)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("IOC Protocol parse args failed with catched error = %+v", r)
		}
	}()
	// http req -> invocation
	rawArguments := make([]interface{}, 0)
	for _, reflectType := range argsType {
		if reflectType.Kind() == reflect.Ptr {
			rawArguments = append(rawArguments, reflect.New(reflectType.Elem()).Interface())
			finalArgument = append(finalArgument, reflect.New(reflectType.Elem()).Interface())
		} else {
			rawArguments = append(rawArguments, reflect.New(reflectType).Interface())
			finalArgument = append(finalArgument, reflect.New(reflectType).Interface())
		}
	}
	if err := json.Unmarshal(data, &rawArguments); err != nil {
		return nil, err
	}
	if len(rawArguments) != len(argsType) {
		errMsg := fmt.Sprintf("IOC Protocol parse args failed, want %d params but %d is given", len(argsType), len(rawArguments))
		logger.Error(errMsg)
		return nil, fmt.Errorf(errMsg)
	}
	for idx, reflectType := range argsType {
		if reflectType.Kind() == reflect.Ptr {
			finalArgument[idx] = rawArguments[idx]
		} else {
			finalArgument[idx] = reflect.ValueOf(rawArguments[idx]).Elem().Interface()
		}
	}
	return finalArgument, nil
}
