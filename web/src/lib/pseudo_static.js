const EduSoho = `
location / {
    index app.php;
    try_files $uri @rewriteapp;
}

location @rewriteapp {
    rewrite ^(.*)$ /app.php/$1 last;
}
`


const EmpireCMS = `
rewrite ^([^\.]*)/listinfo-(.+?)-(.+?)\.html$ $1/e/action/ListInfo/index.php?classid=$2&page=$3 last;
rewrite ^([^\.]*)/showinfo-(.+?)-(.+?)-(.+?)\.html$ $1/e/action/ShowInfo.php?classid=$2&id=$3&page=$4 last;
rewrite ^([^\.]*)/infotype-(.+?)-(.+?)\.html$ $1/e/action/InfoType/index.php?ttid=$2&page=$3 last;
rewrite ^([^\.]*)/tags-(.+?)-(.+?)\.html$ $1/e/tags/index.php?tagname=$2&page=$3 last;
rewrite ^([^\.]*)/comment-(.+?)-(.+?)-(.+?)-(.+?)-(.+?)-(.+?)\.html$  $1/e/pl/index\.php\?doaction=$2&classid=$3&id=$4&page=$5&myorder=$6&tempid=$7 last;
if (!-e $request_filename) {
	return 404;
}
`


const ShopWind = `
location / {
    #Redirect everything that isn't a real file to index.php
    try_files $uri $uri/ /index.php$is_args$args;
}
#If you want a single domain name at the front and back ends
location /admin {
    try_files $uri $uri/ /admin/index.php$is_args$args;
}
location /mobile {
    try_files $uri $uri/ /mobile/index.php$is_args$args;
}
location /api {
    try_files $uri $uri/ /api/index.php$is_args$args;
}
`

const crmeb = `
location / {
    if (!-e $request_filename) {
        rewrite  ^(.*)$  /index.php?s=/$1  last;
        break;
    }
 } 
`

const dabr = `
location / {
    if (!-e $request_filename) {
        rewrite ^/(.*)$ /index.php?q=$1 last;
    }
}
`


const dbshop = `
location /{
    try_files $uri $uri/ /index.php$is_args$args;
}

location ~ \.htaccess{
    deny all;
}
`


const dedecms = `
rewrite "^/list-([0-9]+)\.html$" /plus/list.php?tid=$1 last;
rewrite "^/list-([0-9]+)-([0-9]+)-([0-9]+)\.html$" /plus/list.php?tid=$1&totalresult=$2&PageNo=$3 last;
rewrite "^/view-([0-9]+)-1\.html$" /plus/view.php?arcID=$1 last;
rewrite "^/view-([0-9]+)-([0-9]+)\.html$" /plus/view.php?aid=$1&pageno=$2 last;
rewrite "^/plus/list-([0-9]+)\.html$" /plus/list.php?tid=$1 last;
rewrite "^/plus/list-([0-9]+)-([0-9]+)-([0-9]+)\.html$" /plus/list.php?tid=$1&totalresult=$2&PageNo=$3 last;
rewrite "^/plus/view-([0-9]+)-1\.html$" /plus/view.php?arcID=$1 last;
rewrite "^/plus/view-([0-9]+)-([0-9]+)\.html$" /plus/view.php?aid=$1&pageno=$2 last;
rewrite "^/tags.html$" /tags.php last;
rewrite "^/tag-([0-9]+)-([0-9]+)\.html$" /tags.php?/$1/$2/ last;

`


const discuz = `
location / {
    rewrite ^/archiver/((fid|tid)-[\w\-]+\.html)$ /archiver/index.php?$1 last;
    rewrite ^/forum-([0-9]+)-([0-9]+)\.html$ /forumdisplay.php?fid=$1&page=$2 last;
    rewrite ^/thread-([0-9]+)-([0-9]+)-([0-9]+)\.html$ /viewthread.php?tid=$1&extra=page%3D$3&page=$2 last;
    rewrite ^/space-(username|uid)-(.+)\.html$ /space.php?$1=$2 last;
    rewrite ^/tag-(.+)\.html$ /tag.php?name=$1 last;
}
`


