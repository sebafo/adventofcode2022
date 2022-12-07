package main

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []*File
	size     int
}

type File struct {
	Directory *Directory
	Name      string
	Size      int
}

func (slice *Directory) addSubDirectory(dir *Directory) {
	if slice.Children == nil {
		slice.Children = []*Directory{}
	}
	if slice.containsDir(*dir) {
		return
	}
	slice.Children = append(slice.Children, dir)
}

func (slice *Directory) addFile(file *File) {
	if slice.Files == nil {
		slice.Files = []*File{}
	}
	if slice.containsFile(*file) {
		return
	}
	slice.Files = append(slice.Files, file)
}

func (slice Directory) getSubDirectoryByName(name string) *Directory {
	for _, child := range slice.Children {
		if child.Name == name {
			return child
		}
	}
	panic("Directory not found: " + name)
}

func (slice *Directory) getDirectorySize() int {
	if slice.size != 0 {
		return slice.size
	}

	size := 0
	for _, file := range slice.Files {
		size += file.Size
	}
	for _, child := range slice.Children {
		size += child.getDirectorySize()
	}

	slice.size = size
	return size
}

func (slice Directory) containsDir(val Directory) bool {
	for _, item := range slice.Children {
		if item.Name == val.Name {
			return true
		}
	}
	return false
}

func (slice Directory) containsFile(val File) bool {
	for _, item := range slice.Files {
		if item.Name == val.Name {
			return true
		}
	}
	return false
}
