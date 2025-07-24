# LinkScrapper
a simple Link Scrapper in Go lang

# how to use
create a file named "list.txt" in the same folder where you put the executable 

here you can feed to it as many sites you want but they need to follow that structure:
`https://example.com/?s=;.css.class;https://example.com/`

explaination of the structure:
https://example.com/?s= is the website link, ?s= mean that the purpose of the URL is to search something
.css.class is the class that contain all the URLs you want to obtain
https://example.com is the baseURL you can not include this because its not really needed for 90% of the times but some sites you may want to use it!

pay attention! all the sites are different.

for example if i want to obtain the results from youtube.com it will be:
https://www.youtube.com/results?search_query=;.style-scope.ytd-video-renderer;https://www.youtube.com
since the search link look like this https://www.youtube.com/results?search_query=hello and the CSS this:
```
<div id="dismissible" class="style-scope ytd-video-renderer">
  <a id="thumbnail" href="/watch?v=MxU_HKk3H8k">
  </a>
</div>
```
(the html above is modified to be more clear)