const discuzx = `
rewrite ^([^\.]*)/topic-(.+)\.html$ $1/portal.php?mod=topic&topic=$2 last;
rewrite ^([^\.]*)/article-([0-9]+)-([0-9]+)\.html$ $1/portal.php?mod=view&aid=$2&page=$3 last;
rewrite ^([^\.]*)/forum-(\w+)-([0-9]+)\.html$ $1/forum.php?mod=forumdisplay&fid=$2&page=$3 last;
rewrite ^([^\.]*)/thread-([0-9]+)-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=viewthread&tid=$2&extra=page%3D$4&page=$3 last;
rewrite ^([^\.]*)/group-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=group&fid=$2&page=$3 last;
rewrite ^([^\.]*)/space-(username|uid)-(.+)\.html$ $1/home.php?mod=space&$2=$3 last;
rewrite ^([^\.]*)/blog-([0-9]+)-([0-9]+)\.html$ $1/home.php?mod=space&uid=$2&do=blog&id=$3 last;
rewrite ^([^\.]*)/(fid|tid)-([0-9]+)\.html$ $1/index.php?action=$2&value=$3 last;
rewrite ^([^\.]*)/([a-z]+[a-z0-9_]*)-([a-z0-9_\-]+)\.html$ $1/plugin.php?id=$2:$3 last;
if (!-e $request_filename) {
	return 404;
}
`


const discuzx2 = `
location /bbs/ {
	rewrite ^([^\.]*)/topic-(.+)\.html$ $1/portal.php?mod=topic&topic=$2 last;
	rewrite ^([^\.]*)/article-([0-9]+)-([0-9]+)\.html$ $1/portal.php?mod=view&aid=$2&page=$3 last;
	rewrite ^([^\.]*)/forum-(\w+)-([0-9]+)\.html$ $1/forum.php?mod=forumdisplay&fid=$2&page=$3 last;
	rewrite ^([^\.]*)/thread-([0-9]+)-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=viewthread&tid=$2&extra=page%3D$4&page=$3 last;
	rewrite ^([^\.]*)/group-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=group&fid=$2&page=$3 last;
	rewrite ^([^\.]*)/space-(username|uid)-(.+)\.html$ $1/home.php?mod=space&$2=$3 last;
	rewrite ^([^\.]*)/blog-([0-9]+)-([0-9]+)\.html$ $1/home.php?mod=space&uid=$2&do=blog&id=$3 last;
	rewrite ^([^\.]*)/(fid|tid)-([0-9]+)\.html$ $1/index.php?action=$2&value=$3 last;
	rewrite ^([^\.]*)/([a-z]+[a-z0-9_]*)-([a-z0-9_\-]+)\.html$ $1/plugin.php?id=$2:$3 last;
	if (!-e $request_filename) {
		return 404;
	}
}
`


const discuzx3 = `
location / {
	rewrite ^([^\.]*)/topic-(.+)\.html$ $1/portal.php?mod=topic&topic=$2 last;
	rewrite ^([^\.]*)/article-([0-9]+)-([0-9]+)\.html$ $1/portal.php?mod=view&aid=$2&page=$3 last;
	rewrite ^([^\.]*)/forum-(\w+)-([0-9]+)\.html$ $1/forum.php?mod=forumdisplay&fid=$2&page=$3 last;
	rewrite ^([^\.]*)/thread-([0-9]+)-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=viewthread&tid=$2&extra=page%3D$4&page=$3 last;
	rewrite ^([^\.]*)/group-([0-9]+)-([0-9]+)\.html$ $1/forum.php?mod=group&fid=$2&page=$3 last;
	rewrite ^([^\.]*)/space-(username|uid)-(.+)\.html$ $1/home.php?mod=space&$2=$3 last;
	rewrite ^([^\.]*)/blog-([0-9]+)-([0-9]+)\.html$ $1/home.php?mod=space&uid=$2&do=blog&id=$3 last;
	rewrite ^([^\.]*)/(fid|tid)-([0-9]+)\.html$ $1/index.php?action=$2&value=$3 last;
	rewrite ^([^\.]*)/([a-z]+[a-z0-9_]*)-([a-z0-9_\-]+)\.html$ $1/plugin.php?id=$2:$3 last;
	if (!-e $request_filename) {
			return 404;
	}
}
`


