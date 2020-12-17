// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"fmt"
	"strings"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/rbac_1_0_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	modelplugin "github.com/onosproject/config-models/modelplugin/rbac-1.0.0/rbac_1_0_0"
	"github.com/openconfig/gnmi/proto/gnmi"
	"reflect"
	"regexp"
)

var re *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)

// encodeToGnmiRbac converts OAPI to gNMI.
func encodeToGnmiRbac(
	jsonObj *types.Rbac, needKey bool, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	mp := modelplugin.Device{}
	fmt.Printf("mp %T\n", mp)

	//Property: { Group {[]RbacGroup  map[] [] false <nil> [] false} false false}
	if jsonObj.Group != nil {

	}
	//Property: { Role {[]RbacRole  map[] [] false <nil> [] false} false false}
	if jsonObj.Role != nil {

	}

	//Property: { Group {[]RbacGroup  map[] [] false <nil> [] false} false false}
	if jsonObj.Group != nil {
		for _, item := range *jsonObj.Group {
			item := item //Pinning
			paramsGroup := make([]string, len(params))
			copy(paramsGroup, params)
			paramsGroup = append(paramsGroup, "unknown_id")
			updatesGroup, err :=
				encodeToGnmiRbacGroup(&item, true,
					fmt.Sprintf("%s/%s/{unknown_key}", parentPath, "group"), paramsGroup...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatesGroup...)
		}
	}

	//Property: { Role {[]RbacRole  map[] [] false <nil> [] false} false false}
	if jsonObj.Role != nil {
		for _, item := range *jsonObj.Role {
			item := item //Pinning
			paramsRole := make([]string, len(params))
			copy(paramsRole, params)
			paramsRole = append(paramsRole, "unknown_id")
			updatesRole, err :=
				encodeToGnmiRbacRole(&item, true,
					fmt.Sprintf("%s/%s/{unknown_key}", parentPath, "role"), paramsRole...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatesRole...)
		}
	}

	if needKey {
		reflectKey, err := utils.FindModelPluginObject(mp, "Rbac", params...)
		if err != nil {
			return nil, err
		}
		reflectType := reflectKey.Type()
		reflect2 := reflect.New(reflectType) // Needed so the type can be read to extract list
		reflect2.Elem().Set(*reflectKey)
		keyMap, err := utils.ExtractGnmiListKeyMap(reflect2.Interface())
		if err != nil {
			return nil, err
		}
		for k, v := range keyMap {
			// parentPath = fmt.Sprintf("%s/{%s}", parentPath, k)
			for _, u := range updates {
				if err := utils.ReplaceUnknownKey(u, k, v, "unknown_key", "unknown_id"); err != nil {
					return nil, err
				}
			}
		}
	}
	return updates, nil
}

// encodeToGnmiRbacGroup converts OAPI to gNMI.
func encodeToGnmiRbacGroup(
	jsonObj *types.RbacGroup, needKey bool, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	mp := modelplugin.Device{}
	fmt.Printf("mp %T\n", mp)

	//Property: { Role {[]RbacGroupRole  map[] [] false <nil> [] false} false false}
	if jsonObj.Role != nil {

	}
	//Property: { description {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Description != nil {

		paramsDescription := make([]string, len(params))
		copy(paramsDescription, params)
		stringValDescription := fmt.Sprintf("%v", *jsonObj.Description)
		paramsDescription = append(paramsDescription, stringValDescription)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacGroupDescription", paramsDescription...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/description"), paramsDescription...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}
	//Property: { groupid {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Groupid != nil {

		paramsGroupid := make([]string, len(params))
		copy(paramsGroupid, params)
		stringValGroupid := fmt.Sprintf("%v", *jsonObj.Groupid)
		paramsGroupid = append(paramsGroupid, stringValGroupid)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacGroupGroupid", paramsGroupid...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/groupid"), paramsGroupid...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}

	//Property: { Role {[]RbacGroupRole  map[] [] false <nil> [] false} false false}
	if jsonObj.Role != nil {
		for _, item := range *jsonObj.Role {
			item := item //Pinning
			paramsRole := make([]string, len(params))
			copy(paramsRole, params)
			paramsRole = append(paramsRole, "unknown_id")
			updatesRole, err :=
				encodeToGnmiRbacGroupRole(&item, true,
					fmt.Sprintf("%s/%s/{unknown_key}", parentPath, "role"), paramsRole...)
			if err != nil {
				return nil, err
			}
			updates = append(updates, updatesRole...)
		}
	}

	if needKey {
		reflectKey, err := utils.FindModelPluginObject(mp, "RbacGroup", params...)
		if err != nil {
			return nil, err
		}
		reflectType := reflectKey.Type()
		reflect2 := reflect.New(reflectType) // Needed so the type can be read to extract list
		reflect2.Elem().Set(*reflectKey)
		keyMap, err := utils.ExtractGnmiListKeyMap(reflect2.Interface())
		if err != nil {
			return nil, err
		}
		for k, v := range keyMap {
			// parentPath = fmt.Sprintf("%s/{%s}", parentPath, k)
			for _, u := range updates {
				if err := utils.ReplaceUnknownKey(u, k, v, "unknown_key", "unknown_id"); err != nil {
					return nil, err
				}
			}
		}
	}
	return updates, nil
}

