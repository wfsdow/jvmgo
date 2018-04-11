package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	//将路径list中的每个路径取出，创建对应的Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}

	return nil, nil, errors.New("no find class: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	//	for i, entry := range self {
	//		strs[i] := entry.String()
	//	}
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
