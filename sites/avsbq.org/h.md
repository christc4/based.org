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

	server "arc.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/arc.avsbq.org"
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

	server "mal.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/mal.avsbq.org"
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


	server "lang.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/lang.avsbq.org"
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

	server "lnk.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/lnk.avsbq.org"
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
	server "ico.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/ico.avsbq.org"
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
	server "todo.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/todo.avsbq.org"
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
	server "hbk.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/hbk.avsbq.org"
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

	server "vid.avsbq.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/vid.avsbq.org"
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
	server "akinzon.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/akinzon.org"
		}
		  tls {
		    certificate "/etc/ssl/akinzon.org.pem"
		    key "/etc/ssl/private/akinzon.org.key"
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

	server "pariffin.org" {

		# see https://man.openbsd.org/httpd.conf to enable ssl/tls

		listen on * port 80
		listen on * tls port 443
		connection request timeout 4

		location "/pub/*" {
			root "/werc"
		}

		location found "/*" {
			root "/werc/sites/pariffin.org"
		}
		  tls {
		    certificate "/etc/ssl/pariffin.org.pem"
		    key "/etc/ssl/private/pariffin.org.key"
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
	server "mazumder.org" {
		listen on * port 80
		listen on * tls port 443
		connection request timeout 4
		location "/pub/*" {
			root "/werc"
		}
		location found "/*" {
			root "/werc/sites/mazumder.org"
		}
		  tls {
		    certificate "/etc/ssl/mazumder.org.pem"
		    key "/etc/ssl/private/mazumder.org.key"
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
	https://a.org/}
