package builder

func TaskForEditorAssets(src string, dest string) {
	_PrepareDirectory(dest)
	_CopyDirectoryWithoutSymlink(src, dest)
}
