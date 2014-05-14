{{define "content"}}
<h1>View Code</h1>
  {{with .Solution}}
    <textarea id="sourceCode" name="sourceCode" style="display: none;">{{.Code}}</textarea>
    <table class="CodeRay">
      <tbody>
        <tr>
        <td class="line-numbers"></td>
        <td class="code"><pre class="prettyprint linenums">{{.Code}}</pre></td>
        </tr>
      </tbody>
    </table>
    <p class="tip">Double click to view unformatted code.</p>
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