const drupal = `
if (!-e $request_filename) {
    rewrite ^/(.*)$ /index.php?q=$1 last;
}
`

const ecshop = `
if (!-e $request_filename)
{
    rewrite "^/index\.html" /index.php last;
    rewrite "^/category$" /index.php last;
    rewrite "^/feed-c([0-9]+)\.xml$" /feed.php?cat=$1 last;
    rewrite "^/feed-b([0-9]+)\.xml$" /feed.php?brand=$1 last;
    rewrite "^/feed\.xml$" /feed.php last;
    rewrite "^/category-([0-9]+)-b([0-9]+)-min([0-9]+)-max([0-9]+)-attr([^-]*)-([0-9]+)-(.+)-([a-zA-Z]+)(.*)\.html$" /category.php?id=$1&brand=$2&price_min=$3&price_max=$4&filter_attr=$5&page=$6&sort=$7&order=$8 last;
    rewrite "^/category-([0-9]+)-b([0-9]+)-min([0-9]+)-max([0-9]+)-attr([^-]*)(.*)\.html$" /category.php?id=$1&brand=$2&price_min=$3&price_max=$4&filter_attr=$5 last;
    rewrite "^/category-([0-9]+)-b([0-9]+)-([0-9]+)-(.+)-([a-zA-Z]+)(.*)\.html$" /category.php?id=$1&brand=$2&page=$3&sort=$4&order=$5 last;
    rewrite "^/category-([0-9]+)-b([0-9]+)-([0-9]+)(.*)\.html$" /category.php?id=$1&brand=$2&page=$3 last;
    rewrite "^/category-([0-9]+)-b([0-9]+)(.*)\.html$" /category.php?id=$1&brand=$2 last;
    rewrite "^/category-([0-9]+)(.*)\.html$" /category.php?id=$1 last;
    rewrite "^/goods-([0-9]+)(.*)\.html" /goods.php?id=$1 last;
    rewrite "^/article_cat-([0-9]+)-([0-9]+)-(.+)-([a-zA-Z]+)(.*)\.html$" /article_cat.php?id=$1&page=$2&sort=$3&order=$4 last;
    rewrite "^/article_cat-([0-9]+)-([0-9]+)(.*)\.html$" /article_cat.php?id=$1&page=$2 last;
    rewrite "^/article_cat-([0-9]+)(.*)\.html$" /article_cat.php?id=$1 last;
    rewrite "^/article-([0-9]+)(.*)\.html$" /article.php?id=$1 last;
    rewrite "^/brand-([0-9]+)-c([0-9]+)-([0-9]+)-(.+)-([a-zA-Z]+)\.html" /brand.php?id=$1&cat=$2&page=$3&sort=$4&order=$5 last;
    rewrite "^/brand-([0-9]+)-c([0-9]+)-([0-9]+)(.*)\.html" /brand.php?id=$1&cat=$2&page=$3 last;
    rewrite "^/brand-([0-9]+)-c([0-9]+)(.*)\.html" /brand.php?id=$1&cat=$2 last;
    rewrite "^/brand-([0-9]+)(.*)\.html" /brand.php?id=$1 last;
    rewrite "^/tag-(.*)\.html" /search.php?keywords=$1 last;
    rewrite "^/snatch-([0-9]+)\.html$" /snatch.php?id=$1 last;
    rewrite "^/group_buy-([0-9]+)\.html$" /group_buy.php?act=view&id=$1 last;
    rewrite "^/auction-([0-9]+)\.html$" /auction.php?act=view&id=$1 last;
    rewrite "^/exchange-id([0-9]+)(.*)\.html$" /exchange.php?id=$1&act=view last;
    rewrite "^/exchange-([0-9]+)-min([0-9]+)-max([0-9]+)-([0-9]+)-(.+)-([a-zA-Z]+)(.*)\.html$" /exchange.php?cat_id=$1&integral_min=$2&integral_max=$3&page=$4&sort=$5&order=$6 last;
    rewrite ^/exchange-([0-9]+)-([0-9]+)-(.+)-([a-zA-Z]+)(.*)\.html$" /exchange.php?cat_id=$1&page=$2&sort=$3&order=$4 last;
    rewrite "^/exchange-([0-9]+)-([0-9]+)(.*)\.html$" /exchange.php?cat_id=$1&page=$2 last;
    rewrite "^/exchange-([0-9]+)(.*)\.html$" /exchange.php?cat_id=$1 last;
}
`


