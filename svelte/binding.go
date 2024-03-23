package svelte

//#include "parser.h"
//TSLanguage *tree_sitter_svelte();
import "C"
import (
	sitter "github.com/Cyber-cicco/go-tree-sitter"
	"unsafe"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_svelte())
	return sitter.NewLanguage(ptr)
}
