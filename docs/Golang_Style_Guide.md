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
1. 注释并不是必须要求的，但是鉴于项目的可维护性，需要对重要的函数、对象、方法、变量加以说明。
2. 注释需要以被注释的对象名开始，并以句号结束。
3. 如果对象可数且无明确指定数量的情况下，一律使用单数形式和一般进行时描述。
4. 注释应当只使用"//"符号。
5. 注释一行应当不超过80个字符。
   
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

1. 导入包时，应该按名称或者功能分组，组之间应该空一行。	
		
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

2. 禁止使用相对路径导入（./subpackage），所有导入路径必须符合 go get 标准。
3. 包名应该全部小写。

## 命名

命名一般采用驼峰式命名法。

1. 类型  
	可导出类型必须用大写驼峰命名法，不可导出用小写驼峰命名法。

2. 变量  
  	变量名应该尽量短，尤其对在小作用域内（for、while）的本地变量命名时。
  	
  		Prefer c to lineCount. Prefer i to sliceIndex.
	更具描述性的名字应该用来命名常用变量、类型和全局变量。


## 错误处理
1. 不要使用panic。
2. 必须处理error。如果一个函数返回error，必须对它进行检查和进行相应处理，不能使用 _接收。

## 其他

1. 传递值  
	不要仅仅为了节省几bytes而传递指针作为函数参数。如果函数仅仅需要参数的值，那么就不要参数就不应该是指针。当然这不应用于大的结构体，或者一个可能增长的小结构体。
     
2. 接收器命名  
	一个方法的接收器名字应该反应接收器本身，通常是由一两个对小写字母组成的接收器类型缩写（如用"c"或者"cl"表示"Client"）。不要使用通用的名字（如"me"、"this"、"self"）作为接收器名。
      接收器的名字应当统一，不要在一个方法中用"c"，另一个方法中用"cl"。

3. 接收器类型  
	* 如果接收器是map，func，chan	，不要使用指针。
	* 如果接收器是一个切片(slice)并且方法不对该切片重建或者重新分配内存，不要使用指针。
	* 如果一个方法需要修改接收器，那么接收器必须是一个指针。
	* 如果接收器是一个包含sync.Mutex等相似同步域的结构体，接收器必须是一个指针，以防止拷贝。
	* 如果接收器是一个很大的结构体或者数组，一个指针类型的接收器会更加有效。这个很大可以指的是让人感觉很大。
	* 当一个方法被调用，接收器是一个值类型时，仅仅生成了原接收器的拷贝，方法内对接收器的任何更改都不会应用于原接收器。如果要让任何更改应用于原接收器，那么接收器必须是一个指针。
	* 如果接收器是一个struct，array 或者 slice，并且其任一元素都是一个指针，那么建议使用指针类型接收器，因为这可以让读者更清楚的明白其中的意图。
	* 如果接收器是一个小的值类型array或者struct，并且没有可修改域和指针，或者这个接受器仅仅是基本如int、string的类型，使用值类型接收器会更好。值类型接收器可以有效减少需要回收的垃圾。
	* 最后，当对使用值类型还是指针类型接收器犹豫不决时，请用指针类型。