const emlog = `
location / {
    index index.php index.html;
    if (!-e $request_filename)
    {
        rewrite ^/(.*)$ /index.php last;
    }
}
`


const laravel5 = `
location / {
    try_files $uri $uri/ /index.php?$query_string;
}
`


const maccms = `
rewrite ^/vod-(.*)$ /index.php?m=vod-$1 break;
rewrite ^/art-(.*)$ /index.php?m=art-$1 break;
rewrite ^/gbook-(.*)$ /index.php?m=gbook-$1 break;
rewrite ^/label-(.*)$ /index.php?m=label-$1 break;
rewrite ^/map-(.*)$ /index.php?m=map-$1 break;

`

const mvc = `
location /{
	if (!-e $request_filename) {
	   rewrite  ^(.*)$  /index.php/$1  last;
	   break;
	}
}
`


const niushop = `
location / {
	if (!-e $request_filename) {
		rewrite  ^(.*)$  /index.php?s=$1  last;
		break;
	}
}
`


const pbootcms = `
location / {
    if (!-e $request_filename){
           rewrite ^/index.php(.*)$ /index.php?p=$1 last;
           rewrite ^(.*)$ /index.php?s=$1 last;
    }
}
`


const phpcms = `
location / {
	###以下为PHPCMS 伪静态化rewrite法则
	rewrite ^(.*)show-([0-9]+)-([0-9]+)\.html$ $1/show.php?itemid=$2&page=$3;
	rewrite ^(.*)list-([0-9]+)-([0-9]+)\.html$ $1/list.php?catid=$2&page=$3;
	rewrite ^(.*)show-([0-9]+)\.html$ $1/show.php?specialid=$2;
	####以下为PHPWind 伪静态化rewrite法则
	rewrite ^(.*)-htm-(.*)$ $1.php?$2 last;
	rewrite ^(.*)/simple/([a-z0-9\_]+\.html)$ $1/simple/index.php?$2 last;
}
`


const phpwind = `
location / {
    rewrite ^(.*)-htm-(.*)$ $1.php?$2 last;
    rewrite ^(.*)/simple/([a-z0-9\_]+\.html)$ $1/simple/index.php?$2 last;
}
`


const sablog = `
location / {
	rewrite "^/date/([0-9]{6})/?([0-9]+)?/?$" /index.php?action=article&setdate=$1&page=$2 last;
	rewrite ^/page/([0-9]+)?/?$ /index.php?action=article&page=$1 last;
	rewrite ^/category/([0-9]+)/?([0-9]+)?/?$ /index.php?action=article&cid=$1&page=$2 last;
	rewrite ^/category/([^/]+)/?([0-9]+)?/?$ /index.php?action=article&curl=$1&page=$2 last;
	rewrite ^/(archives|search|article|links)/?$ /index.php?action=$1 last;
	rewrite ^/(comments|tagslist|trackbacks|article)/?([0-9]+)?/?$ /index.php?action=$1&page=$2 last;
	rewrite ^/tag/([^/]+)/?([0-9]+)?/?$ /index.php?action=article&item=$1&page=$2 last;
	rewrite ^/archives/([0-9]+)/?([0-9]+)?/?$ /index.php?action=show&id=$1&page=$2 last;
	rewrite ^/rss/([0-9]+)?/?$ /rss.php?cid=$1 last;
	rewrite ^/rss/([^/]+)/?$ /rss.php?url=$1 last;
	rewrite ^/uid/([0-9]+)/?([0-9]+)?/?$ /index.php?action=article&uid=$1&page=$2 last;
	rewrite ^/user/([^/]+)/?([0-9]+)?/?$ /index.php?action=article&user=$1&page=$2 last;
	rewrite sitemap.xml sitemap.php last;
	rewrite ^(.*)/([0-9a-zA-Z\-\_]+)/?([0-9]+)?/?$ $1/index.php?action=show&alias=$2&page=$3 last;
}
`


