/*
 * trie.go
 * Trie
 *
 * Created by Jim Dovey on 16/07/2010.
 *
 * Copyright (c) 2010 Jim Dovey
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 * Redistributions of source code must retain the above copyright notice,
 * this list of conditions and the following disclaimer.
 *
 * Redistributions in binary form must reproduce the above copyright
 * notice, this list of conditions and the following disclaimer in the
 * documentation and/or other materials provided with the distribution.
 *
 * Neither the name of the project's author nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
 * TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

/*
	The trie package implements a basic character trie type. Instead of using bytes however, it uses
	integer-sized runes as traversal keys.  In Go, this means that each node refers to exactly one Unicode
	character, so the implementation doesn't depend on the particular semantics of UTF-8 byte streams.

	There is an additional specialization, which stores an integer value along with the Unicode character
	on each node.  This is to implement TeX-style hyphenation pattern storage.
*/
package trie

import (
	"strings"
	"container/vector"
	"utf8"
	"sort"
	"fmt"
	"path"
	"rand"
)

// A Trie uses runes rather than characters for indexing, therefore its child key values are integers.
type Trie struct {
	leaf     bool          // whether the node is a leaf (the end of an input string).
	value    interface{}   // the value associated with the string up to this leaf node.
	files map[string] string
	dirs  * vector.Vector
	children map[int]*Trie // a map if len(s) == 0 {
}

type inode struct {
	name        string
	permissions uint64
	size        uint64
	lock        bool
	chunks      *vector.Vector
}

// Creates and returns a new Trie instance.
func NewTrie() *Trie {
	t := new(Trie)
	t.leaf = false
	t.value = nil
	t.files = make(map[string]string)
	t.dirs = new(vector.Vector)
	t.children = make(map[int]*Trie)
	return t
}

func (p *Trie) GetDotString() string {
	rsrc := rand.NewSource(140209)
	rgen := rand.New(rsrc)
	
	outStr := "digraph Trie {"
	vec := new(vector.StringVector)
	
	p.outputDot(vec, 0, -1, rgen)
	
	cnt := vec.Len()
	for i := 0; i < cnt; i++ {
		outStr = strings.Join([]string{outStr, vec.At(i)}, "\n")
	}
	
	return strings.Join([]string{outStr, "}"}, "\n")
}

func (p *Trie) outputDot(vec *vector.StringVector, rune int, serial int64, rgen *rand.Rand) {
	this := make([]byte, 10)
	child := make([]byte, 10)

	utf8.EncodeRune(this, rune)
	
	thisChar := string(this[0])
	
	if serial == -1 {
		thisChar = "root"
	}
	
	
	for childRune, childNode := range p.children {
		utf8.EncodeRune(child, childRune)
		childSerial := rgen.Int63()
		childNodeStr := fmt.Sprintf("\"%s(%d)\"", string(child[0]), childSerial)
		var notation string
		
		if string(child[0]) == "/" {
			notation = fmt.Sprintf("[label=\"%s\" shape=box color=red]", string(child[0]))
		} else {
			notation = fmt.Sprintf("[label=\"%s\"]", string(child[0]))
		}
		vec.Push(fmt.Sprintf("\t%s %s\n\t\"%s(%d)\" -> \"%s(%d)\"", childNodeStr, notation, thisChar, serial, string(child[0]), childSerial))
		childNode.outputDot(vec, childRune, childSerial, rgen)
	}
}

