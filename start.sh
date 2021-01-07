#!/bin/sh
/web-delay -p 8000 &
nginx -g "daemon off;"