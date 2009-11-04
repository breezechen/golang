// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

5l is a modified version of the Plan 9 linker documented at

	http://plan9.bell-labs.com/magic/man2html/1/2l

Its target architecture is the ARM, referred to by these tools as arm.

Major changes include:
	- support for segmented stacks (this feature is implemented here, not in the compilers).


Original options are listed in the link above.

Options new in this version:
-L dir1,dir2,..
	Search for libraries (package files) in the comma-separated list of directories.
	The default is the single location $GOROOT/pkg/$GOOS_arm.


*/
package documentation