// NEEDS FUNCTION TO VALIDATE path_s SYNTAX
func (p *Trie) AddFile(path_s string) bool{
	if len(path_s) == 0 {
		return false
	}
	
	dir, file := path.Split(path_s)
	
	// append the runes to the trie
	leaf := p.find(strings.NewReader(dir))
	if leaf == nil{
		return false
	}
    inode , check := leaf.files[file]
	if check && inode != "" {
		//file already exists
		return false
	} else {
		leaf.files[file] = file
	}
	return true


}
func (p *Trie) DeleteFile(path_s string) bool{
	if len(path_s) == 0{
		return false
	}
	
	dir, file := path.Split(path_s)
	
	leaf := p.find(strings.NewReader(dir))
	if leaf == nil{
		return false
	}
	inode , check := leaf.files[file]
	if !check || inode == "" {
		//file doesn't exist.. (maybe don't return?)
		return false
	}
	leaf.files[file] = ""
	
	//DELETE INODE IN MASTER???
	
	return true



}
func (p *Trie) GetFile(path_s string) (inode string, r bool) {
	if len(path_s) == 0 {
		return "", false
	}
	
	dir, file := path.Split(path_s)
	
	leaf := p.find(strings.NewReader(dir))
	if leaf == nil {
		//bad path_s
		return "", false
	}	
	inode, check := leaf.files[file]
	if !check || inode == ""{
		//bad file name
		return "", false
	}
	return leaf.files[file], true
}
func (p *Trie) AddDir(path_s string) bool {
	if len(path_s) == 0 {
		return false
	}
	
	directory_s, dir_name:= path.Split(path_s)
	
	//check to make sure the dir doesn't already exist
	leaf_test := p.find(strings.NewReader(path_s))
	if leaf_test != nil {
		//dir already exists
		return false
	
	}
	//create the dir
	p.addRunes(strings.NewReader(path_s))
	
	fmt.Printf("BLAH: %s\n", directory_s)
	
	//add dir record to parent dir
	dir := p.find(strings.NewReader(directory_s))
	dir.dirs.Push(dir_name)

	return true

}
func (p *Trie) ReadDir(path_s string) (dirs * vector.Vector, files map[string] string, r bool) {
	if len(path_s) == 0 {
		return nil, nil, false
	}
	path_cor := fmt.Sprintf("%s%s", path_s, "/")
	//CHECK path_s SYNTAX to make sure it ends on a forward slash..
	leaf := p.find(strings.NewReader(path_cor))
	if leaf == nil {
		//bad path_s
		return nil, nil, false
	}
	if leaf.dirs.Len() == 0 {
		return nil, nil, false
	}
	//TRAVERSE FILES STRUCTURE??
	return leaf.dirs, leaf.files, true

}
func (p *Trie) MoveDir(oldpath_s string, newpath_s string) bool {
	movingnode := p.includes(strings.NewReader(oldpath_s))
	if movingnode == nil {
		//invalid old path_s
		return false
	}
	//new_path_cor := fmt.Sprintf("%s%s", newpath_s, "/")
	old_path_cor := fmt.Sprintf("%s%s", oldpath_s, "/")
	overwrite_check := p.find(strings.NewReader(old_path_cor))
	if overwrite_check != nil {
		//new path already exists
		return false
	}
	//parse new path_s and old to get dir nodes
	//also get nodes to insert (1 char less than full path_s)
	//and the name keys of the new and old nodes..
	new_dir_s, new_name := path.Split(newpath_s)
	old_dir_s, old_name := path.Split(oldpath_s)
	
	new_length := len(newpath_s)
	old_length := len(oldpath_s)
	
	old_key := oldpath_s[old_length]
	new_key := newpath_s[new_length]
	
	old_parent_s := oldpath_s[:old_length - 1]
	new_parent_s := newpath_s[:new_length - 1]
	
	//get the new parent node, create if necessary
	newparent := p.addRunes(strings.NewReader(new_parent_s))
	oldparent := p.find(strings.NewReader(old_parent_s))
	if oldparent == nil {
		//what???
		return false
	}
	newdir := p.find(strings.NewReader(new_dir_s))
	if newdir == nil {
		//what the what??
		return false
	}
	
	//uniqueness check in newdir..
	unq_length := newdir.dirs.Len()
	for i := 0 ; i < unq_length ; i++ {
		if newdir.dirs.At(i).(string) == new_name {
			//failed uniquness check
			return false
		}
	}
	
	
	olddir := p.find(strings.NewReader(old_dir_s))
	if olddir == nil {
		//lies
		return false
	}
	
	//first add movingnode to new parent
	newparent.children[int(new_key)] = movingnode
	
	//then add it to the new directory
	newdir.dirs.Push(new_name)
	
	//then remove it from the old parent
	oldparent.children[int(old_key)] = nil
	if(oldparent.dirs.Len() == 0){
		oldparent.leaf = true
	}
	
	//then remove it from the old directory (expensive)
	remove_length := olddir.dirs.Len()
	for i := 0 ; i < remove_length ; i++ {
		if olddir.dirs.At(i).(string) == old_name {
			olddir.dirs.Delete(i)
		}
	}
	return true
	

}
// Internal function: adds items to the trie, reading runes from a strings.Reader.  It returns
// the leaf node at which the addition ends.
func (p *Trie) addRunes(r *strings.Reader) *Trie {
	rune, _, err := r.ReadRune()
	if err != nil {
		p.leaf = true
		return p
	}

	n := p.children[rune]
	if n == nil {
		n = NewTrie()
		p.children[rune] = n
	}

	// recurse to store sub-runes below the new node
	return n.addRunes(r)
}

// Adds a string to the trie. If the string is already present, no additional storage happens. Yay!
func (p *Trie) AddString(s string) {
	if len(s) == 0 {
		return
	}

	// append the runes to the trie -- we're ignoring the value in this invocation
	p.addRunes(strings.NewReader(s))
}

