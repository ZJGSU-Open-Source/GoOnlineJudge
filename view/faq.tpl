{{define "content"}}
<h1>FAQ</h1>
<h2>评分</h2>
<p>试题的解答提交后由评分系统评出即时得分，每一次提交会判决结果会及时通知；系统可能的反馈信息包括：</p>
<table id="FAQ_list">
  <thead>
    <tr>
      <th class="lheader">结果</th>
      <th class="lheader">说明</th>
    </tr>
  </thead>
  <tbody>
        <tr>
          <td class="yellow">Pending</td>
          <td>评测系统还没有评测到这个提交，请稍候</td>
        </tr>

        <tr>
          <td class="blue">Running&Judging</td>
          <td>评测系统正在评测，稍候会有结果</td>
        </tr>

        <tr>
          <td class="purple">Compiler Error</td>
          <td>您提交的代码无法完成编译，点击“编译错误”可以看到编译器输出的错误信息</td>
        </tr>

        <tr>
          <td class="red">Accept</td>
          <td>恭喜！您通过了这道题</td>
        </tr> 

        <tr>
          <td class="green">Print Error</td>
          <td>您的程序输出的格式不符合要求（比如空格和换行与要求不一致）</td>
        </tr> 

        <tr>
          <td class="green">Runtime Error</td>
          <td>您的程序发生段错误，可能是数组越界，堆栈溢出（比如，递归调用层数太多）等情况引起</td>
        </tr> 

        <tr>
          <td class="green">Wrong Error</td>
          <td>您的程序未能对评测系统的数据返回正确的结果</td>
        </tr> 

        <tr>
          <td class="green">Time Limit Exceeded</td>
          <td>您的程序未能在规定时间内运行结束</td>
        </tr> 

        <tr>
          <td class="green">Memory Limit Exceeded</td>
          <td>您的程序使用了超过限制的内存</td>
        </tr> 

        <tr>
          <td class="green">Output Limit Exceeded</td>
          <td>您的程序输出超限</td>
        </tr> 

  </tbody>
</table>
<style>
 
 
.body-classic{
  color:#444;
  font-family:Georgia, Palatino, 'Palatino Linotype', Times, 'Times New Roman', "Hiragino Sans GB", "STXihei", "微软雅黑", serif;
  font-size:16px;
  line-height:1.5em;
  background:#fefefe;
  width: 45em;
  margin: 10px auto;
  padding: 1em;
  outline: 1300px solid #FAFAFA;
}
 
body>:first-child
{
  margin-top:0!important;
}
 
body>:last-child
{
  margin-bottom:0!important;
}
 
blockquote,dl,ol,p,pre,table,ul {
  border: 0;
  margin: 15px 0;
  padding: 0;
}
 
body a {
  color: #4183c4;
  text-decoration: none;
}
 
body a:hover {
  text-decoration: underline;
}
 
body a.absent
{
  color:#c00;
}
 
body a.anchor
{
  display:block;
  padding-left:30px;
  margin-left:-30px;
  cursor:pointer;
  position:absolute;
  top:0;
  left:0;
  bottom:0
}
 
/*h4,h5,h6{ font-weight: bold; }*/
 
.octicon{
  font:normal normal 16px sans-serif;
  width: 1em;
  height: 1em;
  line-height:1;
  display:inline-block;
  text-decoration:none;
  -webkit-font-smoothing:antialiased
}
 
.octicon-link {
  background: url("data:image/svg+xml;utf8,<?xml version='1.0' standalone='no'?> <!DOCTYPE svg PUBLIC '-//W3C//DTD SVG 1.1//EN' 'http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd'> <svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 1024 832'> <metadata>Copyright (C) 2013 by GitHub</metadata> <!-- scale(0.01565557729941) --> <path transform='' d='M768 64h-192s-254 0-256 256c0 22 3 43 8 64h137c-11-19-18-41-18-64 0-128 128-128 128-128h192s128 0 128 128-128 128-128 128 0 64-64 128h64s256 0 256-256-256-256-256-256z m-72 192h-137c11 19 18 41 18 64 0 128-128 128-128 128h-192s-128 0-128-128 128-128 128-128-4-65 66-128h-66s-256 0-256 256 256 256 256 256h192s256 0 256-256c0-22-4-44-8-64z'/> </svg>");
  background-size: contain;
  background-repeat: no-repeat;
  background-position: bottom;
}
 
.octicon-link:before{
  content:'\a0';
}
 
body h1,body h2,body h3,body h4,body h5,body h6{
  margin:1em 0 15px;
  padding:0;
  font-weight:bold;
  line-height:1.7;
  cursor:text;
  position:relative
}
 
body h1 .octicon-link,body h2 .octicon-link,body h3 .octicon-link,body h4 .octicon-link,body h5 .octicon-link,body h6 .octicon-link{
  display:none;
  color:#000
}
 
