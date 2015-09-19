{{define "content"}}
<div class="p-faq mdl-grid">  
  <div class="page padding mdl-cell mdl-cell--8-col mdl-cell--4-col-phone mdl-shadow--2dp">
    <h3>评分</h3>
    <p>用户提交的程序经过Online Judge的即时评测，可能的反馈信息包括：</p>
    <div class="table-area mdl-cell mdl-cell--12-col mdl-cell--4-col-phone mdl-shadow--2dp">
      <table class="go-table">
        <thead>
          <tr>
            <th>结果</th>
            <th>说明</th>
          </tr>
        </thead>
        <tbody>
          <tr>
          	<td class="static-0">Pending</td>
            <td>等待评测。 请稍候</td>
          </tr>
          <tr>
          	<td class="static-1">Running&Judging</td>   
          	<td>正在评测。 请稍候</td>
          </tr>
          <tr>
          	<td class="static-2">Compile Error</td>         
           	<td>编译错误。 程序中有语法错误</td>
          </tr>
          <tr>
          	<td class="static-3">Accepted</td>    
           	<td>答案正确。 恭喜！您通过了这道题</td>
          </tr>
          <tr>
          	<td class="static-4">Presentation Error</td>
          	<td>格式错误。 输出格式不符合要求（比如空格和换行与要求不一致）</td>
          </tr> 
          <tr>
          	<td class="static-5">Runtime Error</td>
          	<td>运行错误。 可能是数组越界，堆栈溢出（比如递归调用层数太多）等情况引起</td>
          </tr> 
          <tr>
          	<td class="static-6">Wrong Answer</td>
          	<td>答案错误。 </td>
          </tr> 
          <tr>
          	<td class="static-7">Time Limit Exceeded</td>
          	<td>时间超限。 程序未能在规定时间内运行结束</td>
          </tr> 
          <tr>
          	<td class="static-8">Memory Limit Exceeded</td>
          	<td>内存超限。 程序使用了超过限制的内存</td>
          </tr> 
          <tr>
          	<td class="static-9">Output Limit Exceeded</td>
          	<td>输出超限。 程序输出的内容超出限制，可能是输出部分陷入无限循环引起</td>
          </tr> 
          <tr>
          	<td class="static-10">System Error</td>
          	<td>系统错误。 评测系统出现了错误，竞赛中将不会记录罚时</td>
          </tr> 
        </tbody>
      </table>
    </div>

    <h3>常见问题</h3>
    <p><b>我应该从哪里读输入，另外应该输出到哪里？</b></p>
    <p>如果没有特别说明，你的程序应该从标准输入（stdin，传统意义上的“键盘”）读入，并输出到标准输出（stdout，传统意义上的“屏幕”），不要使用文件做输入输出。由于系统是在你的程序运行结束后开始检查输出是否是正确的，对于有多组测试数据的输入，可以全部读入之后再输出，也可以处理一组测试数据就输出一组。</p>

    <p><b>为什么我的程序交在这里得到编译错误，而我在自己的机器上已经编译通过了？</b></p>
    <p>本系统所使用的编译器和你在自己机器上使用的可能有区别，请留意几个常见的地方：</p>
    <ul>
      <li>本评测系统运行在 32 位 Linux 系统上，使用的编译器版本和编译参数参见编译器帮助</li>
      <li>Windows平台专有的函数、数据类型、头文件不应使用。如 scanf_s(), _tmain(), _TCHAR, Windows.h 等。</li>
      <li>C 代码 main 函数必须采用 int 作为返回类型，且返回 return 0;.</li>
      <li>C++ 下 64 位整数的类型是 long long，不要使用 __int64</li>
      <li>C/C++ 代码不应包含 stdafx.h。</li>
      <li>Java 代码需使用 Main 作为主类名</li>
      <li>Visual C++ 6.0 和 Turbo C++ 3.0 （及它们的更低版本）有较多违背 C++ 标准 <a href="http://www.iso.org/iso/iso_catalogue/catalogue_ics/catalogue_detail_ics.htm?ics1=35&ics2=60&ics3=&csnumber=50372" target="_blank">ISO/IEC 14882</a> 的地方，不要使用它们来判断 C++ 程序语法上是否有问题</li>
    </ul>

    <p><b>程序的时间和内存占用是如何计算的？</b></p>
    <p>程序的运行时间为程序在所有 CPU 核占用的时间之和，内存占用取程序运行开始到结束占用内存的最大值</p>

    <p><b>为什么同样的程序运行时间和所用内存会不同？</b></p>
    <p>程序运行时间会受到许多因素的影响，尤其是在现代多任务操作系统以及在使用动态库的情况下，多次使用同一输入运行同一程序所需时间和内存有一些不同是正常现象。我们的题目给出的运行限制一般为标准程序的若干倍，也就是说，选用正确的算法和合适的语言，那么运行限制是富余的</p>

    <p><b>不同语言的时间限制和内存限制是相同的吗？</b></p>
    <p>Java代码的时间和内存限制一般是其他编程语言的两倍</p>
    
    <p><b>我提交的代码可以做什么，有什么限制吗？</b></p>
    <p>没有。这里没有系统调用白名单，也没有针对语言限制可使用的包或库。虽然我们比较宽容大度，但还是请不要做不符合道义的事情。如果你需要使用我们系统没有提供的某个语言的某个库，或者需要更改编译参数，可以联系我们</p>

    <h3>编译器列表</h3>
    <hr>
    <h4>gcc for C</h4>
    <li>版本</li>
    <div>
      <div class="code">gcc (Ubuntu 4.8.2-19ubuntu1) 4.8.2</div>
      <div class="code">Copyright (C) 2013 Free Software Foundation, Inc.</div>
      <div class="code">This is free software; see the source for copying conditions. </div>
      <div class="code">There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.</div>
    </div>
    
    <li>编译运行选项</li>
    <div>
      <div class="code">n.c -o Main -Wall -lm --static -std=c99 -DONLINE_JUDGE && ./Main</div>
    </div>
    <hr>

    <h4>g++ for C++</h4>
    <li>版本</li>
    <div>
      <div class="code">g++ (Ubuntu 4.8.2-19ubuntu1) 4.8.2</div>
      <div class="code">Copyright (C) 2013 Free Software Foundation, Inc.</div>
      <div class="code">This is free software; see the source for copying conditions.  </div>
      <div class="code">There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.</div>
    </div>
    <li>编译运行选项</li>
    <div>
      <div class="code">g++ Main.cc -o Main -Wall -lm --static -std=c++0x -DONLINE_JUDGE && ./Main</div>
    </div>
    <hr>

    <h4>OpenJDK for Java</h4>
    <li>版本</li>
    <div>
      <div class="code">java version "1.7.0_55"</div>
      <div class="code">OpenJDK Runtime Environment (IcedTea 2.4.7) (7u55-2.4.7-1ubuntu1)</div>
      <div class="code">OpenJDK Client VM (build 24.51-b03, mixed mode, sharing)</div>
    </div>
    <li>编译运行选项</li>
    <div>
      <div class="code">javac -J-Xms32m -J-Xmx256m Main.java &&/usr/bin/java -Xms128M -Xms512M -DONLINE_JUDGE=true Main</div>
    </div>

    <h3>其他</h3>
    <p>在考试或比赛中遇到其他问题请咨询现场工作人员</p>
    <p>如果对于Online Judge有任何需求或者bug report，请在Github发起一个<a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge/issues/new" target="_blank">issue</a></p><br>

    <h3>维护人员</h3>
    <a href="https://github.com/sakeven" target="_blank">@sakeven</a> 
    <a href="https://github.com/JinweiClarkChao" target="_blank">@JinweiClarkChao</a> 
    <a href="https://github.com/rex-zsd" target="_blank">@rex-zsd</a>

  </div>
</div>

{{end}}
