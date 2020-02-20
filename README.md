## Cache-all routing for Javascript SPA

React.js, Angular.js, Vue.js all support `window.history` as a way to have dynamic front-end URL's served from the same base `index.html` file (plus assets).

How should standard Go routers be setup to make use of a "cache-all" design to serve the index.html for every URL path except the static CSS/Javascript/Media assets (and possible Go API paths)?

## Background

Serving static assets is easy with Go. Assuming your react/angular/vue project is built/bundled into the "build" directory you could use the following to have Go serve them.

    http.FileServer(http.Dir(("build/"))

However, since most single page apps (SPA) use routing (window.history), you can't only serve assets with the same URL path as filesystem path. You also have to serve a "cache-all" `index.html` file to every possible made up URL path since the SPA will handle the rendering for that path.

This project shows how to create cache-all routes while serving other specific routes.