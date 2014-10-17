{{define "content"}}
{{with .Solution}}
<h1>View Code of Problem {{.Pid}}</h1>
    <textarea id="sourceCode" name="sourceCode" style="display: none;" readonly>{{.Code}}</textarea>
    <table class="CodeRay">
      <tbody>
        <tr>
        <td class="line-numbers"></td>
        <td class="code"><pre class="prettyprint linenums">{{.Code}}</pre></td>
        </tr>
      </tbody>
    </table>
    <p class="tip">Double click to view unformatted code.</p>
	<br />
	<a href="/problem/detail?pid={{.Pid}}">Back to problem {{.Pid}}</a>
    <script defer="defer" type="text/javascript">
    //<![CDATA[
      $('.CodeRay .code').dblclick(function() {
        var c = $('.CodeRay .code');
        $('#sourceCode').height(c.height()).width(c.width()).show().focus().select();
        $('.CodeRay').hide();
      });
      $('#sourceCode').blur(function(){$('.CodeRay').show(); $('#sourceCode').hide();}).hide();
    //]]>
    </script>
  {{end}}
{{end}}
