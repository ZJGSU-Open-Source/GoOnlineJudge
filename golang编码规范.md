#golang编码规范
## 版权
1. 如何在程序中使用GNU许可证

无论使用哪种许可证，使用时需要在每个程序的源文件中添加两个元素：一个版权声明和一个复制许可声明，说明该程序使用GNU许可证进行授权。另外在声明版权前应该说明文件的名称以及用途，在复制许可声明之后，最好写上作者的联系信息，使得用户可以联系到你,如果对源文件进行了修改，最好使用简短的信息描述修改的内容。通用的格式如下所示：
        
        <one line to give the program's name and a brief idea of what it does.>
            Copyright (C) <year>  <name of author>
 
            This program is free software: you can redistribute it and/or modify
            it under the terms of the GNU General Public License as published by
            the Free Software Foundation, either version 3 of the License, or
            (at your option) any later version.
 
            This program is distributed in the hope that it will be useful,
            but WITHOUT ANY WARRANTY; without even the implied warranty of
            MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
            GNU General Public License for more details.
 
            You should have received a copy of the GNU General Public License
            along with this program.  If not, see http://www.gnu.org/licenses/.
版权声明中请务必使用英文单词“Copyright”，且其中的年份是完成发布版本的时间，如若1998年完成了发布版本直到1999年才发布出去，那么使用1998年。需要明确说明使用的是GNU GPL的那个版本，是版本2还是版本3。
## 注释
1. 注释并不是必须要求的，但是鉴于项目的可维护性，需要对重要的函数、对象、方法、变量加以说明.
2. 注释需要以被注释的对象名开始，并以句号结束。
3. 如果对象可数且无明确指定数量的情况下，一律使用单数形式和一般进行时描述。
4. 注释应当只使用"//"符号
5. 注释一行应当不超过80个字符
   
		// A Request represents a request to run a command.
		type Request struct { ...
		
		// Encode writes the JSON encoding of req to w.
		func Encode(w io.Writer, req *Request) { ...
		
6. 当对整个包进行文档注释时，需要有一个doc.go文件，用以描述包的功能，包括但不限于以下内容，版权申明，功能描述等。
		
		// Copyright 2009 The Go Authors. All rights reserved.
		// Use of this source code is governed by a BSD-style
		// license that can be found in the LICENSE file.
         
		// Package log implements a simple logging package. It defines a type, Logger,
		// with methods for formatting output. It also has a predefined 'standard'
		// Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and
		// Panic[f|ln], which are easier to use than creating a Logger manually.
		// That logger writes to standard error and prints the date and time
		// of each logged message.
		// The Fatal functions call os.Exit(1) after writing the log message.
		// The Panic functions call panic after writing the log message.
		package log
		
## 包

1. 导入包时，应该按名称或者功能分组，组之间应该空一行	
		
		package main
		
		import (
	          "fmt"
	          "hash/adler32"
	          "os"
	
	          "appengine/user"
	          "appengine/foo"
	
	          "code.google.com/p/x/y"
	          "github.com/foo/bar"
	      )

2. 禁止使用相对路径导入（./subpackage），所有导入路径必须符合 go get 标准
3. 包名应该全部小写

## 命名

命名一般采用驼峰式命名法

1. 类型  
	可导出类型应用大写驼峰命名法，不可导出用小写驼峰命名法。

2. 变量  
  	Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer c to lineCount. Prefer i to sliceIndex.

	The basic rule: the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual things and global variables need more descriptive names.

## 错误处理
1. 不要使用panic
2. 必须处理error。如果一个函数返回error，必须对它进行检查和进行相应处理，不能使用 _接收

## 其他

1. Pass Values  
      Don't pass pointers as function arguments just to save a few bytes. If a function refers to its argument x only as *x throughout, then the argument shouldn't be a pointer. Common instances of this include passing a pointer to a string (*string) or a pointer to an interface value (*io.Reader). In both cases the value itself is a fixed size and can be passed directly. This advice does not apply to large structs, or even small structs that might grow.

2. Receiver Names  
      The name of a method's receiver should be a reflection of its identity; often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client"). Don't use generic names such as "me", "this" or "self", identifiers typical of object-oriented languages that place more emphasis on methods as opposed to functions. The name need not be as descriptive as a that of a method argument, as its role is obvious and serves no documentary purpose. It can be very short as it will appear on almost every line of every method of the type; familiarity admits brevity. Be consistent, too: if you call the receiver "c" in one method, don't call it "cl" in another.

3. Receiver Type

&nbsp;&nbsp;&nbsp;Choosing whether to use a value or pointer receiver on methods can be difficult, especially to new Go programmers. If in doubt, use a pointer, but there are times when a value receiver makes sense, usually for reasons of efficiency, such as for small unchanging structs or values of basic type. Some rules of thumb:

&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a map, func or chan, don't use a pointer to it.  
&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a slice and the method doesn't reslice or reallocate the slice, don't use a pointer to it.  
&nbsp;&nbsp;&nbsp;&nbsp;If the method needs to mutate the receiver, the receiver must be a pointer.
&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a struct that contains a sync.Mutex or similar synchronizing field, the receiver must be a pointer to avoid copying.  
&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a large struct or array, a pointer receiver is more efficient. How large is large? Assume it's equivalent to passing all its elements as arguments to the method. If that feels too large, it's also too large for the receiver.   
&nbsp;&nbsp;&nbsp;&nbsp;Can function or methods, either concurrently or when called from this method, be mutating the receiver? A value type creates a copy of the receiver when the method is invoked, so outside updates will not be applied to this receiver. If changes must be visible in the original receiver, the receiver must be a pointer.  
&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a struct, array or slice and any of its elements is a pointer to something that might be mutating, prefer a pointer receiver, as it will make the intention more clear to the reader.  
&nbsp;&nbsp;&nbsp;&nbsp;If the receiver is a small array or struct that is naturally a value type (for instance, something like the time.Time type), with no mutable fields and no pointers, or is just a simple basic type such as int or string, a value receiver makes sense. A value receiver can reduce the amount of garbage that can be generated; if a value is passed to a value method, an on-stack copy can be used instead of allocating on the heap. (The compiler tries to be smart about avoiding this allocation, but it can't always succeed.) Don't choose a value receiver type for this reason without profiling first.  
&nbsp;&nbsp;&nbsp;Finally, when in doubt, use a pointer receiver.
