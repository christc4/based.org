authority letsencrypt {
  api url "https://acme-v02.api.letsencrypt.org/directory"
  account key "/etc/ssl/private/letsencrypt.key"
}

domain avsbq.org {
  alternative names { www.avsbq.org arc.avsbq.org mal.avsbq.org lang.avsbq.org lnk.avsbq.org ico.avsbq.org todo.avsbq.org hbk.avsbq.org vid.avsbq.org }
  domain key "/etc/ssl/private/avsbq.org.key"
  domain certificate "/etc/ssl/avsbq.org.crt"
  domain full chain certificate "/etc/ssl/avsbq.org.pem"
  sign with letsencrypt
}

domain pariffin.org {
  alternative names { www.pariffin.org }
  domain key "/etc/ssl/private/pariffin.org.key"
  domain certificate "/etc/ssl/pariffin.org.crt"
  domain full chain certificate "/etc/ssl/pariffin.org.pem"
  sign with letsencrypt
}


domain akinzon.org {
  domain key "/etc/ssl/private/akinzon.org.key"
  domain certificate "/etc/ssl/akinzon.org.crt"
  domain full chain certificate "/etc/ssl/akinzon.org.pem"
  sign with letsencrypt
}

domain mazumder.org {
  alternative names { www.mazumder.org }
  domain key "/etc/ssl/private/mazumder.org.key"
  domain certificate "/etc/ssl/mazumder.org.crt"
  domain full chain certificate "/etc/ssl/mazumder.org.pem"
  sign with letsencrypt
}