body h1:hover a.anchor,body h2:hover a.anchor,body h3:hover a.anchor,body h4:hover a.anchor,body h5:hover a.anchor,body h6:hover a.anchor{
  text-decoration:none;
  line-height:1;
  padding-left:0;
  margin-left:-22px;
  top:15%
}
 
body h1:hover a.anchor .octicon-link,body h2:hover a.anchor .octicon-link,body h3:hover a.anchor .octicon-link,body h4:hover a.anchor .octicon-link,body h5:hover a.anchor .octicon-link,body h6:hover a.anchor .octicon-link{
  display:inline-block
}
 
body h1 tt,body h1 code,body h2 tt,body h2 code,body h3 tt,body h3 code,body h4 tt,body h4 code,body h5 tt,body h5 code,body h6 tt,body h6 code{
  font-size:inherit
}
 
body h1{
  font-size:2.5em;
  border-bottom:1px solid #ddd
}
 
body h2{
  font-size:2em;
  border-bottom:1px solid #eee
}
 
body h3{
  font-size:1.5em
}
 
body h4{
  font-size:1.2em
}
 
body h5{
  font-size:1em
}
 
body h6{
  color:#777;
  font-size:1em
}
 
body p,body blockquote,body ul,body ol,body dl,body table,body pre{
  margin:15px 0
}
 
body h1 tt,body h1 code,body h2 tt,body h2 code,body h3 tt,body h3 code,body h4 tt,body h4 code,body h5 tt,body h5 code,body h6 tt,body h6 code
{
  font-size:inherit;
}

body li p.first
{
  display:inline-block;
}
 
body ul,body ol
{
  padding-left:30px;
}
 
body ul.no-list,body ol.no-list
{
  list-style-type:none;
  padding:0;
}
 
body ul ul,body ul ol,body ol ol,body ol ul
{
  margin-bottom:0;
  margin-top:0;
}
 
body dl
{
  padding:0;
}
 
body dl dt
{
  font-size:14px;
  font-style:italic;
  font-weight:700;
  margin-top:15px;
  padding:0;
}
 
body dl dd
{
  margin-bottom:15px;
  padding:0 15px;
}
 
body blockquote
{
  border-left:4px solid #DDD;
  color:#777;
  padding:0 15px;
}
 
body blockquote>:first-child
{
  margin-top:0;
}
 
body blockquote>:last-child
{
  margin-bottom:0;
}
 
body table
{
  display:block;
  overflow:auto;
  width:100%;
  border-collapse: collapse;
  border-spacing: 0;
  padding: 0;
}
 
body table th
{
  font-weight:700;
}
 
body table th,body table td
{
  border:1px solid #ddd;
  padding:6px 13px;
}
 
body table tr
{
  background-color:#fff;
  border-top:1px solid #ccc;
}
 
body table tr:nth-child(2n)
{
  background-color:#f8f8f8;
}
 
body img
{
  -moz-box-sizing:border-box;
  box-sizing:border-box;
  max-width:100%;
}
 
body span.frame
{
  display:block;
  overflow:hidden;
}
 
body span.frame>span
{
  border:1px solid #ddd;
  display:block;
  float:left;
  margin:13px 0 0;
  overflow:hidden;
  padding:7px;
  width:auto;
}
 
body span.frame span img
{
  display:block;
  float:left;
}
 
body span.frame span span
{
  clear:both;
  color:#333;
  display:block;
  padding:5px 0 0;
}
 
body span.align-center
{
  clear:both;
  display:block;
  overflow:hidden;
}
 
body span.align-center>span
{
  display:block;
  margin:13px auto 0;
  overflow:hidden;
  text-align:center;
}
 
body span.align-center span img
{
  margin:0 auto;
  text-align:center;
}
 
body span.align-right
{
  clear:both;
  display:block;
  overflow:hidden;
}
 
body span.align-right>span
{
  display:block;
  margin:13px 0 0;
  overflow:hidden;
  text-align:right;
}
 
body span.align-right span img
{
  margin:0;
  text-align:right;
}
 
body span.float-left
{
  display:block;
  float:left;
  margin-right:13px;
  overflow:hidden;
}
 
body span.float-left span
{
  margin:13px 0 0;
}
 
body span.float-right
{
  display:block;
  float:right;
  margin-left:13px;
  overflow:hidden;
}
 
body span.float-right>span
{
  display:block;
  margin:13px auto 0;
  overflow:hidden;
  text-align:right;
}
 
body code,body tt
{
  background-color:#f8f8f8;
  border:1px solid #ddd;
  border-radius:3px;
  margin:0 2px;
  padding:0 5px;
}
 