// encodeToGnmiRbacGroupRole converts OAPI to gNMI.
func encodeToGnmiRbacGroupRole(
	jsonObj *types.RbacGroupRole, needKey bool, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	mp := modelplugin.Device{}
	fmt.Printf("mp %T\n", mp)

	//Property: { description {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Description != nil {

		paramsDescription := make([]string, len(params))
		copy(paramsDescription, params)
		stringValDescription := fmt.Sprintf("%v", *jsonObj.Description)
		paramsDescription = append(paramsDescription, stringValDescription)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacGroupRoleDescription", paramsDescription...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/description"), paramsDescription...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}
	//Property: { roleid {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Roleid != nil {

		paramsRoleid := make([]string, len(params))
		copy(paramsRoleid, params)
		stringValRoleid := fmt.Sprintf("%v", *jsonObj.Roleid)
		paramsRoleid = append(paramsRoleid, stringValRoleid)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacGroupRoleRoleid", paramsRoleid...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/roleid"), paramsRoleid...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}

	if needKey {
		reflectKey, err := utils.FindModelPluginObject(mp, "RbacGroupRole", params...)
		if err != nil {
			return nil, err
		}
		reflectType := reflectKey.Type()
		reflect2 := reflect.New(reflectType) // Needed so the type can be read to extract list
		reflect2.Elem().Set(*reflectKey)
		keyMap, err := utils.ExtractGnmiListKeyMap(reflect2.Interface())
		if err != nil {
			return nil, err
		}
		for k, v := range keyMap {
			// parentPath = fmt.Sprintf("%s/{%s}", parentPath, k)
			for _, u := range updates {
				if err := utils.ReplaceUnknownKey(u, k, v, "unknown_key", "unknown_id"); err != nil {
					return nil, err
				}
			}
		}
	}
	return updates, nil
}

// encodeToGnmiRbacRole converts OAPI to gNMI.
func encodeToGnmiRbacRole(
	jsonObj *types.RbacRole, needKey bool, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	mp := modelplugin.Device{}
	fmt.Printf("mp %T\n", mp)

	//Property: { Permission {RbacRolePermission  map[] [] false <nil> [] false} false false}
	if jsonObj.Permission != nil {

		submatchallPermission := re.FindAllString("Permission", -1)
		update, err := encodeToGnmiRbacRolePermission(
			jsonObj.Permission, false,
			fmt.Sprintf("%s/%s", parentPath, strings.ToLower(strings.Join(submatchallPermission, "/"))), params...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update...)
	}
	//Property: { description {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Description != nil {

		paramsDescription := make([]string, len(params))
		copy(paramsDescription, params)
		stringValDescription := fmt.Sprintf("%v", *jsonObj.Description)
		paramsDescription = append(paramsDescription, stringValDescription)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacRoleDescription", paramsDescription...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/description"), paramsDescription...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}
	//Property: { roleid {string  map[] [] false <nil> [] false} false false}
	if jsonObj.Roleid != nil {

		paramsRoleid := make([]string, len(params))
		copy(paramsRoleid, params)
		stringValRoleid := fmt.Sprintf("%v", *jsonObj.Roleid)
		paramsRoleid = append(paramsRoleid, stringValRoleid)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacRoleRoleid", paramsRoleid...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/roleid"), paramsRoleid...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}

	if needKey {
		reflectKey, err := utils.FindModelPluginObject(mp, "RbacRole", params...)
		if err != nil {
			return nil, err
		}
		reflectType := reflectKey.Type()
		reflect2 := reflect.New(reflectType) // Needed so the type can be read to extract list
		reflect2.Elem().Set(*reflectKey)
		keyMap, err := utils.ExtractGnmiListKeyMap(reflect2.Interface())
		if err != nil {
			return nil, err
		}
		for k, v := range keyMap {
			// parentPath = fmt.Sprintf("%s/{%s}", parentPath, k)
			for _, u := range updates {
				if err := utils.ReplaceUnknownKey(u, k, v, "unknown_key", "unknown_id"); err != nil {
					return nil, err
				}
			}
		}
	}
	return updates, nil
}

