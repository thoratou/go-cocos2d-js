package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// utilities to cope with the cocos2d-js Class implementation

func computeBackupName(name string) string {
	return "_gococos2d_backup_" + name
}

func findFuncInObjectHierarchy(object js.Object, name string) js.Object {
	if jsFunc := object.Get(name); jsFunc != nil && jsFunc != js.Undefined {
		return jsFunc
	} else if proto := object.Get("__proto__"); proto != nil && proto != js.Undefined {
		return findFuncInObjectHierarchy(proto, name)
	} else {
		return nil
	}
}

func findPureJsObjectInHierarchy(object js.Object) js.Object {
	if child := object.Get("Object"); child != nil && child != js.Undefined {
		return findPureJsObjectInHierarchy(child)
	} else {
		return object
	}
}

func backupFunc(object js.Object, name string) {
	backupName := computeBackupName(name)
	if object.Get(backupName) == js.Undefined {
		jsFunc := findFuncInObjectHierarchy(object, name)
		if jsFunc != nil {
			object.Set(backupName, jsFunc)
		}
	}
}

func BackupFunc(object js.Object, name string) {
	backupFunc(findPureJsObjectInHierarchy(object), name)
}

func superCall(object js.Object, name string, args ...interface{}) js.Object {
	backupName := computeBackupName(name)
	return object.Call(backupName, args)
}

func SuperCall(object js.Object, name string, args ...interface{}) js.Object {
	return superCall(findPureJsObjectInHierarchy(object), name, args)
}
