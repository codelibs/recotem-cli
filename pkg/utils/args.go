package utils

import (
	"strconv"

	"recotem.org/cli/recotem/pkg/openapi"
)

func NilOrString(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

func NilOrInt(s string) *int {
	x, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &x
}

func NilOrFloat32(s string) *float32 {
	x, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return nil
	}
	v := float32(x)
	return &v
}

func NilOrBool(s string) *bool {
	x, err := strconv.ParseBool(s)
	if err != nil {
		return nil
	}
	return &x
}

func NilOrScheme(s string) *openapi.SchemeEnum {
	scheme := NilOrString(s)
	if scheme == nil {
		return nil
	}
	if *scheme == "RG" {
		x := openapi.SchemeEnumRG
		return &x
	} else if *scheme == "TG" {
		x := openapi.SchemeEnumTG
		return &x
	} else if *scheme == "TU" {
		x := openapi.SchemeEnumTU
		return &x
	}
	return nil
}
