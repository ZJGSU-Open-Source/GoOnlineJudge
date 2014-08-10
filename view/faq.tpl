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
body pre code,body pre tt
{
  background-color:transparent;
  border:none;
  margin:0;
  padding:0;
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
<hr>
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