// Adds a string to the trie, with an associated value.  If the string is already present, only
// the value is updated.
func (p *Trie) AddValue(s string, v interface{}) {
	if len(s) == 0 {
		return
	}

	// append the runes to the trie
	leaf := p.addRunes(strings.NewReader(s))
	leaf.value = v
}

// Internal string removal function.  Returns true if this node is empty following the removal.
func (p *Trie) removeRunes(r *strings.Reader) bool {
	rune, _, err := r.ReadRune()
	if err != nil {
		// remove value, remove leaf flag
		p.value = nil
		p.leaf = false
		return len(p.children) == 0
	}

	child, ok := p.children[rune]
	if ok && child.removeRunes(r) {
		// the child is now empty following the removal, so prune it
		p.children[rune] = nil, false
	}

	return len(p.children) == 0
}

// Remove a string from the trie.  Returns true if the Trie is now empty.
func (p *Trie) Remove(s string) bool {
	if len(s) == 0 {
		return len(p.children) == 0
	}

	// remove the runes, returning the final result
	return p.removeRunes(strings.NewReader(s))
}

func (p *Trie) find(r *strings.Reader) *Trie {
	rune, _, err := r.ReadRune()
	if err != nil {
		//if p.leaf {
			return p
		//}
		//return nil
	}

	child, ok := p.children[rune]
	if !ok {
		return nil // no node for this rune was in the trie
	}

	// recurse down to the next node with the remainder of the string
	return child.includes(r)
}




// Internal string inclusion function.
func (p *Trie) includes(r *strings.Reader) *Trie {
	rune, _, err := r.ReadRune()
	if err != nil {
		if p.leaf {
			return p
		}
		return nil
	}

	child, ok := p.children[rune]
	if !ok {
		return nil // no node for this rune was in the trie
	}

	// recurse down to the next node with the remainder of the string
	return child.includes(r)
}

// Test for the inclusion of a particular string in the Trie.
func (p *Trie) Contains(s string) bool {
	if len(s) == 0 {
		return false // empty strings can't be included (how could we add them?)
	}
	return p.includes(strings.NewReader(s)) != nil
}

// Return the value associated with the given string.  Double return: false if the given string was
// not present, true if the string was present.  The value could be both valid and nil.
func (p *Trie) GetValue(s string) (interface{}, bool) {
	if len(s) == 0 {
		return nil, false
	}

	leaf := p.includes(strings.NewReader(s))
	if leaf == nil {
		return nil, false
	}
	return leaf.value, true
}

// Internal output-building function used by Members()
func (p *Trie) buildMembers(prefix string) *vector.StringVector {
	strList := new(vector.StringVector)

	if p.leaf {
		strList.Push(prefix)
	}

	// for each child, go grab all suffixes
	for rune, child := range p.children {
		buf := make([]byte, 4)
		numChars := utf8.EncodeRune(buf, rune)
		strList.AppendVector(child.buildMembers(prefix + string(buf[0:numChars])))
	}

	return strList
}

// Retrieves all member strings, in order.
func (p *Trie) Members() (members *vector.StringVector) {
	members = p.buildMembers(``)
	sort.Sort(members)
	return
}

// Introspection -- counts all the nodes of the entire Trie, NOT including the root node.
func (p *Trie) Size() (sz int) {
	sz = len(p.children)

	for _, child := range p.children {
		sz += child.Size()
	}

	return
}

// Return all anchored substrings of the given string within the Trie.
func (p *Trie) AllSubstrings(s string) *vector.StringVector {
	v := new(vector.StringVector)

	for pos, rune := range s {
		child, ok := p.children[rune]
		if !ok {
			// return whatever we have so far
			break
		}

		// if this is a leaf node, add the string so far to the output vector
		if child.leaf {
			v.Push(s[0:pos])
		}

		p = child
	}

	return v
}

// Return all anchored substrings of the given string within the Trie, with a matching set of
// their associated values.
func (p *Trie) AllSubstringsAndValues(s string) (*vector.StringVector, *vector.Vector) {
	sv := new(vector.StringVector)
	vv := new(vector.Vector)
	
	for pos, rune := range s {
		child, ok := p.children[rune]
		if !ok {
			// return whatever we have so far
			break
		}
		
		// if this is a leaf node, add the string so far and its value
		if child.leaf {
			sv.Push(s[0:pos+utf8.RuneLen(rune)])
			vv.Push(child.value)
		}
		
		p = child
	}
	
	return sv, vv
}