body code
{
  white-space:nowrap;
}
 
 
code,pre{
  font-family:Consolas, "Liberation Mono", Courier, monospace;
  font-size:12px
}
 
body pre>code
{
  background:transparent;
  border:none;
  margin:0;
  padding:0;
  white-space:pre;
}
 
body .highlight pre,body pre
{
  background-color:#f8f8f8;
  border:1px solid #ddd;
  font-size:13px;
  line-height:19px;
  overflow:auto;
  padding:6px 10px;
  border-radius:3px
}
 
body pre code,body pre tt
{
  background-color:transparent;
  border:none;
  margin:0;
  padding:0;
}
 
body .task-list{
  list-style-type:none;
  padding-left:10px
}
 
.task-list-item{
  padding-left:20px
}
 
.task-list-item label{
  font-weight:normal
}
 
.task-list-item.enabled label{
  cursor:pointer
}
 
.task-list-item+.task-list-item{
  margin-top:5px
}
 
.task-list-item-checkbox{
  float:left;
  margin-left:-20px;
  margin-top:7px
}
 
 
.highlight{
  background:#ffffff
}
 
.highlight .c{
  color:#999988;
  font-style:italic
}
 
.highlight .err{
  color:#a61717;
  background-color:#e3d2d2
}
 
.highlight .k{
  font-weight:bold
}
 
.highlight .o{
  font-weight:bold
}
 
.highlight .cm{
  color:#999988;
  font-style:italic
}
 
.highlight .cp{
  color:#999999;
  font-weight:bold
}
 
.highlight .c1{
  color:#999988;
  font-style:italic
}
 
.highlight .cs{
  color:#999999;
  font-weight:bold;
  font-style:italic
}
 
.highlight .gd{
  color:#000000;
  background-color:#ffdddd
}
 
.highlight .gd .x{
  color:#000000;
  background-color:#ffaaaa
}
 
.highlight .ge{
  font-style:italic
}
 
.highlight .gr{
  color:#aa0000
}
 
.highlight .gh{
  color:#999999
}
 
.highlight .gi{
  color:#000000;
  background-color:#ddffdd
}
 
.highlight .gi .x{
  color:#000000;
  background-color:#aaffaa
}
 
.highlight .go{
  color:#888888
}
 
.highlight .gp{
  color:#555555
}
 
.highlight .gs{
  font-weight:bold
}
 
.highlight .gu{
  color:#800080;
  font-weight:bold
}
 
.highlight .gt{
  color:#aa0000
}
 
.highlight .kc{
  font-weight:bold
}
 
.highlight .kd{
  font-weight:bold
}
 
.highlight .kn{
  font-weight:bold
}
 
.highlight .kp{
  font-weight:bold
}
 
.highlight .kr{
  font-weight:bold
}
 
.highlight .kt{
  color:#445588;
  font-weight:bold
}
 
.highlight .m{
  color:#009999
}
 
.highlight .s{
  color:#d14
}
 
.highlight .n{
  color:#333333
}
 
.highlight .na{
  color:#008080
}
 
.highlight .nb{
  color:#0086B3
}
 
.highlight .nc{
  color:#445588;
  font-weight:bold
}
 
.highlight .no{
  color:#008080
}
 
.highlight .ni{
  color:#800080
}
 
.highlight .ne{
  color:#990000;
  font-weight:bold
}
 
.highlight .nf{
  color:#990000;
  font-weight:bold
}
 
.highlight .nn{
  color:#555555
}
 
.highlight .nt{
  color:#000080
}
 
.highlight .nv{
  color:#008080
}
 
.highlight .ow{
  font-weight:bold
}
 
.highlight .w{
  color:#bbbbbb
}
 
.highlight .mf{
  color:#009999
}
 
.highlight .mh{
  color:#009999
}
 
.highlight .mi{
  color:#009999
}
 
.highlight .mo{
  color:#009999
}
 
.highlight .sb{
  color:#d14
}
 
.highlight .sc{
  color:#d14
}
 
.highlight .sd{
  color:#d14
}
 
.highlight .s2{
  color:#d14
}
 
.highlight .se{
  color:#d14
}
 
.highlight .sh{
  color:#d14
}
 
.highlight .si{
  color:#d14
}
 
.highlight .sx{
  color:#d14
}
 
.highlight .sr{
  color:#009926
}
 
.highlight .s1{
  color:#d14
}
 
.highlight .ss{
  color:#990073
}
 
.highlight .bp{
  color:#999999
}
 
.highlight .vc{
  color:#008080
}
 
.highlight .vg{
  color:#008080
}
 
.highlight .vi{
  color:#008080
}
 
.highlight .il{
  color:#009999
}
 
.highlight .gc{
  color:#999;
  background-color:#EAF2F5
}
 
.type-csharp .highlight .k{
  color:#0000FF
}
 
