// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package utils

import (
	aether_2_1_0 "github.com/onosproject/aether-models/models/aether-2.1.x/api"
	"github.com/openconfig/gnmi/proto/gnmi"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func Test_NewGnmiGetRequest(t *testing.T) {
	gnmiGet, err := NewGnmiGetRequest("/rbac/v1.0.0/{target}/rbac/group/role/{roleid}", "internal", "r1")
	assert.NilError(t, err, "unexpected error handling path")

	assert.Equal(t, 1, len(gnmiGet.Path), "expected only one path")
	path0 := gnmiGet.Path[0]
	assert.Equal(t, "internal", path0.Target)
	assert.Equal(t, 3, len(path0.Elem), "expected 3 path elems")
	assert.Equal(t, "rbac", path0.Elem[0].Name)
	assert.Equal(t, "group", path0.Elem[1].Name)
	assert.Equal(t, "role", path0.Elem[2].Name)
	assert.Equal(t, 1, len(path0.Elem[2].Key))
	key2, ok := path0.Elem[2].Key["roleid"]
	assert.Assert(t, ok)
	assert.Equal(t, "r1", key2)
}

func Test_GetResponseUpdate(t *testing.T) {
	path0Elems := make([]*gnmi.PathElem, 0)
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe1"})
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe2"})
	path0Elems = append(path0Elems, &gnmi.PathElem{Name: "pe3"})
	path0 := gnmi.Path{
		Elem:   path0Elems,
		Target: "internal",
	}
	u0 := gnmi.Update{
		Path: &path0,
		Val: &gnmi.TypedValue{
			Value: &gnmi.TypedValue_JsonVal{JsonVal: []byte("{testvalue: 't'}")},
		},
	}
	n0 := gnmi.Notification{
		Update: []*gnmi.Update{
			&u0,
		},
	}

	gr := gnmi.GetResponse{
		Notification: []*gnmi.Notification{
			&n0,
		},
	}

	typedVal, err := GetResponseUpdate(&gr, nil)
	assert.NilError(t, err, "unexpected error")
	jsonVal, ok := typedVal.Value.(*gnmi.TypedValue_JsonVal)
	assert.Assert(t, ok, "expecting to cast to JsonVal")
	assert.Equal(t, "{testvalue: 't'}", string(jsonVal.JsonVal))
}

func Test_buildElems(t *testing.T) {
	pathElems, err := BuildElems(
		"/rbac/v1.0.0/{target}/rbac/role/{roleid}", 4, "role-1")
	assert.NilError(t, err)
	assert.Equal(t, 2, len(pathElems))
	elem0 := pathElems[0]
	assert.Equal(t, "rbac", elem0.Name)
	assert.Equal(t, 0, len(elem0.Key))
	elem1 := pathElems[1]
	assert.Equal(t, "role", elem1.Name)
	assert.Equal(t, 1, len(elem1.Key))
	key1, ok := elem1.Key["roleid"]
	assert.Assert(t, ok)
	assert.Equal(t, "role-1", key1)
}

func Test_updateForElement(t *testing.T) {
	desc := "this is a description"
	gnmiUpdate, err := UpdateForElement(
		&desc, "/test1/test2/{name}", "t1")
	assert.NilError(t, err, "unexpected error")
	assert.Assert(t, gnmiUpdate != nil)
	if gnmiUpdate != nil {
		assert.Equal(t, 2, len(gnmiUpdate.Path.Elem))
		elem0 := gnmiUpdate.Path.Elem[0]
		assert.Equal(t, "test1", elem0.Name)
		assert.Equal(t, 0, len(elem0.Key))
		elem1 := gnmiUpdate.Path.Elem[1]
		assert.Equal(t, "test2", elem1.Name)
		assert.Equal(t, 1, len(elem1.Key))
		key1, ok := elem1.Key["name"]
		assert.Assert(t, ok)
		assert.Equal(t, "t1", key1)
		assert.Equal(t, desc, gnmiUpdate.Val.GetStringVal())
	}
}

func Test_ReplaceUnknownKey(t *testing.T) {
	desc := "this is a description"
	gnmiUpdate, err := UpdateForElement(
		&desc,
		"/test1/test2/{"+UnknownKey+"}", UnknownID)
	assert.NilError(t, err)
	assert.Assert(t, gnmiUpdate != nil)
	if gnmiUpdate != nil {
		keyID, ok := gnmiUpdate.Path.Elem[1].Key[UnknownKey]
		assert.Equal(t, true, ok)
		assert.Equal(t, UnknownID, keyID)
		err = ReplaceUnknownKey(gnmiUpdate, "known_key", "known_value", UnknownKey, UnknownID)
		assert.NilError(t, err, "unexpected error")
		keyID, ok = gnmiUpdate.Path.Elem[1].Key["known_key"]
		assert.Equal(t, true, ok)
		assert.Equal(t, "known_value", keyID)
	}
}

func Test_CreateModelPluginObject_ListInList(t *testing.T) {
	device := new(aether_2_1_0.Device)
	dg1, err := CreateModelPluginObject(device, "SiteSliceDeviceGroupDeviceGroup", "s1", "sl1", "dg1", "dg1-ref")
	assert.NilError(t, err)
	assert.Assert(t, dg1 != nil)

	dg1Obj, ok := dg1.(*string)
	assert.Assert(t, ok)
	assert.Equal(t, "dg1-ref", *dg1Obj)
}

// Test the /Device and /DeviceGroup in 2.1.0
func Test_CreateModelPluginObject_SimilarNameStub(t *testing.T) {
	device := new(aether_2_1_0.Device)
	dg1, err := CreateModelPluginObject(device, "SiteDeviceGroupMbrUplink", "s1", "dg11", "10")
	assert.NilError(t, err)
	assert.Assert(t, dg1 != nil)

	// Can it cope with existing keys
	dg1, err = CreateModelPluginObject(device, "SiteDeviceGroupMbrDownlink", "s1", "dg1", "20")
	assert.NilError(t, err)
	assert.Assert(t, dg1 != nil)

	t1, err := CreateModelPluginObject(device, "TemplateMbrDownlinkBurstSize", "v1", "20")
	assert.NilError(t, err)
	assert.Assert(t, t1 != nil)

	dg1Obj, ok := dg1.(*uint64)
	assert.Assert(t, ok)
	assert.Equal(t, uint64(20), *dg1Obj)

	_, ok = device.Site["s1"]
	assert.Assert(t, ok)
	assert.Equal(t, 1, len(device.Site["s1"].DeviceGroup))
	dgV1, ok := device.Site["s1"].DeviceGroup["dg11"]
	assert.Assert(t, ok)
	assert.Equal(t, uint64(10), *dgV1.Mbr.Uplink)
	assert.Equal(t, uint64(20), *dgV1.Mbr.Downlink)
}

// TODO: uncomment this when it's possible to handle the number structures in the name
//func Test_CreateModelPluginObject_DoubleKey(t *testing.T) {
//	device := new(testdevice_1_0_0.Device)
//	dg1, err := CreateModelPluginObject(device, "Cont1AList5Key1", "k1 10", "k1")
//	assert.NilError(t, err)
//	assert.Assert(t, dg1 != nil)
//
//	dg1, err = CreateModelPluginObject(device, "Cont1AList5Leaf5A", "k1 10", "leaf5a-val")
//	assert.NilError(t, err)
//	assert.Assert(t, dg1 != nil)
//
//	assert.Equal(t, 1, len(device.Cont1A.List5))
//	for k, v := range device.Cont1A.List5 {
//		assert.Equal(t, "{k1 10}", fmt.Sprintf("%v", k))
//		assert.Equal(t, "k1", *v.Key1)
//	}
//
//	leaf5aObj, ok := dg1.(*string)
//	assert.Assert(t, ok)
//	assert.Equal(t, string("leaf5a-val"), *leaf5aObj)
//}

// TODO: uncomment this when it's possible to handle the number structures in the name
//func Test_CreateModelPluginObject_UintSingleKey(t *testing.T) {
//	device := new(testdevice_1_0_0.Device)
//	dg1, err := CreateModelPluginObject(device, "Cont1BStateList2BIndex", "10", "10")
//	assert.NilError(t, err)
//	assert.Assert(t, dg1 != nil)
//
//	dg1, err = CreateModelPluginObject(device, "Cont1BStateList2BLeaf3C", "10", "leaf3c-val")
//	assert.NilError(t, err)
//	assert.Assert(t, dg1 != nil)
//
//	leaf3cObj, ok := dg1.(*string)
//	assert.Assert(t, ok)
//	assert.Equal(t, string("leaf3c-val"), *leaf3cObj)
//}

func Test_ApplEndpoint(t *testing.T) {
	device := new(aether_2_1_0.Device)
	app1, err := CreateModelPluginObject(device, "ApplicationEndpointProtocol", "app1", "ep1", "test-url")
	assert.NilError(t, err)
	assert.Assert(t, app1 != nil)

	dg1Obj, ok := app1.(*string)
	assert.Assert(t, ok)
	assert.Equal(t, "test-url", *dg1Obj)
}

func Test_FindModelPluginObject_Application(t *testing.T) {
	device := new(aether_2_1_0.Device)
	csID := "app1"
	core5gEp := "core5gEp"
	device.Application = map[string]*aether_2_1_0.OnfApplication_Application{
		csID: {
			ApplicationId: &csID,
			Address:       &core5gEp,
		},
	}

	params := []string{csID}

	core5gEpReflect, err := FindModelPluginObject(device, "ApplicationAddress", params...)
	assert.NilError(t, err)
	assert.Assert(t, core5gEpReflect != nil)
	assert.Equal(t, core5gEp, core5gEpReflect.Interface())
}

func Test_FindModelPluginObject_Template(t *testing.T) {
	device := new(aether_2_1_0.Device)
	tID := "t1"
	sst := uint8(123)
	dl := uint64(1000000)
	dlBs := uint32(2000000)
	device.Template = map[string]*aether_2_1_0.OnfTemplate_Template{
		tID: {
			TemplateId: &tID,
			Mbr: &aether_2_1_0.OnfTemplate_Template_Mbr{
				Downlink:          &dl,
				DownlinkBurstSize: &dlBs,
			},
			Sst: &sst,
		},
	}

	params := []string{tID}

	sstReflect, err := FindModelPluginObject(device, "TemplateSst", params...)
	assert.NilError(t, err)
	assert.Assert(t, sstReflect != nil)
	assert.Equal(t, sst, sstReflect.Interface())

	dlReflect, err := FindModelPluginObject(device, "TemplateMbrDownlink", params...)
	assert.NilError(t, err)
	assert.Assert(t, dlReflect != nil)
	assert.Equal(t, dl, dlReflect.Interface())

	// This is an important new case because "DownlinkBurstSize" has the same root as "Downlink"
	dlBsReflect, err := FindModelPluginObject(device, "TemplateMbrDownlinkBurstSize", params...)
	assert.NilError(t, err)
	assert.Assert(t, dlBsReflect != nil)
	assert.Equal(t, dlBs, dlBsReflect.Interface())

}

func Test_findChildByParamName(t *testing.T) {
	mpType := reflect.TypeOf(&aether_2_1_0.OnfSite_Site{})
	pathParts := []string{"Device", "Group", "Device", "Device", "Id"}
	field, skipped, err := findChildByParamNames(mpType, pathParts)
	assert.NilError(t, err)
	assert.Equal(t, "DeviceGroup", field.Name)
	assert.Equal(t, 1, skipped)

}

func Test_findChildByParamName_5GCore(t *testing.T) {
	mpType := reflect.TypeOf(&aether_2_1_0.OnfSite_Site_ConnectivityService{})
	pathParts := []string{"Core", "5G"}
	field, skipped, err := findChildByParamNames(mpType, pathParts)
	assert.NilError(t, err)
	assert.Equal(t, "Core_5G", field.Name)
	assert.Equal(t, 1, skipped)
}

func Test_ExtractGnmiEnumMap(t *testing.T) {
	gnmiElementValue := newSlice()

	name, def, err := ExtractGnmiEnumMap(gnmiElementValue, "SiteSliceConnectivityService", aether_2_1_0.OnfSite_Site_Slice_ConnectivityService_5g)
	assert.NilError(t, err)
	assert.Equal(t, "E_OnfSite_Site_Slice_ConnectivityService", name)
	assert.Equal(t, "5g", def.Name)
}

func Test_ExtractGnmiEnumMapUnset(t *testing.T) {
	gnmiElementValue := newSlice()

	name, def, err := ExtractGnmiEnumMap(gnmiElementValue, "SiteSliceConnectivityService", aether_2_1_0.OnfSite_Site_Slice_ConnectivityService_UNSET)
	assert.NilError(t, err)
	assert.Equal(t, "E_OnfSite_Site_Slice_ConnectivityService", name)
	assert.Assert(t, def == nil)
}

func newSlice() *reflect.Value {
	sliceID := "slice-1"

	slice := aether_2_1_0.OnfSite_Site_Slice{
		SliceId: &sliceID,
	}
	val := reflect.ValueOf(slice)
	return &val
}
