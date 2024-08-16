# Setting up NginX webserver

I followed [landchad.net's](https://landchad.net/basic/nginx/) guide to the letter and it just worked, so here's what I did

`nginx` config files are stored in `/etc/nginx/`,

the two main subdirectories in there, atleast on Debian, are `/etc/nginx/sites-available` and `/etc/nginx/sites-enabled`

1. Create file in `/etc/nginx/sites-available`

`$EDITOR /etc/nginx/sites-available/sitename`

`sitename` is whatever you wish to name your site.

Add the following content,

	server {
		listen 80 ;
		listen [::]:80 ;
		server_name example.org ;
		root /var/www/mysite ;
		index index.html index.htm index.nginx-debian.html ;
		location / {
			try_files $uri $uri/ =404 ;
		}
	}

# Creating directory and site index

`mkdir /var/www/sitename`

Create an index.html file

`nano /var/www/sitename/index.html`


## Enable the site

`ln -s /etc/nginx/sites-available/sitename /etc/nginx/sites-enables`

Now `systemctl reload nginx`

`ufw allow 80 ufw allow 443`

Open the main Nginx config file /etc/nginx/nginx.conf and find the line # server_tokens off;. Uncomment it, and reload Nginx.

## To remove .html in url

location / {
    if ($request_uri ~ ^/(.*)\.html(\?|$)) {
        return 302 /$1;
    }
    try_files $uri $uri.html $uri/ =404;
}


https://stackoverflow.com/questions/38228393/nginx-remove-html-extension