.type-csharp .highlight .kt{
  color:#0000FF
}
 
.type-csharp .highlight .nf{
  color:#000000;
  font-weight:normal
}
 
.type-csharp .highlight .nc{
  color:#2B91AF
}
 
.type-csharp .highlight .nn{
  color:#000000
}
 
.type-csharp .highlight .s{
  color:#A31515
}
 
.type-csharp .highlight .sc{
  color:#A31515
}
</style>


<h2>常见问题</h2>
<p><b>我应该从哪里读输入，另外应该输出到哪里？</b></p>
<p>如果没有特别说明，你的程序应该从标准输入（stdin，传统意义上的“键盘”）读入，并输出到标准输出（stdout，传统意义上的“屏幕”），不要使用文件做输入输出。由于系统是在你的程序运行结束后开始检查输出是否是正确的，对于有多组测试数据的输入，可以全部读入之后再输出，也可以处理一组测试数据就输出一组。</p>
<br>
<p><b>为什么我的程序交在这里得到编译错误，而我在自己的机器上已经编译通过了？</b></p>
<p>本系统所使用的编译器和你在自己机器上使用的可能有区别，请留意几个常见的地方：</p>
<p><li>本系统是 64 位 Linux 系统，使用的编译器版本和编译参数可以参见编译器帮助</li></p>
<p><li>Java 代码需使用 Main 作为主类名</li></p>
<p><li>Visual C++ 6.0 和 Turbo C++ 3.0 （及它们的更低版本）有较多违背 C++ 标准（<a href="http://www.iso.org/iso/iso_catalogue/catalogue_ics/catalogue_detail_ics.htm?ics1=35&ics2=60&ics3=&csnumber=50372" target="new">ISO/IEC 14882</a>）的地方，不要使用它们来判断 C++ 程序语法上是否有问题</li></p>
<p><li>C++ 下 64 位整数的类型是 long long，不要使用 __int64</li></p>
<br>
<p><b>为什么我的程序得到了“返回非零”？</b></p>
<p><li>返回零表示一个程序正常结束，如果没有返回零，则系统认为程序没有正常结束，这时即便输出了正确的内容也不予通过</li></p>
<p><li>C/C++ 代码请确认 int main 函数最终会返回 0，不要声明为 double main 或者 void main</li></p>
<p><li>有异常的语言，请确认程序处理了可能抛出的异常</li></p>
<br>
<p><b>程序的时间和内存占用是如何计算的？</b></p>
<p>程序的运行时间为程序在所有 CPU 核占用的时间之和，内存占用取程序运行开始到结束占用内存的最大值</p>
<br>
<p><b>为什么同样的程序运行时间和所用内存会不同？</b></p>
<p>程序运行时间会受到许多因素的影响，尤其是在现代多任务操作系统以及在使用动态库的情况下，多次使用同一输入运行同一程序所需时间和内存有一些不同是正常现象。我们的题目给出的运行限制一般为标准程序的若干倍，也就是说，选用正确的算法和合适的语言，那么运行限制是富余的</p>
<br>
<p><b>不同语言的时间限制和内存限制是相同的吗？</b></p>
<p>Java代码的时间和内存限制一般是其他编程语言的两倍</p>
<br>
<p><b>我提交的代码可以做什么，有什么限制吗？</b></p>
<p>没有。这里没有系统调用白名单，也没有针对语言限制可使用的包或库。虽然我们比较宽容大度，但还是请不要做不符合道义的事情。如果你需要使用我们系统没有提供的某个语言的某个库，或者需要更改编译参数，可以联系我们</p>

<h2>编译器列表</h2>


<h4>gcc for C</h4>
<li>版本</li>
<pre><code>gcc (Ubuntu 4.8.2-19ubuntu1) 4.8.2
Copyright (C) 2013 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
</code></pre>
<li>编译选项</li>
<pre><code>gcc Main.c -o Main -Wall -lm --static -std=c99 -DONLINE_JUDGE</code></pre>
<hr>


<h4>g++ for C++</h4>
<li>版本</li>
<pre><code>g++ (Ubuntu 4.8.2-19ubuntu1) 4.8.2
Copyright (C) 2013 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
</code></pre>
<li>编译选项</li>
<pre><code>g++ Main.cc -o Main -Wall -lm --static -std=c++0x -DONLINE_JUDGE</code></pre>
<hr>


<h4>gcj for Java</h4>
<li>版本</li>
<pre><code>gcj-4.7 (Debian 4.7.2-3) 4.7.2
Copyright (C) 2012 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
</code></pre>
<li>编译选项</li>
<pre><code>gcj-4.7 --main=Main Main.java -o Main</code></pre>

<h2>其他问题</h2>
<p>在考试或比赛中遇到其他问题请咨询现场工作人员</p>
{{end}}