const seacms = `
location / {
	rewrite ^/frim/index(.+?)\.html$ /list/index.php?$1 last;
	rewrite ^/movie/index(.+?)\.html$ /detail/index.php?$1 last;
	rewrite ^/play/([0-9]+)-([0-9]+)-([0-9]+)\.html$ /video/index.php?$1-$2-$3 last;
	rewrite ^/topic/index(.+?)\.html$ /topic/index.php?$1 last;
	rewrite ^/topiclist/index(.+?).html$ /topiclist/index.php?$1 last;
	rewrite ^/index\.html$ index.php permanent;
	rewrite ^/news\.html$ news/ permanent;
	rewrite ^/part/index(.+?)\.html$ /articlelist/index.php?$1 last;
	rewrite ^/article/index(.+?)\.html$ /article/index.php?$1 last;
}
`


const shopex = `
location / {
    if (!-e $request_filename) {
        rewrite ^/(.+\.(html|xml|json|htm|php|jsp|asp|shtml))$ /index.php?$1 last;
    }
}
`

const thinkphp = `
location ~* (runtime|application)/{
	return 403;
}
location / {
	if (!-e $request_filename){
		rewrite  ^(.*)$  /index.php?s=$1  last;   break;
	}
}
`


const typecho = `
if (!-e $request_filename) {
    rewrite ^(.*)$ /index.php$1 last;
}
`


const typecho2 = `
location /typecho/ {
    if (!-e $request_filename) {
        rewrite ^(.*)$ /typecho/index.php$1 last;
    }
}
`

const wordpress = `
location /
{
	 try_files $uri $uri/ /index.php?$args;
}

rewrite /wp-admin$ $scheme://$host$uri/ permanent;
`

const wp2 = `
rewrite ^.*/files/(.*)$ /wp-includes/ms-files.php?file=$1 last;
if (!-e $request_filename){
	rewrite ^.+?(/wp-.*) $1 last;
	rewrite ^.+?(/.*\.php)$ $1 last;
	rewrite ^ /index.php last;
}
`

const zblog = `
if (-f $request_filename/index.html){
	rewrite (.*) $1/index.html break;
}
if (-f $request_filename/index.php){
	rewrite (.*) $1/index.php;
}
if (!-f $request_filename){
	rewrite (.*) /index.php;
}
`


export default {
    "EduSoho":EduSoho,
    "EmpireCMS":EmpireCMS,
    "ShopWind":ShopWind,
    "crmeb":crmeb,
    "dabr":dabr,
    "dbshop":dbshop,
    "dedecms":dedecms,
    "discuz":discuz,
    "discuzx":discuzx,
    "discuzx2":discuzx2,
    "discuzx3":discuzx3,
    "drupal":drupal,
    "ecshop":ecshop,
    "emlog":emlog,
    "laravel5":laravel5,
    "maccms":maccms,
    "mvc":mvc,
    "niushop":niushop,
    "pbootcms":pbootcms,
    "phpcms":phpcms,
    "phpwind":phpwind,
    "sablog":sablog,
    "seacms":seacms,
    "shopex":shopex,
    "thinkphp":thinkphp,
    "typecho":typecho,
    "typecho2":typecho2,
    "wordpress":wordpress,
    "wp2":wp2,
    "zblog":zblog  
}