// encodeToGnmiRbacRolePermission converts OAPI to gNMI.
func encodeToGnmiRbacRolePermission(
	jsonObj *types.RbacRolePermission, needKey bool, parentPath string, params ...string) (
	[]*gnmi.Update, error) {

	updates := make([]*gnmi.Update, 0)
	mp := modelplugin.Device{}
	fmt.Printf("mp %T\n", mp)

	//Property: { leaf-list-noun {[]string  map[] [] false <nil> [] false} false false}
	if jsonObj.LeafListNoun != nil {

		paramsLeafListNoun := make([]string, len(params))
		copy(paramsLeafListNoun, params)
		paramsLeafListNoun = append(paramsLeafListNoun, *jsonObj.LeafListNoun...)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacRolePermissionNoun", paramsLeafListNoun...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField, fmt.Sprintf("%s%s", parentPath, "/noun"), paramsLeafListNoun...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}
	//Property: { operation {string  map[ALL:ALL CREATE:CREATE READ:READ] [] false <nil> [] false} false false}
	if jsonObj.Operation != nil {

		paramsOperation := make([]string, len(params))
		copy(paramsOperation, params)
		paramsOperation = append(paramsOperation, *jsonObj.Operation)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacRolePermissionOperation", paramsOperation...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField,
			fmt.Sprintf("%s%s", parentPath, "/operation"), paramsOperation...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}
	//Property: { type {string  map[CONFIG:CONFIG GRPC:GRPC] [] false <nil> [] false} false false}
	if jsonObj.Type != nil {

		paramsType := make([]string, len(params))
		copy(paramsType, params)
		paramsType = append(paramsType, *jsonObj.Type)
		mpField, err := utils.CreateModelPluginObject(&mp, "RbacRolePermissionType", paramsType...)
		if err != nil {
			return nil, err
		}
		update, err := utils.UpdateForElement(mpField,
			fmt.Sprintf("%s%s", parentPath, "/type"), paramsType...)
		if err != nil {
			return nil, err
		}
		updates = append(updates, update)

	}

	if needKey {
		reflectKey, err := utils.FindModelPluginObject(mp, "RbacRolePermission", params...)
		if err != nil {
			return nil, err
		}
		reflectType := reflectKey.Type()
		reflect2 := reflect.New(reflectType) // Needed so the type can be read to extract list
		reflect2.Elem().Set(*reflectKey)
		keyMap, err := utils.ExtractGnmiListKeyMap(reflect2.Interface())
		if err != nil {
			return nil, err
		}
		for k, v := range keyMap {
			// parentPath = fmt.Sprintf("%s/{%s}", parentPath, k)
			for _, u := range updates {
				if err := utils.ReplaceUnknownKey(u, k, v, "unknown_key", "unknown_id"); err != nil {
					return nil, err
				}
			}
		}
	}
	return updates, nil
}

//Ignoring Target

//Ignoring RequestBodyRbac

//Ignoring RequestBodyRbacGroup

//Ignoring RequestBodyRbacGroupRole

//Ignoring RequestBodyRbacRole

//Ignoring RequestBodyRbacRolePermission

// Not generating param-types
// Not generating request-bodies
