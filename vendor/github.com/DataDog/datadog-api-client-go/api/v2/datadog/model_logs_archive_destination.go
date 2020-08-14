/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsArchiveDestination - An archive's destination.
type LogsArchiveDestination struct {
	LogsArchiveDestinationAzure *LogsArchiveDestinationAzure
	LogsArchiveDestinationGCS   *LogsArchiveDestinationGCS
	LogsArchiveDestinationS3    *LogsArchiveDestinationS3
}

// LogsArchiveDestinationAzureAsLogsArchiveDestination is a convenience function that returns LogsArchiveDestinationAzure wrapped in LogsArchiveDestination
func LogsArchiveDestinationAzureAsLogsArchiveDestination(v *LogsArchiveDestinationAzure) LogsArchiveDestination {
	return LogsArchiveDestination{LogsArchiveDestinationAzure: v}
}

// LogsArchiveDestinationGCSAsLogsArchiveDestination is a convenience function that returns LogsArchiveDestinationGCS wrapped in LogsArchiveDestination
func LogsArchiveDestinationGCSAsLogsArchiveDestination(v *LogsArchiveDestinationGCS) LogsArchiveDestination {
	return LogsArchiveDestination{LogsArchiveDestinationGCS: v}
}

// LogsArchiveDestinationS3AsLogsArchiveDestination is a convenience function that returns LogsArchiveDestinationS3 wrapped in LogsArchiveDestination
func LogsArchiveDestinationS3AsLogsArchiveDestination(v *LogsArchiveDestinationS3) LogsArchiveDestination {
	return LogsArchiveDestination{LogsArchiveDestinationS3: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogsArchiveDestination) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into LogsArchiveDestinationAzure
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationAzure)
	if err == nil {
		jsonLogsArchiveDestinationAzure, _ := json.Marshal(dst.LogsArchiveDestinationAzure)
		if string(jsonLogsArchiveDestinationAzure) == "{}" { // empty struct
			dst.LogsArchiveDestinationAzure = nil
		} else {
			match++
		}
	} else {
		dst.LogsArchiveDestinationAzure = nil
	}

	// try to unmarshal data into LogsArchiveDestinationGCS
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationGCS)
	if err == nil {
		jsonLogsArchiveDestinationGCS, _ := json.Marshal(dst.LogsArchiveDestinationGCS)
		if string(jsonLogsArchiveDestinationGCS) == "{}" { // empty struct
			dst.LogsArchiveDestinationGCS = nil
		} else {
			match++
		}
	} else {
		dst.LogsArchiveDestinationGCS = nil
	}

	// try to unmarshal data into LogsArchiveDestinationS3
	err = json.Unmarshal(data, &dst.LogsArchiveDestinationS3)
	if err == nil {
		jsonLogsArchiveDestinationS3, _ := json.Marshal(dst.LogsArchiveDestinationS3)
		if string(jsonLogsArchiveDestinationS3) == "{}" { // empty struct
			dst.LogsArchiveDestinationS3 = nil
		} else {
			match++
		}
	} else {
		dst.LogsArchiveDestinationS3 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LogsArchiveDestinationAzure = nil
		dst.LogsArchiveDestinationGCS = nil
		dst.LogsArchiveDestinationS3 = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(LogsArchiveDestination)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(LogsArchiveDestination)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogsArchiveDestination) MarshalJSON() ([]byte, error) {
	if src.LogsArchiveDestinationAzure != nil {
		return json.Marshal(&src.LogsArchiveDestinationAzure)
	}

	if src.LogsArchiveDestinationGCS != nil {
		return json.Marshal(&src.LogsArchiveDestinationGCS)
	}

	if src.LogsArchiveDestinationS3 != nil {
		return json.Marshal(&src.LogsArchiveDestinationS3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogsArchiveDestination) GetActualInstance() interface{} {
	if obj.LogsArchiveDestinationAzure != nil {
		return obj.LogsArchiveDestinationAzure
	}

	if obj.LogsArchiveDestinationGCS != nil {
		return obj.LogsArchiveDestinationGCS
	}

	if obj.LogsArchiveDestinationS3 != nil {
		return obj.LogsArchiveDestinationS3
	}

	// all schemas are nil
	return nil
}

type NullableLogsArchiveDestination struct {
	value *LogsArchiveDestination
	isSet bool
}

func (v NullableLogsArchiveDestination) Get() *LogsArchiveDestination {
	return v.value
}

func (v *NullableLogsArchiveDestination) Set(val *LogsArchiveDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveDestination(val *LogsArchiveDestination) *NullableLogsArchiveDestination {
	return &NullableLogsArchiveDestination{value: val, isSet: true}
}

func (v NullableLogsArchiveDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}