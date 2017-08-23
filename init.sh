#! /bin/sh 
case "$1" in 
    start) 
        echo "Starting spinnakerdemo" 
        /usr/local/bin/spinnakerdemo
        ;; 
    stop) 
        echo -n "Shutting down spinnakerdemo." 
        killproc -TERM /usr/local/bin/spinnakerdemo 
        ;; 
    *) 
        echo "Usage: $0 {start|stop}" 
        exit 1 
esac 
exit 0 
