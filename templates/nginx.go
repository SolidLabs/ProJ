package templates

var PhpFpmDefault = `server {
    server_name %s;
    root %s;

    location / {
        try_files $uri /app.php$is_args$args; # try to serve file directly, fallback to app.php
    }
    location ~ ^/(app_dev|config|app)\.php(/|$) {
        fastcgi_pass unix:/var/run/php-fpm.sock;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $realpath_root;
    }

    location ~ \.php$ {
        return 404;
    }

    error_log /var/log/nginx/project_error.log;
    access_log /var/log/nginx/project_access.log;
}

`
