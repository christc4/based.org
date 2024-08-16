# HTTPD

This is my httpd.conf

	server "avsbq.org" {
		listen on * port 80
		listen on * tls port 443
		connection request timeout 4
		location "/pub/*" {
			root "/werc"
		}
		location found "/*" {
			root "/werc/sites/avsbq.org"
		}
		  tls {
		    certificate "/etc/ssl/avsbq.org.pem"
		    key "/etc/ssl/private/avsbq.org.key"
		  }
		 location "/.well-known/acme-challenge/*" {
		    root "/acme"
		    request strip 2
		  }
		location not found "/*" {
			root "/"
			fastcgi {
				param PLAN9 "/usr/local/plan9"
				param DOCUMENT_ROOT "/werc/bin"
				param SCRIPT_FILENAME "/werc/bin/werc.rc"
				socket "/run/slowcgi.sock"
			}
		}
		gzip-static
	}

	types {
		include "/usr/share/misc/mime.types"
	}

# ACME-CLIENT


	authority letsencrypt {
	  api url "https://acme..."
	  account key "/etc/ssl/private/letsencrypt.key"
	}

	domain avsbq.org {
	  alternative names { www.avsbq.org arc.avsbq.org mal.avsbq.org lang.avsbq.org lnk.avsbq.org ico.avsbq.org todo.avsbq.org hbk.avsbq.org vid.avsbq.org }
	  domain key "/etc/ssl/private/avsbq.org.key"
	  domain certificate "/etc/ssl/avsbq.org.crt"
	  domain full chain certificate "/etc/ssl/avsbq.org.pem"
	  sign with letsencrypt
	}

# COMMANDS

`httpd -n`

checks if your configuration is okay

`acme-client -vF avsbq.org`

request a new certificate or renew



