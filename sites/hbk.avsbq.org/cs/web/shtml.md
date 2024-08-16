# Enable SSI to handle .shtml

### For nginx

What to put in your .conf

	server {
		server_name avsbq.org ;
		root /var/www/avsbq.org ;
		index index.html index.shtml ;
		location / {
		if ($request_url ~ ^/(.*)\.shtml(\?|$)) {
		return 302 /$1;
		}
			try_files $uri $uri.shtml $uri/ =404 ;
			ssi on ;
		}


How you would handle reused bits of a .shtml site

	<#include file="/headers">
	<#include file="/footers">

.shtml has been largely superseded by other mechanisms, this site used to be 100% powered by it.

You may need to enable .shtml extensions in your .htaccess file

	AddType text/html .shtml
	AddHandler server-parsed .shtml
