<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, user-scalable=no, width=device-width">
    <title>{{.Title}}</title>

    <link rel="stylesheet" href="/static/css/fonts.css">
    <link rel="stylesheet" href="/static/css/material.css">
    <link rel="stylesheet" href="/static/css/GoOnlineJudge.css">
    
    <script src="/static/js/jquery.js"></script>
    <script src="/static/js/material.js"></script>
    
  </head>
  <body>

    <!-- Always shows a header, even in smaller screens. -->
    <div class="mdl-layout mdl-js-layout mdl-layout--fixed-header">

      <header class="mdl-layout__header">
        <div class="mdl-layout__header-row">
          <!-- Title -->
          <span class="mdl-layout-title">
            <a href="/" class="mdl-color-text--white">ZJGSU Online Judge</a>
          </span>
          <!-- Add spacer, to align navigation to the right -->
          <div class="mdl-layout-spacer"></div>
          <!-- Navigation. We hide it in small screens. -->
          <nav class="mdl-navigation mdl-layout--large-screen-only">
            <a class="mdl-navigation__link" href="/">Home</a>
            <a class="mdl-navigation__link" href="/admin">Admin</a>
            <a class="mdl-navigation__link" href="/admin/notice">Notice</a>
            <a class="mdl-navigation__link" href="/admin/news">News</a>
            <a class="mdl-navigation__link" href="/admin/problems">Problems</a>
            <a class="mdl-navigation__link" href="/admin/contests">Contests</a>
            <a class="mdl-navigation__link" href="/admin/users">Users</a>
            {{if .IsCurrentUser}}
            <a class="mdl-navigation__link" href="/settings">[{{.CurrentUser}}]</a>
            <a class="mdl-navigation__link J_signout" href="#">[Sign Out]</a>
          {{end}}
          </nav>
        </div>
      </header>
      <div class="mdl-layout__drawer">
        <span class="mdl-layout-title">ZJGSU OJ</span>
        <nav class="mdl-navigation">
          <a class="mdl-navigation__link" href="/">Home</a>
          <a class="mdl-navigation__link" href="/admin">Admin</a>
          <a class="mdl-navigation__link" href="/admin/notice">Notice</a>
          <a class="mdl-navigation__link" href="/admin/news">News</a>
          <a class="mdl-navigation__link" href="/admin/problems">Problems</a>
          <a class="mdl-navigation__link" href="/admin/contests/">Contests</a>
          <a class="mdl-navigation__link" href="/admin/users">Users</a>
          <div class="mdl-layout--small-screen-only">
            {{if .IsCurrentUser}}
              {{if .IsShowAdmin}}<a href="/admin/" class="mdl-navigation__link">[Admin]</a>{{end}}
              {{if .IsShowTeacher}}<a href="/admin/" class="mdl-navigation__link">[Teacher]</a>{{end}}
              <a href="/settings" class="mdl-navigation__link">[{{.CurrentUser}}]</a>
              <a href="#" class="mdl-navigation__link J_signout">[Sign Out]</a>
            {{else}}
              {{if .IsUserSignIn}}{{else}}<a href="/sess" class="mdl-navigation__link">[Sign In]</a>{{end}}
              {{if .IsUserSignUp}}{{else}}<a href="/users/new" class="mdl-navigation__link">[Sign Up]</a>{{end}}
            {{end}}
          </div>
        </nav>
      </div>

      <main class="mdl-layout__content J_main">
        
        <div class="page-content"><!-- Your content goes here -->
          {{template "content" .}}
          
        </div>

        
        <footer class="mdl-mega-footer">
          <div class="mdl-mega-footer__middle-section">

            <div class="mdl-mega-footer__drop-down-section">
              <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
              <h1 class="mdl-mega-footer__heading">FAQ</h1>
              <ul class="mdl-mega-footer__link-list">
                <li><a href="/faq">Q&A</a></li>
              </ul>
            </div>
            <div class="mdl-mega-footer__drop-down-section">
              <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
              <h1 class="mdl-mega-footer__heading">About Us</h1>
              <ul class="mdl-mega-footer__link-list">
                <li><a href="/osc">OSC</a></li>
                <li><a href="https://github.com/ZJGSU-Open-Source/GoOnlineJudge" target="_blank">GitHub</a></li>
              </ul>
            </div>
            <div class="mdl-mega-footer__drop-down-section">
              <input class="mdl-mega-footer__heading-checkbox" type="checkbox" checked>
              <h1 class="mdl-mega-footer__heading">Developer</h1>
              <ul class="mdl-mega-footer__link-list">
                <li><a href="https://github.com/memelee" target="_blank">@memelee</a></li>
                <li><a href="https://github.com/sakeven" target="_blank">@sakeven</a></li>
                <li><a href="https://github.com/JinweiClarkChao" target="_blank">@JinweiClarkChao</a></li>
                <li><a href="https://github.com/rex-zsd" target="_blank">@rex-zsd</a></li>
              </ul>
            </div>

          </div>

          <div class="mdl-mega-footer__bottom-section">
            <div class="mdl-logo">ZJGSU Online Judge Version 15.05.22</div>
            <ul class="mdl-mega-footer__link-list">
              <li><a href="#">Copyright © 2013-2015 ZJGSU ACM Club</a></li>
            </ul>
          </div>

        </footer>
      </main>
      <div class="warn J_warn"></div>
    </div>
  </body>
  <script src="/static/js/GoOnlineJudge.js"></script>
</html>